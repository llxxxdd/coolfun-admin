// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"coolfun-admin/data/ent/dictionary"
	"coolfun-admin/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DictionaryDelete is the builder for deleting a Dictionary entity.
type DictionaryDelete struct {
	config
	hooks    []Hook
	mutation *DictionaryMutation
}

// Where appends a list predicates to the DictionaryDelete builder.
func (dd *DictionaryDelete) Where(ps ...predicate.Dictionary) *DictionaryDelete {
	dd.mutation.Where(ps...)
	return dd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dd *DictionaryDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, DictionaryMutation](ctx, dd.sqlExec, dd.mutation, dd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (dd *DictionaryDelete) ExecX(ctx context.Context) int {
	n, err := dd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dd *DictionaryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: dictionary.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: dictionary.FieldID,
			},
		},
	}
	if ps := dd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, dd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	dd.mutation.done = true
	return affected, err
}

// DictionaryDeleteOne is the builder for deleting a single Dictionary entity.
type DictionaryDeleteOne struct {
	dd *DictionaryDelete
}

// Exec executes the deletion query.
func (ddo *DictionaryDeleteOne) Exec(ctx context.Context) error {
	n, err := ddo.dd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{dictionary.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ddo *DictionaryDeleteOne) ExecX(ctx context.Context) {
	ddo.dd.ExecX(ctx)
}
