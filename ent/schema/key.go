package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Key holds the schema definition for the Key entity.
type Key struct {
	ent.Schema
}

// Fields of the Key.
func (Key) Fields() []ent.Field {
	return []ent.Field{
		field.String("type"),        // "age", "ed25519", "x25519"
		field.String("public_key"),  // string
		field.String("private_key"), // encrypted string
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now),
	}
}

// Edges of the Key.
func (Key) Edges() []ent.Edge {
	return nil
}
