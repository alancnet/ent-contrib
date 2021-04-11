// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/contrib/entproto/internal/entprototest/ent/validmessage"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ValidMessageUpdate is the builder for updating ValidMessage entities.
type ValidMessageUpdate struct {
	config
	hooks    []Hook
	mutation *ValidMessageMutation
}

// Where adds a new predicate for the ValidMessageUpdate builder.
func (vmu *ValidMessageUpdate) Where(ps ...predicate.ValidMessage) *ValidMessageUpdate {
	vmu.mutation.predicates = append(vmu.mutation.predicates, ps...)
	return vmu
}

// SetName sets the "name" field.
func (vmu *ValidMessageUpdate) SetName(s string) *ValidMessageUpdate {
	vmu.mutation.SetName(s)
	return vmu
}

// SetTs sets the "ts" field.
func (vmu *ValidMessageUpdate) SetTs(t time.Time) *ValidMessageUpdate {
	vmu.mutation.SetTs(t)
	return vmu
}

// SetUUID sets the "uuid" field.
func (vmu *ValidMessageUpdate) SetUUID(u uuid.UUID) *ValidMessageUpdate {
	vmu.mutation.SetUUID(u)
	return vmu
}

// SetU8 sets the "u8" field.
func (vmu *ValidMessageUpdate) SetU8(u uint8) *ValidMessageUpdate {
	vmu.mutation.ResetU8()
	vmu.mutation.SetU8(u)
	return vmu
}

// AddU8 adds u to the "u8" field.
func (vmu *ValidMessageUpdate) AddU8(u uint8) *ValidMessageUpdate {
	vmu.mutation.AddU8(u)
	return vmu
}

// Mutation returns the ValidMessageMutation object of the builder.
func (vmu *ValidMessageUpdate) Mutation() *ValidMessageMutation {
	return vmu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (vmu *ValidMessageUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(vmu.hooks) == 0 {
		affected, err = vmu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ValidMessageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			vmu.mutation = mutation
			affected, err = vmu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(vmu.hooks) - 1; i >= 0; i-- {
			mut = vmu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vmu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (vmu *ValidMessageUpdate) SaveX(ctx context.Context) int {
	affected, err := vmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (vmu *ValidMessageUpdate) Exec(ctx context.Context) error {
	_, err := vmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vmu *ValidMessageUpdate) ExecX(ctx context.Context) {
	if err := vmu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (vmu *ValidMessageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   validmessage.Table,
			Columns: validmessage.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: validmessage.FieldID,
			},
		},
	}
	if ps := vmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vmu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: validmessage.FieldName,
		})
	}
	if value, ok := vmu.mutation.Ts(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: validmessage.FieldTs,
		})
	}
	if value, ok := vmu.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: validmessage.FieldUUID,
		})
	}
	if value, ok := vmu.mutation.U8(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: validmessage.FieldU8,
		})
	}
	if value, ok := vmu.mutation.AddedU8(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: validmessage.FieldU8,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, vmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{validmessage.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ValidMessageUpdateOne is the builder for updating a single ValidMessage entity.
type ValidMessageUpdateOne struct {
	config
	hooks    []Hook
	mutation *ValidMessageMutation
}

// SetName sets the "name" field.
func (vmuo *ValidMessageUpdateOne) SetName(s string) *ValidMessageUpdateOne {
	vmuo.mutation.SetName(s)
	return vmuo
}

// SetTs sets the "ts" field.
func (vmuo *ValidMessageUpdateOne) SetTs(t time.Time) *ValidMessageUpdateOne {
	vmuo.mutation.SetTs(t)
	return vmuo
}

// SetUUID sets the "uuid" field.
func (vmuo *ValidMessageUpdateOne) SetUUID(u uuid.UUID) *ValidMessageUpdateOne {
	vmuo.mutation.SetUUID(u)
	return vmuo
}

// SetU8 sets the "u8" field.
func (vmuo *ValidMessageUpdateOne) SetU8(u uint8) *ValidMessageUpdateOne {
	vmuo.mutation.ResetU8()
	vmuo.mutation.SetU8(u)
	return vmuo
}

// AddU8 adds u to the "u8" field.
func (vmuo *ValidMessageUpdateOne) AddU8(u uint8) *ValidMessageUpdateOne {
	vmuo.mutation.AddU8(u)
	return vmuo
}

// Mutation returns the ValidMessageMutation object of the builder.
func (vmuo *ValidMessageUpdateOne) Mutation() *ValidMessageMutation {
	return vmuo.mutation
}

// Save executes the query and returns the updated ValidMessage entity.
func (vmuo *ValidMessageUpdateOne) Save(ctx context.Context) (*ValidMessage, error) {
	var (
		err  error
		node *ValidMessage
	)
	if len(vmuo.hooks) == 0 {
		node, err = vmuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ValidMessageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			vmuo.mutation = mutation
			node, err = vmuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(vmuo.hooks) - 1; i >= 0; i-- {
			mut = vmuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vmuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (vmuo *ValidMessageUpdateOne) SaveX(ctx context.Context) *ValidMessage {
	node, err := vmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (vmuo *ValidMessageUpdateOne) Exec(ctx context.Context) error {
	_, err := vmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vmuo *ValidMessageUpdateOne) ExecX(ctx context.Context) {
	if err := vmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (vmuo *ValidMessageUpdateOne) sqlSave(ctx context.Context) (_node *ValidMessage, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   validmessage.Table,
			Columns: validmessage.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: validmessage.FieldID,
			},
		},
	}
	id, ok := vmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing ValidMessage.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := vmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vmuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: validmessage.FieldName,
		})
	}
	if value, ok := vmuo.mutation.Ts(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: validmessage.FieldTs,
		})
	}
	if value, ok := vmuo.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: validmessage.FieldUUID,
		})
	}
	if value, ok := vmuo.mutation.U8(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: validmessage.FieldU8,
		})
	}
	if value, ok := vmuo.mutation.AddedU8(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: validmessage.FieldU8,
		})
	}
	_node = &ValidMessage{config: vmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, vmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{validmessage.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
