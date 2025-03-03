package crawler

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"net/http"
	"net/url"
	"strings"
)

type Crawler2 struct {
	CurUrl    string
	Urlhost   string
	LocalDir  string
	LocalFile string
	UrlCrawl  bool
	transp    *http.Transport
	colly     *colly.Collector
}

type Crawler2Result struct {
}

type Crawler2Links struct {
	visited bool
}

func NewCrawler2(path string, disk bool, file string, dir string) *Crawler2 {
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

func (c *Crawler2) checkLink(link string) bool {
	crawllinks := []string{"about", "doctors", "staff", "location", "dentists", "meet", "office", "dr", "info", "phone"}
	myl, err := url.Parse(link)
	//fmt.Printf("%s\n", link)
	if myl[0] == '/' {
		if len(myl) > 2 {
			if myl[1] != '/' {
				myl = fmt.Sprintf("https://%s%s", c.Urlhost, myl)
			}
		}
	}
	parsedUrl, myerr := url.Parse(myl)
	if myerr != nil {
		return false
	}
	fmt.Printf("%v\n", parsedUrl)
	if parsedUrl.Hostname() != c.CurUrl {
		fmt.Printf("Error Parsing URL(2) %s\n", myl)
		fmt.Printf("%s %s\n", parsedUrl.Host, c.CurUrl)
		return false
	}
	tl := strings.ToLower(myl)
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

func (c *Crawler2) Crawl() {

	var cLinks map[string]Crawler2Links = make(map[string]Crawler2Links)
	c.colly = colly.NewCollector(
		colly.MaxDepth(2),
		//colly.Async()
	)
	if !c.UrlCrawl {
		c.colly.WithTransport(c.transp)
	}

	c.colly.OnHTML("name, practice, phone, address, tel, addy, location, dr, info", func(e *colly.HTMLElement) {
		fmt.Printf("%v:%s\n", e.Name, c.cleanText(e.Text))
	})

	c.colly.OnHTML("a", func(e *colly.HTMLElement) {
		fmt.Printf("%s::%v\n", e.Name, e.DOM.Nodes[0])
	})

	c.colly.OnHTML("div[phone]", func(e *colly.HTMLElement) {
		fmt.Printf("Phone:%s\n", e.Name)
	})

	c.colly.OnHTML("a[phone]", func(e *colly.HTMLElement) {
		fmt.Printf("Phone:%s:%v\n", e.Name, e.DOM.Nodes)
	})

	c.colly.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if _, ok := cLinks[link]; !ok {
			if c.checkLink(link) {
				if link[0] == '/' {
					link = fmt.Sprintf("%s%s", c.CurUrl, link)
				}
				cLinks[link] = Crawler2Links{visited: false}
				fmt.Printf("Link: %s\n", link)
			}
		}
	})

	c.colly.OnHTML("div", func(r *colly.HTMLElement) {
		//fmt.Printf("%v\n", r.Attr)
	})
	// Find the parent element, then find children by selector

	var err error
	if c.UrlCrawl {
		err = c.colly.Visit(c.CurUrl)
	} else {
		err = c.colly.Visit("file://" + c.LocalDir + "/" + c.LocalFile)
	}
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("There were %d links found.\n", len(cLinks))
	fmt.Printf("Links are:  %v\n", cLinks)
}
