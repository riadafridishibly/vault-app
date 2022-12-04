package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// File holds the schema definition for the File entity.
type File struct {
	ent.Schema
}

func (File) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the File.
func (File) Fields() []ent.Field {
	return []ent.Field{
		field.String("src_path"), // where was file actually stored
		field.String("current_path"),
		field.String("filetype"),
	}
}

// Edges of the File.
func (File) Edges() []ent.Edge {
	return nil
}
