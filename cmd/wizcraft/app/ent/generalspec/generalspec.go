// Code generated by ent, DO NOT EDIT.

package generalspec

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the generalspec type in the database.
	Label = "general_spec"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// EdgeService holds the string denoting the service edge name in mutations.
	EdgeService = "service"
	// EdgeDatabase holds the string denoting the database edge name in mutations.
	EdgeDatabase = "database"
	// EdgeApispec holds the string denoting the apispec edge name in mutations.
	EdgeApispec = "apispec"
	// EdgeProject holds the string denoting the project edge name in mutations.
	EdgeProject = "project"
	// Table holds the table name of the generalspec in the database.
	Table = "general_specs"
	// ServiceTable is the table that holds the service relation/edge.
	ServiceTable = "general_specs"
	// ServiceInverseTable is the table name for the Service entity.
	// It exists in this package in order to avoid circular dependency with the "service" package.
	ServiceInverseTable = "services"
	// ServiceColumn is the table column denoting the service relation/edge.
	ServiceColumn = "general_spec_service"
	// DatabaseTable is the table that holds the database relation/edge.
	DatabaseTable = "general_specs"
	// DatabaseInverseTable is the table name for the Database entity.
	// It exists in this package in order to avoid circular dependency with the "database" package.
	DatabaseInverseTable = "databases"
	// DatabaseColumn is the table column denoting the database relation/edge.
	DatabaseColumn = "general_spec_database"
	// ApispecTable is the table that holds the apispec relation/edge.
	ApispecTable = "general_specs"
	// ApispecInverseTable is the table name for the APISpec entity.
	// It exists in this package in order to avoid circular dependency with the "apispec" package.
	ApispecInverseTable = "api_specs"
	// ApispecColumn is the table column denoting the apispec relation/edge.
	ApispecColumn = "general_spec_apispec"
	// ProjectTable is the table that holds the project relation/edge.
	ProjectTable = "general_specs"
	// ProjectInverseTable is the table name for the Project entity.
	// It exists in this package in order to avoid circular dependency with the "project" package.
	ProjectInverseTable = "projects"
	// ProjectColumn is the table column denoting the project relation/edge.
	ProjectColumn = "general_spec_project"
)

// Columns holds all SQL columns for generalspec fields.
var Columns = []string{
	FieldID,
	FieldUUID,
	FieldName,
	FieldType,
	FieldStatus,
	FieldDescription,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "general_specs"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"general_spec_service",
	"general_spec_database",
	"general_spec_apispec",
	"general_spec_project",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultUUID holds the default value on creation for the "uuid" field.
	DefaultUUID func() uuid.UUID
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus string
)

// OrderOption defines the ordering options for the GeneralSpec queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUUID orders the results by the uuid field.
func ByUUID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUUID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByServiceField orders the results by service field.
func ByServiceField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newServiceStep(), sql.OrderByField(field, opts...))
	}
}

// ByDatabaseField orders the results by database field.
func ByDatabaseField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDatabaseStep(), sql.OrderByField(field, opts...))
	}
}

// ByApispecField orders the results by apispec field.
func ByApispecField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newApispecStep(), sql.OrderByField(field, opts...))
	}
}

// ByProjectField orders the results by project field.
func ByProjectField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProjectStep(), sql.OrderByField(field, opts...))
	}
}
func newServiceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ServiceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, ServiceTable, ServiceColumn),
	)
}
func newDatabaseStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DatabaseInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, DatabaseTable, DatabaseColumn),
	)
}
func newApispecStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ApispecInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, ApispecTable, ApispecColumn),
	)
}
func newProjectStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProjectInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, ProjectTable, ProjectColumn),
	)
}
