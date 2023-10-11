package routes

import (
	"github.com/Micah-Shallom/pkg/controllers"
	"github.com/gorilla/mux"
)

var MoviesRouter = func(router *mux.Router) {
	// movies = append(movies, Movie{ID: "1", Isbn: "456545", Title: "Movie One", Director: &Director{FirstName: "Shallom", LastName: "Micah Bawa"}})
	// movies = append(movies, Movie{ID: "2", Isbn: "133243", Title: "Movie Two", Director: &Director{FirstName: "Theophilus", LastName: "Micah Bawa"}})

	router.HandleFunc("/movies", controllers.GetMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", controllers.GetMovie).Methods("GET")
	router.HandleFunc("/movies", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", controllers.UpdateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", controllers.DeleteMovie).Methods("DELETE")
}