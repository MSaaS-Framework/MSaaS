package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type APISpec struct {
	ent.Schema
}

func (APISpec) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.JSON("openapi_spec", []byte{}), // OPENAPI 스펙을 저장
	}
}

func (APISpec) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("service", Service.Type).Ref("apispec").Unique(),
		edge.From("project", Project.Type).Ref("apispecs").Unique(),
	}
}
