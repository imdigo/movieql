let movies = [
  {
    id: 0,
    name: "DoLim",
    score: 3
  },
  {
    id: 1,
    name: "HM",
    score: 4
  },
  {
    id: 2,
    name: "BS",
    score: 2
  },
  {
    id: 3,
    name: "YJ",
    score: 1
  },
  {
    id: 4,
    name: "MJ",
    score: 5
  },
];

export const getMovies = () => movies;

export const getById = id => {
  const filteredMovies = movies.filter(movie => movie.id === id);
  return filteredMovies[0];
};

export const deleteMovie = (id) => {
  const cleanedMovies = movies.filter(movie => movie.id !== id);
  if (movies.length > cleanedMovies.length) {
    movies = cleanedMovies;
    return true;
  } else {
    return false;
  }
};

export const addMovie = (name, score) => {
  const newMovie = {
    id: `${movies.length + 1}`,
    name,
    score,
  };
  movies.push(newMovie)
  return newMovie;
};