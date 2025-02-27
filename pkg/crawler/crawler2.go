package crawler

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

type Crawler2 struct {
	CurUrl string
	colly  *colly.Collector
}

type Crawler2Result struct {
}

type Crawler2Links struct {
	visited bool
}

func NewCrawler2(curUrl string) *Crawler2 {
	return &Crawler2{CurUrl: curUrl}
}

func (c *Crawler2) Crawl() {
	var cLinks map[string]Crawler2Links = make(map[string]Crawler2Links)
	c.colly = colly.NewCollector(
		colly.MaxDepth(2),
		//colly.Async()
	)

	c.colly.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if val, ok := cLinks[link]; !ok {
			cLinks[link] = Crawler2Links{visited: false}
			fmt.Printf("Link: %s\n", link)
		} else {
			fmt.Printf("Duplicate Link: %s:%v\n", link, val)
		}
	})
	c.colly.Visit(c.CurUrl)
	fmt.Printf("There were %d links found.\n", len(cLinks))
	fmt.Printf("Links are:  %v\n", cLinks)
}
