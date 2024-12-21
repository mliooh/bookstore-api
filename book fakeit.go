package bookstore_api

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math/rand"
)

// define book model

type Book struct {
	ID       uint   `gorm:"primary_key"`
	Title    string `json:"title" validate:"required"`
	Author   string `json:"author" validate:"required"`
	Genre    string `json:"genre"`
	Quantity uint   `json:"quantity"`
}

// initialise database
var db *gorm.DB

func loadDatabase() {
	var db, err = gorm.Open(sqlite.Open("bookstore.db"), &gorm.Config{})
	if err != nil {
		fmt.Printf("failed to connect database")
	}

	// Create tables
	db.AutoMigrate(&Book{})

	//	gofakeit create data

	var genres = []string{"Fiction", "Non-fiction", "Fantasy", "Sci-Fi", "Romance", "Thriller"}

	for i := 0; i < 200; i++ {
		book := Book{

			Title:    gofakeit.BookTitle(),
			Author:   gofakeit.Name(),
			Genre:    genres[rand.Intn(len(genres))],
			Quantity: uint(rand.Intn(50) + 1),
		}
		db.Create(&book)
	}
	println(fmt.Sprintf("Bookstore loaded %d books", 200))
}
