// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"coolfun-admin/data/ent/menu"
	"coolfun-admin/data/ent/menuparam"
	"coolfun-admin/data/ent/predicate"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MenuParamUpdate is the builder for updating MenuParam entities.
type MenuParamUpdate struct {
	config
	hooks    []Hook
	mutation *MenuParamMutation
}

// Where appends a list predicates to the MenuParamUpdate builder.
func (mpu *MenuParamUpdate) Where(ps ...predicate.MenuParam) *MenuParamUpdate {
	mpu.mutation.Where(ps...)
	return mpu
}

// SetUpdatedAt sets the "updated_at" field.
func (mpu *MenuParamUpdate) SetUpdatedAt(t time.Time) *MenuParamUpdate {
	mpu.mutation.SetUpdatedAt(t)
	return mpu
}

// SetType sets the "type" field.
func (mpu *MenuParamUpdate) SetType(s string) *MenuParamUpdate {
	mpu.mutation.SetType(s)
	return mpu
}

// SetKey sets the "key" field.
func (mpu *MenuParamUpdate) SetKey(s string) *MenuParamUpdate {
	mpu.mutation.SetKey(s)
	return mpu
}

// SetValue sets the "value" field.
func (mpu *MenuParamUpdate) SetValue(s string) *MenuParamUpdate {
	mpu.mutation.SetValue(s)
	return mpu
}

// SetMenusID sets the "menus" edge to the Menu entity by ID.
func (mpu *MenuParamUpdate) SetMenusID(id uint64) *MenuParamUpdate {
	mpu.mutation.SetMenusID(id)
	return mpu
}

// SetNillableMenusID sets the "menus" edge to the Menu entity by ID if the given value is not nil.
func (mpu *MenuParamUpdate) SetNillableMenusID(id *uint64) *MenuParamUpdate {
	if id != nil {
		mpu = mpu.SetMenusID(*id)
	}
	return mpu
}

// SetMenus sets the "menus" edge to the Menu entity.
func (mpu *MenuParamUpdate) SetMenus(m *Menu) *MenuParamUpdate {
	return mpu.SetMenusID(m.ID)
}

// Mutation returns the MenuParamMutation object of the builder.
func (mpu *MenuParamUpdate) Mutation() *MenuParamMutation {
	return mpu.mutation
}

