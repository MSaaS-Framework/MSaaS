// Code generated by ent, DO NOT EDIT.

package user

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// FieldActive holds the string denoting the active field in the database.
	FieldActive = "active"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldLastLogin holds the string denoting the last_login field in the database.
	FieldLastLogin = "last_login"
	// FieldProfileImage holds the string denoting the profile_image field in the database.
	FieldProfileImage = "profile_image"
	// FieldGithubTokenHash holds the string denoting the github_token_hash field in the database.
	FieldGithubTokenHash = "github_token_hash"
	// EdgePermissions holds the string denoting the permissions edge name in mutations.
	EdgePermissions = "permissions"
	// EdgeProjects holds the string denoting the projects edge name in mutations.
	EdgeProjects = "projects"
	// Table holds the table name of the user in the database.
	Table = "users"
	// PermissionsTable is the table that holds the permissions relation/edge.
	PermissionsTable = "user_general_spec_permissions"
	// PermissionsInverseTable is the table name for the UserGeneralSpecPermissions entity.
	// It exists in this package in order to avoid circular dependency with the "usergeneralspecpermissions" package.
	PermissionsInverseTable = "user_general_spec_permissions"
	// PermissionsColumn is the table column denoting the permissions relation/edge.
	PermissionsColumn = "user_permissions"
	// ProjectsTable is the table that holds the projects relation/edge. The primary key declared below.
	ProjectsTable = "user_projects"
	// ProjectsInverseTable is the table name for the Project entity.
	// It exists in this package in order to avoid circular dependency with the "project" package.
	ProjectsInverseTable = "projects"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldEmail,
	FieldPassword,
	FieldRole,
	FieldActive,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldLastLogin,
	FieldProfileImage,
	FieldGithubTokenHash,
}

var (
	// ProjectsPrimaryKey and ProjectsColumn2 are the table columns denoting the
	// primary key for the projects relation (M2M).
	ProjectsPrimaryKey = []string{"user_id", "project_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
	// DefaultActive holds the default value on creation for the "active" field.
	DefaultActive bool
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Role defines the type for the "role" enum field.
type Role string

// RoleDeveloper is the default value of the Role enum.
const DefaultRole = RoleDeveloper

// Role values.
const (
	RoleAdmin     Role = "admin"
	RoleDeveloper Role = "developer"
	RoleEngineer  Role = "engineer"
)

func (r Role) String() string {
	return string(r)
}

// RoleValidator is a validator for the "role" field enum values. It is called by the builders before save.
func RoleValidator(r Role) error {
	switch r {
	case RoleAdmin, RoleDeveloper, RoleEngineer:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for role field: %q", r)
	}
}

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByRole orders the results by the role field.
func ByRole(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRole, opts...).ToFunc()
}

// ByActive orders the results by the active field.
func ByActive(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldActive, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByLastLogin orders the results by the last_login field.
func ByLastLogin(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastLogin, opts...).ToFunc()
}

// ByProfileImage orders the results by the profile_image field.
func ByProfileImage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProfileImage, opts...).ToFunc()
}

// ByGithubTokenHash orders the results by the github_token_hash field.
func ByGithubTokenHash(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGithubTokenHash, opts...).ToFunc()
}

// ByPermissionsCount orders the results by permissions count.
func ByPermissionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPermissionsStep(), opts...)
	}
}

// ByPermissions orders the results by permissions terms.
func ByPermissions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPermissionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByProjectsCount orders the results by projects count.
func ByProjectsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newProjectsStep(), opts...)
	}
}

// ByProjects orders the results by projects terms.
func ByProjects(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProjectsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newPermissionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PermissionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PermissionsTable, PermissionsColumn),
	)
}
func newProjectsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProjectsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, ProjectsTable, ProjectsPrimaryKey...),
	)
}
