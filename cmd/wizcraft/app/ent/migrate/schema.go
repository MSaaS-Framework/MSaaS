// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// APISpecsColumns holds the columns for the "api_specs" table.
	APISpecsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "openapi_spec", Type: field.TypeJSON},
		{Name: "general_spec_apispec", Type: field.TypeInt, Unique: true},
		{Name: "service_apispec", Type: field.TypeUUID, Unique: true},
	}
	// APISpecsTable holds the schema information for the "api_specs" table.
	APISpecsTable = &schema.Table{
		Name:       "api_specs",
		Columns:    APISpecsColumns,
		PrimaryKey: []*schema.Column{APISpecsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "api_specs_general_specs_apispec",
				Columns:    []*schema.Column{APISpecsColumns[2]},
				RefColumns: []*schema.Column{GeneralSpecsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "api_specs_services_apispec",
				Columns:    []*schema.Column{APISpecsColumns[3]},
				RefColumns: []*schema.Column{ServicesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// DatabasesColumns holds the columns for the "databases" table.
	DatabasesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "connection_path", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
		{Name: "db_type", Type: field.TypeString},
		{Name: "general_spec_database", Type: field.TypeInt, Unique: true},
		{Name: "service_databases", Type: field.TypeUUID, Nullable: true},
	}
	// DatabasesTable holds the schema information for the "databases" table.
	DatabasesTable = &schema.Table{
		Name:       "databases",
		Columns:    DatabasesColumns,
		PrimaryKey: []*schema.Column{DatabasesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "databases_general_specs_database",
				Columns:    []*schema.Column{DatabasesColumns[4]},
				RefColumns: []*schema.Column{GeneralSpecsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "databases_services_databases",
				Columns:    []*schema.Column{DatabasesColumns[5]},
				RefColumns: []*schema.Column{ServicesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// GeneralSpecsColumns holds the columns for the "general_specs" table.
	GeneralSpecsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uuid", Type: field.TypeUUID},
		{Name: "project_uuid", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "type", Type: field.TypeString},
		{Name: "status", Type: field.TypeString, Default: "created"},
		{Name: "description", Type: field.TypeString},
		{Name: "project_general_specs", Type: field.TypeUUID},
	}
	// GeneralSpecsTable holds the schema information for the "general_specs" table.
	GeneralSpecsTable = &schema.Table{
		Name:       "general_specs",
		Columns:    GeneralSpecsColumns,
		PrimaryKey: []*schema.Column{GeneralSpecsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "general_specs_projects_general_specs",
				Columns:    []*schema.Column{GeneralSpecsColumns[7]},
				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ProjectsColumns holds the columns for the "projects" table.
	ProjectsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
	}
	// ProjectsTable holds the schema information for the "projects" table.
	ProjectsTable = &schema.Table{
		Name:       "projects",
		Columns:    ProjectsColumns,
		PrimaryKey: []*schema.Column{ProjectsColumns[0]},
	}
	// ServicesColumns holds the columns for the "services" table.
	ServicesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "general_spec_service", Type: field.TypeInt, Unique: true},
	}
	// ServicesTable holds the schema information for the "services" table.
	ServicesTable = &schema.Table{
		Name:       "services",
		Columns:    ServicesColumns,
		PrimaryKey: []*schema.Column{ServicesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "services_general_specs_service",
				Columns:    []*schema.Column{ServicesColumns[1]},
				RefColumns: []*schema.Column{GeneralSpecsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "name", Type: field.TypeString, Default: "unknown"},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"admin", "developer", "engineer"}, Default: "developer"},
		{Name: "active", Type: field.TypeBool, Default: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "last_login", Type: field.TypeTime, Nullable: true},
		{Name: "profile_image", Type: field.TypeString, Nullable: true},
		{Name: "github_token_hash", Type: field.TypeString, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// UserGeneralSpecPermissionsColumns holds the columns for the "user_general_spec_permissions" table.
	UserGeneralSpecPermissionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "permission", Type: field.TypeEnum, Enums: []string{"read", "write", "manage"}, Default: "read"},
		{Name: "general_spec_permissions", Type: field.TypeInt},
		{Name: "user_permissions", Type: field.TypeUUID},
	}
	// UserGeneralSpecPermissionsTable holds the schema information for the "user_general_spec_permissions" table.
	UserGeneralSpecPermissionsTable = &schema.Table{
		Name:       "user_general_spec_permissions",
		Columns:    UserGeneralSpecPermissionsColumns,
		PrimaryKey: []*schema.Column{UserGeneralSpecPermissionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_general_spec_permissions_general_specs_permissions",
				Columns:    []*schema.Column{UserGeneralSpecPermissionsColumns[2]},
				RefColumns: []*schema.Column{GeneralSpecsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "user_general_spec_permissions_users_permissions",
				Columns:    []*schema.Column{UserGeneralSpecPermissionsColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UserProjectsColumns holds the columns for the "user_projects" table.
	UserProjectsColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "project_id", Type: field.TypeUUID},
	}
	// UserProjectsTable holds the schema information for the "user_projects" table.
	UserProjectsTable = &schema.Table{
		Name:       "user_projects",
		Columns:    UserProjectsColumns,
		PrimaryKey: []*schema.Column{UserProjectsColumns[0], UserProjectsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_projects_user_id",
				Columns:    []*schema.Column{UserProjectsColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_projects_project_id",
				Columns:    []*schema.Column{UserProjectsColumns[1]},
				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		APISpecsTable,
		DatabasesTable,
		GeneralSpecsTable,
		ProjectsTable,
		ServicesTable,
		UsersTable,
		UserGeneralSpecPermissionsTable,
		UserProjectsTable,
	}
)

func init() {
	APISpecsTable.ForeignKeys[0].RefTable = GeneralSpecsTable
	APISpecsTable.ForeignKeys[1].RefTable = ServicesTable
	DatabasesTable.ForeignKeys[0].RefTable = GeneralSpecsTable
	DatabasesTable.ForeignKeys[1].RefTable = ServicesTable
	GeneralSpecsTable.ForeignKeys[0].RefTable = ProjectsTable
	ServicesTable.ForeignKeys[0].RefTable = GeneralSpecsTable
	UserGeneralSpecPermissionsTable.ForeignKeys[0].RefTable = GeneralSpecsTable
	UserGeneralSpecPermissionsTable.ForeignKeys[1].RefTable = UsersTable
	UserProjectsTable.ForeignKeys[0].RefTable = UsersTable
	UserProjectsTable.ForeignKeys[1].RefTable = ProjectsTable
}
