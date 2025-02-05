package sqldb

import "gorm.io/gorm"

type Areacodes struct {
	gorm.Model
	ID    uint   `gorm:"primary_key"`
	Code  string `gorm:"column:code"`
	State uint   `gorm:"column:state"`
}
