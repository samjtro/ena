package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/user"
	"strconv"
	"strings"
	"text/template"

	badger "github.com/dgraph-io/badger/v3"
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
	sendEmail        bool
	typeFlag         string
	siteFlag         string
	keywordFlag      string
	fromFlag         string
	fromPasswordFlag string
	toFlag           string
	pageSizeFlag     int
	daysSinceFlag    int
	utcDiffFlag      string
	results          []Articles
	resultsForEmail  []Articles
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

	keywordList = strings.Split(keywordFlag, ",")
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

	typeFlag = os.Getenv("Type")
	siteFlag = os.Getenv("site")
	keywordFlag = os.Getenv("keyword")
	utcDiffFlag = os.Getenv("utcdiff")
	fromFlag = os.Getenv("from")
	fromPasswordFlag = os.Getenv("frompassword")
	toFlag = os.Getenv("to")
	pageSizeFlag, err = strconv.Atoi(os.Getenv("pagesize"))

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
	db, err := badger.Open(badger.DefaultOptions("/tmp/badger"))

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	for _, keyword := range keywordList {
		Scrape(db, keyword)
	}

	SendEmail(emailContents.String())
}
