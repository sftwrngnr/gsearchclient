package crawler

import (
	"fmt"
	"log"
	"os"

	"github.com/playwright-community/playwright-go"
)

func Crawl(url string, fname string, procfunc Filterfunc) (err error) {
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
	/*
		err := playwright.Install()
		if err != nil {
			log.Fatal(err)
		}

	*/

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

	fmt.Printf("%v\n", content)
	if procfunc != nil {
		myerr := procfunc(content)
		if myerr != nil {
			log.Printf("Filterfunc failed: %v", myerr)
		}
	}

	var myData interface{}
	test, err := page.Evaluate("div", &myData)
	if err != nil {
		log.Printf("could not evaluate test: %v", err)
	} else {
		fmt.Printf("%v\n", test)
	}
	entries, err := page.Locator("<div").All()
	fmt.Printf("%v\n", entries)
	if err != nil {
		log.Printf("could not get entries: %v", err)
	}
	return err
	/*
		for i, entry := range entries {

			title, err := entry.Locator("td.title > span > a").TextContent()
			if err != nil {
				log.Fatalf("could not get text content: %v", err)
			}
			fmt.Printf("%d: %s\n", i+1, title)
		}

	*/
}
