package graphql

import (
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
	paramString := ""
	for paramKey, paramValue := range params {
		paramString += ("&" + paramKey + "=" + paramValue)
	}
	return paramString
}

// GetMovies gets movies from yts movie api.
func GetMovies(rating float64, limit int) {
	// 	const {
	//     data: {
	//       data: { movies }
	//     }
	//   } = await axios(ListMoviesURL, {
	//     params: {
	//       limit,
	//       minimum_rating: rating
	//     }
	//   });
	//   return movies;

	params := makeParams(map[string]string{
		"limit":          string(limit),
		"minimum_rating": fmt.Sprintf("%f", rating),
	})
	resp, _ := http.Get(ListMoviesURL + params)
	robots, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	fmt.Printf("%s\n", robots)

}
