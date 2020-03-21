package graphql

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// BaseURL of yts
const BaseURL = "https://yts-proxy.now.sh/"

// ListMoviesURL is api url of listing movie on yts
const ListMoviesURL = BaseURL + "list_movies.json"

// MovieDetailsURL is api url of movie details on yts
const MovieDetailsURL = BaseURL + "movie_details.json"

// MovieSuggestionsURL is api url of suggestion list on yts
const MovieSuggestionsURL = BaseURL + "movie_suggestions.json"

func makeParams(params map[string]string) string {
	paramString := "?"
	for paramKey, paramValue := range params {
		if paramValue != "" {
			paramString += ("&" + paramKey + "=" + paramValue)
		}
	}
	return paramString
}

// ID .

// Movie Data .
type Movie struct {
	ID               int32     `json:"id"`
	Title            string    `json:"title"`
	Rating           float64   `json:"rating"`
	DescriptionIntro string    `json:"summary"`
	Language         string    `json:"language"`
	MediumCoverImage string    `json:"medium_cover_image"`
	Genres           []*string `json:"genres"`
}

// MoviesData .
type MoviesData struct {
	Data struct {
		Movies []*Movie `json:"movies"`
	} `json:"data"`
}

// MovieData .
type MovieData struct {
	Data struct {
		Movie *Movie `json:"movie"`
	} `json:"data"`
}

// GetMovies gets movies from yts movie api.
func GetMovies(rating *float64, limit *int32) []*Movie {
	params := ""
	if rating != nil || limit != nil {
		params = makeParams(map[string]string{
			"limit":          fmt.Sprintf("%d", *limit),
			"minimum_rating": fmt.Sprintf("%.1f", *rating),
		})
	}

	resp, err := http.Get(ListMoviesURL + params)
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	var d MoviesData
	err = json.Unmarshal(body, &d)
	if err != nil {
		panic(err)
	}
	return d.Data.Movies
}

// GetMovie from yts using ID
func GetMovie(id int32) *Movie {
	params := makeParams(map[string]string{
		"movie_id": fmt.Sprintf("%d", id),
	})
	resp, err := http.Get(MovieDetailsURL + params)
	fmt.Println(MovieDetailsURL + params)
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	var d MovieData
	err = json.Unmarshal(body, &d)
	if err != nil {
		panic(err)
	}
	return d.Data.Movie
}

// GetSuggestions from yts using ID
func GetSuggestions(id int32) []*Movie {
	params := makeParams(map[string]string{
		"movie_id": fmt.Sprintf("%d", id),
	})
	resp, err := http.Get(MovieSuggestionsURL + params)
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	var d MoviesData
	err = json.Unmarshal(body, &d)
	if err != nil {
		panic(err)
	}
	return d.Data.Movies
}
