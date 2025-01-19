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
		// Database와의 관계 (1:N)
		edge.To("databases", Database.Type),
		// APISpec과의 관계 (1:1)
		edge.To("apispec", APISpec.Type).Unique(),
		// GeneralSpec과의 관계 (1:1)
		edge.From("generalspec", GeneralSpec.Type).
			Ref("service").
			Unique().
			Required(),
	}
}
