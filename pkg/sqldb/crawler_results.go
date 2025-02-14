package sqldb

import "gorm.io/gorm"

type CrawlerResults struct {
	gorm.Model
	ID           uint   `json:"id"`
	Title        string `json:"title"`
	Link         string `json:"link"`
	Display_link string `json:"display_link"`
}
