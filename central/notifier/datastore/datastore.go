package datastore

import (
	"context"

	"github.com/stackrox/rox/central/notifier/datastore/internal/store"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
)

// DataStore provides storage functionality for notifiers.
//go:generate mockgen-wrapper
type DataStore interface {
	GetNotifier(ctx context.Context, id string) (*storage.Notifier, bool, error)
	GetNotifiers(ctx context.Context, request *v1.GetNotifiersRequest) ([]*storage.Notifier, error)
	AddNotifier(ctx context.Context, notifier *storage.Notifier) (string, error)
	UpdateNotifier(ctx context.Context, notifier *storage.Notifier) error
	RemoveNotifier(ctx context.Context, id string) error
}

// New returns a new Store instance using the provided bolt DB instance.
func New(storage store.Store) DataStore {
	return &datastoreImpl{
		storage: storage,
	}
}
