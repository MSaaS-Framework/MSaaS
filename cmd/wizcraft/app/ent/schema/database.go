package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Database holds the schema definition for the Database entity.
type Database struct {
	ent.Schema
}

func (Database) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("connection_path").NotEmpty(), // 접속 경로
		field.String("password").Sensitive(),       // 비밀번호
		field.String("db_type").NotEmpty(),         // 데이터베이스 종류
	}
}

func (Database) Edges() []ent.Edge {
	return []ent.Edge{
		// Service와의 관계 (N:1)
		edge.From("service", Service.Type).
			Ref("databases").
			Unique(),
		// GeneralSpec과의 관계 (1:1)
		edge.From("generalspec", GeneralSpec.Type).
			Ref("database").
			Unique().
			Required(),
	}
}
