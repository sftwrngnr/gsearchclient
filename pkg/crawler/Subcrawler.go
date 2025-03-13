package crawler

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	fp "github.com/amonsat/fullname_parser"
	"github.com/sftwrngnr/gsearchclient/pkg/sqldb"
	"github.com/sftwrngnr/gsearchclient/pkg/system"
	"log"
	"math/rand"
	"os"
	"slices"
	"strings"
	"time"
)

type CPage struct {
	Fname string
	PgRef string
	Succ  bool
}
type Subcrawler struct {
	Email    []string
	Dentists []string
	pcrawled int
	pcrawlf  int
	CPages   []CPage
	Procfunc Filterfunc
}

func (s *Subcrawler) SCRawler(clist []string) {
	getRString := func() string {
		var rval string = ""
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for i := 0; i < 8; i++ {
			rval += fmt.Sprintf("%c", 65+r.Int()%26)
		}
		return rval
	}

	myfn := getRString()
	for i, url := range clist {
		myfname := fmt.Sprintf("/tmp/%s_%d.html", myfn, i)
		fmt.Printf("Crawling %s saving to %s\n", url, myfname)
		// Store file info to crawlerinfo
		turl := url
		if (!strings.Contains(turl, "https://")) && (!strings.Contains(turl, "http://")) {
			turl = "https://" + turl
		}

		err := Crawl(turl, myfname, s)
		// Wait for 60 seconds
		if err != nil {
			log.Printf("Error crawling %s: %s\n", url, err)
		}
		s.CPages = append(s.CPages, CPage{Fname: myfname, PgRef: turl, Succ: err == nil})
		//time.Sleep(0.5 * time.Second)
	}

}

func (sc *Subcrawler) SCCallback(document *goquery.Document) error {
	document.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			if strings.Contains(href, "mailto") {
				mystr := strings.TrimSpace(href)
				maila := strings.Split(mystr, ":")
				if len(maila) > 1 {
					if !slices.Contains(sc.Email, maila[1]) {
						sc.Email = append(sc.Email, maila[1])
					}
				}
			}
		}
	})
	document.Find("h1, h2, h3, h4, strong").Each(func(i int, s *goquery.Selection) {
		//fmt.Printf("%d, %s\n", i, s.Text())
		mystr := strings.TrimSpace(s.Text())
		if strings.Contains(strings.ToLower(mystr), "dds") || strings.Contains(strings.ToLower(mystr), "dmd") || strings.Contains(strings.ToLower(mystr), "msd") {
			sc.appendDentists(mystr)
		} else if strings.Contains(strings.ToLower(mystr), "dr ") || strings.Contains(strings.ToLower(mystr), "dr.") {
			tstr := mystr
			if len(mystr) > 60 {
				tstr = tstr[:60]
			}
			name := fp.ParseFullname(tstr)
			fmt.Printf("%s::%v\n", tstr, name)
			sc.appendDentists(mystr)
		}

	})
	return nil
}

func (sc *Subcrawler) appendDentists(mystr string) {
	//`minNLen := 2 // minum number of characters required for a name
	//suffixes := []string{"DDS", "DMD, MS", "MD"}
	mysLine := sc.sanitize(mystr)
	sc.Dentists = append(sc.Dentists, strings.Split(mysLine, ";")...)
}

func (sc *Subcrawler) sanitize(str string) string {
	rval := ""
	for _, r := range str {
		ar := r
		if r == 10 || r == 13 {
			ar = ';'
		} else if r < 32 {
			ar = 32
		}
		rval += string(ar)
	}
	return str
}
func (sc *Subcrawler) TransferDataToDB(cid uint) {
	fmt.Printf("Transfering data to DB\n")
	for _, dta := range sc.CPages {
		// Load raw data
		myCP := &sqldb.Crawlerpage{Crid: cid, Pageref: dta.PgRef, Filename: dta.Fname, Filedata: sc.getFileData(dta.Fname)}
		err := system.GetSystemParams().Dbc.CreateCrawlerPage(myCP)
		if err != nil {
			log.Printf("Error creating crawler page: %s\n", err)
		}
	}
	// Build email list if necessary
	tEmail := ""

	if len(sc.Email) > 1 {
		tEmail = strings.Join(sc.Email, ",")
	} else if len(sc.Email) == 1 {
		tEmail = sc.Email[0]
	}
	for _, dta := range sc.Dentists {
		myD := &sqldb.Simpleresult{Cid: cid, Dentist: dta, Email: tEmail}
		err := system.GetSystemParams().Dbc.CreateSimpleResults(myD)
		if err != nil {
			log.Printf("Error creating simpler results: %s\n", err)
		}
	}
}

func (sc *Subcrawler) getFileData(fn string) string {
	f, err := os.ReadFile(fn)
	if err != nil {
		fmt.Printf("Error opening file %s: %s\n", fn, err)
		return ""
	}
	err = os.Remove(fn)
	if err != nil {
		fmt.Printf("Error removing file %s: %s\n", fn, err)
	}
	return string(f)
}
