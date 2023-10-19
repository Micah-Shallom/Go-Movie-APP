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
	body , jsonerr := json.Marshal(movie)
	if jsonerr != nil {
		fmt.Println(err)
	}
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
	newmovie := &models.Movie{}
	w.Header().Set("Content-Type", "application/json")
	utils.ParseBody(r, newmovie)
	params := mux.Vars(r)
	id := params["id"]
	movieID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error while parsing", err)
	}

	movie, db := models.GetMovieByID(movieID)
	if newmovie.Isbn != "" {
		movie.Isbn = newmovie.Isbn
	}

	if newmovie.Title != "" {
		movie.Title = newmovie.Title
	}

	if newmovie.Director.FirstName != "" {
		movie.Director.FirstName = newmovie.Director.FirstName
	}

	if newmovie.Director.LastName != "" {
		movie.Director.LastName = newmovie.Director.LastName
	}

	db.Save(&movie)
	body, jsonerr := json.Marshal(movie)

	if jsonerr != nil {
		fmt.Println(jsonerr)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	movieID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error whle parsing", err)
	}
	movie := models.DeleteBook(movieID)
	body, jsonerr := json.Marshal(movie)
	if jsonerr != nil {
		fmt.Println(jsonerr)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}