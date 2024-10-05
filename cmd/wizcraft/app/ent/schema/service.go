// ent/schema/service.go

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Service holds the schema definition for the Service entity.
type Service struct {
	ent.Schema
}

// Fields of the Service.
func (Service) Fields() []ent.Field {
	return []ent.Field{
		// UUID 필드를 추가
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
	}
}

func (Service) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("databases", Database.Type),       // 여러 Database와 연결
		edge.To("apispec", APISpec.Type).Unique(), // 하나의 APISpec과 연결
	}
}
