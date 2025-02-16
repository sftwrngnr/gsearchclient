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
}

func (dbc *DBConnData) GetKeywords() (keywords []Keywords, err error) {
	err = dbc.DB.Find(&keywords).Error
	return
}

func (dbc *DBConnData) GetKeywordList(kwds *[]Keywords) error {
	return dbc.DB.Find(kwds).Error
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
