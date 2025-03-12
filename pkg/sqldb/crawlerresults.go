package sqldb

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Crawlerresults struct {
	gorm.Model
	ID             uint      `gorm:"primary_key"`
	Queryid        uint      `gorm:"column:queryid"`
	Url            string    `gorm:"column:url"`
	Crawled        bool      `gorm:"column:crawled"`
	Pagescrawled   uint      `gorm:"column:pagescrawled"`
	Crawldepth     uint      `gorm:"column:crawldepth"`
	Totalduration  float64   `gorm:"column:totalduration"`
	Alloweddomains []byte    `gorm:"type:jsonb:columnn:alloweddomains"`
	Success        bool      `gorm:"column:success"`
	Crawldate      time.Time `gorm:"column:crawldate"`
	Pagecrawlsucc  float64   `gorm:"column:pagecrawlsucc"`
	Profile        uint      `gorm:"column:profile"`
	Urlimportdate  time.Time `gorm:"column:urlimportdate"`
	Status         uint      `gorm:"column:status"`
	Crawler        uint      `gorm:"column:crawler"`
	Campaign       uint      `gorm:"column:campaign"`
}

func (dbc *DBConnData) TransferQryUrls(qid uint, url string, alldomains []byte, crawler uint, campaign uint) (err error) {
	cr := Crawlerresults{Queryid: qid, Url: url, Urlimportdate: time.Now(), Alloweddomains: alldomains,
		Crawler: crawler, Campaign: campaign}
	err = dbc.DB.Create(&cr).Error
	if err != nil {
		fmt.Printf("Crawleresults error while transferring %d URLs to %s::%s\n", qid, url, err.Error())
	}
	return

}

func (dbc *DBConnData) GetUrlsToCrawl(campaign uint, crawler uint) (cres []Crawlerresults, err error) {
	mysqlqry := "campaign = ? and crawler = ? and crawled = false"
	if crawler == 0 {
		// Now crawler specified, get results for campaign
		mysqlqry = "campaign = ? and crawled = false"
		fmt.Printf("")
		err = dbc.DB.Where(mysqlqry, campaign).Find(&cres).Error
		return
	}
	err = dbc.DB.Where(mysqlqry, campaign, crawler).Find(&cres).Error
	return
}

func (dbc *DBConnData) UpdateCrawlerresults(cres *Crawlerresults) (err error) {
	err = dbc.DB.Save(cres).Error
	return
}
