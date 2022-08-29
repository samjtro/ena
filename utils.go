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

var (
	headlineURLs []HeadlineURL
	result       Articles
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

func Scrape(db *badger.DB, keyword string) {
	ScrapeHelper(keyword)

	if typeFlag == "keyword" || typeFlag == "kw" {
		var data string

		for i, x := range results {
			if typeFlag == x.Keyword {
				data += x.HeadlineURLs[i].URL
			}
		}

		hash := Hash(data)
		err := CheckSimilarity(db, keywordFlag, hash)

		if err != nil {
			AddKeyValue(db, keywordFlag, hash)

			if err := tmpl.Execute(&emailContents, results); err != nil {
				log.Fatal(err)
			}

			SendEmail(emailContents.String())
		}
	} else if typeFlag == "multi-keyword" || typeFlag == "mkw" {
		for _, x := range keywordList {
			var data string

			for i, y := range results {
				if x == y.Keyword {
					data += y.HeadlineURLs[i].URL
				}
			}

			hash := Hash(data)
			err := CheckSimilarity(db, x, hash)

			if err != nil {
				sendEmail = true
				AddKeyValue(db, x, hash)

				for _, z := range results {
					if z.Keyword == x {
						resultsForEmail = append(resultsForEmail, z)
					}
				}
			}
		}
	}

	result.HeadlineURLs = headlineURLs
}

func ScrapeHelper(keyword string) {
	headlineURL := HeadlineURL{}

	if result.Keyword == "" {
		result.Keyword = keyword
	} else if result.Keyword != keyword {
		results = append(results, result)
		result = Articles{}
		result.Keyword = keyword
	}

	if siteFlag == "google" || siteFlag == "g" {
		c.OnHTML("h3.ipQwMb.ekueJc.RD0gLb", func(e *colly.HTMLElement) {
			headlineURL.Headline = e.Text

			unformattedLink := e.ChildAttr("a[href]", "href")
			link := "https://news.google.com" + unformattedLink[1:]

			headlineURL.URL = link
			headlineURLs = append(headlineURLs, headlineURL)
		})

		c.Visit(fmt.Sprintf(googleNewsURLList.Keyword, keyword, daysSinceFlag))
	} else if siteFlag == "prnewswire" || siteFlag == "prn" {
		c.OnHTML("a.news-release", func(e *colly.HTMLElement) {
			headlineURL.Headline = e.Text

			unformattedLink := e.Attr("href")
			link := "https://www.prnewswire.com/" + unformattedLink

			headlineURL.URL = link
			headlineURLs = append(headlineURLs, headlineURL)
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

func Hash(in string) string {
	h := sha256.New()
	h.Write([]byte(in))

	return string(hex.EncodeToString(h.Sum(nil)))
}
