package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Heart holds the schema definition for the Heart entity.
type Heart struct {
	ent.Schema
}

// Fields of the Heart.
func (Heart) Fields() []ent.Field {
	return []ent.Field{
		field.Int("event_id").Unique(),
		field.String("writer"),
		field.Bool("is_heart").Default(false),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Heart.
func (Heart) Edges() []ent.Edge {
	return nil
}
