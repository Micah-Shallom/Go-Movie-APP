package main

import (
	"fmt"
	"net/http"
	"errors"
	"log"
	"github.com/Micah-Shallom/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	routes.MoviesRouter(r)
	http.Handle("/", r)

	fmt.Printf("Starting Server on Port 8000")
	err := http.ListenAndServe(":8000", r)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("The Server has closed")
	} else if err != nil {
		fmt.Printf("The Server failed to start")
		log.Fatal(err)
	}
}
