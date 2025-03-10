package crawler

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"slices"
	"strings"
	"time"
)

type Subcrawler struct {
	email    []string
	dentists []string
}

func SCRawler(clist []string) {

	hasher := sha1.New()
	var sc Subcrawler
	for i, url := range clist {
		hasher.Write([]byte(url))
		sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
		myfname := fmt.Sprintf("/tmp/%s_%d.html", sha[0:8], i)
		fmt.Printf("Crawling %s saving to %s\n", url, myfname)
		turl := url
		if !strings.HasPrefix(turl, "http://") || !strings.HasPrefix(turl, "https://") {
			turl = "https://" + turl
		}

		err := Crawl(turl, myfname, &sc, SCCallback)
		// Wiat for 60 seconds
		if err != nil {
			log.Printf("Error crawling %s: %s\n", url, err)
		}
		time.Sleep(5 * time.Second)
	}

}

func SCCallback(document *goquery.Document, sc *Subcrawler) error {
	document.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			if strings.Contains(href, "mailto") {
				mystr := strings.TrimSpace(href)
				maila := strings.Split(mystr, ":")
				if len(maila) > 1 {
					if !slices.Contains(sc.email, maila[1]) {
						sc.email = append(sc.email, maila[1])
					}
				}
			}
		}
	})
	document.Find("h1, h2, h3").Each(func(i int, s *goquery.Selection) {
		//fmt.Printf("%d, %s\n", i, s.Text())
		mystr := strings.TrimSpace(s.Text())
		if strings.Contains(strings.ToLower(mystr), "dds") || strings.Contains(strings.ToLower(mystr), "dmd") {
			sc.dentists = appendDentists(sc.dentists, mystr)
		}
	})
	return nil
}

func appendDentists(dentists []string, mystr string) []string {
	//`minNLen := 2 // minum number of characters required for a name
	//suffixes := []string{"DDS", "DMD, MS", "MD"}
	mysLine := sanitize(mystr)
	myLine := strings.Split(mysLine, ";")
	return myLine
}

func sanitize(str string) string {
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

func parse(str string, suffixes []string) []string {
	checkSfx := func(instr string) bool {
		for _, sfx := range suffixes {
			if strings.Contains(strings.ToUpper(instr), sfx) {
				return true
			}
		}
		return false
	}
	var rval []string = make([]string, 0)
	var lname bool = false
	splitstr := strings.Split(str, ",")
	myName := ""

	for _, k := range splitstr {
		if checkSfx(k) {
			fmt.Printf("Found %s\n", k)
			myName += "," + k
		} else {
			if strings.Index(k, ".") == -1 {
				if !lname {
					myName = k
					lname = true
				} else {
					myName += k
					lname = false
				}
			} else {
				myName += k
			}
		}
	}
	return rval
}
