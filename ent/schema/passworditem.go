package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// PasswordItem holds the schema definition for the PasswordItem entity.
type PasswordItem struct {
	ent.Schema
}

// Fields of the PasswordItem.
func (PasswordItem) Fields() []ent.Field {
	return []ent.Field{
		field.String("avatar").Nillable(),
		field.String("description").Nillable(),
		field.String("site_name").Nillable(),
		field.String("site_url").Nillable(),
		field.String("username").Nillable(),
		field.String("username_type").Nillable(), // email, username
		field.String("password").Nillable(),
		field.JSON("tags", []string{}).Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now),
	}
}

// Edges of the PasswordItem.
func (PasswordItem) Edges() []ent.Edge {
	return nil
}
