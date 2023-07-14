// Code generated by pg-bindings generator. DO NOT EDIT.

package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/stackrox/rox/central/metrics"
	"github.com/stackrox/rox/central/role/resources"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/logging"
	ops "github.com/stackrox/rox/pkg/metrics"
	"github.com/stackrox/rox/pkg/postgres"
	pkgSchema "github.com/stackrox/rox/pkg/postgres/schema"
	pgSearch "github.com/stackrox/rox/pkg/search/postgres"
	"github.com/stackrox/rox/pkg/sync"
	"gorm.io/gorm"
)

const (
	baseTable = "compliance_operator_rule_v2"
	storeName = "ComplianceOperatorRuleV2"

	// using copyFrom, we may not even want to batch.  It would probably be simpler
	// to deal with failures if we just sent it all.  Something to think about as we
	// proceed and move into more e2e and larger performance testing
	batchSize = 10000
)

var (
	log            = logging.LoggerForModule()
	schema         = pkgSchema.ComplianceOperatorRuleV2Schema
	targetResource = resources.ComplianceOperator
)

type storeType = storage.ComplianceOperatorRuleV2

// Store is the interface to interact with the storage for storage.ComplianceOperatorRuleV2
type Store interface {
	Upsert(ctx context.Context, obj *storeType) error
	UpsertMany(ctx context.Context, objs []*storeType) error
	Delete(ctx context.Context, name string) error
	DeleteByQuery(ctx context.Context, q *v1.Query) error
	DeleteMany(ctx context.Context, identifiers []string) error

	Count(ctx context.Context) (int, error)
	Exists(ctx context.Context, name string) (bool, error)

	Get(ctx context.Context, name string) (*storeType, bool, error)
	GetMany(ctx context.Context, identifiers []string) ([]*storeType, []int, error)
	GetIDs(ctx context.Context) ([]string, error)

	Walk(ctx context.Context, fn func(obj *storeType) error) error
}

type storeImpl struct {
	*pgSearch.GenericStore[storeType, *storeType]
	mutex sync.RWMutex
}

// New returns a new Store instance using the provided sql instance.
func New(db postgres.DB) Store {
	return &storeImpl{
		GenericStore: pgSearch.NewGenericStore[storeType, *storeType](
			db,
			schema,
			pkGetter,
			insertIntoComplianceOperatorRuleV2,
			copyFromComplianceOperatorRuleV2,
			metricsSetAcquireDBConnDuration,
			metricsSetPostgresOperationDurationTime,
			pgSearch.GloballyScopedUpsertChecker[storeType, *storeType](targetResource),
			targetResource,
		),
	}
}

// region Helper functions

func pkGetter(obj *storeType) string {
	return obj.GetName()
}

func metricsSetPostgresOperationDurationTime(start time.Time, op ops.Op) {
	metrics.SetPostgresOperationDurationTime(start, op, storeName)
}

func metricsSetAcquireDBConnDuration(start time.Time, op ops.Op) {
	metrics.SetAcquireDBConnDuration(start, op, storeName)
}

func insertIntoComplianceOperatorRuleV2(_ context.Context, batch *pgx.Batch, obj *storage.ComplianceOperatorRuleV2) error {

	serialized, marshalErr := obj.Marshal()
	if marshalErr != nil {
		return marshalErr
	}

	values := []interface{}{
		// parent primary keys start
		obj.GetName(),
		obj.GetVersion(),
		obj.GetRuleVersion(),
		obj.GetRuleType(),
		obj.GetSeverity(),
		serialized,
	}

	finalStr := "INSERT INTO compliance_operator_rule_v2 (Name, Version, RuleVersion, RuleType, Severity, serialized) VALUES($1, $2, $3, $4, $5, $6) ON CONFLICT(Name) DO UPDATE SET Name = EXCLUDED.Name, Version = EXCLUDED.Version, RuleVersion = EXCLUDED.RuleVersion, RuleType = EXCLUDED.RuleType, Severity = EXCLUDED.Severity, serialized = EXCLUDED.serialized"
	batch.Queue(finalStr, values...)

	return nil
}

func copyFromComplianceOperatorRuleV2(ctx context.Context, s pgSearch.Deleter, tx *postgres.Tx, objs ...*storage.ComplianceOperatorRuleV2) error {

	inputRows := [][]interface{}{}

	var err error

	// This is a copy so first we must delete the rows and re-add them
	// Which is essentially the desired behaviour of an upsert.
	var deletes []string

	copyCols := []string{
		"name",
		"version",
		"ruleversion",
		"ruletype",
		"severity",
		"serialized",
	}

	for idx, obj := range objs {
		// Todo: ROX-9499 Figure out how to more cleanly template around this issue.
		log.Debugf("This is here for now because there is an issue with pods_TerminatedInstances where the obj "+
			"in the loop is not used as it only consists of the parent ID and the index.  Putting this here as a stop gap "+
			"to simply use the object.  %s", obj)

		serialized, marshalErr := obj.Marshal()
		if marshalErr != nil {
			return marshalErr
		}

		inputRows = append(inputRows, []interface{}{
			obj.GetName(),
			obj.GetVersion(),
			obj.GetRuleVersion(),
			obj.GetRuleType(),
			obj.GetSeverity(),
			serialized,
		})

		// Add the ID to be deleted.
		deletes = append(deletes, obj.GetName())

		// if we hit our batch size we need to push the data
		if (idx+1)%batchSize == 0 || idx == len(objs)-1 {
			// copy does not upsert so have to delete first.  parent deletion cascades so only need to
			// delete for the top level parent

			if err := s.DeleteMany(ctx, deletes); err != nil {
				return err
			}
			// clear the inserts and vals for the next batch
			deletes = nil

			_, err = tx.CopyFrom(ctx, pgx.Identifier{"compliance_operator_rule_v2"}, copyCols, pgx.CopyFromRows(inputRows))

			if err != nil {
				return err
			}

			// clear the input rows for the next batch
			inputRows = inputRows[:0]
		}
	}

	return err
}

// endregion Helper functions

// region Used for testing

// CreateTableAndNewStore returns a new Store instance for testing.
func CreateTableAndNewStore(ctx context.Context, db postgres.DB, gormDB *gorm.DB) Store {
	pkgSchema.ApplySchemaForTable(ctx, gormDB, baseTable)
	return New(db)
}

// Destroy drops the tables associated with the target object type.
func Destroy(ctx context.Context, db postgres.DB) {
	dropTableComplianceOperatorRuleV2(ctx, db)
}

func dropTableComplianceOperatorRuleV2(ctx context.Context, db postgres.DB) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS compliance_operator_rule_v2 CASCADE")

}

// endregion Used for testing
