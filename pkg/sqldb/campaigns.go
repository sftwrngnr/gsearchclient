package sqldb

import "gorm.io/gorm"

type Campaigns struct {
	gorm.Model
	ID      uint   `gorm:"primaryKey"`
	Company string `gorm:"column:company"`
	Name    string `gorm:"column:name"`
}

func (dbc *DBConnData) GetCompanyCampaigns(company uint) (campaigns []Campaigns, err error) {
	err = dbc.DB.Where("company = ?", company).Find(&campaigns).Error
	return
}
