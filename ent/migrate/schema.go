// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// FilesColumns holds the columns for the "files" table.
	FilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "src_path", Type: field.TypeString},
		{Name: "current_path", Type: field.TypeString},
		{Name: "filetype", Type: field.TypeString},
	}
	// FilesTable holds the schema information for the "files" table.
	FilesTable = &schema.Table{
		Name:       "files",
		Columns:    FilesColumns,
		PrimaryKey: []*schema.Column{FilesColumns[0]},
	}
	// KeysColumns holds the columns for the "keys" table.
	KeysColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "type", Type: field.TypeString},
		{Name: "public_key", Type: field.TypeString},
		{Name: "private_key", Type: field.TypeString},
		{Name: "is_active", Type: field.TypeBool},
		{Name: "references", Type: field.TypeInt},
	}
	// KeysTable holds the schema information for the "keys" table.
	KeysTable = &schema.Table{
		Name:       "keys",
		Columns:    KeysColumns,
		PrimaryKey: []*schema.Column{KeysColumns[0]},
	}
	// MasterPasswordsColumns holds the columns for the "master_passwords" table.
	MasterPasswordsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "password_hash", Type: field.TypeString},
	}
	// MasterPasswordsTable holds the schema information for the "master_passwords" table.
	MasterPasswordsTable = &schema.Table{
		Name:       "master_passwords",
		Columns:    MasterPasswordsColumns,
		PrimaryKey: []*schema.Column{MasterPasswordsColumns[0]},
	}
	// PasswordItemsColumns holds the columns for the "password_items" table.
	PasswordItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "avatar", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "site_name", Type: field.TypeString, Nullable: true},
		{Name: "site_url", Type: field.TypeString, Nullable: true},
		{Name: "username", Type: field.TypeString, Nullable: true},
		{Name: "username_type", Type: field.TypeString, Nullable: true},
		{Name: "password", Type: field.TypeString, Nullable: true},
		{Name: "tags", Type: field.TypeJSON, Nullable: true},
	}
	// PasswordItemsTable holds the schema information for the "password_items" table.
	PasswordItemsTable = &schema.Table{
		Name:       "password_items",
		Columns:    PasswordItemsColumns,
		PrimaryKey: []*schema.Column{PasswordItemsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		FilesTable,
		KeysTable,
		MasterPasswordsTable,
		PasswordItemsTable,
	}
)

func init() {
}
