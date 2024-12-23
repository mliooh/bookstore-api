package main

//
//import (
//	"fmt"
//	"github.com/brianvoe/gofakeit/v6"
//	"gorm.io/driver/sqlite"
//	"gorm.io/gorm"
//	"math/rand"
//)
//
//// initialise database
//var db *gorm.DB
//
//func loadDatabase() {
//	var err error
//	db, err = gorm.Open(sqlite.Open("bookstore.db"), &gorm.Config{})
//	if err != nil {
//		fmt.Printf("failed to connect database")
//	}
//
//	// Create tables
//	db.AutoMigrate(&Book{})
//	println("Database and tables created")
//
//}
//
//func createBooks() {
//	//	gofakeit create data
//
//	var genres = []string{"Fiction", "Non-fiction", "Fantasy", "Sci-Fi", "Romance", "Thriller"}
//
//	for i := 0; i < 200; i++ {
//		book := Book{
//
//			Title:    gofakeit.BookTitle(),
//			Author:   gofakeit.Name(),
//			Genre:    genres[rand.Intn(len(genres))],
//			Quantity: uint(rand.Intn(50) + 1),
//		}
//		db.Create(&book)
//	}
//	println(fmt.Sprintf("Bookstore loaded %d books", 200))
//}
