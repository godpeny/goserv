// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// APIResponsesColumns holds the columns for the "api_responses" table.
	APIResponsesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "code", Type: field.TypeInt},
		{Name: "type", Type: field.TypeString, Default: "unknown"},
		{Name: "message", Type: field.TypeString, Default: "unknown"},
	}
	// APIResponsesTable holds the schema information for the "api_responses" table.
	APIResponsesTable = &schema.Table{
		Name:        "api_responses",
		Columns:     APIResponsesColumns,
		PrimaryKey:  []*schema.Column{APIResponsesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// ProjectsColumns holds the columns for the "projects" table.
	ProjectsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// ProjectsTable holds the schema information for the "projects" table.
	ProjectsTable = &schema.Table{
		Name:        "projects",
		Columns:     ProjectsColumns,
		PrimaryKey:  []*schema.Column{ProjectsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "age", Type: field.TypeInt},
		{Name: "name", Type: field.TypeString, Default: "unknown"},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:        "users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		APIResponsesTable,
		ProjectsTable,
		UsersTable,
	}
)

func init() {
}
