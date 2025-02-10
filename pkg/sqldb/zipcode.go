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

func (dbc *DBConnData) GetZipsForState(abbrv string) ([]Zipcode, error) {
	rval := []Zipcode{}
	myStId, err := dbc.GetStateId(abbrv)
	if err != nil {
		return rval, err
	}
	fmt.Printf("myStId: %v\n", myStId)
	myzcq := Zipcode{State: myStId}
	err = dbc.DB.Find(&rval, &myzcq).Error
	fmt.Printf("There are %d records returned\n", len(rval))
	return rval, err
}
