package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Quote struct {
	ID     int    `json:"id"`
	Author string `json:"author"`
	Text   string `json:"text"`
}

func main() {
	r := gin.Default()

	// Routes
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Quote App!",
		})
	})

	r.GET("/quote", getRandomQuote)
	r.POST("/quote", insertQuote)

	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func getRandomQuote(c *gin.Context) {
	// Retrieve data from the API or database
	quote := fetchQuoteFromAPI()

	c.JSON(http.StatusOK, quote)
}

func insertQuote(c *gin.Context) {
	// Get quote data from request
	var quote Quote
	err := c.ShouldBindJSON(&quote)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Store the quote in the database
	err = saveQuoteToDatabase(quote)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Quote inserted successfully!",
	})
}

func fetchQuoteFromAPI() Quote {
	// Perform API request to get a random quote
	// Implement your code here to fetch the quote from the API
	// Example:
	// resp, err := http.Get("https://animechan.vercel.app/api/random")
	// ...

	// Parse the response and extract the quote data
	// Example:
	// var quote Quote
	// err = json.NewDecoder(resp.Body).Decode(&quote)
	// ...

	// Return the quote
	// return quote

	// Placeholder implementation (replace with actual API request)
	return Quote{
		ID:     1,
		Author: "Author Name",
		Text:   "This is a random quote.",
	}
}

func saveQuoteToDatabase(quote Quote) error {
	// Connect to the database
	db, err := sql.Open("postgres", "postgres://username:password@localhost/quote_db?sslmode=disable")
	if err != nil {
		return err
	}
	defer db.Close()

	// Insert the quote into the database
	_, err = db.Exec("INSERT INTO quotes (author, text) VALUES ($1, $2)", quote.Author, quote.Text)
	if err != nil {
		return err
	}

	return nil
}
