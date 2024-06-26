// Code generated by pg-bindings generator. DO NOT EDIT.

package schema

import (
	"reflect"
	"time"

	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/walker"
	"github.com/stackrox/rox/pkg/sac/resources"
	"github.com/stackrox/rox/pkg/search"
	"github.com/stackrox/rox/pkg/search/postgres/mapping"
)

var (
	// CreateTableSecuredUnitsStmt holds the create statement for table `secured_units`.
	CreateTableSecuredUnitsStmt = &postgres.CreateStmts{
		GormModel: (*SecuredUnits)(nil),
		Children:  []*postgres.CreateStmts{},
	}

	// SecuredUnitsSchema is the go schema for table `secured_units`.
	SecuredUnitsSchema = func() *walker.Schema {
		schema := GetSchemaForTable("secured_units")
		if schema != nil {
			return schema
		}
		schema = walker.Walk(reflect.TypeOf((*storage.SecuredUnits)(nil)), "secured_units")
		schema.SetOptionsMap(search.Walk(v1.SearchCategory_ADMINISTRATION_USAGE, "securedunits", (*storage.SecuredUnits)(nil)))
		schema.ScopingResource = resources.Administration
		RegisterTable(schema, CreateTableSecuredUnitsStmt)
		mapping.RegisterCategoryToTable(v1.SearchCategory_ADMINISTRATION_USAGE, schema)
		return schema
	}()
)

const (
	// SecuredUnitsTableName specifies the name of the table in postgres.
	SecuredUnitsTableName = "secured_units"
)

// SecuredUnits holds the Gorm model for Postgres table `secured_units`.
type SecuredUnits struct {
	ID          string     `gorm:"column:id;type:uuid;primaryKey"`
	Timestamp   *time.Time `gorm:"column:timestamp;type:timestamp;unique"`
	NumNodes    int64      `gorm:"column:numnodes;type:bigint"`
	NumCPUUnits int64      `gorm:"column:numcpuunits;type:bigint"`
	Serialized  []byte     `gorm:"column:serialized;type:bytea"`
}
