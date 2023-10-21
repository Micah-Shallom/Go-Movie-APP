package main

import (
	"fmt"
	"net/http"
	"errors"
	"log"
	"github.com/Micah-Shallom/pkg/routes"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

func main() {
	r := mux.NewRouter()

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
    methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
    origins := handlers.AllowedOrigins([]string{"*"})

	routes.MoviesRouter(r)
	http.Handle("/", r)

	fmt.Printf("Starting Server on Port 8000\n")
	err := http.ListenAndServe(":8000", handlers.CORS(headers,methods, origins)(r))
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("The Server has closed")
	} else if err != nil {
		fmt.Printf("The Server failed to start")
		log.Fatal(err)
	}
}
