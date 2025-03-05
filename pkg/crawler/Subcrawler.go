package crawler

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"log"
	"strings"
	"time"
)

func SCRawler(clist []string) {

	hasher := sha1.New()
	for i, url := range clist {
		hasher.Write([]byte(url))
		sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
		myfname := fmt.Sprintf("/tmp/%s_%d.html", sha[0:8], i)
		fmt.Printf("Crawling %s saving to %s\n", url, myfname)
		turl := url
		if !strings.HasPrefix(turl, "http://") || !strings.HasPrefix(turl, "https://") {
			turl = "https://" + turl
		}
		err := Crawl(turl, myfname)
		// Wiat for 60 seconds
		if err != nil {
			log.Printf("Error crawling %s: %s\n", url, err)
		}
		time.Sleep(5 * time.Second)
	}

}
