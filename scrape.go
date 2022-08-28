package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

// Scrape() executes different c.OnHTML() functions based on the set siteFlag.
func Scrape(keyword string) {
	if siteFlag == "google" || siteFlag == "g" {
		var result Article

		c.OnHTML("h3.ipQwMb.ekueJc.RD0gLb", func(e *colly.HTMLElement) {
			result.Headline = e.Text

			unformattedLink := e.ChildAttr("a[href]", "href")
			link := "https://news.google.com" + unformattedLink[1:]

			result.URL = link
			results = append(results, result)
		})

		c.Visit(fmt.Sprintf(googleNewsURLList.Keyword, keyword, daysSinceFlag))
	} else if siteFlag == "prnewswire" || siteFlag == "prn" {
		var result Article

		c.OnHTML("a.news-release", func(e *colly.HTMLElement) {
			result.Headline += e.Text

			unformattedLink := e.Attr("href")
			link := "https://www.prnewswire.com/" + unformattedLink

			result.URL = link
			results = append(results, result)
		})

		c.Visit(fmt.Sprintf(prNewsWireURLList.Keyword, keyword))
	} else {
		log.Fatalf("Incorrect Site Flag (--s). Should have been either 'google' or 'prnewswire'. Try again with a different value, or simply use the default 'prnewswire' by avoiding setting --s entirely.")
	}
}
