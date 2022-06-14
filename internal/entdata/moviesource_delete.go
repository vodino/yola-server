// Code generated by entc, DO NOT EDIT.

package entdata

import (
	"context"
	"fmt"
	"yola/internal/entdata/moviesource"
	"yola/internal/entdata/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MovieSourceDelete is the builder for deleting a MovieSource entity.
type MovieSourceDelete struct {
	config
	hooks    []Hook
	mutation *MovieSourceMutation
}

// Where appends a list predicates to the MovieSourceDelete builder.
func (msd *MovieSourceDelete) Where(ps ...predicate.MovieSource) *MovieSourceDelete {
	msd.mutation.Where(ps...)
	return msd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (msd *MovieSourceDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(msd.hooks) == 0 {
		affected, err = msd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MovieSourceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			msd.mutation = mutation
			affected, err = msd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(msd.hooks) - 1; i >= 0; i-- {
			if msd.hooks[i] == nil {
				return 0, fmt.Errorf("entdata: uninitialized hook (forgotten import entdata/runtime?)")
			}
			mut = msd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, msd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (msd *MovieSourceDelete) ExecX(ctx context.Context) int {
	n, err := msd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (msd *MovieSourceDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: moviesource.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: moviesource.FieldID,
			},
		},
	}
	if ps := msd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, msd.driver, _spec)
}

// MovieSourceDeleteOne is the builder for deleting a single MovieSource entity.
type MovieSourceDeleteOne struct {
	msd *MovieSourceDelete
}

// Exec executes the deletion query.
func (msdo *MovieSourceDeleteOne) Exec(ctx context.Context) error {
	n, err := msdo.msd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{moviesource.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (msdo *MovieSourceDeleteOne) ExecX(ctx context.Context) {
	msdo.msd.ExecX(ctx)
}