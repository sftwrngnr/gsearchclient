package crawler

import (
	"fmt"
	"github.com/sftwrngnr/gsearchclient/pkg/system"
	"net/url"
)

func TransferURLS() (urls []string, err error) {
	mydbr, terr := system.GetSystemParams().Dbc.GetUrls(false)
	if terr != nil {
		err = terr
		return
	}
	for _, r := range mydbr {
		turl := r.Url
		parsedUrl, myerr := url.Parse(turl)
		if myerr != nil {
			fmt.Printf("%s\n", err.Error())
			continue
		}
		urls = append(urls, fmt.Sprintf("%s://%s", parsedUrl.Scheme, parsedUrl.Host))
	}
	return
}
