package postgres

import (
	"context"
	"time"

	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/pkg/auth/permissions"
	ops "github.com/stackrox/rox/pkg/metrics"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/pgutils"
	"github.com/stackrox/rox/pkg/postgres/walker"
	"github.com/stackrox/rox/pkg/search"
)

const (
	cursorBatchSize = 50
)

// PermissionChecker is a permission checker that could be used by GenericStore
type PermissionChecker interface {
	CountAllowed(ctx context.Context) (bool, error)
	ExistsAllowed(ctx context.Context) (bool, error)
	WalkAllowed(ctx context.Context) (bool, error)
}

type primaryKeyGetter[T any, PT unmarshaler[T]] func(obj PT) string
type durationTimeSetter func(start time.Time, op ops.Op)

// GenericStore implements subset of Store interface for resources with single ID.
type GenericStore[T any, PT unmarshaler[T]] struct {
	db                               postgres.DB
	schema                           *walker.Schema
	pkGetter                         primaryKeyGetter[T, PT]
	setAcquireDBConnDuration         durationTimeSetter
	setPostgresOperationDurationTime durationTimeSetter
	permissionChecker                PermissionChecker
	targetResource                   permissions.ResourceMetadata
}

// NewGenericStore returns new subStore implementation for given resource.
// subStore implements subset of Store operations.
func NewGenericStore[T any, PT unmarshaler[T]](
	db postgres.DB,
	schema *walker.Schema,
	pkGetter primaryKeyGetter[T, PT],
	setAcquireDBConnDuration durationTimeSetter,
	setPostgresOperationDurationTime durationTimeSetter,
	targetResource permissions.ResourceMetadata,
) *GenericStore[T, PT] {
	return &GenericStore[T, PT]{
		db:                               db,
		schema:                           schema,
		pkGetter:                         pkGetter,
		setAcquireDBConnDuration:         setAcquireDBConnDuration,
		setPostgresOperationDurationTime: setPostgresOperationDurationTime,
		targetResource:                   targetResource,
	}
}

// NewGenericStoreWithPermissionChecker returns new subStore implementation for given resource.
// subStore implements subset of Store operations.
func NewGenericStoreWithPermissionChecker[T any, PT unmarshaler[T]](
	db postgres.DB,
	schema *walker.Schema,
	pkGetter primaryKeyGetter[T, PT],
	setAcquireDBConnDuration durationTimeSetter,
	setPostgresOperationDurationTime durationTimeSetter,
	checker PermissionChecker,
) *GenericStore[T, PT] {
	return &GenericStore[T, PT]{
		db:                               db,
		schema:                           schema,
		pkGetter:                         pkGetter,
		setAcquireDBConnDuration:         setAcquireDBConnDuration,
		setPostgresOperationDurationTime: setPostgresOperationDurationTime,
		permissionChecker:                checker,
	}
}

func (s *GenericStore[T, PT]) hasPermissionsChecker() bool {
	return s.permissionChecker != nil
}

// AcquireConn returns Acquires new connection from DB.
func (s *GenericStore[T, PT]) AcquireConn(ctx context.Context, op ops.Op) (*postgres.Conn, error) {
	defer s.setAcquireDBConnDuration(time.Now(), op)
	conn, err := s.db.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// Exists tells whether the ID exists in the store.
func (s *GenericStore[T, PT]) Exists(ctx context.Context, id string) (bool, error) {
	defer s.setPostgresOperationDurationTime(time.Now(), ops.Exists)

	var sacQueryFilter *v1.Query
	if s.hasPermissionsChecker() {
		if ok, err := s.permissionChecker.ExistsAllowed(ctx); err != nil {
			return false, err
		} else if !ok {
			return false, nil
		}
	} else {
		filter, err := GetReadSACQuery(ctx, s.targetResource)
		if err != nil {
			return false, err
		}
		sacQueryFilter = filter
	}

	q := search.ConjunctionQuery(
		sacQueryFilter,
		search.NewQueryBuilder().AddDocIDs(id).ProtoQuery(),
	)

	count, err := RunCountRequestForSchema(ctx, s.schema, q, s.db)
	// With joins and multiple paths to the scoping resources, it can happen that the Count query for an object identifier
	// returns more than 1, despite the fact that the identifier is unique in the table.
	return count > 0, err
}

// Count returns the number of objects in the store.
func (s *GenericStore[T, PT]) Count(ctx context.Context) (int, error) {
	defer s.setPostgresOperationDurationTime(time.Now(), ops.Count)

	var sacQueryFilter *v1.Query
	if s.hasPermissionsChecker() {
		if ok, err := s.permissionChecker.CountAllowed(ctx); err != nil || !ok {
			return 0, err
		}
	} else {
		filter, err := GetReadSACQuery(ctx, s.targetResource)
		if err != nil {
			return 0, err
		}
		sacQueryFilter = filter
	}

	return RunCountRequestForSchema(ctx, s.schema, sacQueryFilter, s.db)
}

// Walk iterates over all the objects in the store and applies the closure.
func (s *GenericStore[T, PT]) Walk(ctx context.Context, fn func(obj PT) error) error {
	var sacQueryFilter *v1.Query
	if s.hasPermissionsChecker() {
		if ok, err := s.permissionChecker.WalkAllowed(ctx); err != nil || !ok {
			return err
		}
	} else {
		filter, err := GetReadSACQuery(ctx, s.targetResource)
		if err != nil {
			return err
		}
		sacQueryFilter = filter
	}
	fetcher, closer, err := RunCursorQueryForSchema[T, PT](ctx, s.schema, sacQueryFilter, s.db)
	if err != nil {
		return err
	}
	defer closer()
	for {
		rows, err := fetcher(cursorBatchSize)
		if err != nil {
			return pgutils.ErrNilIfNoRows(err)
		}
		for _, data := range rows {
			if err := fn(data); err != nil {
				return err
			}
		}
		if len(rows) != cursorBatchSize {
			break
		}
	}
	return nil
}
