package crawler

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"os"
	"strings"

	"github.com/playwright-community/playwright-go"
)

func Crawl(url string, fname string, sc *Subcrawler) (err error) {
	var tfile string = "/tmp/index.hthml"
	var pw *playwright.Playwright
	var browser playwright.Browser
	shutdownpw := func() {
		if err = browser.Close(); err != nil {
			log.Fatalf("could not close browser: %v", err)
		}
		if err = pw.Stop(); err != nil {
			log.Fatalf("could not stop Playwright: %v", err)
		}

	}
	fmt.Printf("Crawling %s\n", url)
	pw, err = playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	browser, err = pw.Chromium.Launch()
	defer shutdownpw()
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
		return
	}
	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	if _, err = page.Goto(url); err != nil {
		log.Printf("could not goto: %v", err)
		return
	}
	fmt.Printf("Opened page %s\n", url)

	content, cerr := page.Content()
	if cerr != nil {
		log.Printf("could not get content: %v", cerr)
		return
	}

	if fname != "" {
		tfile = fname
	}
	f, ferr := os.OpenFile(tfile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if ferr != nil {
		log.Fatalf("could not open file: %v", ferr)
	}
	defer f.Close()

	_, ferr = f.WriteString(string(content))
	if ferr != nil {
		log.Fatalf("could not write to file: %v", ferr)
	}

	//fmt.Printf("%v\n", content)
	if sc.Procfunc != nil {
		doc, derr := goquery.NewDocumentFromReader(strings.NewReader(content))
		if derr != nil {
			log.Printf("could not create document: %v", derr)
			return
		}

		myerr := sc.Procfunc(doc)
		if myerr != nil {
			log.Printf("Filterfunc failed: %v", myerr)
		}
	} else {
		err = errors.New("No processor function available for crawling")
		log.Printf("%s\n", err.Error())

	}

	return err
}
