package sqldb

import "gorm.io/gorm"

type CrawlerResults struct {
	gorm.Model
	ID             uint    `gorm:"primary_key"`
	Queryid        uint    `gorm:"column:queryid"`
	Url            string  `gorm:"column:url"`
	Pagescrawled   uint    `gorm:"column:pagescrawled"`
	Crawldepth     uint    `gorm:"column:crawldepth"`
	Totalduration  float64 `gorm:"column:totalduration"`
	Alloweddomains []byte  `gorm:"type:jsonb:columnn:alloweddomains"`
	Link           string  `json:"column:link"`
	Display_link   string  `json:"column:display_link"`
}

func (dbc *DBConnData) TransferQryUrls(qid uint, url string) (err error) {
	dbc.DB.Create(&CrawlerResults{Queryid: qid, Url: url})
	return

}
