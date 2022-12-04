package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// PasswordItem holds the schema definition for the PasswordItem entity.
type PasswordItem struct {
	ent.Schema
}

func (PasswordItem) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the PasswordItem.
func (PasswordItem) Fields() []ent.Field {
	return []ent.Field{
		field.String("avatar").Nillable().Optional(),
		field.String("description").Nillable().Optional(),
		field.String("site_name").Nillable().Optional(),
		field.String("site_url").Nillable().Optional(),
		field.String("username").Nillable().Optional(),
		field.String("username_type").Nillable().Optional(), // email, username
		field.String("password").Nillable().Optional(),
		field.JSON("tags", []string{}).Optional(),
	}
}

// Edges of the PasswordItem.
func (PasswordItem) Edges() []ent.Edge {
	return nil
}
