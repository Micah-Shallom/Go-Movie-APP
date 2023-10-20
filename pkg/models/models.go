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
	Director *Director `json:"director" gorm:"foreignKey:DirectorID; constraint:OnDelete:CASCADE" `
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

// func DeleteMovie(ID int) Movie {
// 	var movie Movie
// 	db.Where("ID=?", ID).Delete(movie)
// 	fmt.Println(movie)
// 	return movie
// }

func DeleteMovie(ID int) (*Movie, error) {
	var movie Movie
	if err := db.Preload("Director").Where("ID=?", ID).First(&movie).Error; err != nil {
		return nil, err
	}
	if err := db.Delete(&movie).Error; err != nil {
		return nil, err
	}
	fmt.Println(movie)
	return &movie, nil
}
