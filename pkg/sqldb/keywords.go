package sqldb

import "gorm.io/gorm"

type Keywords struct {
	gorm.Model
	ID      uint   `gorm:"primary_key"`
	Keyword string `gorm:"column:keyword"`
}

func (dbc *DBConnData) GetKeywords() (keywords []Keywords, err error) {
	err = dbc.DB.Find(&keywords).Error
	return
}
