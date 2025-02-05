package sqldb

import "gorm.io/gorm"

type Cityareacodes struct {
	gorm.Model
	ID       uint `gorm:"primary_key"`
	Areacode uint `gorm:"column:areacode"`
	City     uint `gorm:"column:city"`
}
