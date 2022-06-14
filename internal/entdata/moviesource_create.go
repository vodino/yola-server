// Code generated by entc, DO NOT EDIT.

package entdata

import (
	"context"
	"errors"
	"fmt"
	"yola/internal/entdata/moviesource"
	"yola/internal/entdata/schema"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MovieSourceCreate is the builder for creating a MovieSource entity.
type MovieSourceCreate struct {
	config
	mutation *MovieSourceMutation
	hooks    []Hook
}

// SetMangaSerieSearchURL sets the "manga_serie_search_url" field.
func (msc *MovieSourceCreate) SetMangaSerieSearchURL(s string) *MovieSourceCreate {
	msc.mutation.SetMangaSerieSearchURL(s)
	return msc
}

// SetNillableMangaSerieSearchURL sets the "manga_serie_search_url" field if the given value is not nil.
func (msc *MovieSourceCreate) SetNillableMangaSerieSearchURL(s *string) *MovieSourceCreate {
	if s != nil {
		msc.SetMangaSerieSearchURL(*s)
	}
	return msc
}

// SetMangaFilmSearchURL sets the "manga_film_search_url" field.
func (msc *MovieSourceCreate) SetMangaFilmSearchURL(s string) *MovieSourceCreate {
	msc.mutation.SetMangaFilmSearchURL(s)
	return msc
}

// SetNillableMangaFilmSearchURL sets the "manga_film_search_url" field if the given value is not nil.
func (msc *MovieSourceCreate) SetNillableMangaFilmSearchURL(s *string) *MovieSourceCreate {
	if s != nil {
		msc.SetMangaFilmSearchURL(*s)
	}
	return msc
}

// SetSerieSearchURL sets the "serie_search_url" field.
func (msc *MovieSourceCreate) SetSerieSearchURL(s string) *MovieSourceCreate {
	msc.mutation.SetSerieSearchURL(s)
	return msc
}

// SetNillableSerieSearchURL sets the "serie_search_url" field if the given value is not nil.
func (msc *MovieSourceCreate) SetNillableSerieSearchURL(s *string) *MovieSourceCreate {
	if s != nil {
		msc.SetSerieSearchURL(*s)
	}
	return msc
}

// SetFilmSearchURL sets the "film_search_url" field.
func (msc *MovieSourceCreate) SetFilmSearchURL(s string) *MovieSourceCreate {
	msc.mutation.SetFilmSearchURL(s)
	return msc
}

// SetNillableFilmSearchURL sets the "film_search_url" field if the given value is not nil.
func (msc *MovieSourceCreate) SetNillableFilmSearchURL(s *string) *MovieSourceCreate {
	if s != nil {
		msc.SetFilmSearchURL(*s)
	}
	return msc
}

// SetMangaSerieLatestURL sets the "manga_serie_latest_url" field.
func (msc *MovieSourceCreate) SetMangaSerieLatestURL(s string) *MovieSourceCreate {
	msc.mutation.SetMangaSerieLatestURL(s)
	return msc
}

// SetNillableMangaSerieLatestURL sets the "manga_serie_latest_url" field if the given value is not nil.
func (msc *MovieSourceCreate) SetNillableMangaSerieLatestURL(s *string) *MovieSourceCreate {
	if s != nil {
		msc.SetMangaSerieLatestURL(*s)
	}
	return msc
}

// SetMangaFilmLatestURL sets the "manga_film_latest_url" field.
func (msc *MovieSourceCreate) SetMangaFilmLatestURL(s string) *MovieSourceCreate {
	msc.mutation.SetMangaFilmLatestURL(s)
	return msc
}

// SetNillableMangaFilmLatestURL sets the "manga_film_latest_url" field if the given value is not nil.
func (msc *MovieSourceCreate) SetNillableMangaFilmLatestURL(s *string) *MovieSourceCreate {
	if s != nil {
		msc.SetMangaFilmLatestURL(*s)
	}
	return msc
}

// SetSerieLatestURL sets the "serie_latest_url" field.
func (msc *MovieSourceCreate) SetSerieLatestURL(s string) *MovieSourceCreate {
	msc.mutation.SetSerieLatestURL(s)
	return msc
}

// SetNillableSerieLatestURL sets the "serie_latest_url" field if the given value is not nil.
func (msc *MovieSourceCreate) SetNillableSerieLatestURL(s *string) *MovieSourceCreate {
	if s != nil {
		msc.SetSerieLatestURL(*s)
	}
	return msc
}

// SetFilmLatestURL sets the "film_latest_url" field.
func (msc *MovieSourceCreate) SetFilmLatestURL(s string) *MovieSourceCreate {
	msc.mutation.SetFilmLatestURL(s)
	return msc
}

// SetNillableFilmLatestURL sets the "film_latest_url" field if the given value is not nil.
func (msc *MovieSourceCreate) SetNillableFilmLatestURL(s *string) *MovieSourceCreate {
	if s != nil {
		msc.SetFilmLatestURL(*s)
	}
	return msc
}

// SetMangaSerieLatestPostSelector sets the "manga_serie_latest_post_selector" field.
func (msc *MovieSourceCreate) SetMangaSerieLatestPostSelector(sps *schema.MoviePostSelector) *MovieSourceCreate {
	msc.mutation.SetMangaSerieLatestPostSelector(sps)
	return msc
}

// SetMangaFilmLatestPostSelector sets the "manga_film_latest_post_selector" field.
func (msc *MovieSourceCreate) SetMangaFilmLatestPostSelector(sps *schema.MoviePostSelector) *MovieSourceCreate {
	msc.mutation.SetMangaFilmLatestPostSelector(sps)
	return msc
}

// SetSerieLatestPostSelector sets the "serie_latest_post_selector" field.
func (msc *MovieSourceCreate) SetSerieLatestPostSelector(sps *schema.MoviePostSelector) *MovieSourceCreate {
	msc.mutation.SetSerieLatestPostSelector(sps)
	return msc
}

// SetFilmLatestPostSelector sets the "film_latest_post_selector" field.
func (msc *MovieSourceCreate) SetFilmLatestPostSelector(sps *schema.MoviePostSelector) *MovieSourceCreate {
	msc.mutation.SetFilmLatestPostSelector(sps)
	return msc
}

// SetMangaSerieSearchPostSelector sets the "manga_serie_search_post_selector" field.
func (msc *MovieSourceCreate) SetMangaSerieSearchPostSelector(sps *schema.MoviePostSelector) *MovieSourceCreate {
	msc.mutation.SetMangaSerieSearchPostSelector(sps)
	return msc
}

// SetMangaFilmSearchPostSelector sets the "manga_film_search_post_selector" field.
func (msc *MovieSourceCreate) SetMangaFilmSearchPostSelector(sps *schema.MoviePostSelector) *MovieSourceCreate {
	msc.mutation.SetMangaFilmSearchPostSelector(sps)
	return msc
}

// SetSerieSearchPostSelector sets the "serie_search_post_selector" field.
func (msc *MovieSourceCreate) SetSerieSearchPostSelector(sps *schema.MoviePostSelector) *MovieSourceCreate {
	msc.mutation.SetSerieSearchPostSelector(sps)
	return msc
}

// SetFilmSearchPostSelector sets the "film_search_post_selector" field.
func (msc *MovieSourceCreate) SetFilmSearchPostSelector(sps *schema.MoviePostSelector) *MovieSourceCreate {
	msc.mutation.SetFilmSearchPostSelector(sps)
	return msc
}

// SetMangaSerieArticleSelector sets the "manga_serie_article_selector" field.
func (msc *MovieSourceCreate) SetMangaSerieArticleSelector(sas *schema.MovieArticleSelector) *MovieSourceCreate {
	msc.mutation.SetMangaSerieArticleSelector(sas)
	return msc
}

// SetMangaFilmArticleSelector sets the "manga_film_article_selector" field.
func (msc *MovieSourceCreate) SetMangaFilmArticleSelector(sas *schema.MovieArticleSelector) *MovieSourceCreate {
	msc.mutation.SetMangaFilmArticleSelector(sas)
	return msc
}

// SetSerieArticleSelector sets the "serie_article_selector" field.
func (msc *MovieSourceCreate) SetSerieArticleSelector(sas *schema.MovieArticleSelector) *MovieSourceCreate {
	msc.mutation.SetSerieArticleSelector(sas)
	return msc
}

// SetFilmArticleSelector sets the "film_article_selector" field.
func (msc *MovieSourceCreate) SetFilmArticleSelector(sas *schema.MovieArticleSelector) *MovieSourceCreate {
	msc.mutation.SetFilmArticleSelector(sas)
	return msc
}

// SetStatus sets the "status" field.
func (msc *MovieSourceCreate) SetStatus(b bool) *MovieSourceCreate {
	msc.mutation.SetStatus(b)
	return msc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (msc *MovieSourceCreate) SetNillableStatus(b *bool) *MovieSourceCreate {
	if b != nil {
		msc.SetStatus(*b)
	}
	return msc
}

// SetName sets the "name" field.
func (msc *MovieSourceCreate) SetName(s string) *MovieSourceCreate {
	msc.mutation.SetName(s)
	return msc
}

// SetURL sets the "url" field.
func (msc *MovieSourceCreate) SetURL(s string) *MovieSourceCreate {
	msc.mutation.SetURL(s)
	return msc
}

// Mutation returns the MovieSourceMutation object of the builder.
func (msc *MovieSourceCreate) Mutation() *MovieSourceMutation {
	return msc.mutation
}

// Save creates the MovieSource in the database.
func (msc *MovieSourceCreate) Save(ctx context.Context) (*MovieSource, error) {
	var (
		err  error
		node *MovieSource
	)
	msc.defaults()
	if len(msc.hooks) == 0 {
		if err = msc.check(); err != nil {
			return nil, err
		}
		node, err = msc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MovieSourceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = msc.check(); err != nil {
				return nil, err
			}
			msc.mutation = mutation
			if node, err = msc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(msc.hooks) - 1; i >= 0; i-- {
			if msc.hooks[i] == nil {
				return nil, fmt.Errorf("entdata: uninitialized hook (forgotten import entdata/runtime?)")
			}
			mut = msc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, msc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (msc *MovieSourceCreate) SaveX(ctx context.Context) *MovieSource {
	v, err := msc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (msc *MovieSourceCreate) Exec(ctx context.Context) error {
	_, err := msc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (msc *MovieSourceCreate) ExecX(ctx context.Context) {
	if err := msc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (msc *MovieSourceCreate) defaults() {
	if _, ok := msc.mutation.Status(); !ok {
		v := moviesource.DefaultStatus
		msc.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (msc *MovieSourceCreate) check() error {
	if _, ok := msc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`entdata: missing required field "MovieSource.status"`)}
	}
	if _, ok := msc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`entdata: missing required field "MovieSource.name"`)}
	}
	if _, ok := msc.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New(`entdata: missing required field "MovieSource.url"`)}
	}
	return nil
}

func (msc *MovieSourceCreate) sqlSave(ctx context.Context) (*MovieSource, error) {
	_node, _spec := msc.createSpec()
	if err := sqlgraph.CreateNode(ctx, msc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (msc *MovieSourceCreate) createSpec() (*MovieSource, *sqlgraph.CreateSpec) {
	var (
		_node = &MovieSource{config: msc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: moviesource.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: moviesource.FieldID,
			},
		}
	)
	if value, ok := msc.mutation.MangaSerieSearchURL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: moviesource.FieldMangaSerieSearchURL,
		})
		_node.MangaSerieSearchURL = &value
	}
	if value, ok := msc.mutation.MangaFilmSearchURL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: moviesource.FieldMangaFilmSearchURL,
		})
		_node.MangaFilmSearchURL = &value
	}
	if value, ok := msc.mutation.SerieSearchURL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: moviesource.FieldSerieSearchURL,
		})
		_node.SerieSearchURL = &value
	}
	if value, ok := msc.mutation.FilmSearchURL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: moviesource.FieldFilmSearchURL,
		})
		_node.FilmSearchURL = &value
	}
	if value, ok := msc.mutation.MangaSerieLatestURL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: moviesource.FieldMangaSerieLatestURL,
		})
		_node.MangaSerieLatestURL = &value
	}
	if value, ok := msc.mutation.MangaFilmLatestURL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: moviesource.FieldMangaFilmLatestURL,
		})
		_node.MangaFilmLatestURL = &value
	}
	if value, ok := msc.mutation.SerieLatestURL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: moviesource.FieldSerieLatestURL,
		})
		_node.SerieLatestURL = &value
	}
	if value, ok := msc.mutation.FilmLatestURL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: moviesource.FieldFilmLatestURL,
		})
		_node.FilmLatestURL = &value
	}
	if value, ok := msc.mutation.MangaSerieLatestPostSelector(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: moviesource.FieldMangaSerieLatestPostSelector,
		})
		_node.MangaSerieLatestPostSelector = value
	}
	if value, ok := msc.mutation.MangaFilmLatestPostSelector(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: moviesource.FieldMangaFilmLatestPostSelector,
		})
		_node.MangaFilmLatestPostSelector = value
	}
	if value, ok := msc.mutation.SerieLatestPostSelector(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: moviesource.FieldSerieLatestPostSelector,
		})
		_node.SerieLatestPostSelector = value
	}
	if value, ok := msc.mutation.FilmLatestPostSelector(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: moviesource.FieldFilmLatestPostSelector,
		})
		_node.FilmLatestPostSelector = value
	}
	if value, ok := msc.mutation.MangaSerieSearchPostSelector(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: moviesource.FieldMangaSerieSearchPostSelector,
		})
		_node.MangaSerieSearchPostSelector = value
	}
	if value, ok := msc.mutation.MangaFilmSearchPostSelector(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: moviesource.FieldMangaFilmSearchPostSelector,
		})
		_node.MangaFilmSearchPostSelector = value
	}
	if value, ok := msc.mutation.SerieSearchPostSelector(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: moviesource.FieldSerieSearchPostSelector,
		})
		_node.SerieSearchPostSelector = value
	}
	if value, ok := msc.mutation.FilmSearchPostSelector(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: moviesource.FieldFilmSearchPostSelector,
		})
		_node.FilmSearchPostSelector = value
	}
	if value, ok := msc.mutation.MangaSerieArticleSelector(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: moviesource.FieldMangaSerieArticleSelector,
		})
		_node.MangaSerieArticleSelector = value
	}
	if value, ok := msc.mutation.MangaFilmArticleSelector(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: moviesource.FieldMangaFilmArticleSelector,
		})
		_node.MangaFilmArticleSelector = value
	}
	if value, ok := msc.mutation.SerieArticleSelector(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: moviesource.FieldSerieArticleSelector,
		})
		_node.SerieArticleSelector = value
	}
	if value, ok := msc.mutation.FilmArticleSelector(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: moviesource.FieldFilmArticleSelector,
		})
		_node.FilmArticleSelector = value
	}
	if value, ok := msc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: moviesource.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := msc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: moviesource.FieldName,
		})
		_node.Name = value
	}
	if value, ok := msc.mutation.URL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: moviesource.FieldURL,
		})
		_node.URL = value
	}
	return _node, _spec
}

// MovieSourceCreateBulk is the builder for creating many MovieSource entities in bulk.
type MovieSourceCreateBulk struct {
	config
	builders []*MovieSourceCreate
}

// Save creates the MovieSource entities in the database.
func (mscb *MovieSourceCreateBulk) Save(ctx context.Context) ([]*MovieSource, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mscb.builders))
	nodes := make([]*MovieSource, len(mscb.builders))
	mutators := make([]Mutator, len(mscb.builders))
	for i := range mscb.builders {
		func(i int, root context.Context) {
			builder := mscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MovieSourceMutation)
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
					_, err = mutators[i+1].Mutate(root, mscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mscb *MovieSourceCreateBulk) SaveX(ctx context.Context) []*MovieSource {
	v, err := mscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mscb *MovieSourceCreateBulk) Exec(ctx context.Context) error {
	_, err := mscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mscb *MovieSourceCreateBulk) ExecX(ctx context.Context) {
	if err := mscb.Exec(ctx); err != nil {
		panic(err)
	}
}
