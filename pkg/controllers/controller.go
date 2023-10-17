package controllers

import (
	"net/http"
)

func GetMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

}

	
func GetMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

}
func CreateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

}
func UpdateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

}
func DeleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

}