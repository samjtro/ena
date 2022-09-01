package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"os/user"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/gocolly/colly"
	"github.com/joho/godotenv"
	gomail "gopkg.in/gomail.v2"
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

	tmpl = template.Must(template.ParseFiles(fmt.Sprintf("/home/%s/sn/tmpl/template.html", currentUser.Username)))
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

	if err := tmpl.Execute(&emailContents, results); err != nil {
		log.Fatal(err)
	}

	SendEmail(emailContents.String())
}

func Now(t time.Time) string {
	str := fmt.Sprintf("%d-%d-%dT%d:%d:%d%s",
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
		utcDiffFlag)

	return str
}

func Scrape(keyword string) {
	var headlineURLs []HeadlineURL

	result := Articles{
		Keyword: keyword,
	}

	if siteFlag == "google" || siteFlag == "g" {
		c.OnHTML("h3.ipQwMb.ekueJc.RD0gLb", func(e *colly.HTMLElement) {
			headlineURL := HeadlineURL{}
			headlineURL.Headline = e.Text

			unformattedLink := e.ChildAttr("a[href]", "href")
			link := "https://news.google.com" + unformattedLink[1:]

			headlineURL.URL = link
			headlineURLs = append(headlineURLs, headlineURL)
		})

		c.Visit(fmt.Sprintf(googleNewsURLList.Keyword, keyword, daysSinceFlag))
		result.HeadlineURLs = headlineURLs
		results = append(results, result)
	} else if siteFlag == "prnewswire" || siteFlag == "prn" {
		c.OnHTML("a.news-release", func(e *colly.HTMLElement) {
			headlineURL := HeadlineURL{}
			headlineURL.Headline = e.Text

			unformattedLink := e.Attr("href")
			link := "https://www.prnewswire.com/" + unformattedLink

			headlineURL.URL = link
			headlineURLs = append(headlineURLs, headlineURL)
		})

		c.Visit(fmt.Sprintf(prNewsWireURLList.Keyword, keyword))
		result.HeadlineURLs = headlineURLs
		results = append(results, result)
	} else {
		log.Fatalf("Incorrect Site Flag (--s). Should have been either 'google' or 'prnewswire'. Try again with a different value, or simply use the default 'prnewswire' by avoiding setting --s entirely.")
	}
}

func SendEmail(message string) {
	msg := gomail.NewMessage()
	msg.SetHeader("From", fromFlag)
	msg.SetHeader("To", toFlag)
	msg.SetHeader("Subject", fmt.Sprintf("[sn] Keyword: %s, Site: %s - %s", keywordFlag, siteFlag, Now(time.Now())))
	msg.SetBody("text/html", message)
	handler := gomail.NewDialer("smtp.gmail.com", 587, fromFlag, fromPasswordFlag)

	if err := handler.DialAndSend(msg); err != nil {
		log.Fatalf("Error: %s", err.Error())
	} else {
		fmt.Println("** Email Sent **")
	}
}

func Hash(in string) string {
	h := sha256.New()
	h.Write([]byte(in))

	return string(hex.EncodeToString(h.Sum(nil)))
}
