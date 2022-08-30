package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"github.com/gocolly/colly"
	gomail "gopkg.in/gomail.v2"
)

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
