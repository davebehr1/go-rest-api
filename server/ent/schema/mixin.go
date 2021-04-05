package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Mixin holds the schema definition for the Mixin entity.
type TimeMixin struct {
	mixin.Schema
}

// Fields of the Mixin.
func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("publishedAt").Default(time.Now).UpdateDefault(time.Now),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Mixin.
func (TimeMixin) Edges() []ent.Edge {
	return nil
}
