// Code generated by ent, DO NOT EDIT.

package database

import (
	"MSaaS-Framework/MSaaS/cmd/wizcraft/app/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Database {
	return predicate.Database(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Database {
	return predicate.Database(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Database {
	return predicate.Database(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Database {
	return predicate.Database(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Database {
	return predicate.Database(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Database {
	return predicate.Database(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Database {
	return predicate.Database(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Database {
	return predicate.Database(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Database {
	return predicate.Database(sql.FieldLTE(FieldID, id))
}

// ConnectionPath applies equality check predicate on the "connection_path" field. It's identical to ConnectionPathEQ.
func ConnectionPath(v string) predicate.Database {
	return predicate.Database(sql.FieldEQ(FieldConnectionPath, v))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.Database {
	return predicate.Database(sql.FieldEQ(FieldPassword, v))
}

// DbType applies equality check predicate on the "db_type" field. It's identical to DbTypeEQ.
func DbType(v string) predicate.Database {
	return predicate.Database(sql.FieldEQ(FieldDbType, v))
}

// ConnectionPathEQ applies the EQ predicate on the "connection_path" field.
func ConnectionPathEQ(v string) predicate.Database {
	return predicate.Database(sql.FieldEQ(FieldConnectionPath, v))
}

// ConnectionPathNEQ applies the NEQ predicate on the "connection_path" field.
func ConnectionPathNEQ(v string) predicate.Database {
	return predicate.Database(sql.FieldNEQ(FieldConnectionPath, v))
}

// ConnectionPathIn applies the In predicate on the "connection_path" field.
func ConnectionPathIn(vs ...string) predicate.Database {
	return predicate.Database(sql.FieldIn(FieldConnectionPath, vs...))
}

// ConnectionPathNotIn applies the NotIn predicate on the "connection_path" field.
func ConnectionPathNotIn(vs ...string) predicate.Database {
	return predicate.Database(sql.FieldNotIn(FieldConnectionPath, vs...))
}

// ConnectionPathGT applies the GT predicate on the "connection_path" field.
func ConnectionPathGT(v string) predicate.Database {
	return predicate.Database(sql.FieldGT(FieldConnectionPath, v))
}

// ConnectionPathGTE applies the GTE predicate on the "connection_path" field.
func ConnectionPathGTE(v string) predicate.Database {
	return predicate.Database(sql.FieldGTE(FieldConnectionPath, v))
}

// ConnectionPathLT applies the LT predicate on the "connection_path" field.
func ConnectionPathLT(v string) predicate.Database {
	return predicate.Database(sql.FieldLT(FieldConnectionPath, v))
}

// ConnectionPathLTE applies the LTE predicate on the "connection_path" field.
func ConnectionPathLTE(v string) predicate.Database {
	return predicate.Database(sql.FieldLTE(FieldConnectionPath, v))
}

// ConnectionPathContains applies the Contains predicate on the "connection_path" field.
func ConnectionPathContains(v string) predicate.Database {
	return predicate.Database(sql.FieldContains(FieldConnectionPath, v))
}

// ConnectionPathHasPrefix applies the HasPrefix predicate on the "connection_path" field.
func ConnectionPathHasPrefix(v string) predicate.Database {
	return predicate.Database(sql.FieldHasPrefix(FieldConnectionPath, v))
}

// ConnectionPathHasSuffix applies the HasSuffix predicate on the "connection_path" field.
func ConnectionPathHasSuffix(v string) predicate.Database {
	return predicate.Database(sql.FieldHasSuffix(FieldConnectionPath, v))
}

// ConnectionPathEqualFold applies the EqualFold predicate on the "connection_path" field.
func ConnectionPathEqualFold(v string) predicate.Database {
	return predicate.Database(sql.FieldEqualFold(FieldConnectionPath, v))
}

// ConnectionPathContainsFold applies the ContainsFold predicate on the "connection_path" field.
func ConnectionPathContainsFold(v string) predicate.Database {
	return predicate.Database(sql.FieldContainsFold(FieldConnectionPath, v))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.Database {
	return predicate.Database(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.Database {
	return predicate.Database(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.Database {
	return predicate.Database(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.Database {
	return predicate.Database(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.Database {
	return predicate.Database(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.Database {
	return predicate.Database(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.Database {
	return predicate.Database(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.Database {
	return predicate.Database(sql.FieldLTE(FieldPassword, v))
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.Database {
	return predicate.Database(sql.FieldContains(FieldPassword, v))
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.Database {
	return predicate.Database(sql.FieldHasPrefix(FieldPassword, v))
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.Database {
	return predicate.Database(sql.FieldHasSuffix(FieldPassword, v))
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.Database {
	return predicate.Database(sql.FieldEqualFold(FieldPassword, v))
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.Database {
	return predicate.Database(sql.FieldContainsFold(FieldPassword, v))
}

// DbTypeEQ applies the EQ predicate on the "db_type" field.
func DbTypeEQ(v string) predicate.Database {
	return predicate.Database(sql.FieldEQ(FieldDbType, v))
}

// DbTypeNEQ applies the NEQ predicate on the "db_type" field.
func DbTypeNEQ(v string) predicate.Database {
	return predicate.Database(sql.FieldNEQ(FieldDbType, v))
}

// DbTypeIn applies the In predicate on the "db_type" field.
func DbTypeIn(vs ...string) predicate.Database {
	return predicate.Database(sql.FieldIn(FieldDbType, vs...))
}

// DbTypeNotIn applies the NotIn predicate on the "db_type" field.
func DbTypeNotIn(vs ...string) predicate.Database {
	return predicate.Database(sql.FieldNotIn(FieldDbType, vs...))
}

// DbTypeGT applies the GT predicate on the "db_type" field.
func DbTypeGT(v string) predicate.Database {
	return predicate.Database(sql.FieldGT(FieldDbType, v))
}

// DbTypeGTE applies the GTE predicate on the "db_type" field.
func DbTypeGTE(v string) predicate.Database {
	return predicate.Database(sql.FieldGTE(FieldDbType, v))
}

// DbTypeLT applies the LT predicate on the "db_type" field.
func DbTypeLT(v string) predicate.Database {
	return predicate.Database(sql.FieldLT(FieldDbType, v))
}

// DbTypeLTE applies the LTE predicate on the "db_type" field.
func DbTypeLTE(v string) predicate.Database {
	return predicate.Database(sql.FieldLTE(FieldDbType, v))
}

// DbTypeContains applies the Contains predicate on the "db_type" field.
func DbTypeContains(v string) predicate.Database {
	return predicate.Database(sql.FieldContains(FieldDbType, v))
}

// DbTypeHasPrefix applies the HasPrefix predicate on the "db_type" field.
func DbTypeHasPrefix(v string) predicate.Database {
	return predicate.Database(sql.FieldHasPrefix(FieldDbType, v))
}

// DbTypeHasSuffix applies the HasSuffix predicate on the "db_type" field.
func DbTypeHasSuffix(v string) predicate.Database {
	return predicate.Database(sql.FieldHasSuffix(FieldDbType, v))
}

// DbTypeEqualFold applies the EqualFold predicate on the "db_type" field.
func DbTypeEqualFold(v string) predicate.Database {
	return predicate.Database(sql.FieldEqualFold(FieldDbType, v))
}

// DbTypeContainsFold applies the ContainsFold predicate on the "db_type" field.
func DbTypeContainsFold(v string) predicate.Database {
	return predicate.Database(sql.FieldContainsFold(FieldDbType, v))
}

// HasService applies the HasEdge predicate on the "service" edge.
func HasService() predicate.Database {
	return predicate.Database(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ServiceTable, ServiceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasServiceWith applies the HasEdge predicate on the "service" edge with a given conditions (other predicates).
func HasServiceWith(preds ...predicate.Service) predicate.Database {
	return predicate.Database(func(s *sql.Selector) {
		step := newServiceStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProject applies the HasEdge predicate on the "project" edge.
func HasProject() predicate.Database {
	return predicate.Database(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProjectTable, ProjectColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProjectWith applies the HasEdge predicate on the "project" edge with a given conditions (other predicates).
func HasProjectWith(preds ...predicate.Project) predicate.Database {
	return predicate.Database(func(s *sql.Selector) {
		step := newProjectStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasGeneralspec applies the HasEdge predicate on the "generalspec" edge.
func HasGeneralspec() predicate.Database {
	return predicate.Database(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, GeneralspecTable, GeneralspecColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGeneralspecWith applies the HasEdge predicate on the "generalspec" edge with a given conditions (other predicates).
func HasGeneralspecWith(preds ...predicate.GeneralSpec) predicate.Database {
	return predicate.Database(func(s *sql.Selector) {
		step := newGeneralspecStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Database) predicate.Database {
	return predicate.Database(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Database) predicate.Database {
	return predicate.Database(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Database) predicate.Database {
	return predicate.Database(sql.NotPredicates(p))
}
