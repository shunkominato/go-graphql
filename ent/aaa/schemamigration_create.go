// Code generated by ent, DO NOT EDIT.

package aaa

import (
	"context"
	"fmt"
	"server/ent/aaa/schemamigration"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SchemaMigrationCreate is the builder for creating a SchemaMigration entity.
type SchemaMigrationCreate struct {
	config
	mutation *SchemaMigrationMutation
	hooks    []Hook
}

// SetID sets the "id" field.
func (smc *SchemaMigrationCreate) SetID(s string) *SchemaMigrationCreate {
	smc.mutation.SetID(s)
	return smc
}

// Mutation returns the SchemaMigrationMutation object of the builder.
func (smc *SchemaMigrationCreate) Mutation() *SchemaMigrationMutation {
	return smc.mutation
}

// Save creates the SchemaMigration in the database.
func (smc *SchemaMigrationCreate) Save(ctx context.Context) (*SchemaMigration, error) {
	return withHooks[*SchemaMigration, SchemaMigrationMutation](ctx, smc.sqlSave, smc.mutation, smc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (smc *SchemaMigrationCreate) SaveX(ctx context.Context) *SchemaMigration {
	v, err := smc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (smc *SchemaMigrationCreate) Exec(ctx context.Context) error {
	_, err := smc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (smc *SchemaMigrationCreate) ExecX(ctx context.Context) {
	if err := smc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (smc *SchemaMigrationCreate) check() error {
	return nil
}

func (smc *SchemaMigrationCreate) sqlSave(ctx context.Context) (*SchemaMigration, error) {
	if err := smc.check(); err != nil {
		return nil, err
	}
	_node, _spec := smc.createSpec()
	if err := sqlgraph.CreateNode(ctx, smc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected SchemaMigration.ID type: %T", _spec.ID.Value)
		}
	}
	smc.mutation.id = &_node.ID
	smc.mutation.done = true
	return _node, nil
}

func (smc *SchemaMigrationCreate) createSpec() (*SchemaMigration, *sqlgraph.CreateSpec) {
	var (
		_node = &SchemaMigration{config: smc.config}
		_spec = sqlgraph.NewCreateSpec(schemamigration.Table, sqlgraph.NewFieldSpec(schemamigration.FieldID, field.TypeString))
	)
	if id, ok := smc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	return _node, _spec
}

// SchemaMigrationCreateBulk is the builder for creating many SchemaMigration entities in bulk.
type SchemaMigrationCreateBulk struct {
	config
	builders []*SchemaMigrationCreate
}

// Save creates the SchemaMigration entities in the database.
func (smcb *SchemaMigrationCreateBulk) Save(ctx context.Context) ([]*SchemaMigration, error) {
	specs := make([]*sqlgraph.CreateSpec, len(smcb.builders))
	nodes := make([]*SchemaMigration, len(smcb.builders))
	mutators := make([]Mutator, len(smcb.builders))
	for i := range smcb.builders {
		func(i int, root context.Context) {
			builder := smcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SchemaMigrationMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, smcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, smcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, smcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (smcb *SchemaMigrationCreateBulk) SaveX(ctx context.Context) []*SchemaMigration {
	v, err := smcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (smcb *SchemaMigrationCreateBulk) Exec(ctx context.Context) error {
	_, err := smcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (smcb *SchemaMigrationCreateBulk) ExecX(ctx context.Context) {
	if err := smcb.Exec(ctx); err != nil {
		panic(err)
	}
}
