package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// WiseSaying holds the schema definition for the WiseSaying entity.
type WiseSaying struct {
	ent.Schema
}

// Fields of the User.
func (WiseSaying) Fields() []ent.Field {
	return []ent.Field{
		field.String("wise_saying"),
		field.String("created_by"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the WiseSaying.
func (WiseSaying) Edges() []ent.Edge {
	return nil
}
