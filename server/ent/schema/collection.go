package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Collection holds the schema definition for the Collection entity.
type Collection struct {
	ent.Schema
}

// Fields of the Collection.
func (Collection) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the Collection.
func (Collection) Edges() []ent.Edge {
	return nil
}
