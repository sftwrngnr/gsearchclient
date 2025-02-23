package crawler

import (
	"encoding/json"
	"fmt"
	"github.com/sftwrngnr/gsearchclient/pkg/system"
	"net/url"
	"strings"
)

func TransferURLS() (urls []string, err error) {
	alldomains := func(url string) []byte {
		myDomains := AllowedDomains{}
		myDomains.Domains = append(myDomains.Domains, url)
		pdom := strings.Index(url, ".")
		if pdom != -1 {
			pdom++
			myDomains.Domains = append(myDomains.Domains, url[pdom:])
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
		fmt.Printf("Processing %s\n", r.Url)
		turl := r.Url
		parsedUrl, myerr := url.Parse(turl)
		if myerr != nil {
			fmt.Printf("%s\n", err.Error())
			continue
		}
		myUrl := fmt.Sprintf("%s://%s", parsedUrl.Scheme, parsedUrl.Host)
		urls = append(urls, myUrl)
		err = system.GetSystemParams().Dbc.TransferQryUrls(1, myUrl, alldomains(parsedUrl.Host), 6, 1)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
		}
		r.Transferred = true
		err = system.GetSystemParams().Dbc.UpdateRec(&r)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
		}
	}
	return
}
