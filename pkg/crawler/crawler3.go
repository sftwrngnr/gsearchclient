package crawler

import (
	"fmt"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"net/http"
	"net/url"
	"strings"
)

type Crawler3 struct {
	CurUrl   string
	Urlhost  string
	LocalDir string
	UrlCrawl bool
	transp   *http.Transport
}

type Crawler3Result struct {
}

type Crawler3Links struct {
	visited bool
}

func NewCrawler3(path string, disk bool, dir string) *Crawler3 {
	rval := &Crawler3{CurUrl: path, UrlCrawl: !disk}
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

func (c *Crawler3) checkIgnoreUrls(inl string) bool {
	myIgnl := []string{"youtube.com", "instagram.com", "pinterest.com", "yelp.com"}
	for _, v := range myIgnl {
		if strings.Contains(inl, v) {
			return true
		}
	}
	return false
}

func (c *Crawler3) checkLink(link string) bool {
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

func (c *Crawler3) cleanText(text string) string {
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

func (c *Crawler3) Crawl(sc *Subcrawler) {

	var cLinks map[string]Crawler3Links = make(map[string]Crawler3Links)
	geziyor.NewGeziyor(&geziyor.Options{
		StartRequestsFunc: func(g *geziyor.Geziyor) {
			g.GetRendered(c.Urlhost, g.Opt.ParseFunc)
		},
		ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {
			fmt.Println(string(r.Body))
		},
		//BrowserEndpoint: "ws://localhost:3000",
	}).Start()

	fmt.Printf("There were %d links found.\n", len(cLinks))
	myKeys := make([]string, 0)
	for k, _ := range cLinks {
		myKeys = append(myKeys, k)
	}

	sc.SCRawler(myKeys)
}
