package sqldb

import "gorm.io/gorm"

type Ignoreurl struct {
	gorm.Model
	ID         uint   `gorm:"primary_key"`
	Url        string `gorm:"column:url"`
	Subdomains string `gorm:"column:subdomains"`
}

func (dbc *DBConnData) GetIgnoreUrls() (urls []Ignoreurl, err error) {
	err = dbc.DB.Find(&urls).Error
	return
}
