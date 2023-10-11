package main

import (
	"fmt"
	"net/http"

	"github.com/Micah-Shallom/pkg/config"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	http.Handle("/", r)
	config.Connect()
	d := config.GetDB()
	fmt.Println(d)

	fmt.Printf("Starting Server on Port 8000")
	// err := http.ListenAndServe(":8000", r)
	// if errors.Is(err, http.ErrServerClosed) {
	// 	fmt.Printf("The Server has closed")
	// } else if err != nil {
	// 	fmt.Printf("The Server failed to start")
	// 	log.Fatal(err)
	// }
}