package sqldb

import "gorm.io/gorm"

type Cities struct {
	gorm.Model
	ID    uint   `gorm:"primary_key"`
	Name  string `gorm:"column name"`
	State uint   `gorm:"column state"`
}

func (dbc *DBConnData) GetCitiesForState(abbrv string) (cities []Cities, err error) {
	stid, serr := dbc.GetStateId(abbrv)
	if serr != nil {
		err = serr
		return
	}

	err = dbc.DB.Where("state = ?", stid).Find(&cities).Error
	return
}
