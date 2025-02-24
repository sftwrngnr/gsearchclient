package sqldb

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"column name"`
}

func (dbc *DBConnData) GetCompanyList() (comp []Company, err error) {
	err = dbc.DB.Find(&comp).Error
	return
}