// ClearMenus clears the "menus" edge to the Menu entity.
func (mpu *MenuParamUpdate) ClearMenus() *MenuParamUpdate {
	mpu.mutation.ClearMenus()
	return mpu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mpu *MenuParamUpdate) Save(ctx context.Context) (int, error) {
	mpu.defaults()
	return withHooks[int, MenuParamMutation](ctx, mpu.sqlSave, mpu.mutation, mpu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mpu *MenuParamUpdate) SaveX(ctx context.Context) int {
	affected, err := mpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mpu *MenuParamUpdate) Exec(ctx context.Context) error {
	_, err := mpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mpu *MenuParamUpdate) ExecX(ctx context.Context) {
	if err := mpu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mpu *MenuParamUpdate) defaults() {
	if _, ok := mpu.mutation.UpdatedAt(); !ok {
		v := menuparam.UpdateDefaultUpdatedAt()
		mpu.mutation.SetUpdatedAt(v)
	}
}

func (mpu *MenuParamUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   menuparam.Table,
			Columns: menuparam.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: menuparam.FieldID,
			},
		},
	}
	if ps := mpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mpu.mutation.UpdatedAt(); ok {
		_spec.SetField(menuparam.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := mpu.mutation.GetType(); ok {
		_spec.SetField(menuparam.FieldType, field.TypeString, value)
	}
	if value, ok := mpu.mutation.Key(); ok {
		_spec.SetField(menuparam.FieldKey, field.TypeString, value)
	}
	if value, ok := mpu.mutation.Value(); ok {
		_spec.SetField(menuparam.FieldValue, field.TypeString, value)
	}
	if mpu.mutation.MenusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   menuparam.MenusTable,
			Columns: []string{menuparam.MenusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: menu.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mpu.mutation.MenusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   menuparam.MenusTable,
			Columns: []string{menuparam.MenusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: menu.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{menuparam.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mpu.mutation.done = true
	return n, nil
}

// MenuParamUpdateOne is the builder for updating a single MenuParam entity.
type MenuParamUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MenuParamMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (mpuo *MenuParamUpdateOne) SetUpdatedAt(t time.Time) *MenuParamUpdateOne {
	mpuo.mutation.SetUpdatedAt(t)
	return mpuo
}

// SetType sets the "type" field.
func (mpuo *MenuParamUpdateOne) SetType(s string) *MenuParamUpdateOne {
	mpuo.mutation.SetType(s)
	return mpuo
}

// SetKey sets the "key" field.
func (mpuo *MenuParamUpdateOne) SetKey(s string) *MenuParamUpdateOne {
	mpuo.mutation.SetKey(s)
	return mpuo
}

// SetValue sets the "value" field.
func (mpuo *MenuParamUpdateOne) SetValue(s string) *MenuParamUpdateOne {
	mpuo.mutation.SetValue(s)
	return mpuo
}

// SetMenusID sets the "menus" edge to the Menu entity by ID.
func (mpuo *MenuParamUpdateOne) SetMenusID(id uint64) *MenuParamUpdateOne {
	mpuo.mutation.SetMenusID(id)
	return mpuo
}

// SetNillableMenusID sets the "menus" edge to the Menu entity by ID if the given value is not nil.
func (mpuo *MenuParamUpdateOne) SetNillableMenusID(id *uint64) *MenuParamUpdateOne {
	if id != nil {
		mpuo = mpuo.SetMenusID(*id)
	}
	return mpuo
}

// SetMenus sets the "menus" edge to the Menu entity.
func (mpuo *MenuParamUpdateOne) SetMenus(m *Menu) *MenuParamUpdateOne {
	return mpuo.SetMenusID(m.ID)
}

// Mutation returns the MenuParamMutation object of the builder.
func (mpuo *MenuParamUpdateOne) Mutation() *MenuParamMutation {
	return mpuo.mutation
}

// ClearMenus clears the "menus" edge to the Menu entity.
func (mpuo *MenuParamUpdateOne) ClearMenus() *MenuParamUpdateOne {
	mpuo.mutation.ClearMenus()
	return mpuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mpuo *MenuParamUpdateOne) Select(field string, fields ...string) *MenuParamUpdateOne {
	mpuo.fields = append([]string{field}, fields...)
	return mpuo
}

// Save executes the query and returns the updated MenuParam entity.
func (mpuo *MenuParamUpdateOne) Save(ctx context.Context) (*MenuParam, error) {
	mpuo.defaults()
	return withHooks[*MenuParam, MenuParamMutation](ctx, mpuo.sqlSave, mpuo.mutation, mpuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mpuo *MenuParamUpdateOne) SaveX(ctx context.Context) *MenuParam {
	node, err := mpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mpuo *MenuParamUpdateOne) Exec(ctx context.Context) error {
	_, err := mpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mpuo *MenuParamUpdateOne) ExecX(ctx context.Context) {
	if err := mpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mpuo *MenuParamUpdateOne) defaults() {
	if _, ok := mpuo.mutation.UpdatedAt(); !ok {
		v := menuparam.UpdateDefaultUpdatedAt()
		mpuo.mutation.SetUpdatedAt(v)
	}
}

func (mpuo *MenuParamUpdateOne) sqlSave(ctx context.Context) (_node *MenuParam, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   menuparam.Table,
			Columns: menuparam.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: menuparam.FieldID,
			},
		},
	}
	id, ok := mpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MenuParam.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := mpuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, menuparam.FieldID)
		for _, f := range fields {
			if !menuparam.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != menuparam.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mpuo.mutation.UpdatedAt(); ok {
		_spec.SetField(menuparam.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := mpuo.mutation.GetType(); ok {
		_spec.SetField(menuparam.FieldType, field.TypeString, value)
	}
	if value, ok := mpuo.mutation.Key(); ok {
		_spec.SetField(menuparam.FieldKey, field.TypeString, value)
	}
	if value, ok := mpuo.mutation.Value(); ok {
		_spec.SetField(menuparam.FieldValue, field.TypeString, value)
	}
	if mpuo.mutation.MenusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   menuparam.MenusTable,
			Columns: []string{menuparam.MenusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: menu.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mpuo.mutation.MenusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   menuparam.MenusTable,
			Columns: []string{menuparam.MenusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: menu.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &MenuParam{config: mpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{menuparam.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	mpuo.mutation.done = true
	return _node, nil
}
