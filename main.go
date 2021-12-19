package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	crawl()
}

func crawl() {
	c := colly.NewCollector(
		colly.AllowedDomains("imdb.com", "www.imdb.com"),
	)
	infoCollector := c.Clone()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL.String())
	})

	infoCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting Profile URL: ", r.URL.String())
	})

	c.Visit("https://www.imdb.com/search/name/?birth_monthday=12-20")
}
