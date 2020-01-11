package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

const (
	OpenWeatherApiBaseURL = "https://api.openweathermap.org/data/2.5/"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Infof("Successfully read .env file")

	APIKey := os.Getenv("OPENWEATHER_API_KEY")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {

		url := OpenWeatherApiBaseURL + `weather?q=London&appid=` + APIKey
		resp, err := http.Get(url)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to reach external API",
			})
			return
		}
		defer resp.Body.Close()

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to read from external API's response",
			})
			return
		}

		c.JSON(200, gin.H{
			"data": string(data),
		})
	})

	r.Run("localhost:8080")
}
