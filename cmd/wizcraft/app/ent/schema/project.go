package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Project holds the schema definition for the Project entity.
type Project struct {
	ent.Schema
}

// Fields of the Project.
func (Project) Fields() []ent.Field {
	return []ent.Field{
		// UUID 필드를 추가
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
	}
}

// Edges of the Project.
func (Project) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("services", Service.Type),
		edge.To("databases", Database.Type),
		edge.To("apispecs", APISpec.Type),
		edge.From("generalspec", GeneralSpec.Type).
			Ref("project").Unique(),
	}
}
