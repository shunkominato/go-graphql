// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"story.com/story/app/ent/todo"
	"story.com/story/app/ent/user"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetProvider sets the "provider" field.
func (uc *UserCreate) SetProvider(s string) *UserCreate {
	uc.mutation.SetProvider(s)
	return uc
}

// SetUID sets the "uid" field.
func (uc *UserCreate) SetUID(s string) *UserCreate {
	uc.mutation.SetUID(s)
	return uc
}

// SetEncryptedPassword sets the "encrypted_password" field.
func (uc *UserCreate) SetEncryptedPassword(s string) *UserCreate {
	uc.mutation.SetEncryptedPassword(s)
	return uc
}

// SetResetPasswordToken sets the "reset_password_token" field.
func (uc *UserCreate) SetResetPasswordToken(s string) *UserCreate {
	uc.mutation.SetResetPasswordToken(s)
	return uc
}

// SetNillableResetPasswordToken sets the "reset_password_token" field if the given value is not nil.
func (uc *UserCreate) SetNillableResetPasswordToken(s *string) *UserCreate {
	if s != nil {
		uc.SetResetPasswordToken(*s)
	}
	return uc
}

// SetResetPasswordSentAt sets the "reset_password_sent_at" field.
func (uc *UserCreate) SetResetPasswordSentAt(t time.Time) *UserCreate {
	uc.mutation.SetResetPasswordSentAt(t)
	return uc
}

// SetNillableResetPasswordSentAt sets the "reset_password_sent_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableResetPasswordSentAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetResetPasswordSentAt(*t)
	}
	return uc
}

// SetAllowPasswordChange sets the "allow_password_change" field.
func (uc *UserCreate) SetAllowPasswordChange(b bool) *UserCreate {
	uc.mutation.SetAllowPasswordChange(b)
	return uc
}

// SetNillableAllowPasswordChange sets the "allow_password_change" field if the given value is not nil.
func (uc *UserCreate) SetNillableAllowPasswordChange(b *bool) *UserCreate {
	if b != nil {
		uc.SetAllowPasswordChange(*b)
	}
	return uc
}

// SetRememberCreatedAt sets the "remember_created_at" field.
func (uc *UserCreate) SetRememberCreatedAt(t time.Time) *UserCreate {
	uc.mutation.SetRememberCreatedAt(t)
	return uc
}

// SetNillableRememberCreatedAt sets the "remember_created_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableRememberCreatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetRememberCreatedAt(*t)
	}
	return uc
}

// SetConfirmationToken sets the "confirmation_token" field.
func (uc *UserCreate) SetConfirmationToken(s string) *UserCreate {
	uc.mutation.SetConfirmationToken(s)
	return uc
}

// SetNillableConfirmationToken sets the "confirmation_token" field if the given value is not nil.
func (uc *UserCreate) SetNillableConfirmationToken(s *string) *UserCreate {
	if s != nil {
		uc.SetConfirmationToken(*s)
	}
	return uc
}

// SetConfirmedAt sets the "confirmed_at" field.
func (uc *UserCreate) SetConfirmedAt(t time.Time) *UserCreate {
	uc.mutation.SetConfirmedAt(t)
	return uc
}

// SetNillableConfirmedAt sets the "confirmed_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableConfirmedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetConfirmedAt(*t)
	}
	return uc
}

// SetConfirmationSentAt sets the "confirmation_sent_at" field.
func (uc *UserCreate) SetConfirmationSentAt(t time.Time) *UserCreate {
	uc.mutation.SetConfirmationSentAt(t)
	return uc
}

// SetNillableConfirmationSentAt sets the "confirmation_sent_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableConfirmationSentAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetConfirmationSentAt(*t)
	}
	return uc
}

// SetUnconfirmedEmail sets the "unconfirmed_email" field.
func (uc *UserCreate) SetUnconfirmedEmail(s string) *UserCreate {
	uc.mutation.SetUnconfirmedEmail(s)
	return uc
}

// SetNillableUnconfirmedEmail sets the "unconfirmed_email" field if the given value is not nil.
func (uc *UserCreate) SetNillableUnconfirmedEmail(s *string) *UserCreate {
	if s != nil {
		uc.SetUnconfirmedEmail(*s)
	}
	return uc
}

// SetName sets the "name" field.
func (uc *UserCreate) SetName(s string) *UserCreate {
	uc.mutation.SetName(s)
	return uc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uc *UserCreate) SetNillableName(s *string) *UserCreate {
	if s != nil {
		uc.SetName(*s)
	}
	return uc
}

