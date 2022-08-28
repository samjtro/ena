package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/gocolly/colly"
	"github.com/joho/godotenv"
	gomail "gopkg.in/gomail.v2"
)

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
	Headline string
	URL      string
}

func init() {
	err := godotenv.Load("config.env")

	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	typeFlag = os.Getenv("type")
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

	msg := gomail.NewMessage()
	msg.SetHeader("From", fromFlag)
	msg.SetHeader("To", toFlag)
	msg.SetHeader("Subject", fmt.Sprintf("[sn] Keyword: %s, Site: %s - %s", keywordFlag, siteFlag, Now(time.Now())))

	tmpl := template.Must(template.ParseFiles("template.html"))

	if err := tmpl.Execute(&tpl, results); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	msgString := tpl.String()
	msg.SetBody("text/html", msgString)
	handler := gomail.NewDialer("smtp.gmail.com", 587, fromFlag, fromPasswordFlag)

	if err := handler.DialAndSend(msg); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
