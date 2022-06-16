package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"
	"time"
	"yola/internal"
	"yola/internal/entdata/moviesource"
	"yola/internal/entdata/schema"
	"yola/internal/source"
)

func (h *Handler) FilmSearchPost(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	params := internal.Params(r.Form)
	query, _ := params.String("query")
	page, _ := params.Int("page")

	movieSources := h.MovieSource.Query().Where(moviesource.Status(true)).AllX(ctx)
	sources := source.ParseMovieSources[source.FilmSource](movieSources)

	moviePosts := make([]schema.MoviePost, 0)
	group := new(sync.WaitGroup)
	for _, s := range sources {
		group.Add(1)
		go func(source source.FilmSource) {
			posts := source.FilmSearchPost(ctx, query, page)
			moviePosts = append(moviePosts, posts...)
			group.Done()
		}(s)
	}
	group.Wait()

	response := internal.Shuffle(moviePosts)
	json.NewEncoder(w).Encode(response)
}
