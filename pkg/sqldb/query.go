package sqldb

import (
	"fmt"
	"gorm.io/gorm"
)

type Query struct {
	gorm.Model
	ID          uint   `gorm:"primary_key"`
	State       uint   `gorm:"column:state"`
	Querystring string `gorm:"column:query_string"`
}

func (dbc *DBConnData) SaveQueryData(stateid uint, kwds []Keywords,
	zcl []Zipcode, acl []string, qstring string) (queryid uint, err error) {
	myQry := &Query{State: stateid, Querystring: qstring}
	err = dbc.DB.Create(myQry).Error
	if err != nil {
		fmt.Printf("Blew chow in create query %s", err)
		return
	}
	queryid = myQry.ID
	err = dbc.AddQueryKeywords(myQry.ID, kwds)
	if err != nil {
		fmt.Printf("Blew chow in add keywords %s", err)
		return
	}
	err = dbc.AddQueryZipodes(myQry.ID, zcl)
	if err != nil {
		fmt.Printf("Blew chow in add zipcodes %s", err)
		return
	}
	err = dbc.AddQueryAreacodes(myQry.ID, acl)
	if err != nil {
		fmt.Printf("Blew chow in add areacodes %s", err)
		return
	}
	return
}
