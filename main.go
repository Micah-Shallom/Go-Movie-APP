package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}


func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "456545", Title: "Movie One", Director: &Director{FirstName: "Shallom", LastName: "Micah Bawa"}})
	movies = append(movies, Movie{ID: "2", Isbn: "133243", Title: "Movie Two", Director: &Director{FirstName: "Theophilus", LastName: "Micah Bawa"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting Server on Port 8000")

	err := http.ListenAndServe(":8000", r)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("The Server has closed")
	}else if err != nil {
		fmt.Printf("The Server failed to start")
		log.Fatal(err)
	}

}