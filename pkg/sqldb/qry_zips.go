package sqldb

import (
	"fmt"
	"gorm.io/gorm"
)

type qry_zips struct {
	gorm.Model
	ID       uint `gorm:"primary_key"`
	Query_id uint `gorm:"column:query_id"`
	Zip_id   uint `gorm:"column:zip_id"`
}

func (dbc *DBConnData) AddQueryZipodes(query_id uint, acs []Zipcode) (rval error) {
	for i, k := range acs {
		fmt.Printf("Creating record %d for %s\n", i, k.Zipcode)
		myqz := &qry_zips{Query_id: query_id, Zip_id: k.ID}
		rval = dbc.DB.Create(myqz).Error
		if rval != nil {
			break
		}
	}
	return
}
