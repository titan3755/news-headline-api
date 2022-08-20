package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"math/rand"
	"time"
)

func main() {
	data := ""
	go webScraper(&data)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": data,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

func webScraper(data *string) { 
	for range time.NewTicker(time.Second * 1).C { 
		*data = strconv.Itoa(rand.Int())
	}
}
