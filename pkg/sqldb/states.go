package sqldb

import "gorm.io/gorm"

type States struct {
	gorm.Model
	ID      uint   `gorm:"primary_key"`
	Abbrev  string `gorm:"column:abbrev"`
	Name    string `gorm:"column:name"`
	Capitol string `gorm:"column:capitol"`
	Region  string `gorm:"column:region"`
}

func (dbc *DBConnData) GetAllStates() ([]States, error) {
	var rval []States
	err := dbc.DB.Find(&rval).Error
	return rval, err
}

//func (dbc *DBConnData) GetState(id string) (States, error) {}

func (dbc *DBConnData) GetStateId(abbrev string) (uint, error) {
	var state States
	err := dbc.DB.First(&state, "abbrev = ?", abbrev).Error
	return state.ID, err

}
