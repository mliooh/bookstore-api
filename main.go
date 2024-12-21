package bookstore_api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "net/http"
)

//get books function

func getBooks(c *gin.Context) {
	var books []Book
	db.Find(&books)
	c.IndentedJSON(http.StatusOK, books)
}

func addBooks(c *gin.Context) {

}

func main() {
	loadDatabase()
	router := gin.Default()
	router.GET("/books", getBooks)
	router.Run("localhost:8080")
	//	get books function

}
