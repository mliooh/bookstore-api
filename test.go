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
	ID       string `json:"id" gorm:"primaryKey"`
	Author   string `json:"author"`
	Title    string `json:"title"`
	Genre    string `json:"genre"`
	Quantity int    `json:"quantity"`
}

var db *gorm.DB

// Check if books already exist
func booksExist() bool {
	var count int64
	db.Model(&Book{}).Count(&count)
	return count > 0
}

func loadDatabase() {
	var err error
	db, err = gorm.Open(sqlite.Open("bookstore.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// Create tables
	db.AutoMigrate(&Book{})
	fmt.Println("Database and tables created")
}

// gofakeit data creation

func createBooks() {
	if booksExist() {
		fmt.Println("Books already exist, skipping data seeding.")
		return
	}

	// Generate sample books
	var genres = []string{"Fiction", "Non-fiction", "Fantasy", "Sci-Fi", "Romance", "Thriller"}
	for i := 0; i < 200; i++ {
		book := Book{
			Title:    gofakeit.BookTitle(),
			Author:   gofakeit.Name(),
			Genre:    genres[rand.Intn(len(genres))],
			Quantity: rand.Intn(50) + 1,
		}
		db.Create(&book)
	}
	fmt.Println("Books created successfully.")
}

// fetch books

func getBooks(c *gin.Context) {
	var books []Book
	db.Find(&books)
	c.JSON(http.StatusOK, books)
}

// add new books

func addBook(c *gin.Context) {
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&book)
	c.JSON(http.StatusCreated, book)
}

// delete books

func deleteBook(c *gin.Context) {
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Delete(&book)
	c.JSON(http.StatusOK, book)
}

func main() {
	// Initialize database and seed data
	loadDatabase()
	createBooks()

	// Launch API
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", addBook)
	router.DELETE("/books", deleteBook)
	fmt.Println("Starting server at http://localhost:8081")
	err := router.Run("localhost:8081")
	if err != nil {
		return
	}
}
