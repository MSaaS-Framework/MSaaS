package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// UserGeneralSpecPermissions holds the schema definition for the UserGeneralSpecPermissions entity.
type UserGeneralSpecPermissions struct {
	ent.Schema
}

// Fields of the UserGeneralSpecPermissions.
func (UserGeneralSpecPermissions) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.Enum("permission").Values("read", "write", "manage").Default("read"),
	}
}

// Edges of the UserGeneralSpecPermissions.
func (UserGeneralSpecPermissions) Edges() []ent.Edge {
	return []ent.Edge{
		// User와의 관계
		edge.From("user", User.Type).
			Ref("permissions").
			Unique().
			Required(),
		// GeneralSpec와의 관계
		edge.From("general_spec", GeneralSpec.Type).
			Ref("permissions").
			Unique().
			Required(),
	}
}
