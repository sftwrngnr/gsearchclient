package sqldb

import "gorm.io/gorm"

type States struct {
	gorm.Model
	ID      uint   `gorm:"primary_key"`
	Abbrev  string `gorm:"column:abbrev"`
	Name    string `gorm:"column:name"`
	Capitol string `gorm:"column:capitol"`
	Region  string `gorm:"column:region"`
}
