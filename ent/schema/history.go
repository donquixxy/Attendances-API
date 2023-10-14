package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// History holds the schema definition for the History entity.
type History struct {
	ent.Schema
}

// Fields of the History.
func (History) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.String("id_user"),
		field.Int("type"),
		field.Time("time"),
		field.Time("created_at"),
		field.Time("updated_at"),
	}
}

// Edges of the History.
func (History) Edges() []ent.Edge {
	return nil
}
