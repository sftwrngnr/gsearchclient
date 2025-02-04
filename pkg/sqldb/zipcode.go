package sqldb

import "gorm.io/gorm"

type Zipcode struct {
	gorm.Model
	ID         uint   `gorm:"primary_key"`
	Zipcode    string `gorm:"column:zipcode"`
	City       uint   `gorm:"column:city"`
	State      uint   `gorm:"column:state"`
	Population uint   `gorm:"column:population"`
}
