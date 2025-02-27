package crawler

import (
	"fmt"
	"log"

	"github.com/playwright-community/playwright-go"
)

func Crawl(url string) {
	fmt.Printf("Crawling %s\n", url)
	/*
		err := playwright.Install()
		if err != nil {
			log.Fatal(err)
		}

	*/

	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	browser, err := pw.Chromium.Launch()
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	if _, err = page.Goto(url); err != nil {
		log.Fatalf("could not goto: %v", err)
	}
	fmt.Printf("Opened page %s\n", url)

	content, cerr := page.Content()
	if cerr != nil {
		log.Fatalf("could not get content: %v", cerr)
	}
	fmt.Printf("%v\n", content)
	var myData interface{}
	test, err := page.Evaluate("div", &myData)
	if err != nil {
		log.Fatalf("could not evaluate test: %v", err)
	}
	fmt.Printf("%v\n", test)
	entries, err := page.Locator("<div").All()
	fmt.Printf("%v\n", entries)
	if err != nil {
		log.Fatalf("could not get entries: %v", err)
	}
	/*
		for i, entry := range entries {

			title, err := entry.Locator("td.title > span > a").TextContent()
			if err != nil {
				log.Fatalf("could not get text content: %v", err)
			}
			fmt.Printf("%d: %s\n", i+1, title)
		}

	*/
	if err = browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
	if err = pw.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}
}
