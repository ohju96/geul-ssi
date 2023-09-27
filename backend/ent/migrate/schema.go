// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// EventsColumns holds the columns for the "events" table.
	EventsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "contents", Type: field.TypeString},
		{Name: "writer", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// EventsTable holds the schema information for the "events" table.
	EventsTable = &schema.Table{
		Name:       "events",
		Columns:    EventsColumns,
		PrimaryKey: []*schema.Column{EventsColumns[0]},
	}
	// HeartsColumns holds the columns for the "hearts" table.
	HeartsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "event_id", Type: field.TypeInt, Unique: true},
		{Name: "writer", Type: field.TypeString},
		{Name: "is_heart", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
	}
	// HeartsTable holds the schema information for the "hearts" table.
	HeartsTable = &schema.Table{
		Name:       "hearts",
		Columns:    HeartsColumns,
		PrimaryKey: []*schema.Column{HeartsColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "nickname", Type: field.TypeString, Unique: true},
		{Name: "secret_code", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// WiseSayingsColumns holds the columns for the "wise_sayings" table.
	WiseSayingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "wise_saying", Type: field.TypeString},
		{Name: "created_by", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// WiseSayingsTable holds the schema information for the "wise_sayings" table.
	WiseSayingsTable = &schema.Table{
		Name:       "wise_sayings",
		Columns:    WiseSayingsColumns,
		PrimaryKey: []*schema.Column{WiseSayingsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		EventsTable,
		HeartsTable,
		UsersTable,
		WiseSayingsTable,
	}
)

func init() {
}