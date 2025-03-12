package crawler

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"net/http"
	"net/url"
	"strings"
)

type Crawler2 struct {
	CrawlId  uint
	CurUrl   string
	Urlhost  string
	LocalDir string
	UrlCrawl bool
	transp   *http.Transport
	colly    *colly.Collector
}

type Crawler2Result struct {
}

type Crawler2Links struct {
	visited bool
}

func NewCrawler2(path string, disk bool, dir string) *Crawler2 {
	rval := &Crawler2{CurUrl: path, UrlCrawl: !disk}
	if disk {
		rval.transp = &http.Transport{}

		rval.transp.RegisterProtocol("file", http.NewFileTransport(http.Dir(dir)))

	}
	myurl, err := url.Parse(path)
	if err != nil {
		return nil
	}
	rval.Urlhost = myurl.Host
	return rval
}

func (c *Crawler2) checkIgnoreUrls(inl string) bool {
	myIgnl := []string{"youtube.com", "instagram.com", "pinterest.com"}
	for _, v := range myIgnl {
		if strings.Contains(inl, v) {
			return true
		}
	}
	return false
}

func (c *Crawler2) checkLink(link string) bool {
	crawllinks := []string{"about", "doctors", "staff", "location", "Dentists", "meet", "office", "dr", "info", "phone"}
	myl, err := url.Parse(link)
	if err != nil {
		return false
	}
	tl := strings.ToLower(myl.String())
	if c.checkIgnoreUrls(tl) {
		return false
	}
	for _, l := range crawllinks {
		if strings.Contains(tl, l) {
			return true
		}
	}
	return false
}

func (c *Crawler2) cleanText(text string) string {
	slist := strings.Split(text, "\n")
	var rval string
	for _, l := range slist {
		ts := strings.TrimSpace(l)
		if len(ts) > 0 {
			rval = rval + fmt.Sprintf("%s\n", ts)
		}
	}

	return rval
}

func (c *Crawler2) Crawl(sc *Subcrawler) {

	var cLinks map[string]Crawler2Links = make(map[string]Crawler2Links)
	c.colly = colly.NewCollector(
		colly.MaxDepth(2),
		//colly.Async()
	)
	if !c.UrlCrawl {
		c.colly.WithTransport(c.transp)
	}

	c.colly.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if _, ok := cLinks[link]; !ok {
			if c.checkLink(link) {
				if link[0] == '/' {
					link = fmt.Sprintf("%s%s", c.Urlhost, link)
				}

				cLinks[link] = Crawler2Links{visited: false}
				//sfmt.Printf("Link: %s\n", link)
			}
		}
	})

	c.colly.OnHTML("div", func(r *colly.HTMLElement) {
		//fmt.Printf("%v\n", r.Attr)
	})
	// Find the parent element, then find children by selector

	var err error
	if c.UrlCrawl {
		// Save index
		err = c.colly.Visit(c.CurUrl)
		c.colly.Wait()
	} else {
		err = c.colly.Visit("file://" + c.LocalDir + "/" + "index.html")
	}
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("There were %d links found.\n", len(cLinks))
	myKeys := make([]string, 0)
	for k, _ := range cLinks {
		myKeys = append(myKeys, k)
	}
	sc.SCRawler(myKeys)
	// Transfer sc data
	sc.TransferDataToDB(c.CrawlId)
}
