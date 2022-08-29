package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	badger "github.com/dgraph-io/badger/v3"
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

// Scrape() executes different c.OnHTML() functions based on the set siteFlag.
func Scrape(keyword string) {
	result := Article{
		Keyword: keyword,
	}

	if siteFlag == "google" || siteFlag == "g" {
		c.OnHTML("h3.ipQwMb.ekueJc.RD0gLb", func(e *colly.HTMLElement) {
			result.HeadlineURL.Headline = e.Text

			unformattedLink := e.ChildAttr("a[href]", "href")
			link := "https://news.google.com" + unformattedLink[1:]

			result.HeadlineURL.URL = link
			results = append(results, result)
		})

		c.Visit(fmt.Sprintf(googleNewsURLList.Keyword, keyword, daysSinceFlag))
	} else if siteFlag == "prnewswire" || siteFlag == "prn" {
		c.OnHTML("a.news-release", func(e *colly.HTMLElement) {
			result.HeadlineURL.Headline = e.Text

			unformattedLink := e.Attr("href")
			link := "https://www.prnewswire.com/" + unformattedLink

			result.HeadlineURL.URL = link
			results = append(results, result)
		})

		c.Visit(fmt.Sprintf(prNewsWireURLList.Keyword, keyword))
	} else if siteFlag == "all" {
		result1 := result
		c.OnHTML("h3.ipQwMb.ekueJc.RD0gLb", func(e *colly.HTMLElement) {
			result.HeadlineURL.Headline = e.Text

			unformattedLink := e.ChildAttr("a[href]", "href")
			link := "https://news.google.com" + unformattedLink[1:]

			result.HeadlineURL.URL = link
			results = append(results, result)
		})

		c.Visit(fmt.Sprintf(googleNewsURLList.Keyword, keyword, daysSinceFlag))

		c.OnHTML("a.news-release", func(e *colly.HTMLElement) {
			result1.HeadlineURL.Headline = e.Text

			unformattedLink := e.Attr("href")
			link := "https://www.prnewswire.com/" + unformattedLink

			result1.HeadlineURL.URL = link
			results = append(results, result1)
		})

		c.Visit(fmt.Sprintf(prNewsWireURLList.Keyword, keyword))
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
	}
}

func AddKeyValue(db *badger.DB, k, v string) {
	err := db.Update(func(txn *badger.Txn) error {
		e := badger.NewEntry([]byte(k), []byte(v))
		err := txn.SetEntry(e)
		return err
	})

	if err != nil {
		log.Fatal(err)
	}
}

func Hash(in string) string {
	h := sha256.New()
	h.Write([]byte(in))

	return string(hex.EncodeToString(h.Sum(nil)))
}
