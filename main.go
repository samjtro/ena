package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

var (
	typeFlag     string
	siteFlag     string
	keywordFlag  string
	pageSizeFlag int
	/*sectorFlag   string
	roundsFlag   int*/
	results []string

	prNewsWireURLList = PRNewsWireURLs{
		Keyword: "https://www.prnewswire.com/search/all/?keyword=%s",
		/*BusinessTech:      "https://www.prnewswire.com/news-releases/business-technology-latest-news/business-technology-latest-news-list/?page=%d&pagesize=%d",
		GeneralBusiness:   "https://www.prnewswire.com/news-releases/general-business-latest-news/general-business-latest-news-list/?page=%d&pagesize=%d",
		FinancialServices: "https://www.prnewswire.com/news-releases/financial-services-latest-news/financial-services-latest-news-list/?page=%d&pagesize=%d",*/
	}

	googleNewsURLList = GoogleNewsURLs{
		Keyword: "https://news.google.com/search?q=%s",
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

func init() {
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

	flag.StringVar(&typeFlag, "t", "kw", "Type of search. Options: 'keyword' 'kw'. Default is 'kw'")
	flag.StringVar(&siteFlag, "s", "google", "Site to search. Options: 'prnewswire' 'google'. Default is 'prnewswire'")
	flag.StringVar(&keywordFlag, "kw", "apple-inc", "Keyword(s) you'd like to search for. If you have multiple, seperate them with the ',' identifier (NO spaces). For multi-word strings, seperate them as I just did (e.g. word1-word2-...) Not neccesary for a sector search. ")
	flag.IntVar(&pageSizeFlag, "psize", 100, "# of entries per page. Default is 100.")
	/*flag.StringVar(&sectorFlag, "s", "business-tech", "Sector for sector search. Options: 'business-tech' 'general-business' 'financial-services'")
	flag.IntVar(&roundsFlag, "r", 1, "# of pages to iterate through. Default is 1.")*/

	flag.Parse()
}

func main() {
	if typeFlag == "multi-keyword" || typeFlag == "mkw" {
		keywordList := strings.Split(keywordFlag, ",")

		for _, keyword := range keywordList {
			fmt.Printf("Searching for Keyword {%s}\n", keyword)
			Scrape(keyword)
		}
	} else if typeFlag == "keyword" || typeFlag == "kw" {
		Scrape(keywordFlag)
	}

	for i := 0; i < len(results); i += 2 {
		y := strings.Split(results[i], ",/|")
		fmt.Printf("Headline:  %s\nURL:  %s\n\n", y[0+i], y[1+i])
	}
}

func Scrape(keyword string) {
	if siteFlag == "google" || siteFlag == "g" {
		var result string

		c.OnHTML("h3.ipQwMb.ekueJc.RD0gLb", func(e *colly.HTMLElement) {
			result += e.Text // headline
			unformattedLink := e.ChildAttr("a[href]", "href")
			link := "https://news.google.com" + unformattedLink[1:]
			result += ",/|" + link + ",/|" // add the link with ,/| in between it and the headline for later Splitting
			results = append(results, result)
		})

		c.Visit(fmt.Sprintf(googleNewsURLList.Keyword, keyword))
	} else if siteFlag == "prnewswire" || siteFlag == "prn" {
		var result string

		c.OnHTML("a.news-release", func(e *colly.HTMLElement) {
			result += e.Text // headline
			unformattedLink := e.Attr("href")
			link := "https://www.prnewswire.com/" + unformattedLink
			result += ",/|" + link + ",/|" // add the link with ,/| in between it and the headline for later Splitting
			results = append(results, result)
		})

		c.Visit(fmt.Sprintf(prNewsWireURLList.Keyword, keyword))
	} else {
		log.Fatalf("Incorrect Site Flag (--s). Should have been either 'google' or 'prnewswire'. Try again with a different value, or simply use the default 'prnewswire' by avoiding setting --s entirely.")
	}
}
