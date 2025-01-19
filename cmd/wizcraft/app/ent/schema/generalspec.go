package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// GeneralSpec holds the schema definition for the GeneralSpec entity.
type GeneralSpec struct {
	ent.Schema
}

// Fields of the GeneralSpec.
func (GeneralSpec) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid", uuid.UUID{}).
			Default(uuid.New),
		field.UUID("project_uuid", uuid.UUID{}),
		field.String("name"),
		field.String("type"),
		field.String("status").Default("created"),
		field.String("description"),
	}
}

// Edges of the GeneralSpec.
func (GeneralSpec) Edges() []ent.Edge {
	return []ent.Edge{
		// Project와의 관계
		edge.From("project", Project.Type).
			Ref("general_specs").
			Unique().
			Required(),
		edge.To("service", Service.Type).Unique(),
		edge.To("database", Database.Type).Unique(),
		edge.To("apispec", APISpec.Type).Unique(),
		edge.To("permissions", UserGeneralSpecPermissions.Type),
	}
}
