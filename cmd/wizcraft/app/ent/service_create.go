// Code generated by ent, DO NOT EDIT.

package ent

import (
	"MSaaS-Framework/MSaaS/cmd/wizcraft/app/ent/apispec"
	"MSaaS-Framework/MSaaS/cmd/wizcraft/app/ent/database"
	"MSaaS-Framework/MSaaS/cmd/wizcraft/app/ent/service"
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ServiceCreate is the builder for creating a Service entity.
type ServiceCreate struct {
	config
	mutation *ServiceMutation
	hooks    []Hook
}

// SetID sets the "id" field.
func (sc *ServiceCreate) SetID(u uuid.UUID) *ServiceCreate {
	sc.mutation.SetID(u)
	return sc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sc *ServiceCreate) SetNillableID(u *uuid.UUID) *ServiceCreate {
	if u != nil {
		sc.SetID(*u)
	}
	return sc
}

// AddDatabaseIDs adds the "databases" edge to the Database entity by IDs.
func (sc *ServiceCreate) AddDatabaseIDs(ids ...uuid.UUID) *ServiceCreate {
	sc.mutation.AddDatabaseIDs(ids...)
	return sc
}

// AddDatabases adds the "databases" edges to the Database entity.
func (sc *ServiceCreate) AddDatabases(d ...*Database) *ServiceCreate {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return sc.AddDatabaseIDs(ids...)
}

// SetApispecID sets the "apispec" edge to the APISpec entity by ID.
func (sc *ServiceCreate) SetApispecID(id uuid.UUID) *ServiceCreate {
	sc.mutation.SetApispecID(id)
	return sc
}

// SetNillableApispecID sets the "apispec" edge to the APISpec entity by ID if the given value is not nil.
func (sc *ServiceCreate) SetNillableApispecID(id *uuid.UUID) *ServiceCreate {
	if id != nil {
		sc = sc.SetApispecID(*id)
	}
	return sc
}

// SetApispec sets the "apispec" edge to the APISpec entity.
func (sc *ServiceCreate) SetApispec(a *APISpec) *ServiceCreate {
	return sc.SetApispecID(a.ID)
}

// Mutation returns the ServiceMutation object of the builder.
func (sc *ServiceCreate) Mutation() *ServiceMutation {
	return sc.mutation
}

// Save creates the Service in the database.
func (sc *ServiceCreate) Save(ctx context.Context) (*Service, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *ServiceCreate) SaveX(ctx context.Context) *Service {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *ServiceCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *ServiceCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *ServiceCreate) defaults() {
	if _, ok := sc.mutation.ID(); !ok {
		v := service.DefaultID()
		sc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *ServiceCreate) check() error {
	return nil
}

func (sc *ServiceCreate) sqlSave(ctx context.Context) (*Service, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *ServiceCreate) createSpec() (*Service, *sqlgraph.CreateSpec) {
	var (
		_node = &Service{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(service.Table, sqlgraph.NewFieldSpec(service.FieldID, field.TypeUUID))
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if nodes := sc.mutation.DatabasesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   service.DatabasesTable,
			Columns: []string{service.DatabasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(database.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.ApispecIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   service.ApispecTable,
			Columns: []string{service.ApispecColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(apispec.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ServiceCreateBulk is the builder for creating many Service entities in bulk.
type ServiceCreateBulk struct {
	config
	err      error
	builders []*ServiceCreate
}

// Save creates the Service entities in the database.
func (scb *ServiceCreateBulk) Save(ctx context.Context) ([]*Service, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Service, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ServiceMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *ServiceCreateBulk) SaveX(ctx context.Context) []*Service {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *ServiceCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *ServiceCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
