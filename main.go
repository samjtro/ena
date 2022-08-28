package main

import (
	"bytes"
	"errors"
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

type Article struct {
	Keyword     string
	HeadlineURL HeadlineURL
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
	results          []Article
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
	var emailContents bytes.Buffer
	var results1 []Article
	keywordList := strings.Split(keywordFlag, ",")

	if typeFlag == "multi-keyword" || typeFlag == "mkw" {
		for _, keyword := range keywordList {
			Scrape(keyword)
		}
	} else if typeFlag == "keyword" || typeFlag == "kw" {
		Scrape(keywordFlag)
	}

	db, err := badger.Open(badger.DefaultOptions("/tmp/badger"))

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if typeFlag == "keyword" || typeFlag == "kw" {
		var data string

		for _, x := range results {
			if typeFlag == x.Keyword {
				data += x.HeadlineURL.URL
			}
		}

		hash := Hash(data)

		err = db.View(func(txn *badger.Txn) error {
			it := txn.NewIterator(badger.DefaultIteratorOptions)

			defer it.Close()

			for it.Rewind(); it.Valid(); it.Next() {
				item := it.Item()

				var key, val []byte
				item.KeyCopy(key)
				item.ValueCopy(val)

				if string(key) == keywordFlag {
					if string(val) == hash {
						return nil
					} else {
						return errors.New("")
					}
				} else {
					return errors.New("")
				}
			}

			return nil
		})

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

			for _, y := range results {
				if x == y.Keyword {
					data += y.HeadlineURL.URL
				}
			}

			hash := Hash(data)

			err = db.View(func(txn *badger.Txn) error {
				it := txn.NewIterator(badger.DefaultIteratorOptions)

				defer it.Close()

				for it.Rewind(); it.Valid(); it.Next() {
					item := it.Item()

					var key, val []byte
					item.KeyCopy(key)
					item.ValueCopy(val)

					if string(key) == keywordFlag {
						if string(val) == hash {
							return nil
						} else {
							return errors.New("")
						}
					} else {
						return errors.New("")
					}
				}

				return nil
			})

			if err != nil {
				sendEmail = true
				AddKeyValue(db, keywordFlag, hash)

				for _, z := range results {
					if z.Keyword == x {
						results1 = append(results1, z)
					}
				}
			}
		}

		if sendEmail {
			if err := tmpl.Execute(&emailContents, results1); err != nil {
				log.Fatal(err)
			}

			SendEmail(emailContents.String())
		}
	}
}
