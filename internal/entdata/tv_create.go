// Code generated by entc, DO NOT EDIT.

package entdata

import (
	"context"
	"errors"
	"fmt"
	"yola/internal/entdata/tv"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TvCreate is the builder for creating a Tv entity.
type TvCreate struct {
	config
	mutation *TvMutation
	hooks    []Hook
}

// SetLogo sets the "logo" field.
func (tc *TvCreate) SetLogo(s string) *TvCreate {
	tc.mutation.SetLogo(s)
	return tc
}

// SetVideo sets the "video" field.
func (tc *TvCreate) SetVideo(s string) *TvCreate {
	tc.mutation.SetVideo(s)
	return tc
}

// SetTitle sets the "title" field.
func (tc *TvCreate) SetTitle(s string) *TvCreate {
	tc.mutation.SetTitle(s)
	return tc
}

// SetStatus sets the "status" field.
func (tc *TvCreate) SetStatus(b bool) *TvCreate {
	tc.mutation.SetStatus(b)
	return tc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tc *TvCreate) SetNillableStatus(b *bool) *TvCreate {
	if b != nil {
		tc.SetStatus(*b)
	}
	return tc
}

// SetCountry sets the "country" field.
func (tc *TvCreate) SetCountry(s string) *TvCreate {
	tc.mutation.SetCountry(s)
	return tc
}

// SetNillableCountry sets the "country" field if the given value is not nil.
func (tc *TvCreate) SetNillableCountry(s *string) *TvCreate {
	if s != nil {
		tc.SetCountry(*s)
	}
	return tc
}

// SetDescription sets the "description" field.
func (tc *TvCreate) SetDescription(s string) *TvCreate {
	tc.mutation.SetDescription(s)
	return tc
}

// SetLanguage sets the "language" field.
func (tc *TvCreate) SetLanguage(s string) *TvCreate {
	tc.mutation.SetLanguage(s)
	return tc
}

// SetNillableLanguage sets the "language" field if the given value is not nil.
func (tc *TvCreate) SetNillableLanguage(s *string) *TvCreate {
	if s != nil {
		tc.SetLanguage(*s)
	}
	return tc
}

// Mutation returns the TvMutation object of the builder.
func (tc *TvCreate) Mutation() *TvMutation {
	return tc.mutation
}

// Save creates the Tv in the database.
func (tc *TvCreate) Save(ctx context.Context) (*Tv, error) {
	var (
		err  error
		node *Tv
	)
	tc.defaults()
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TvMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			if node, err = tc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			if tc.hooks[i] == nil {
				return nil, fmt.Errorf("entdata: uninitialized hook (forgotten import entdata/runtime?)")
			}
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TvCreate) SaveX(ctx context.Context) *Tv {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TvCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TvCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TvCreate) defaults() {
	if _, ok := tc.mutation.Status(); !ok {
		v := tv.DefaultStatus
		tc.mutation.SetStatus(v)
	}
	if _, ok := tc.mutation.Country(); !ok {
		v := tv.DefaultCountry
		tc.mutation.SetCountry(v)
	}
	if _, ok := tc.mutation.Language(); !ok {
		v := tv.DefaultLanguage
		tc.mutation.SetLanguage(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TvCreate) check() error {
	if _, ok := tc.mutation.Logo(); !ok {
		return &ValidationError{Name: "logo", err: errors.New(`entdata: missing required field "Tv.logo"`)}
	}
	if v, ok := tc.mutation.Logo(); ok {
		if err := tv.LogoValidator(v); err != nil {
			return &ValidationError{Name: "logo", err: fmt.Errorf(`entdata: validator failed for field "Tv.logo": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Video(); !ok {
		return &ValidationError{Name: "video", err: errors.New(`entdata: missing required field "Tv.video"`)}
	}
	if v, ok := tc.mutation.Video(); ok {
		if err := tv.VideoValidator(v); err != nil {
			return &ValidationError{Name: "video", err: fmt.Errorf(`entdata: validator failed for field "Tv.video": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`entdata: missing required field "Tv.title"`)}
	}
	if v, ok := tc.mutation.Title(); ok {
		if err := tv.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`entdata: validator failed for field "Tv.title": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`entdata: missing required field "Tv.status"`)}
	}
	if _, ok := tc.mutation.Country(); !ok {
		return &ValidationError{Name: "country", err: errors.New(`entdata: missing required field "Tv.country"`)}
	}
	if _, ok := tc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`entdata: missing required field "Tv.description"`)}
	}
	if v, ok := tc.mutation.Description(); ok {
		if err := tv.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`entdata: validator failed for field "Tv.description": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Language(); !ok {
		return &ValidationError{Name: "language", err: errors.New(`entdata: missing required field "Tv.language"`)}
	}
	return nil
}

func (tc *TvCreate) sqlSave(ctx context.Context) (*Tv, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (tc *TvCreate) createSpec() (*Tv, *sqlgraph.CreateSpec) {
	var (
		_node = &Tv{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: tv.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tv.FieldID,
			},
		}
	)
	if value, ok := tc.mutation.Logo(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tv.FieldLogo,
		})
		_node.Logo = value
	}
	if value, ok := tc.mutation.Video(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tv.FieldVideo,
		})
		_node.Video = value
	}
	if value, ok := tc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tv.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := tc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: tv.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := tc.mutation.Country(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tv.FieldCountry,
		})
		_node.Country = value
	}
	if value, ok := tc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tv.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := tc.mutation.Language(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tv.FieldLanguage,
		})
		_node.Language = value
	}
	return _node, _spec
}

// TvCreateBulk is the builder for creating many Tv entities in bulk.
type TvCreateBulk struct {
	config
	builders []*TvCreate
}

// Save creates the Tv entities in the database.
func (tcb *TvCreateBulk) Save(ctx context.Context) ([]*Tv, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Tv, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TvMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TvCreateBulk) SaveX(ctx context.Context) []*Tv {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TvCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TvCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
