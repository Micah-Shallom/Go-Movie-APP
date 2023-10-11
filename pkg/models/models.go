package models

import "github.com/Micah-Shallom/pkg/config"

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func init() {
	config.Connect()
	db := config.GetDB()

	_ = db
}