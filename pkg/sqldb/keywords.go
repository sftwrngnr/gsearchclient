package sqldb

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type Keywords struct {
	gorm.Model
	ID      uint   `gorm:"primary_key"`
	Keyword string `gorm:"column:keyword"`
	Req     bool   `gorm:"column:req"`
}

func (dbc *DBConnData) GetKeywords() (keywrds []Keywords, err error) {
	err = dbc.DB.Find(&keywrds).Error
	return
}

func (dbc *DBConnData) GetReqKeywords() (reqd []string, err error) {
	var keywords []Keywords

	err = dbc.DB.Where("req = true").Find(&keywords).Error
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	for _, kw := range keywords {
		reqd = append(reqd, kw.Keyword)
	}
	fmt.Printf("%v\n", reqd)
	return
}

func (dbc *DBConnData) GetMatchingKeywords(kwl []uint, KList *[]Keywords) (err error) {
	if len(kwl) == 0 {
		err = errors.New("At least one keyword must be selected.")
		return
	}
	err = dbc.DB.Where("id in ?", kwl).Find(KList).Error
	fmt.Printf("There are %d records returned\n", len(*KList))
	return
}

func (dbc *DBConnData) DeleteKeywords() error {
	return dbc.DB.Exec("delete from keywords").Error
}
