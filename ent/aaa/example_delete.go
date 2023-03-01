// Code generated by ent, DO NOT EDIT.

package aaa

import (
	"context"
	"server/ent/aaa/example"
	"server/ent/aaa/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ExampleDelete is the builder for deleting a Example entity.
type ExampleDelete struct {
	config
	hooks    []Hook
	mutation *ExampleMutation
}

// Where appends a list predicates to the ExampleDelete builder.
func (ed *ExampleDelete) Where(ps ...predicate.Example) *ExampleDelete {
	ed.mutation.Where(ps...)
	return ed
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ed *ExampleDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, ExampleMutation](ctx, ed.sqlExec, ed.mutation, ed.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ed *ExampleDelete) ExecX(ctx context.Context) int {
	n, err := ed.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ed *ExampleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(example.Table, sqlgraph.NewFieldSpec(example.FieldID, field.TypeInt))
	if ps := ed.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ed.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ed.mutation.done = true
	return affected, err
}

// ExampleDeleteOne is the builder for deleting a single Example entity.
type ExampleDeleteOne struct {
	ed *ExampleDelete
}

// Where appends a list predicates to the ExampleDelete builder.
func (edo *ExampleDeleteOne) Where(ps ...predicate.Example) *ExampleDeleteOne {
	edo.ed.mutation.Where(ps...)
	return edo
}

// Exec executes the deletion query.
func (edo *ExampleDeleteOne) Exec(ctx context.Context) error {
	n, err := edo.ed.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{example.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (edo *ExampleDeleteOne) ExecX(ctx context.Context) {
	if err := edo.Exec(ctx); err != nil {
		panic(err)
	}
}
