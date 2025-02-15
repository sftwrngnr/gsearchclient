package sqldb

import "gorm.io/gorm"

type qry_ac struct {
	gorm.Model
	ID       uint `gorm:"primary_key"`
	Query_id uint `gorm:"column:query_id"`
	Qry_ac   uint `gorm:"column:qry_ac"`
}
