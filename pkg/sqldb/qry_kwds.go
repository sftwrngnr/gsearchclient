package sqldb

import (
	"fmt"
	"gorm.io/gorm"
)

type qry_kwds struct {
	gorm.Model
	ID         uint `gorm:"primary_key"`
	Query_id   uint `gorm:"column:query_id"`
	Keyword_id uint `gorm:"column:keyword_id"`
}

func (dbc *DBConnData) AddQueryKeywords(query_id uint, kwds []Keywords) (rval error) {
	for i, k := range kwds {
		fmt.Printf("Creating record %d for %s\n", i, k.Keyword)
		myqryk := &qry_kwds{Query_id: query_id, Keyword_id: k.ID}
		rval = dbc.DB.Create(myqryk).Error
		if rval != nil {
			break
		}
	}
	return
}
