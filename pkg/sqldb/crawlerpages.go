package sqldb

import "gorm.io/gorm"

type CrawlerPage struct {
	gorm.Model
	ID       uint   `gorm:"primary_key"`
	Crid     uint   `gorm:"column:crid"`
	Pageref  string `gorm:"column:pageref"`
	Filename string `gorm:"column:filename"`
	Filedata string `gorm:"column:filedata"`
}

func (dbc *DBConnData) CreateCrawlerPage(cp *Crawlerpage) error {
	return dbc.DB.Create(cp).Error
}
