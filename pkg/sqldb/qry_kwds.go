package sqldb

import "gorm.io/gorm"

type qry_kwds struct {
	gorm.Model
	ID         uint `gorm:"primary_key"`
	Query_id   uint `gorm:"column:query_id"`
	Keyword_id uint `gorm:"column:keyword_id"`
}
