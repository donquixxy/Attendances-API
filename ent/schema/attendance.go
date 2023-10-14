package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Attendance holds the schema definition for the Attendance entity.
type Attendance struct {
	ent.Schema
}

// Fields of the Attendance.
func (Attendance) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.String("id_user"),
		field.Int("type").NonNegative(),
		field.Time("created_at"),
		field.Time("updated_at"),
	}
}

// Edges of the Attendance.
func (Attendance) Edges() []ent.Edge {
	return nil
}
