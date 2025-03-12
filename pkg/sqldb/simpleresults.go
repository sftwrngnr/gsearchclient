package sqldb

import "gorm.io/gorm"

type Simpleresult struct {
	gorm.Model
	ID      uint   `gorm:"primary_key"`
	Cid     uint   `gorm:column:"cid"`
	Dentist string `gorm:"column:dentist"`
	Email   string `gorm:"column:email"`
}

func (dbc *DBConnData) CreateSimpleResults(sr *Simpleresult) error {
	return dbc.DB.Create(sr).Error
}
