package sqldb

import "gorm.io/gorm"

type Cities struct {
	gorm.Model
	ID    uint   `gorm:"primary_key"`
	Name  string `gorm:"column name"`
	State uint   `gorm:"column state"`
}
