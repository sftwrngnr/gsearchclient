package http

import (
	"bufio"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/sftwrngnr/gsearchclient/pkg/crawler"
	"log"
	"os"
)

func TestSubCrawler(sl []string) {
	fmt.Println("TestSubCrawler")
	var sc crawler.Subcrawler
	for _, s := range sl {
		// Open and test
		fr, err := os.Open(s)
		if err != nil {
			fmt.Println(err)
			continue
		}
		rdr := bufio.NewReader(fr)
		doc, derr := goquery.NewDocumentFromReader(rdr)
		if derr != nil {
			log.Printf("could not create document: %v", derr)
			return
		}
		myerr := crawler.SCCallback(doc, &sc)
		if myerr != nil {
			fmt.Println(myerr)
		}
	}
	fmt.Printf("Crawl results: %v\n", sc)
}
