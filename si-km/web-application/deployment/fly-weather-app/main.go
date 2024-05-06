package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WeatherResponse struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Name string `json:"name"`
}

func main() {
	router := gin.Default()

	router.GET("/weather", func(c *gin.Context) {
		response, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=London&units=metric&appid=your_api_key")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		data, _ := ioutil.ReadAll(response.Body)

		var weatherResponse WeatherResponse
		json.Unmarshal(data, &weatherResponse)

		c.JSON(http.StatusOK, gin.H{"city": weatherResponse.Name, "temperature": weatherResponse.Main.Temp})
	})

	router.Run(":8080")
}
