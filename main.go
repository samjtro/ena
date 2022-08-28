package main

import (
	"bytes"
	"log"
	"os"
	"strconv"
	"strings"
	"text/template"

	"github.com/gocolly/colly"
	"github.com/joho/godotenv"
)

type PRNewsWireURLs struct {
	Keyword           string
	BusinessTech      string
	GeneralBusiness   string
	FinancialServices string
}

type GoogleNewsURLs struct {
	Keyword string
}

type Article struct {
	Keyword     string
	HeadlineURL map[string]string
}

var (
	typeFlag         string
	siteFlag         string
	keywordFlag      string
	fromFlag         string
	fromPasswordFlag string
	toFlag           string
	pageSizeFlag     int
	daysSinceFlag    int
	utcDiffFlag      string
	results          []Article
	tpl              bytes.Buffer

	prNewsWireURLList = PRNewsWireURLs{
		Keyword:           "https://www.prnewswire.com/search/all/?keyword=%s",
		BusinessTech:      "https://www.prnewswire.com/news-releases/business-technology-latest-news/business-technology-latest-news-list/?page=%d&pagesize=%d",
		GeneralBusiness:   "https://www.prnewswire.com/news-releases/general-business-latest-news/general-business-latest-news-list/?page=%d&pagesize=%d",
		FinancialServices: "https://www.prnewswire.com/news-releases/financial-services-latest-news/financial-services-latest-news-list/?page=%d&pagesize=%d",
	}

	googleNewsURLList = GoogleNewsURLs{
		Keyword: "https://news.google.com/search?q=%s%%20when%%3A%dd",
	}

	c = colly.NewCollector()
)

func init() {
	err := godotenv.Load("config.env")

	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	typeFlag = os.Getenv("Type")
	siteFlag = os.Getenv("site")
	keywordFlag = os.Getenv("keyword")
	utcDiffFlag = os.Getenv("utcdiff")
	fromFlag = os.Getenv("from")
	fromPasswordFlag = os.Getenv("frompassword")
	toFlag = os.Getenv("to")
	pageSizeFlag, err = strconv.Atoi(os.Getenv("pagesize"))

	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	daysSinceFlag, err = strconv.Atoi(os.Getenv("dayssince"))

	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Error: ", err)
	})
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

	tmpl := template.Must(template.ParseFiles("template.html"))

	if err := tmpl.Execute(&tpl, results); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	SendEmail(tpl.String())
}