// SetNickname sets the "nickname" field.
func (uc *UserCreate) SetNickname(s string) *UserCreate {
	uc.mutation.SetNickname(s)
	return uc
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (uc *UserCreate) SetNillableNickname(s *string) *UserCreate {
	if s != nil {
		uc.SetNickname(*s)
	}
	return uc
}

// SetImage sets the "image" field.
func (uc *UserCreate) SetImage(s string) *UserCreate {
	uc.mutation.SetImage(s)
	return uc
}

// SetNillableImage sets the "image" field if the given value is not nil.
func (uc *UserCreate) SetNillableImage(s *string) *UserCreate {
	if s != nil {
		uc.SetImage(*s)
	}
	return uc
}

// SetEmail sets the "email" field.
func (uc *UserCreate) SetEmail(s string) *UserCreate {
	uc.mutation.SetEmail(s)
	return uc
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uc *UserCreate) SetNillableEmail(s *string) *UserCreate {
	if s != nil {
		uc.SetEmail(*s)
	}
	return uc
}

// SetTokens sets the "tokens" field.
func (uc *UserCreate) SetTokens(s []string) *UserCreate {
	uc.mutation.SetTokens(s)
	return uc
}

// SetCreatedAt sets the "created_at" field.
func (uc *UserCreate) SetCreatedAt(t time.Time) *UserCreate {
	uc.mutation.SetCreatedAt(t)
	return uc
}

// SetUpdatedAt sets the "updated_at" field.
func (uc *UserCreate) SetUpdatedAt(t time.Time) *UserCreate {
	uc.mutation.SetUpdatedAt(t)
	return uc
}

// SetID sets the "id" field.
func (uc *UserCreate) SetID(i int) *UserCreate {
	uc.mutation.SetID(i)
	return uc
}

// AddTodoIDs adds the "todos" edge to the Todo entity by IDs.
func (uc *UserCreate) AddTodoIDs(ids ...int) *UserCreate {
	uc.mutation.AddTodoIDs(ids...)
	return uc
}

// AddTodos adds the "todos" edges to the Todo entity.
func (uc *UserCreate) AddTodos(t ...*Todo) *UserCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uc.AddTodoIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	return withHooks[*User, UserMutation](ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.Provider(); !ok {
		return &ValidationError{Name: "provider", err: errors.New(`ent: missing required field "User.provider"`)}
	}
	if _, ok := uc.mutation.UID(); !ok {
		return &ValidationError{Name: "uid", err: errors.New(`ent: missing required field "User.uid"`)}
	}
	if _, ok := uc.mutation.EncryptedPassword(); !ok {
		return &ValidationError{Name: "encrypted_password", err: errors.New(`ent: missing required field "User.encrypted_password"`)}
	}
	if _, ok := uc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "User.created_at"`)}
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "User.updated_at"`)}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	)
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uc.mutation.Provider(); ok {
		_spec.SetField(user.FieldProvider, field.TypeString, value)
		_node.Provider = value
	}
	if value, ok := uc.mutation.UID(); ok {
		_spec.SetField(user.FieldUID, field.TypeString, value)
		_node.UID = value
	}
	if value, ok := uc.mutation.EncryptedPassword(); ok {
		_spec.SetField(user.FieldEncryptedPassword, field.TypeString, value)
		_node.EncryptedPassword = value
	}
	if value, ok := uc.mutation.ResetPasswordToken(); ok {
		_spec.SetField(user.FieldResetPasswordToken, field.TypeString, value)
		_node.ResetPasswordToken = value
	}
	if value, ok := uc.mutation.ResetPasswordSentAt(); ok {
		_spec.SetField(user.FieldResetPasswordSentAt, field.TypeTime, value)
		_node.ResetPasswordSentAt = value
	}
	if value, ok := uc.mutation.AllowPasswordChange(); ok {
		_spec.SetField(user.FieldAllowPasswordChange, field.TypeBool, value)
		_node.AllowPasswordChange = value
	}
	if value, ok := uc.mutation.RememberCreatedAt(); ok {
		_spec.SetField(user.FieldRememberCreatedAt, field.TypeTime, value)
		_node.RememberCreatedAt = value
	}
	if value, ok := uc.mutation.ConfirmationToken(); ok {
		_spec.SetField(user.FieldConfirmationToken, field.TypeString, value)
		_node.ConfirmationToken = value
	}
	if value, ok := uc.mutation.ConfirmedAt(); ok {
		_spec.SetField(user.FieldConfirmedAt, field.TypeTime, value)
		_node.ConfirmedAt = value
	}
	if value, ok := uc.mutation.ConfirmationSentAt(); ok {
		_spec.SetField(user.FieldConfirmationSentAt, field.TypeTime, value)
		_node.ConfirmationSentAt = value
	}
	if value, ok := uc.mutation.UnconfirmedEmail(); ok {
		_spec.SetField(user.FieldUnconfirmedEmail, field.TypeString, value)
		_node.UnconfirmedEmail = value
	}
	if value, ok := uc.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := uc.mutation.Nickname(); ok {
		_spec.SetField(user.FieldNickname, field.TypeString, value)
		_node.Nickname = value
	}
	if value, ok := uc.mutation.Image(); ok {
		_spec.SetField(user.FieldImage, field.TypeString, value)
		_node.Image = value
	}
	if value, ok := uc.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := uc.mutation.Tokens(); ok {
		_spec.SetField(user.FieldTokens, field.TypeJSON, value)
		_node.Tokens = value
	}
	if value, ok := uc.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := uc.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := uc.mutation.TodosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TodosTable,
			Columns: []string{user.TodosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: todo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
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
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
