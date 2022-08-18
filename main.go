package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func main() {
	urlList := []string{"https://www.prnewswire.com/news-releases/business-technology-latest-news/business-technology-latest-news-list/", "https://www.prnewswire.com/news-releases/general-business-latest-news/general-business-latest-news-list/", "https://www.prnewswire.com/news-releases/financial-services-latest-news/financial-services-latest-news-list/"}

	for _, x := range urlList {
		Scrape(x)
	}
}

func Scrape(url string) {
	c := colly.NewCollector()
	var headline, body, author, thumbnail string

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Error: ", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Connection Established")
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished!")
	})

	c.OnHTML("div.col-sm-8.col-lg-9.pull-left.card", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	c.Visit(url)
}
