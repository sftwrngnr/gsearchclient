package sqldb

import (
	"fmt"
	"gorm.io/gorm"
)

type qry_ac struct {
	gorm.Model
	ID       uint   `gorm:"primary_key"`
	Query_id uint   `gorm:"column:query_id"`
	Qry_ac   string `gorm:"column:qry_ac"`
}

func (dbc *DBConnData) AddQueryAreacodes(query_id uint, acs []string) (rval error) {
	for i, k := range acs {
		fmt.Printf("Creating record %d for %s\n", i, k)
		myqryac := &qry_ac{Query_id: query_id, Qry_ac: k}
		rval = dbc.DB.Create(myqryac).Error
		if rval != nil {
			break
		}
	}
	return
}
