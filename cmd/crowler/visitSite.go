package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/YasserLoukniti/crowlerGo/pkg/protocols"
)

func visitSite(site protocols.Site, results chan<- string) {
	// Request the HTML page.
	res, err := http.Get(site.Domain)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		band, ok := s.Attr("href")
		if ok {
			fmt.Printf("Review %d: %s\n", i, band)

		}
		file := protocols.File{
			Name:     band,
			Url:      site.Domain + "/" + band,
			Lastseen: time.Now(),
			SiteId:   site.Id,
		}
		go createFileRequest(file, results)

	})
}
