package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Micah-Shallom/pkg/models"
	"github.com/Micah-Shallom/pkg/utils"
	"github.com/gorilla/mux"
)

func GetMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Action: GetMovies")
	w.Header().Set("Content-Type", "application/json")
	movies := models.GetAllMovies()
	body, err := json.Marshal(movies)
	if err != nil {
		fmt.Println("Error marshaling movies:", err)
		return
	}
	fmt.Println("Successfully got all movies")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Action: GetMovie")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	ID := params["id"]
	movieID, err := strconv.Atoi(ID)
	if err != nil {
		fmt.Println("Error parsing ID:", err)
		return
	}
	movie, _ := models.GetMovieByID(movieID)
	body, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("Error marshaling movie:", err)
		return
	}
	fmt.Println("Successfully got movie by ID")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Action: CreateMovie")
	w.Header().Set("Content-Type", "application/json")
	CreateMovie := &models.Movie{}
	utils.ParseBody(r, CreateMovie)
	movie := CreateMovie.CreateMovie()
	body, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("Error marshaling new movie:", err)
		return
	}
	fmt.Println("Successfully created new movie")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Action: UpdateMovie")
	w.Header().Set("Content-Type", "application/json")
	newMovie := &models.Movie{}
	utils.ParseBody(r, newMovie)
	params := mux.Vars(r)
	id := params["id"]
	movieID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error parsing ID:", err)
		return
	}
	movie, db := models.GetMovieByID(movieID)

	if newMovie.Isbn != "" {
		movie.Isbn = newMovie.Isbn
	}
	if newMovie.Title != "" {
		movie.Title = newMovie.Title
	}
	if newMovie.Director.FirstName != "" {
		movie.Director.FirstName = newMovie.Director.FirstName
	}
	if newMovie.Director.LastName != "" {
		movie.Director.LastName = newMovie.Director.LastName
	}
	db.Save(&movie)
	body, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("Error marshaling updated movie:", err)
		return
	}
	fmt.Println("Successfully updated movie")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Action: DeleteMovie")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	movieID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error parsing ID:", err)
		return
	}
	movie, _ := models.DeleteMovie(movieID)
	body, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("Error marshaling deleted movie:", err)
		return
	}
	fmt.Println("Successfully deleted movie")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
