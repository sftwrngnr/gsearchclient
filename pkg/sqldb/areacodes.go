package sqldb

import (
	"fmt"
	"gorm.io/gorm"
)

type Areacodes struct {
	gorm.Model
	ID        uint    `gorm:"primary_key"`
	Code      string  `gorm:"column:code"`
	State     uint    `gorm:"column:state"`
	Latitude  float32 `gorm:"column:latitude"`
	Longitude float32 `gorm:"column:longitude"`
}

type ConsolidateAreacodes struct {
	Code   string
	IDList []uint
}

func (dbc *DBConnData) GetAreaCodesForState(abbrv string) ([]ConsolidateAreacodes, error) {
	rval := []Areacodes{}
	myStId, err := dbc.GetStateId(abbrv)
	if err != nil {
		return nil, err
	}
	fmt.Printf("myStId: %v\n", myStId)
	myarq := Areacodes{State: myStId}
	err = dbc.DB.Find(&rval, &myarq).Error
	fmt.Printf("There are %d records returned\n", len(rval))
	return consolidate(rval), err
}

func consolidate(inval []Areacodes) []ConsolidateAreacodes {
	rVal := []ConsolidateAreacodes{}
	if len(inval) > 0 {
		var newCa ConsolidateAreacodes
		for _, ac := range inval {
			if newCa.Code != ac.Code {
				if newCa.Code != "" {
					rVal = append(rVal, newCa)
				}
				newCa.Code = ac.Code
				newCa.IDList = []uint{}
				newCa.IDList = append(newCa.IDList, ac.ID)
			} else {
				newCa.IDList = append(newCa.IDList, ac.ID)
			}
		}
		rVal = append(rVal, newCa)

	}
	return rVal
}
