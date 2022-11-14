// Code generated by ent, DO NOT EDIT.

package passworditem

import (
	"time"
)

const (
	// Label holds the string label denoting the passworditem type in the database.
	Label = "password_item"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAvatar holds the string denoting the avatar field in the database.
	FieldAvatar = "avatar"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldSiteName holds the string denoting the site_name field in the database.
	FieldSiteName = "site_name"
	// FieldSiteURL holds the string denoting the site_url field in the database.
	FieldSiteURL = "site_url"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldUsernameType holds the string denoting the username_type field in the database.
	FieldUsernameType = "username_type"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldTags holds the string denoting the tags field in the database.
	FieldTags = "tags"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// Table holds the table name of the passworditem in the database.
	Table = "password_items"
)

// Columns holds all SQL columns for passworditem fields.
var Columns = []string{
	FieldID,
	FieldAvatar,
	FieldDescription,
	FieldSiteName,
	FieldSiteURL,
	FieldUsername,
	FieldUsernameType,
	FieldPassword,
	FieldTags,
	FieldCreatedAt,
	FieldUpdatedAt,
}

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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
)
