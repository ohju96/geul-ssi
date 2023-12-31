// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"geulSsi/ent/predicate"
	"geulSsi/ent/wisesaying"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WiseSayingDelete is the builder for deleting a WiseSaying entity.
type WiseSayingDelete struct {
	config
	hooks    []Hook
	mutation *WiseSayingMutation
}

// Where appends a list predicates to the WiseSayingDelete builder.
func (wsd *WiseSayingDelete) Where(ps ...predicate.WiseSaying) *WiseSayingDelete {
	wsd.mutation.Where(ps...)
	return wsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wsd *WiseSayingDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, wsd.sqlExec, wsd.mutation, wsd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (wsd *WiseSayingDelete) ExecX(ctx context.Context) int {
	n, err := wsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wsd *WiseSayingDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(wisesaying.Table, sqlgraph.NewFieldSpec(wisesaying.FieldID, field.TypeInt))
	if ps := wsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, wsd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	wsd.mutation.done = true
	return affected, err
}

// WiseSayingDeleteOne is the builder for deleting a single WiseSaying entity.
type WiseSayingDeleteOne struct {
	wsd *WiseSayingDelete
}

// Where appends a list predicates to the WiseSayingDelete builder.
func (wsdo *WiseSayingDeleteOne) Where(ps ...predicate.WiseSaying) *WiseSayingDeleteOne {
	wsdo.wsd.mutation.Where(ps...)
	return wsdo
}

// Exec executes the deletion query.
func (wsdo *WiseSayingDeleteOne) Exec(ctx context.Context) error {
	n, err := wsdo.wsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{wisesaying.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wsdo *WiseSayingDeleteOne) ExecX(ctx context.Context) {
	if err := wsdo.Exec(ctx); err != nil {
		panic(err)
	}
}
