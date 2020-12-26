package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Useroom holds the schema definition for the Useroom entity.
type Useroom struct {
	ent.Schema
}

// Fields of the Useroom.
func (Useroom) Fields() []ent.Field {
	return []ent.Field{
		field.Time("added_time"),
	}
}

// Edges of the Useroom.
func (Useroom) Edges() []ent.Edge {
   return nil
}


