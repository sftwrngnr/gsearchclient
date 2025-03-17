package sqldb

import (
	"fmt"
	"gorm.io/gorm"
)

type Zipcode struct {
	gorm.Model
	ID         uint    `gorm:"primary_key"`
	Zipcode    string  `gorm:"column:zipcode"`
	City       uint    `gorm:"column:city"`
	State      uint    `gorm:"column:state"`
	Population uint    `gorm:"column:population"`
	Latitude   float32 `gorm:"column:latitude"`
	Longitude  float32 `gorm:"column:longitude"`
}

func (dbc *DBConnData) GetZipsForState(abbrv string) (zips []Zipcode, err error) {
	var myStid uint
	myStid, err = dbc.GetStateId(abbrv)
	if err != nil {
		return
	}
	myzcq := Zipcode{State: myStid}
	err = dbc.DB.Find(&zips, &myzcq).Error
	fmt.Printf("There are %d records returned\n", len(zips))
	return
}

func (dbc *DBConnData) GetZipcodesForList(zc []string, zcl *[]Zipcode) (err error) {
	err = dbc.DB.Where("zipcode IN ?", zc).Find(&zcl).Error
	fmt.Printf("There are %d records returned\n", len(*zcl))
	return
}

func (dbc *DBConnData) Top10Zipcodes(stid uint) (zips []Zipcode, err error) {
	err = dbc.DB.Where("state = ?", stid).Order("population DESC").Limit(10).Find(&zips).Error
	fmt.Printf("There are %d records returned\n", len(zips))
	return
}
