// Code generated by ent, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"story.com/story/app/ent/model/arinternalmetadatum"
	"story.com/story/app/ent/model/predicate"
)

// ArInternalMetadatumUpdate is the builder for updating ArInternalMetadatum entities.
type ArInternalMetadatumUpdate struct {
	config
	hooks    []Hook
	mutation *ArInternalMetadatumMutation
}

// Where appends a list predicates to the ArInternalMetadatumUpdate builder.
func (aimu *ArInternalMetadatumUpdate) Where(ps ...predicate.ArInternalMetadatum) *ArInternalMetadatumUpdate {
	aimu.mutation.Where(ps...)
	return aimu
}

// SetValue sets the "value" field.
func (aimu *ArInternalMetadatumUpdate) SetValue(s string) *ArInternalMetadatumUpdate {
	aimu.mutation.SetValue(s)
	return aimu
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (aimu *ArInternalMetadatumUpdate) SetNillableValue(s *string) *ArInternalMetadatumUpdate {
	if s != nil {
		aimu.SetValue(*s)
	}
	return aimu
}

// ClearValue clears the value of the "value" field.
func (aimu *ArInternalMetadatumUpdate) ClearValue() *ArInternalMetadatumUpdate {
	aimu.mutation.ClearValue()
	return aimu
}

// SetCreatedAt sets the "created_at" field.
func (aimu *ArInternalMetadatumUpdate) SetCreatedAt(t time.Time) *ArInternalMetadatumUpdate {
	aimu.mutation.SetCreatedAt(t)
	return aimu
}

// SetUpdatedAt sets the "updated_at" field.
func (aimu *ArInternalMetadatumUpdate) SetUpdatedAt(t time.Time) *ArInternalMetadatumUpdate {
	aimu.mutation.SetUpdatedAt(t)
	return aimu
}

// Mutation returns the ArInternalMetadatumMutation object of the builder.
func (aimu *ArInternalMetadatumUpdate) Mutation() *ArInternalMetadatumMutation {
	return aimu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (aimu *ArInternalMetadatumUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, ArInternalMetadatumMutation](ctx, aimu.sqlSave, aimu.mutation, aimu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aimu *ArInternalMetadatumUpdate) SaveX(ctx context.Context) int {
	affected, err := aimu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (aimu *ArInternalMetadatumUpdate) Exec(ctx context.Context) error {
	_, err := aimu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aimu *ArInternalMetadatumUpdate) ExecX(ctx context.Context) {
	if err := aimu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (aimu *ArInternalMetadatumUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(arinternalmetadatum.Table, arinternalmetadatum.Columns, sqlgraph.NewFieldSpec(arinternalmetadatum.FieldID, field.TypeString))
	if ps := aimu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aimu.mutation.Value(); ok {
		_spec.SetField(arinternalmetadatum.FieldValue, field.TypeString, value)
	}
	if aimu.mutation.ValueCleared() {
		_spec.ClearField(arinternalmetadatum.FieldValue, field.TypeString)
	}
	if value, ok := aimu.mutation.CreatedAt(); ok {
		_spec.SetField(arinternalmetadatum.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := aimu.mutation.UpdatedAt(); ok {
		_spec.SetField(arinternalmetadatum.FieldUpdatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, aimu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{arinternalmetadatum.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	aimu.mutation.done = true
	return n, nil
}

// ArInternalMetadatumUpdateOne is the builder for updating a single ArInternalMetadatum entity.
type ArInternalMetadatumUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ArInternalMetadatumMutation
}

// SetValue sets the "value" field.
func (aimuo *ArInternalMetadatumUpdateOne) SetValue(s string) *ArInternalMetadatumUpdateOne {
	aimuo.mutation.SetValue(s)
	return aimuo
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (aimuo *ArInternalMetadatumUpdateOne) SetNillableValue(s *string) *ArInternalMetadatumUpdateOne {
	if s != nil {
		aimuo.SetValue(*s)
	}
	return aimuo
}

// ClearValue clears the value of the "value" field.
func (aimuo *ArInternalMetadatumUpdateOne) ClearValue() *ArInternalMetadatumUpdateOne {
	aimuo.mutation.ClearValue()
	return aimuo
}

// SetCreatedAt sets the "created_at" field.
func (aimuo *ArInternalMetadatumUpdateOne) SetCreatedAt(t time.Time) *ArInternalMetadatumUpdateOne {
	aimuo.mutation.SetCreatedAt(t)
	return aimuo
}

// SetUpdatedAt sets the "updated_at" field.
func (aimuo *ArInternalMetadatumUpdateOne) SetUpdatedAt(t time.Time) *ArInternalMetadatumUpdateOne {
	aimuo.mutation.SetUpdatedAt(t)
	return aimuo
}

// Mutation returns the ArInternalMetadatumMutation object of the builder.
func (aimuo *ArInternalMetadatumUpdateOne) Mutation() *ArInternalMetadatumMutation {
	return aimuo.mutation
}

// Where appends a list predicates to the ArInternalMetadatumUpdate builder.
func (aimuo *ArInternalMetadatumUpdateOne) Where(ps ...predicate.ArInternalMetadatum) *ArInternalMetadatumUpdateOne {
	aimuo.mutation.Where(ps...)
	return aimuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (aimuo *ArInternalMetadatumUpdateOne) Select(field string, fields ...string) *ArInternalMetadatumUpdateOne {
	aimuo.fields = append([]string{field}, fields...)
	return aimuo
}

// Save executes the query and returns the updated ArInternalMetadatum entity.
func (aimuo *ArInternalMetadatumUpdateOne) Save(ctx context.Context) (*ArInternalMetadatum, error) {
	return withHooks[*ArInternalMetadatum, ArInternalMetadatumMutation](ctx, aimuo.sqlSave, aimuo.mutation, aimuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aimuo *ArInternalMetadatumUpdateOne) SaveX(ctx context.Context) *ArInternalMetadatum {
	node, err := aimuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aimuo *ArInternalMetadatumUpdateOne) Exec(ctx context.Context) error {
	_, err := aimuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aimuo *ArInternalMetadatumUpdateOne) ExecX(ctx context.Context) {
	if err := aimuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (aimuo *ArInternalMetadatumUpdateOne) sqlSave(ctx context.Context) (_node *ArInternalMetadatum, err error) {
	_spec := sqlgraph.NewUpdateSpec(arinternalmetadatum.Table, arinternalmetadatum.Columns, sqlgraph.NewFieldSpec(arinternalmetadatum.FieldID, field.TypeString))
	id, ok := aimuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`model: missing "ArInternalMetadatum.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := aimuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, arinternalmetadatum.FieldID)
		for _, f := range fields {
			if !arinternalmetadatum.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != arinternalmetadatum.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := aimuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aimuo.mutation.Value(); ok {
		_spec.SetField(arinternalmetadatum.FieldValue, field.TypeString, value)
	}
	if aimuo.mutation.ValueCleared() {
		_spec.ClearField(arinternalmetadatum.FieldValue, field.TypeString)
	}
	if value, ok := aimuo.mutation.CreatedAt(); ok {
		_spec.SetField(arinternalmetadatum.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := aimuo.mutation.UpdatedAt(); ok {
		_spec.SetField(arinternalmetadatum.FieldUpdatedAt, field.TypeTime, value)
	}
	_node = &ArInternalMetadatum{config: aimuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aimuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{arinternalmetadatum.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	aimuo.mutation.done = true
	return _node, nil
}
