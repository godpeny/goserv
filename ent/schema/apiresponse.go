package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// APIResponse holds the schema definition for the APIResponse entity.
type APIResponse struct {
	ent.Schema
}

// Fields of the APIResponse.
func (APIResponse) Fields() []ent.Field {
	return []ent.Field{
		field.Int("code").
			Positive(),
		field.String("type").
			Default("unknown"),
		field.String("message").
			Default("unknown"),
	}
}

// Edges of the APIResponse.
func (APIResponse) Edges() []ent.Edge {
	return nil
}
