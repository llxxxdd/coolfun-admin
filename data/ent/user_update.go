// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"coolfun-admin/data/ent/predicate"
	"coolfun-admin/data/ent/token"
	"coolfun-admin/data/ent/user"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUpdatedAt sets the "updated_at" field.
func (uu *UserUpdate) SetUpdatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// SetStatus sets the "status" field.
func (uu *UserUpdate) SetStatus(u uint8) *UserUpdate {
	uu.mutation.ResetStatus()
	uu.mutation.SetStatus(u)
	return uu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (uu *UserUpdate) SetNillableStatus(u *uint8) *UserUpdate {
	if u != nil {
		uu.SetStatus(*u)
	}
	return uu
}

// AddStatus adds u to the "status" field.
func (uu *UserUpdate) AddStatus(u int8) *UserUpdate {
	uu.mutation.AddStatus(u)
	return uu
}

// ClearStatus clears the value of the "status" field.
func (uu *UserUpdate) ClearStatus() *UserUpdate {
	uu.mutation.ClearStatus()
	return uu
}

// SetUsername sets the "username" field.
func (uu *UserUpdate) SetUsername(s string) *UserUpdate {
	uu.mutation.SetUsername(s)
	return uu
}

// SetPassword sets the "password" field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetNickname sets the "nickname" field.
func (uu *UserUpdate) SetNickname(s string) *UserUpdate {
	uu.mutation.SetNickname(s)
	return uu
}

// SetSideMode sets the "side_mode" field.
func (uu *UserUpdate) SetSideMode(s string) *UserUpdate {
	uu.mutation.SetSideMode(s)
	return uu
}

// SetNillableSideMode sets the "side_mode" field if the given value is not nil.
func (uu *UserUpdate) SetNillableSideMode(s *string) *UserUpdate {
	if s != nil {
		uu.SetSideMode(*s)
	}
	return uu
}

// ClearSideMode clears the value of the "side_mode" field.
func (uu *UserUpdate) ClearSideMode() *UserUpdate {
	uu.mutation.ClearSideMode()
	return uu
}

// SetBaseColor sets the "base_color" field.
func (uu *UserUpdate) SetBaseColor(s string) *UserUpdate {
	uu.mutation.SetBaseColor(s)
	return uu
}

// SetNillableBaseColor sets the "base_color" field if the given value is not nil.
func (uu *UserUpdate) SetNillableBaseColor(s *string) *UserUpdate {
	if s != nil {
		uu.SetBaseColor(*s)
	}
	return uu
}

// ClearBaseColor clears the value of the "base_color" field.
func (uu *UserUpdate) ClearBaseColor() *UserUpdate {
	uu.mutation.ClearBaseColor()
	return uu
}

// SetActiveColor sets the "active_color" field.
func (uu *UserUpdate) SetActiveColor(s string) *UserUpdate {
	uu.mutation.SetActiveColor(s)
	return uu
}

// SetNillableActiveColor sets the "active_color" field if the given value is not nil.
func (uu *UserUpdate) SetNillableActiveColor(s *string) *UserUpdate {
	if s != nil {
		uu.SetActiveColor(*s)
	}
	return uu
}

// ClearActiveColor clears the value of the "active_color" field.
func (uu *UserUpdate) ClearActiveColor() *UserUpdate {
	uu.mutation.ClearActiveColor()
	return uu
}

// SetRoleID sets the "role_id" field.
func (uu *UserUpdate) SetRoleID(u uint64) *UserUpdate {
	uu.mutation.ResetRoleID()
	uu.mutation.SetRoleID(u)
	return uu
}

// SetNillableRoleID sets the "role_id" field if the given value is not nil.
func (uu *UserUpdate) SetNillableRoleID(u *uint64) *UserUpdate {
	if u != nil {
		uu.SetRoleID(*u)
	}
	return uu
}

// AddRoleID adds u to the "role_id" field.
func (uu *UserUpdate) AddRoleID(u int64) *UserUpdate {
	uu.mutation.AddRoleID(u)
	return uu
}

// ClearRoleID clears the value of the "role_id" field.
func (uu *UserUpdate) ClearRoleID() *UserUpdate {
	uu.mutation.ClearRoleID()
	return uu
}

// SetMobile sets the "mobile" field.
func (uu *UserUpdate) SetMobile(s string) *UserUpdate {
	uu.mutation.SetMobile(s)
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uu *UserUpdate) SetNillableEmail(s *string) *UserUpdate {
	if s != nil {
		uu.SetEmail(*s)
	}
	return uu
}

// ClearEmail clears the value of the "email" field.
func (uu *UserUpdate) ClearEmail() *UserUpdate {
	uu.mutation.ClearEmail()
	return uu
}

// SetWecom sets the "wecom" field.
func (uu *UserUpdate) SetWecom(s string) *UserUpdate {
	uu.mutation.SetWecom(s)
	return uu
}

// SetNillableWecom sets the "wecom" field if the given value is not nil.
func (uu *UserUpdate) SetNillableWecom(s *string) *UserUpdate {
	if s != nil {
		uu.SetWecom(*s)
	}
	return uu
}

// ClearWecom clears the value of the "wecom" field.
func (uu *UserUpdate) ClearWecom() *UserUpdate {
	uu.mutation.ClearWecom()
	return uu
}

// SetAvatar sets the "avatar" field.
func (uu *UserUpdate) SetAvatar(s string) *UserUpdate {
	uu.mutation.SetAvatar(s)
	return uu
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAvatar(s *string) *UserUpdate {
	if s != nil {
		uu.SetAvatar(*s)
	}
	return uu
}

// ClearAvatar clears the value of the "avatar" field.
func (uu *UserUpdate) ClearAvatar() *UserUpdate {
	uu.mutation.ClearAvatar()
	return uu
}

// SetTokenID sets the "token" edge to the Token entity by ID.
func (uu *UserUpdate) SetTokenID(id uint64) *UserUpdate {
	uu.mutation.SetTokenID(id)
	return uu
}

// SetNillableTokenID sets the "token" edge to the Token entity by ID if the given value is not nil.
func (uu *UserUpdate) SetNillableTokenID(id *uint64) *UserUpdate {
	if id != nil {
		uu = uu.SetTokenID(*id)
	}
	return uu
}

// SetToken sets the "token" edge to the Token entity.
func (uu *UserUpdate) SetToken(t *Token) *UserUpdate {
	return uu.SetTokenID(t.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearToken clears the "token" edge to the Token entity.
func (uu *UserUpdate) ClearToken() *UserUpdate {
	uu.mutation.ClearToken()
	return uu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	uu.defaults()
	return withHooks[int, UserMutation](ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uu *UserUpdate) defaults() {
	if _, ok := uu.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uu.mutation.SetUpdatedAt(v)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uu.mutation.Status(); ok {
		_spec.SetField(user.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := uu.mutation.AddedStatus(); ok {
		_spec.AddField(user.FieldStatus, field.TypeUint8, value)
	}
	if uu.mutation.StatusCleared() {
		_spec.ClearField(user.FieldStatus, field.TypeUint8)
	}
	if value, ok := uu.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uu.mutation.Nickname(); ok {
		_spec.SetField(user.FieldNickname, field.TypeString, value)
	}
	if value, ok := uu.mutation.SideMode(); ok {
		_spec.SetField(user.FieldSideMode, field.TypeString, value)
	}
	if uu.mutation.SideModeCleared() {
		_spec.ClearField(user.FieldSideMode, field.TypeString)
	}
	if value, ok := uu.mutation.BaseColor(); ok {
		_spec.SetField(user.FieldBaseColor, field.TypeString, value)
	}
	if uu.mutation.BaseColorCleared() {
		_spec.ClearField(user.FieldBaseColor, field.TypeString)
	}
	if value, ok := uu.mutation.ActiveColor(); ok {
		_spec.SetField(user.FieldActiveColor, field.TypeString, value)
	}
	if uu.mutation.ActiveColorCleared() {
		_spec.ClearField(user.FieldActiveColor, field.TypeString)
	}
	if value, ok := uu.mutation.RoleID(); ok {
		_spec.SetField(user.FieldRoleID, field.TypeUint64, value)
	}
	if value, ok := uu.mutation.AddedRoleID(); ok {
		_spec.AddField(user.FieldRoleID, field.TypeUint64, value)
	}
	if uu.mutation.RoleIDCleared() {
		_spec.ClearField(user.FieldRoleID, field.TypeUint64)
	}
	if value, ok := uu.mutation.Mobile(); ok {
		_spec.SetField(user.FieldMobile, field.TypeString, value)
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if uu.mutation.EmailCleared() {
		_spec.ClearField(user.FieldEmail, field.TypeString)
	}
	if value, ok := uu.mutation.Wecom(); ok {
		_spec.SetField(user.FieldWecom, field.TypeString, value)
	}
	if uu.mutation.WecomCleared() {
		_spec.ClearField(user.FieldWecom, field.TypeString)
	}
	if value, ok := uu.mutation.Avatar(); ok {
		_spec.SetField(user.FieldAvatar, field.TypeString, value)
	}
	if uu.mutation.AvatarCleared() {
		_spec.ClearField(user.FieldAvatar, field.TypeString)
	}
	if uu.mutation.TokenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.TokenTable,
			Columns: []string{user.TokenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: token.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.TokenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.TokenTable,
			Columns: []string{user.TokenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: token.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (uuo *UserUpdateOne) SetUpdatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// SetStatus sets the "status" field.
func (uuo *UserUpdateOne) SetStatus(u uint8) *UserUpdateOne {
	uuo.mutation.ResetStatus()
	uuo.mutation.SetStatus(u)
	return uuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableStatus(u *uint8) *UserUpdateOne {
	if u != nil {
		uuo.SetStatus(*u)
	}
	return uuo
}

// AddStatus adds u to the "status" field.
func (uuo *UserUpdateOne) AddStatus(u int8) *UserUpdateOne {
	uuo.mutation.AddStatus(u)
	return uuo
}

// ClearStatus clears the value of the "status" field.
func (uuo *UserUpdateOne) ClearStatus() *UserUpdateOne {
	uuo.mutation.ClearStatus()
	return uuo
}

// SetUsername sets the "username" field.
func (uuo *UserUpdateOne) SetUsername(s string) *UserUpdateOne {
	uuo.mutation.SetUsername(s)
	return uuo
}

// SetPassword sets the "password" field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetNickname sets the "nickname" field.
func (uuo *UserUpdateOne) SetNickname(s string) *UserUpdateOne {
	uuo.mutation.SetNickname(s)
	return uuo
}

// SetSideMode sets the "side_mode" field.
func (uuo *UserUpdateOne) SetSideMode(s string) *UserUpdateOne {
	uuo.mutation.SetSideMode(s)
	return uuo
}

// SetNillableSideMode sets the "side_mode" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableSideMode(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetSideMode(*s)
	}
	return uuo
}

// ClearSideMode clears the value of the "side_mode" field.
func (uuo *UserUpdateOne) ClearSideMode() *UserUpdateOne {
	uuo.mutation.ClearSideMode()
	return uuo
}

// SetBaseColor sets the "base_color" field.
func (uuo *UserUpdateOne) SetBaseColor(s string) *UserUpdateOne {
	uuo.mutation.SetBaseColor(s)
	return uuo
}

// SetNillableBaseColor sets the "base_color" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableBaseColor(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetBaseColor(*s)
	}
	return uuo
}

// ClearBaseColor clears the value of the "base_color" field.
func (uuo *UserUpdateOne) ClearBaseColor() *UserUpdateOne {
	uuo.mutation.ClearBaseColor()
	return uuo
}

// SetActiveColor sets the "active_color" field.
func (uuo *UserUpdateOne) SetActiveColor(s string) *UserUpdateOne {
	uuo.mutation.SetActiveColor(s)
	return uuo
}

// SetNillableActiveColor sets the "active_color" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableActiveColor(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetActiveColor(*s)
	}
	return uuo
}

// ClearActiveColor clears the value of the "active_color" field.
func (uuo *UserUpdateOne) ClearActiveColor() *UserUpdateOne {
	uuo.mutation.ClearActiveColor()
	return uuo
}

// SetRoleID sets the "role_id" field.
func (uuo *UserUpdateOne) SetRoleID(u uint64) *UserUpdateOne {
	uuo.mutation.ResetRoleID()
	uuo.mutation.SetRoleID(u)
	return uuo
}

// SetNillableRoleID sets the "role_id" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableRoleID(u *uint64) *UserUpdateOne {
	if u != nil {
		uuo.SetRoleID(*u)
	}
	return uuo
}

// AddRoleID adds u to the "role_id" field.
func (uuo *UserUpdateOne) AddRoleID(u int64) *UserUpdateOne {
	uuo.mutation.AddRoleID(u)
	return uuo
}

// ClearRoleID clears the value of the "role_id" field.
func (uuo *UserUpdateOne) ClearRoleID() *UserUpdateOne {
	uuo.mutation.ClearRoleID()
	return uuo
}

// SetMobile sets the "mobile" field.
func (uuo *UserUpdateOne) SetMobile(s string) *UserUpdateOne {
	uuo.mutation.SetMobile(s)
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableEmail(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetEmail(*s)
	}
	return uuo
}

// ClearEmail clears the value of the "email" field.
func (uuo *UserUpdateOne) ClearEmail() *UserUpdateOne {
	uuo.mutation.ClearEmail()
	return uuo
}

// SetWecom sets the "wecom" field.
func (uuo *UserUpdateOne) SetWecom(s string) *UserUpdateOne {
	uuo.mutation.SetWecom(s)
	return uuo
}

// SetNillableWecom sets the "wecom" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableWecom(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetWecom(*s)
	}
	return uuo
}

// ClearWecom clears the value of the "wecom" field.
func (uuo *UserUpdateOne) ClearWecom() *UserUpdateOne {
	uuo.mutation.ClearWecom()
	return uuo
}

// SetAvatar sets the "avatar" field.
func (uuo *UserUpdateOne) SetAvatar(s string) *UserUpdateOne {
	uuo.mutation.SetAvatar(s)
	return uuo
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAvatar(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetAvatar(*s)
	}
	return uuo
}

// ClearAvatar clears the value of the "avatar" field.
func (uuo *UserUpdateOne) ClearAvatar() *UserUpdateOne {
	uuo.mutation.ClearAvatar()
	return uuo
}

// SetTokenID sets the "token" edge to the Token entity by ID.
func (uuo *UserUpdateOne) SetTokenID(id uint64) *UserUpdateOne {
	uuo.mutation.SetTokenID(id)
	return uuo
}

// SetNillableTokenID sets the "token" edge to the Token entity by ID if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableTokenID(id *uint64) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetTokenID(*id)
	}
	return uuo
}

// SetToken sets the "token" edge to the Token entity.
func (uuo *UserUpdateOne) SetToken(t *Token) *UserUpdateOne {
	return uuo.SetTokenID(t.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearToken clears the "token" edge to the Token entity.
func (uuo *UserUpdateOne) ClearToken() *UserUpdateOne {
	uuo.mutation.ClearToken()
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	uuo.defaults()
	return withHooks[*User, UserMutation](ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uuo *UserUpdateOne) defaults() {
	if _, ok := uuo.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uuo.mutation.SetUpdatedAt(v)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.Status(); ok {
		_spec.SetField(user.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := uuo.mutation.AddedStatus(); ok {
		_spec.AddField(user.FieldStatus, field.TypeUint8, value)
	}
	if uuo.mutation.StatusCleared() {
		_spec.ClearField(user.FieldStatus, field.TypeUint8)
	}
	if value, ok := uuo.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Nickname(); ok {
		_spec.SetField(user.FieldNickname, field.TypeString, value)
	}
	if value, ok := uuo.mutation.SideMode(); ok {
		_spec.SetField(user.FieldSideMode, field.TypeString, value)
	}
	if uuo.mutation.SideModeCleared() {
		_spec.ClearField(user.FieldSideMode, field.TypeString)
	}
	if value, ok := uuo.mutation.BaseColor(); ok {
		_spec.SetField(user.FieldBaseColor, field.TypeString, value)
	}
	if uuo.mutation.BaseColorCleared() {
		_spec.ClearField(user.FieldBaseColor, field.TypeString)
	}
	if value, ok := uuo.mutation.ActiveColor(); ok {
		_spec.SetField(user.FieldActiveColor, field.TypeString, value)
	}
	if uuo.mutation.ActiveColorCleared() {
		_spec.ClearField(user.FieldActiveColor, field.TypeString)
	}
	if value, ok := uuo.mutation.RoleID(); ok {
		_spec.SetField(user.FieldRoleID, field.TypeUint64, value)
	}
	if value, ok := uuo.mutation.AddedRoleID(); ok {
		_spec.AddField(user.FieldRoleID, field.TypeUint64, value)
	}
	if uuo.mutation.RoleIDCleared() {
		_spec.ClearField(user.FieldRoleID, field.TypeUint64)
	}
	if value, ok := uuo.mutation.Mobile(); ok {
		_spec.SetField(user.FieldMobile, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if uuo.mutation.EmailCleared() {
		_spec.ClearField(user.FieldEmail, field.TypeString)
	}
	if value, ok := uuo.mutation.Wecom(); ok {
		_spec.SetField(user.FieldWecom, field.TypeString, value)
	}
	if uuo.mutation.WecomCleared() {
		_spec.ClearField(user.FieldWecom, field.TypeString)
	}
	if value, ok := uuo.mutation.Avatar(); ok {
		_spec.SetField(user.FieldAvatar, field.TypeString, value)
	}
	if uuo.mutation.AvatarCleared() {
		_spec.ClearField(user.FieldAvatar, field.TypeString)
	}
	if uuo.mutation.TokenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.TokenTable,
			Columns: []string{user.TokenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: token.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.TokenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.TokenTable,
			Columns: []string{user.TokenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: token.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
