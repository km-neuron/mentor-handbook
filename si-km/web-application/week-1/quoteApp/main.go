package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Quotes struct {
	Quote     string `json:"q"`
	Speaker   string `json:"a"`
	QuoteHTML string `json:"h"`
}

func main() {
	r := gin.Default()

	// Handle GET request to /random endpoint
	r.GET("/random", func(c *gin.Context) {
		// Create HTTP client instance
		client := &http.Client{}

		// Create new request
		req, err := http.NewRequest("GET", "https://zenquotes.io/api/random", nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Send the request
		resp, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer resp.Body.Close()

		// Read the response body
		responseData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Unmarshal the response body into Quotes struct
		var quote []Quotes
		if err := json.Unmarshal(responseData, &quote); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"quote": quote[0].Quote})
	})

	// Serve HTML page with quote
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "<html><head><title>My quote</title></head><body>"+quote[0].QuoteHTML+"</body></html>")
	})

	// Run the server
	fmt.Println("Starting web server at localhost:8080")
	r.Run(":8080")
}
