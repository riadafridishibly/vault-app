package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Key holds the schema definition for the Key entity.
type Key struct {
	ent.Schema
}

func (Key) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the Key.
func (Key) Fields() []ent.Field {
	return []ent.Field{
		field.String("type"),        // "age", "ed25519", "x25519"
		field.String("public_key"),  // string
		field.String("private_key"), // encrypted string
		field.Bool("is_active"),
		field.Int("references"), // How many times the key is used to encrypt
	}
}

// Edges of the Key.
func (Key) Edges() []ent.Edge {
	return nil
}
