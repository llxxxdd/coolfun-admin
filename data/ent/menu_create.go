// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"coolfun-admin/data/ent/menu"
	"coolfun-admin/data/ent/menuparam"
	"coolfun-admin/data/ent/role"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MenuCreate is the builder for creating a Menu entity.
type MenuCreate struct {
	config
	mutation *MenuMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (mc *MenuCreate) SetCreatedAt(t time.Time) *MenuCreate {
	mc.mutation.SetCreatedAt(t)
	return mc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mc *MenuCreate) SetNillableCreatedAt(t *time.Time) *MenuCreate {
	if t != nil {
		mc.SetCreatedAt(*t)
	}
	return mc
}

// SetUpdatedAt sets the "updated_at" field.
func (mc *MenuCreate) SetUpdatedAt(t time.Time) *MenuCreate {
	mc.mutation.SetUpdatedAt(t)
	return mc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mc *MenuCreate) SetNillableUpdatedAt(t *time.Time) *MenuCreate {
	if t != nil {
		mc.SetUpdatedAt(*t)
	}
	return mc
}

// SetParentID sets the "parent_id" field.
func (mc *MenuCreate) SetParentID(u uint64) *MenuCreate {
	mc.mutation.SetParentID(u)
	return mc
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (mc *MenuCreate) SetNillableParentID(u *uint64) *MenuCreate {
	if u != nil {
		mc.SetParentID(*u)
	}
	return mc
}

// SetMenuLevel sets the "menu_level" field.
func (mc *MenuCreate) SetMenuLevel(u uint32) *MenuCreate {
	mc.mutation.SetMenuLevel(u)
	return mc
}

// SetMenuType sets the "menu_type" field.
func (mc *MenuCreate) SetMenuType(u uint32) *MenuCreate {
	mc.mutation.SetMenuType(u)
	return mc
}

// SetPath sets the "path" field.
func (mc *MenuCreate) SetPath(s string) *MenuCreate {
	mc.mutation.SetPath(s)
	return mc
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (mc *MenuCreate) SetNillablePath(s *string) *MenuCreate {
	if s != nil {
		mc.SetPath(*s)
	}
	return mc
}

// SetName sets the "name" field.
func (mc *MenuCreate) SetName(s string) *MenuCreate {
	mc.mutation.SetName(s)
	return mc
}

// SetRedirect sets the "redirect" field.
func (mc *MenuCreate) SetRedirect(s string) *MenuCreate {
	mc.mutation.SetRedirect(s)
	return mc
}

// SetNillableRedirect sets the "redirect" field if the given value is not nil.
func (mc *MenuCreate) SetNillableRedirect(s *string) *MenuCreate {
	if s != nil {
		mc.SetRedirect(*s)
	}
	return mc
}

// SetComponent sets the "component" field.
func (mc *MenuCreate) SetComponent(s string) *MenuCreate {
	mc.mutation.SetComponent(s)
	return mc
}

// SetNillableComponent sets the "component" field if the given value is not nil.
func (mc *MenuCreate) SetNillableComponent(s *string) *MenuCreate {
	if s != nil {
		mc.SetComponent(*s)
	}
	return mc
}

// SetOrderNo sets the "order_no" field.
func (mc *MenuCreate) SetOrderNo(u uint32) *MenuCreate {
	mc.mutation.SetOrderNo(u)
	return mc
}

// SetNillableOrderNo sets the "order_no" field if the given value is not nil.
func (mc *MenuCreate) SetNillableOrderNo(u *uint32) *MenuCreate {
	if u != nil {
		mc.SetOrderNo(*u)
	}
	return mc
}

// SetDisabled sets the "disabled" field.
func (mc *MenuCreate) SetDisabled(b bool) *MenuCreate {
	mc.mutation.SetDisabled(b)
	return mc
}

// SetNillableDisabled sets the "disabled" field if the given value is not nil.
func (mc *MenuCreate) SetNillableDisabled(b *bool) *MenuCreate {
	if b != nil {
		mc.SetDisabled(*b)
	}
	return mc
}

// SetTitle sets the "title" field.
func (mc *MenuCreate) SetTitle(s string) *MenuCreate {
	mc.mutation.SetTitle(s)
	return mc
}

// SetIcon sets the "icon" field.
func (mc *MenuCreate) SetIcon(s string) *MenuCreate {
	mc.mutation.SetIcon(s)
	return mc
}

// SetHideMenu sets the "hide_menu" field.
func (mc *MenuCreate) SetHideMenu(b bool) *MenuCreate {
	mc.mutation.SetHideMenu(b)
	return mc
}

// SetNillableHideMenu sets the "hide_menu" field if the given value is not nil.
func (mc *MenuCreate) SetNillableHideMenu(b *bool) *MenuCreate {
	if b != nil {
		mc.SetHideMenu(*b)
	}
	return mc
}

// SetHideBreadcrumb sets the "hide_breadcrumb" field.
func (mc *MenuCreate) SetHideBreadcrumb(b bool) *MenuCreate {
	mc.mutation.SetHideBreadcrumb(b)
	return mc
}

// SetNillableHideBreadcrumb sets the "hide_breadcrumb" field if the given value is not nil.
func (mc *MenuCreate) SetNillableHideBreadcrumb(b *bool) *MenuCreate {
	if b != nil {
		mc.SetHideBreadcrumb(*b)
	}
	return mc
}

// SetCurrentActiveMenu sets the "current_active_menu" field.
func (mc *MenuCreate) SetCurrentActiveMenu(s string) *MenuCreate {
	mc.mutation.SetCurrentActiveMenu(s)
	return mc
}

// SetNillableCurrentActiveMenu sets the "current_active_menu" field if the given value is not nil.
func (mc *MenuCreate) SetNillableCurrentActiveMenu(s *string) *MenuCreate {
	if s != nil {
		mc.SetCurrentActiveMenu(*s)
	}
	return mc
}

// SetIgnoreKeepAlive sets the "ignore_keep_alive" field.
func (mc *MenuCreate) SetIgnoreKeepAlive(b bool) *MenuCreate {
	mc.mutation.SetIgnoreKeepAlive(b)
	return mc
}

// SetNillableIgnoreKeepAlive sets the "ignore_keep_alive" field if the given value is not nil.
func (mc *MenuCreate) SetNillableIgnoreKeepAlive(b *bool) *MenuCreate {
	if b != nil {
		mc.SetIgnoreKeepAlive(*b)
	}
	return mc
}

// SetHideTab sets the "hide_tab" field.
func (mc *MenuCreate) SetHideTab(b bool) *MenuCreate {
	mc.mutation.SetHideTab(b)
	return mc
}

// SetNillableHideTab sets the "hide_tab" field if the given value is not nil.
func (mc *MenuCreate) SetNillableHideTab(b *bool) *MenuCreate {
	if b != nil {
		mc.SetHideTab(*b)
	}
	return mc
}

// SetFrameSrc sets the "frame_src" field.
func (mc *MenuCreate) SetFrameSrc(s string) *MenuCreate {
	mc.mutation.SetFrameSrc(s)
	return mc
}

// SetNillableFrameSrc sets the "frame_src" field if the given value is not nil.
func (mc *MenuCreate) SetNillableFrameSrc(s *string) *MenuCreate {
	if s != nil {
		mc.SetFrameSrc(*s)
	}
	return mc
}

// SetCarryParam sets the "carry_param" field.
func (mc *MenuCreate) SetCarryParam(b bool) *MenuCreate {
	mc.mutation.SetCarryParam(b)
	return mc
}

// SetNillableCarryParam sets the "carry_param" field if the given value is not nil.
func (mc *MenuCreate) SetNillableCarryParam(b *bool) *MenuCreate {
	if b != nil {
		mc.SetCarryParam(*b)
	}
	return mc
}

// SetHideChildrenInMenu sets the "hide_children_in_menu" field.
func (mc *MenuCreate) SetHideChildrenInMenu(b bool) *MenuCreate {
	mc.mutation.SetHideChildrenInMenu(b)
	return mc
}

// SetNillableHideChildrenInMenu sets the "hide_children_in_menu" field if the given value is not nil.
func (mc *MenuCreate) SetNillableHideChildrenInMenu(b *bool) *MenuCreate {
	if b != nil {
		mc.SetHideChildrenInMenu(*b)
	}
	return mc
}

// SetAffix sets the "affix" field.
func (mc *MenuCreate) SetAffix(b bool) *MenuCreate {
	mc.mutation.SetAffix(b)
	return mc
}

// SetNillableAffix sets the "affix" field if the given value is not nil.
func (mc *MenuCreate) SetNillableAffix(b *bool) *MenuCreate {
	if b != nil {
		mc.SetAffix(*b)
	}
	return mc
}

// SetDynamicLevel sets the "dynamic_level" field.
func (mc *MenuCreate) SetDynamicLevel(u uint32) *MenuCreate {
	mc.mutation.SetDynamicLevel(u)
	return mc
}

// SetNillableDynamicLevel sets the "dynamic_level" field if the given value is not nil.
func (mc *MenuCreate) SetNillableDynamicLevel(u *uint32) *MenuCreate {
	if u != nil {
		mc.SetDynamicLevel(*u)
	}
	return mc
}

// SetRealPath sets the "real_path" field.
func (mc *MenuCreate) SetRealPath(s string) *MenuCreate {
	mc.mutation.SetRealPath(s)
	return mc
}

// SetNillableRealPath sets the "real_path" field if the given value is not nil.
func (mc *MenuCreate) SetNillableRealPath(s *string) *MenuCreate {
	if s != nil {
		mc.SetRealPath(*s)
	}
	return mc
}

// SetID sets the "id" field.
func (mc *MenuCreate) SetID(u uint64) *MenuCreate {
	mc.mutation.SetID(u)
	return mc
}

// AddRoleIDs adds the "roles" edge to the Role entity by IDs.
func (mc *MenuCreate) AddRoleIDs(ids ...uint64) *MenuCreate {
	mc.mutation.AddRoleIDs(ids...)
	return mc
}

// AddRoles adds the "roles" edges to the Role entity.
func (mc *MenuCreate) AddRoles(r ...*Role) *MenuCreate {
	ids := make([]uint64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return mc.AddRoleIDs(ids...)
}

// SetParent sets the "parent" edge to the Menu entity.
func (mc *MenuCreate) SetParent(m *Menu) *MenuCreate {
	return mc.SetParentID(m.ID)
}

// AddChildIDs adds the "children" edge to the Menu entity by IDs.
func (mc *MenuCreate) AddChildIDs(ids ...uint64) *MenuCreate {
	mc.mutation.AddChildIDs(ids...)
	return mc
}

// AddChildren adds the "children" edges to the Menu entity.
func (mc *MenuCreate) AddChildren(m ...*Menu) *MenuCreate {
	ids := make([]uint64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mc.AddChildIDs(ids...)
}

// AddParamIDs adds the "params" edge to the MenuParam entity by IDs.
func (mc *MenuCreate) AddParamIDs(ids ...uint64) *MenuCreate {
	mc.mutation.AddParamIDs(ids...)
	return mc
}

// AddParams adds the "params" edges to the MenuParam entity.
func (mc *MenuCreate) AddParams(m ...*MenuParam) *MenuCreate {
	ids := make([]uint64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mc.AddParamIDs(ids...)
}

// Mutation returns the MenuMutation object of the builder.
func (mc *MenuCreate) Mutation() *MenuMutation {
	return mc.mutation
}

// Save creates the Menu in the database.
func (mc *MenuCreate) Save(ctx context.Context) (*Menu, error) {
	mc.defaults()
	return withHooks[*Menu, MenuMutation](ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MenuCreate) SaveX(ctx context.Context) *Menu {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MenuCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MenuCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MenuCreate) defaults() {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		v := menu.DefaultCreatedAt()
		mc.mutation.SetCreatedAt(v)
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		v := menu.DefaultUpdatedAt()
		mc.mutation.SetUpdatedAt(v)
	}
	if _, ok := mc.mutation.Path(); !ok {
		v := menu.DefaultPath
		mc.mutation.SetPath(v)
	}
	if _, ok := mc.mutation.Redirect(); !ok {
		v := menu.DefaultRedirect
		mc.mutation.SetRedirect(v)
	}
	if _, ok := mc.mutation.Component(); !ok {
		v := menu.DefaultComponent
		mc.mutation.SetComponent(v)
	}
	if _, ok := mc.mutation.OrderNo(); !ok {
		v := menu.DefaultOrderNo
		mc.mutation.SetOrderNo(v)
	}
	if _, ok := mc.mutation.Disabled(); !ok {
		v := menu.DefaultDisabled
		mc.mutation.SetDisabled(v)
	}
	if _, ok := mc.mutation.HideMenu(); !ok {
		v := menu.DefaultHideMenu
		mc.mutation.SetHideMenu(v)
	}
	if _, ok := mc.mutation.HideBreadcrumb(); !ok {
		v := menu.DefaultHideBreadcrumb
		mc.mutation.SetHideBreadcrumb(v)
	}
	if _, ok := mc.mutation.CurrentActiveMenu(); !ok {
		v := menu.DefaultCurrentActiveMenu
		mc.mutation.SetCurrentActiveMenu(v)
	}
	if _, ok := mc.mutation.IgnoreKeepAlive(); !ok {
		v := menu.DefaultIgnoreKeepAlive
		mc.mutation.SetIgnoreKeepAlive(v)
	}
	if _, ok := mc.mutation.HideTab(); !ok {
		v := menu.DefaultHideTab
		mc.mutation.SetHideTab(v)
	}
	if _, ok := mc.mutation.FrameSrc(); !ok {
		v := menu.DefaultFrameSrc
		mc.mutation.SetFrameSrc(v)
	}
	if _, ok := mc.mutation.CarryParam(); !ok {
		v := menu.DefaultCarryParam
		mc.mutation.SetCarryParam(v)
	}
	if _, ok := mc.mutation.HideChildrenInMenu(); !ok {
		v := menu.DefaultHideChildrenInMenu
		mc.mutation.SetHideChildrenInMenu(v)
	}
	if _, ok := mc.mutation.Affix(); !ok {
		v := menu.DefaultAffix
		mc.mutation.SetAffix(v)
	}
	if _, ok := mc.mutation.DynamicLevel(); !ok {
		v := menu.DefaultDynamicLevel
		mc.mutation.SetDynamicLevel(v)
	}
	if _, ok := mc.mutation.RealPath(); !ok {
		v := menu.DefaultRealPath
		mc.mutation.SetRealPath(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MenuCreate) check() error {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Menu.created_at"`)}
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Menu.updated_at"`)}
	}
	if _, ok := mc.mutation.MenuLevel(); !ok {
		return &ValidationError{Name: "menu_level", err: errors.New(`ent: missing required field "Menu.menu_level"`)}
	}
	if _, ok := mc.mutation.MenuType(); !ok {
		return &ValidationError{Name: "menu_type", err: errors.New(`ent: missing required field "Menu.menu_type"`)}
	}
	if _, ok := mc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Menu.name"`)}
	}
	if _, ok := mc.mutation.OrderNo(); !ok {
		return &ValidationError{Name: "order_no", err: errors.New(`ent: missing required field "Menu.order_no"`)}
	}
	if _, ok := mc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Menu.title"`)}
	}
	if _, ok := mc.mutation.Icon(); !ok {
		return &ValidationError{Name: "icon", err: errors.New(`ent: missing required field "Menu.icon"`)}
	}
	return nil
}

func (mc *MenuCreate) sqlSave(ctx context.Context) (*Menu, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *MenuCreate) createSpec() (*Menu, *sqlgraph.CreateSpec) {
	var (
		_node = &Menu{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: menu.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: menu.FieldID,
			},
		}
	)
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mc.mutation.CreatedAt(); ok {
		_spec.SetField(menu.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := mc.mutation.UpdatedAt(); ok {
		_spec.SetField(menu.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := mc.mutation.MenuLevel(); ok {
		_spec.SetField(menu.FieldMenuLevel, field.TypeUint32, value)
		_node.MenuLevel = value
	}
	if value, ok := mc.mutation.MenuType(); ok {
		_spec.SetField(menu.FieldMenuType, field.TypeUint32, value)
		_node.MenuType = value
	}
	if value, ok := mc.mutation.Path(); ok {
		_spec.SetField(menu.FieldPath, field.TypeString, value)
		_node.Path = value
	}
	if value, ok := mc.mutation.Name(); ok {
		_spec.SetField(menu.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := mc.mutation.Redirect(); ok {
		_spec.SetField(menu.FieldRedirect, field.TypeString, value)
		_node.Redirect = value
	}
	if value, ok := mc.mutation.Component(); ok {
		_spec.SetField(menu.FieldComponent, field.TypeString, value)
		_node.Component = value
	}
	if value, ok := mc.mutation.OrderNo(); ok {
		_spec.SetField(menu.FieldOrderNo, field.TypeUint32, value)
		_node.OrderNo = value
	}
	if value, ok := mc.mutation.Disabled(); ok {
		_spec.SetField(menu.FieldDisabled, field.TypeBool, value)
		_node.Disabled = value
	}
	if value, ok := mc.mutation.Title(); ok {
		_spec.SetField(menu.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := mc.mutation.Icon(); ok {
		_spec.SetField(menu.FieldIcon, field.TypeString, value)
		_node.Icon = value
	}
	if value, ok := mc.mutation.HideMenu(); ok {
		_spec.SetField(menu.FieldHideMenu, field.TypeBool, value)
		_node.HideMenu = value
	}
	if value, ok := mc.mutation.HideBreadcrumb(); ok {
		_spec.SetField(menu.FieldHideBreadcrumb, field.TypeBool, value)
		_node.HideBreadcrumb = value
	}
	if value, ok := mc.mutation.CurrentActiveMenu(); ok {
		_spec.SetField(menu.FieldCurrentActiveMenu, field.TypeString, value)
		_node.CurrentActiveMenu = value
	}
	if value, ok := mc.mutation.IgnoreKeepAlive(); ok {
		_spec.SetField(menu.FieldIgnoreKeepAlive, field.TypeBool, value)
		_node.IgnoreKeepAlive = value
	}
	if value, ok := mc.mutation.HideTab(); ok {
		_spec.SetField(menu.FieldHideTab, field.TypeBool, value)
		_node.HideTab = value
	}
	if value, ok := mc.mutation.FrameSrc(); ok {
		_spec.SetField(menu.FieldFrameSrc, field.TypeString, value)
		_node.FrameSrc = value
	}
	if value, ok := mc.mutation.CarryParam(); ok {
		_spec.SetField(menu.FieldCarryParam, field.TypeBool, value)
		_node.CarryParam = value
	}
	if value, ok := mc.mutation.HideChildrenInMenu(); ok {
		_spec.SetField(menu.FieldHideChildrenInMenu, field.TypeBool, value)
		_node.HideChildrenInMenu = value
	}
	if value, ok := mc.mutation.Affix(); ok {
		_spec.SetField(menu.FieldAffix, field.TypeBool, value)
		_node.Affix = value
	}
	if value, ok := mc.mutation.DynamicLevel(); ok {
		_spec.SetField(menu.FieldDynamicLevel, field.TypeUint32, value)
		_node.DynamicLevel = value
	}
	if value, ok := mc.mutation.RealPath(); ok {
		_spec.SetField(menu.FieldRealPath, field.TypeString, value)
		_node.RealPath = value
	}
	if nodes := mc.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   menu.RolesTable,
			Columns: menu.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: role.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   menu.ParentTable,
			Columns: []string{menu.ParentColumn},
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
		_node.ParentID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   menu.ChildrenTable,
			Columns: []string{menu.ChildrenColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.ParamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   menu.ParamsTable,
			Columns: []string{menu.ParamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: menuparam.FieldID,
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

// MenuCreateBulk is the builder for creating many Menu entities in bulk.
type MenuCreateBulk struct {
	config
	builders []*MenuCreate
}

// Save creates the Menu entities in the database.
func (mcb *MenuCreateBulk) Save(ctx context.Context) ([]*Menu, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Menu, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MenuMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
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
					nodes[i].ID = uint64(id)
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MenuCreateBulk) SaveX(ctx context.Context) []*Menu {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MenuCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MenuCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}
