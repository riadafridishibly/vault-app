package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// MasterPassword holds the schema definition for the MasterPassword entity.
type MasterPassword struct {
	ent.Schema
}

// Fields of the MasterPassword.
func (MasterPassword) Fields() []ent.Field {
	return []ent.Field{
		field.String("password_hash").NotEmpty(),
	}
}

// Edges of the MasterPassword.
func (MasterPassword) Edges() []ent.Edge {
	return nil
}
