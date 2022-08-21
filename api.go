package main

import (
	"log"
	"math/rand"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

type Headline struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Sdesc string `json:"sdesc"`
}

func main() {
	data := make([]Headline, 0)
	go webScraper(&data)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to News-headline scraper API! Visit '/headlines' route to get the responses.",
		})
	})
	r.GET("/headlines", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "List of recent news headlines -->",
			"data": data,
		})
	}) 
	r.Run() // listen and serve on 0.0.0.0:8080
}

func webScraper(data *[]Headline) { 
	
	webScraperLogic := func () {
		c := colly.NewCollector()
		c.OnHTML(".zox-s-title2", func (element *colly.HTMLElement) {
			headlineId := rand.Int()
			headlineTitle := element.Text
			headlineDesc := ""
			headline := Headline{
				Id: headlineId,
				Title: headlineTitle,
				Sdesc: headlineDesc,
			}
			*data = append(*data, headline)
		})
		c.OnScraped(func (r *colly.Response) {
			log.Println("Scraped data from website!")
		})
		c.Visit("https://boldtv.com/")
	}
	webScraperLogic()
	for range time.NewTicker(time.Minute * 1).C { 
		webScraperLogic()
		// enc := json.NewEncoder(os.Stdout)
		// enc.SetIndent("", " ")
		// enc.Encode(*data)
	}

}
