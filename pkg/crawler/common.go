package crawler

import "github.com/PuerkitoBio/goquery"

type AllowedDomains struct {
	Domains []string `json:"domains"`
}

type Filterfunc func(document *goquery.Document, sc *Subcrawler) error
