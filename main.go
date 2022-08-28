package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

var (
	typeFlag      string
	siteFlag      string
	keywordFlag   string
	pageSizeFlag  int
	daysSinceFlag int
	/*sectorFlag   string
	roundsFlag   int*/
	results []Article

	prNewsWireURLList = PRNewsWireURLs{
		Keyword: "https://www.prnewswire.com/search/all/?keyword=%s",
		/*BusinessTech:      "https://www.prnewswire.com/news-releases/business-technology-latest-news/business-technology-latest-news-list/?page=%d&pagesize=%d",
		GeneralBusiness:   "https://www.prnewswire.com/news-releases/general-business-latest-news/general-business-latest-news-list/?page=%d&pagesize=%d",
		FinancialServices: "https://www.prnewswire.com/news-releases/financial-services-latest-news/financial-services-latest-news-list/?page=%d&pagesize=%d",*/
	}

	googleNewsURLList = GoogleNewsURLs{
		Keyword: "https://news.google.com/search?q=%s%%20when%%3A%dd",
	}

	c = colly.NewCollector()
)

type PRNewsWireURLs struct {
	Keyword string
	/*BusinessTech
	GeneralBusiness
	FinancialServices*/
}

type GoogleNewsURLs struct {
	Keyword string
}

type Article struct {
	Headline string
	URL      string
}

func init() {
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Error: ", err)
	})

	flag.StringVar(&typeFlag, "t", "kw", "Type of search. Options: 'multi-keyword' 'mkw' 'keyword' 'kw'. Default is 'kw'")
	flag.StringVar(&siteFlag, "s", "google", "Site to search. Options: 'prnewswire' 'google'. Default is 'google'")
	flag.StringVar(&keywordFlag, "kw", "apple", "Keyword(s) you'd like to search for. If you have multiple, seperate them with the ',' identifier (NO spaces). For multi-word strings, seperate them as I just did (e.g. word1-word2-...) Not neccesary for a sector search. ")
	flag.IntVar(&pageSizeFlag, "ps", 100, "# of entries per page. Default is 100.")
	flag.IntVar(&daysSinceFlag, "d", 7, "How many days before present to return articles worth of. Default is 7, e.g. T-7 days worth of articles.")
	/*flag.StringVar(&sectorFlag, "s", "business-tech", "Sector for sector search. Options: 'business-tech' 'general-business' 'financial-services'")
	flag.IntVar(&roundsFlag, "r", 1, "# of pages to iterate through. Default is 1.")*/

	flag.Parse()
}

func main() {
	if typeFlag == "multi-keyword" || typeFlag == "mkw" {
		keywordList := strings.Split(keywordFlag, ",")

		for _, keyword := range keywordList {
			Scrape(keyword)
		}
	} else if typeFlag == "keyword" || typeFlag == "kw" {
		Scrape(keywordFlag)
	}

	for _, x := range results {
		fmt.Printf("Headline:  %s\nURL:  %s\n\n", x.Headline, x.URL)
	}
}
