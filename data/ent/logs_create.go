// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"coolfun-admin/data/ent/logs"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LogsCreate is the builder for creating a Logs entity.
type LogsCreate struct {
	config
	mutation *LogsMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (lc *LogsCreate) SetCreatedAt(t time.Time) *LogsCreate {
	lc.mutation.SetCreatedAt(t)
	return lc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lc *LogsCreate) SetNillableCreatedAt(t *time.Time) *LogsCreate {
	if t != nil {
		lc.SetCreatedAt(*t)
	}
	return lc
}

// SetUpdatedAt sets the "updated_at" field.
func (lc *LogsCreate) SetUpdatedAt(t time.Time) *LogsCreate {
	lc.mutation.SetUpdatedAt(t)
	return lc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (lc *LogsCreate) SetNillableUpdatedAt(t *time.Time) *LogsCreate {
	if t != nil {
		lc.SetUpdatedAt(*t)
	}
	return lc
}

// SetType sets the "type" field.
func (lc *LogsCreate) SetType(s string) *LogsCreate {
	lc.mutation.SetType(s)
	return lc
}

// SetMethod sets the "method" field.
func (lc *LogsCreate) SetMethod(s string) *LogsCreate {
	lc.mutation.SetMethod(s)
	return lc
}

// SetAPI sets the "api" field.
func (lc *LogsCreate) SetAPI(s string) *LogsCreate {
	lc.mutation.SetAPI(s)
	return lc
}

// SetSuccess sets the "success" field.
func (lc *LogsCreate) SetSuccess(b bool) *LogsCreate {
	lc.mutation.SetSuccess(b)
	return lc
}

// SetReqContent sets the "req_content" field.
func (lc *LogsCreate) SetReqContent(s string) *LogsCreate {
	lc.mutation.SetReqContent(s)
	return lc
}

// SetNillableReqContent sets the "req_content" field if the given value is not nil.
func (lc *LogsCreate) SetNillableReqContent(s *string) *LogsCreate {
	if s != nil {
		lc.SetReqContent(*s)
	}
	return lc
}

// SetRespContent sets the "resp_content" field.
func (lc *LogsCreate) SetRespContent(s string) *LogsCreate {
	lc.mutation.SetRespContent(s)
	return lc
}

// SetNillableRespContent sets the "resp_content" field if the given value is not nil.
func (lc *LogsCreate) SetNillableRespContent(s *string) *LogsCreate {
	if s != nil {
		lc.SetRespContent(*s)
	}
	return lc
}

// SetIP sets the "ip" field.
func (lc *LogsCreate) SetIP(s string) *LogsCreate {
	lc.mutation.SetIP(s)
	return lc
}

// SetNillableIP sets the "ip" field if the given value is not nil.
func (lc *LogsCreate) SetNillableIP(s *string) *LogsCreate {
	if s != nil {
		lc.SetIP(*s)
	}
	return lc
}

// SetUserAgent sets the "user_agent" field.
func (lc *LogsCreate) SetUserAgent(s string) *LogsCreate {
	lc.mutation.SetUserAgent(s)
	return lc
}

// SetNillableUserAgent sets the "user_agent" field if the given value is not nil.
func (lc *LogsCreate) SetNillableUserAgent(s *string) *LogsCreate {
	if s != nil {
		lc.SetUserAgent(*s)
	}
	return lc
}

// SetOperator sets the "operator" field.
func (lc *LogsCreate) SetOperator(s string) *LogsCreate {
	lc.mutation.SetOperator(s)
	return lc
}

// SetNillableOperator sets the "operator" field if the given value is not nil.
func (lc *LogsCreate) SetNillableOperator(s *string) *LogsCreate {
	if s != nil {
		lc.SetOperator(*s)
	}
	return lc
}

// SetTime sets the "time" field.
func (lc *LogsCreate) SetTime(i int) *LogsCreate {
	lc.mutation.SetTime(i)
	return lc
}

// SetNillableTime sets the "time" field if the given value is not nil.
func (lc *LogsCreate) SetNillableTime(i *int) *LogsCreate {
	if i != nil {
		lc.SetTime(*i)
	}
	return lc
}

// SetID sets the "id" field.
func (lc *LogsCreate) SetID(u uint64) *LogsCreate {
	lc.mutation.SetID(u)
	return lc
}

// Mutation returns the LogsMutation object of the builder.
func (lc *LogsCreate) Mutation() *LogsMutation {
	return lc.mutation
}

// Save creates the Logs in the database.
func (lc *LogsCreate) Save(ctx context.Context) (*Logs, error) {
	lc.defaults()
	return withHooks[*Logs, LogsMutation](ctx, lc.sqlSave, lc.mutation, lc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (lc *LogsCreate) SaveX(ctx context.Context) *Logs {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lc *LogsCreate) Exec(ctx context.Context) error {
	_, err := lc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lc *LogsCreate) ExecX(ctx context.Context) {
	if err := lc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lc *LogsCreate) defaults() {
	if _, ok := lc.mutation.CreatedAt(); !ok {
		v := logs.DefaultCreatedAt()
		lc.mutation.SetCreatedAt(v)
	}
	if _, ok := lc.mutation.UpdatedAt(); !ok {
		v := logs.DefaultUpdatedAt()
		lc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lc *LogsCreate) check() error {
	if _, ok := lc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Logs.created_at"`)}
	}
	if _, ok := lc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Logs.updated_at"`)}
	}
	if _, ok := lc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Logs.type"`)}
	}
	if _, ok := lc.mutation.Method(); !ok {
		return &ValidationError{Name: "method", err: errors.New(`ent: missing required field "Logs.method"`)}
	}
	if _, ok := lc.mutation.API(); !ok {
		return &ValidationError{Name: "api", err: errors.New(`ent: missing required field "Logs.api"`)}
	}
	if _, ok := lc.mutation.Success(); !ok {
		return &ValidationError{Name: "success", err: errors.New(`ent: missing required field "Logs.success"`)}
	}
	return nil
}

func (lc *LogsCreate) sqlSave(ctx context.Context) (*Logs, error) {
	if err := lc.check(); err != nil {
		return nil, err
	}
	_node, _spec := lc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	lc.mutation.id = &_node.ID
	lc.mutation.done = true
	return _node, nil
}

func (lc *LogsCreate) createSpec() (*Logs, *sqlgraph.CreateSpec) {
	var (
		_node = &Logs{config: lc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: logs.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: logs.FieldID,
			},
		}
	)
	if id, ok := lc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := lc.mutation.CreatedAt(); ok {
		_spec.SetField(logs.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := lc.mutation.UpdatedAt(); ok {
		_spec.SetField(logs.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := lc.mutation.GetType(); ok {
		_spec.SetField(logs.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := lc.mutation.Method(); ok {
		_spec.SetField(logs.FieldMethod, field.TypeString, value)
		_node.Method = value
	}
	if value, ok := lc.mutation.API(); ok {
		_spec.SetField(logs.FieldAPI, field.TypeString, value)
		_node.API = value
	}
	if value, ok := lc.mutation.Success(); ok {
		_spec.SetField(logs.FieldSuccess, field.TypeBool, value)
		_node.Success = value
	}
	if value, ok := lc.mutation.ReqContent(); ok {
		_spec.SetField(logs.FieldReqContent, field.TypeString, value)
		_node.ReqContent = value
	}
	if value, ok := lc.mutation.RespContent(); ok {
		_spec.SetField(logs.FieldRespContent, field.TypeString, value)
		_node.RespContent = value
	}
	if value, ok := lc.mutation.IP(); ok {
		_spec.SetField(logs.FieldIP, field.TypeString, value)
		_node.IP = value
	}
	if value, ok := lc.mutation.UserAgent(); ok {
		_spec.SetField(logs.FieldUserAgent, field.TypeString, value)
		_node.UserAgent = value
	}
	if value, ok := lc.mutation.Operator(); ok {
		_spec.SetField(logs.FieldOperator, field.TypeString, value)
		_node.Operator = value
	}
	if value, ok := lc.mutation.Time(); ok {
		_spec.SetField(logs.FieldTime, field.TypeInt, value)
		_node.Time = value
	}
	return _node, _spec
}

// LogsCreateBulk is the builder for creating many Logs entities in bulk.
type LogsCreateBulk struct {
	config
	builders []*LogsCreate
}

// Save creates the Logs entities in the database.
func (lcb *LogsCreateBulk) Save(ctx context.Context) ([]*Logs, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lcb.builders))
	nodes := make([]*Logs, len(lcb.builders))
	mutators := make([]Mutator, len(lcb.builders))
	for i := range lcb.builders {
		func(i int, root context.Context) {
			builder := lcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LogsMutation)
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
					_, err = mutators[i+1].Mutate(root, lcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, lcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lcb *LogsCreateBulk) SaveX(ctx context.Context) []*Logs {
	v, err := lcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lcb *LogsCreateBulk) Exec(ctx context.Context) error {
	_, err := lcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcb *LogsCreateBulk) ExecX(ctx context.Context) {
	if err := lcb.Exec(ctx); err != nil {
		panic(err)
	}
}
