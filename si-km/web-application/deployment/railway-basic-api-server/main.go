package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Quotes struct {
	Quote     string `json:"q"`
	Speaker   string `json:"a"`
	QuoteHTML string `json:"h"`
}

var client = &http.Client{}

func fetchRandomQuote() (*Quotes, error) {
	req, err := http.NewRequest("GET", "https://zenquotes.io/api/random", nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var quotes []Quotes
	if err := json.Unmarshal(responseData, &quotes); err != nil {
		return nil, err
	}

	return &quotes[0], nil
}

func main() {
	r := gin.Default()

	r.GET("/random", func(c *gin.Context) {
		quote, err := fetchRandomQuote()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"quote": quote.Quote})
	})

	r.GET("/", func(c *gin.Context) {
		quote, err := fetchRandomQuote()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(quote.QuoteHTML))
	})

	fmt.Println("Starting web server at localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
