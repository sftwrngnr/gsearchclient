package crawler

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

func SCRawler(clist []string) {

	hasher := sha1.New()
	for i, url := range clist {
		hasher.Write([]byte(url))
		sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
		myfname := fmt.Sprintf("/tmp/%s_%d.html", sha[0:8], i)
		fmt.Printf("Crawling %s saving to %s\n", url, myfname)
		Crawl(url, myfname)
	}

}
