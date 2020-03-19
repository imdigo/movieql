package graphql

import (
	"strconv"

	"github.com/graph-gophers/graphql-go"
)

// RootResolver .
type RootResolver struct{}

// MoviesArgs is arg type of Movies query resolver
type MoviesArgs struct {
	Rating float64
	Limit  int32
}

// MovieArgs is arg type of Movie query resolver
type MovieArgs struct {
	ID graphql.ID
}

// Movies query resolver
func (*RootResolver) Movies(args MoviesArgs) []*MovieResolver {
	movies := GetMovies(args.Rating, args.Limit)
	movieResolvers := make([]*MovieResolver, 0)
	for _, m := range movies {
		movieResolvers = append(movieResolvers, &MovieResolver{m})
	}
	return movieResolvers
}

// Movie query resolver
func (*RootResolver) Movie(args MovieArgs) *MovieResolver {
	movie := GetMovie(args.ID)
	movieResolver := MovieResolver{movie}
	return &movieResolver
}

// MovieResolver struct
type MovieResolver struct {
	Movie *Movie
}

// ID of Movie
func (r *MovieResolver) ID() graphql.ID {
	id := strconv.Itoa(r.Movie.ID) // first, convert to string
	return graphql.ID(id)          // coerce the string to the graphql.ID type
}

// Title of Movie
func (r *MovieResolver) Title() string {
	return r.Movie.Title
}

// Rating of Movie
func (r *MovieResolver) Rating() float64 {
	return r.Movie.Rating
}

// DescriptionIntro of Movie
func (r *MovieResolver) DescriptionIntro() string {
	return r.Movie.DescriptionIntro
}

// Language of Movie
func (r *MovieResolver) Language() *string {
	return &r.Movie.Language
}

// MediumCoverImage of Movie
func (r *MovieResolver) MediumCoverImage() *string {
	return &r.Movie.MediumCoverImage
}

// Genres of Movie
func (r *MovieResolver) Genres() *[]*string {
	return &r.Movie.Genres
}
