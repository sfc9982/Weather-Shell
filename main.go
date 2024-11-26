package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	city := "230110" // Harbin HRB Haerbin Asia/Harbin

	r.GET("/weather", func(c *gin.Context) {
		weather := motd_weather(city)
		c.String(http.StatusOK, weather)
	})

	r.Run(":8848")
}
