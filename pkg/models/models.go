package models

import (
	"fmt"
	"github.com/Micah-Shallom/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Movie struct {
	gorm.Model
	// ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
	DirectorID uint		`json:"director_id"` //foreign key
}
type Director struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Director{})
	db.AutoMigrate(&Movie{})
	db.Model(&Movie{}).AddForeignKey("director_id", "directors(id)", "CASCADE", "CASCADE")
	fmt.Println("Database Connection Successful")
}


func GetAllMovies() []Movie {
	var movies []Movie
	db.Preload("Director").Find(&movies)
	return movies
}

func GetMovieByID(ID int) (*Movie, *gorm.DB) {
	var movie Movie
	db := db.Where("ID=?", ID).Find(&movie)
	return &movie, db
}


func (m *Movie) CreateMovie () *Movie {
	db.NewRecord(m)
	db.Create(&m)
	return m
}

func DeleteMovie(ID int) Movie {
	var movie Movie
	db.Where("ID=?", ID).Delete(movie)
	return movie
}