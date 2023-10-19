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

func GetMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	movies := models.GetAllMovies()
	body, err := json.Marshal(movies)
	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func GetMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	ID := params["id"]
	movieID, err := strconv.Atoi(ID); 
	if err != nil {
		fmt.Println(err)
	}
	movie, _ := models.GetMovieByID(movieID)
	body , _ := json.Marshal(movie)
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func CreateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	CreateMovie := &models.Movie{}
	utils.ParseBody(r, CreateMovie)
	movie := CreateMovie.CreateMovie()
	body, err := json.Marshal(movie)
	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request){
	updatemovie := &models.Movie{}
	w.Header().Set("Content-Type", "application/json")
	utils.ParseBody(r, updatemovie)
	params := mux.Vars(r)
	id := params["id"]
	movieID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error converting str to integer", err)
	}

	movie, db := models.GetMovieByID(movieID)
	if updatemovie.Isbn != "" {
		movie.Isbn = updatemovie.Isbn
	}
	if updatemovie.Title != "" {
		movie.Title = updatemovie.Title
	}
	if updatemovie.Director.FirstName != "" {
		movie.Director.FirstName = updatemovie.Director.FirstName
	}
	if updatemovie.Director.LastName != "" {
		movie.Director.LastName = updatemovie.Director.LastName
	}
	db.Save(&movie)
	body, _ := json.Marshal(movie)
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
}