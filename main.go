package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
	"math/rand"
	"net/http"
)

type Book struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Author   string `json:"author"`
	Title    string `json:"title"`
	Genre    string `json:"genre"`
	Quantity int    `json:"quantity"`
}

// initialise database
var db *gorm.DB

func loadDatabase() {

	var err error
	db, err = gorm.Open(sqlite.Open("file:bookstore.db?cache=shared&_fk=1"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Book{})
	fmt.Println("Database and tables created")

	/*var db, err = gorm.Open(sqlite.Open("bookstore.db"), &gorm.Config{})

	if err != nil {
		fmt.Printf("failed to connect database")
	}

	// Create tables
	db.AutoMigrate(&Book{})
	println("Database and tables created")*/

}

func createBooks() {
	//	gofakeit create data

	var count int64
	db.Model(&Book{}).Count(&count)
	if count > 0 {
		println("Book already exists")
		return
	}
	var genres = []string{"Fiction", "Non-fiction", "Fantasy", "Sci-Fi", "Romance", "Thriller"}

	for i := 0; i < 200; i++ {
		book := Book{

			Title:    gofakeit.BookTitle(),
			Author:   gofakeit.Name(),
			Genre:    genres[rand.Intn(len(genres))],
			Quantity: int(uint(rand.Intn(50) + 1)),
		}
		db.Create(&book)
	}
	println(fmt.Sprintf("Bookstore loaded %d books", 200))
}

//get books function

func getBooks(c *gin.Context) {
	var books []Book
	db.Find(&books)
	c.IndentedJSON(http.StatusOK, books)
}

func addBooks(c *gin.Context) {
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&book)
	c.IndentedJSON(http.StatusCreated, book)

}

func deleteBooks(c *gin.Context) {
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Delete(&book)
	c.IndentedJSON(http.StatusOK, book)
}

func main() {
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	loadDatabase()
	createBooks()
	router.GET("/books", getBooks)
	//router.GET("/books/:id", getBooks)
	router.POST("/books", addBooks)
	router.DELETE("/books/:id", deleteBooks)
	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
	//	get books function

}
