package main

import (
	"log"
	"math/rand"
	"time"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

type Headline struct {
	Id int `json:"id"`
	Title string `json:"title"`
	SDesc string `json:"sdesc"`
	Url string `json:"url"`
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
		var temp []Headline
		c.OnHTML(".media__content", func (element *colly.HTMLElement) {
			headlineId := rand.Int()
			headlineTitle := strings.Trim(element.ChildText(".media__link"), "\n ")
			headlineSDesc := strings.Trim(element.ChildText(".media__summary"), "\n ")
			headlineUrl := element.ChildAttr(".media__link", "href")
			headline := Headline{
				Id: headlineId,
				Title: headlineTitle,
				SDesc: headlineSDesc,
				Url: headlineUrl,
			}
			temp = append(temp, headline)
			*data = temp
		})
		c.OnScraped(func (_ *colly.Response) {
			log.Println("Scraped data from website!")
		})
		c.Visit("https://www.bbc.com/")
	}
	webScraperLogic()
	for range time.NewTicker(time.Minute * 1).C { 
		webScraperLogic()
		// enc := json.NewEncoder(os.Stdout)
		// enc.SetIndent("", " ")
		// enc.Encode(*data)
	}

}
