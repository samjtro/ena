package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
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

type HeadlineURL struct {
	Headline string
	URL      string
}

type Articles struct {
	Keyword      string
	HeadlineURLs []HeadlineURL
}

var (
	siteFlag         string
	keywordFlag      string
	fromFlag         string
	fromPasswordFlag string
	toFlag           string
	daysSinceFlag    int
	utcDiffFlag      string
	results          []Articles
	emailContents    bytes.Buffer
	tmpl             *template.Template

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
	currentUser, err := user.Current()

	if err != nil {
		log.Fatalf(err.Error())
	}

	tmpl = template.Must(template.ParseFiles(fmt.Sprintf("/home/%s/sn/template.html", currentUser.Username)))
	err = godotenv.Load(fmt.Sprintf("/home/%s/config.env", currentUser.Username))

	if err != nil {
		log.Fatal(err)
	}

	siteFlag = os.Getenv("site")
	keywordFlag = os.Getenv("keyword")
	utcDiffFlag = os.Getenv("utcdiff")
	fromFlag = os.Getenv("from")
	fromPasswordFlag = os.Getenv("frompassword")
	toFlag = os.Getenv("to")

	if err != nil {
		log.Fatal(err)
	}

	daysSinceFlag, err = strconv.Atoi(os.Getenv("dayssince"))

	if err != nil {
		log.Fatal(err)
	}

	c.OnError(func(_ *colly.Response, err error) {
		log.Fatal(err)
	})
}

func main() {
	keywordList := strings.Split(keywordFlag, ",")

	for _, keyword := range keywordList {
		Scrape(keyword)
	}

	if err := tmpl.Execute(&emailContents, results); err != nil {
		log.Fatal(err)
	}

	path := filepath.Join("var", "www", "html", "index.html")
	file, err := os.Create(path)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	if err := tmpl.Execute(&emailContents, results); err != nil {
		log.Fatal(err)
	}

	SendEmail(emailContents.String())
}
