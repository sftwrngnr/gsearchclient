package crawler

import (
	"encoding/json"
	"fmt"
	"github.com/sftwrngnr/gsearchclient/pkg/system"
	"net/url"
	"strconv"
	"strings"
)

func TransferURLS(mymap map[string][]string) (urls []string, err error) {
	var (
		//comp int
		camp  int
		crawl int
		iurls []string
	)
	/*
		comp, err = strconv.Atoi(mymap["Company"][0])
		if err != nil {
			return
		}

	*/
	ignurls, ierr := system.GetSystemParams().Dbc.GetIgnoreUrls()
	if ierr != nil {
		err = ierr
		return
	}
	for _, iurl := range ignurls {
		iurls = append(iurls, iurl.Url)
	}
	camp, err = strconv.Atoi(mymap["Campaign"][0])
	if err != nil {
		return
	}
	crawl, err = strconv.Atoi(mymap["Crawler"][0])
	if err != nil {
		return
	}

	alldomains := func(url string) []byte {
		myDomains := AllowedDomains{}
		turl := url
		if !strings.Contains(strings.ToLower(url), "www.") {
			turl = "www." + url
			fmt.Printf("%s\n", turl)
		}
		myDomains.Domains = append(myDomains.Domains, turl)
		pdom := strings.Index(turl, ".")
		if pdom != -1 {
			if strings.Index(turl, "www.") != -1 {
				pdom++
				myDomains.Domains = append(myDomains.Domains, turl[pdom:])
			}
		}
		fmt.Printf("%v\n", myDomains.Domains)
		rval, err := json.Marshal(myDomains)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
		return rval
	}
	mydbr, terr := system.GetSystemParams().Dbc.GetUrls(false)
	if terr != nil {
		err = terr
		return
	}
	for _, r := range mydbr {
		if FoundIgnoreUrl(r.Url, iurls) {
			continue
		}
		fmt.Printf("Processing %s\n", r.Url)
		turl := r.Url
		parsedUrl, myerr := url.Parse(turl)
		if myerr != nil {
			fmt.Printf("%s\n", err.Error())
			continue
		}
		myUrl := fmt.Sprintf("%s://%s", parsedUrl.Scheme, parsedUrl.Host)
		urls = append(urls, myUrl)
		err = system.GetSystemParams().Dbc.TransferQryUrls(1, myUrl, alldomains(parsedUrl.Host), uint(crawl), uint(camp))
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			continue
		}
		r.Transferred = true
		err = system.GetSystemParams().Dbc.UpdateRec(&r)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
		}
	}
	return
}

func FoundIgnoreUrl(url string, urls []string) bool {
	for _, u := range urls {
		if strings.Contains(url, u) {
			return true
		}
	}
	return false
}
