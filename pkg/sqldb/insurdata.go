package sqldb

import "gorm.io/gorm"

type insurdata struct {
	gorm.Model
	ID       uint   `gorm:"primary_key"`
	Name     string `gorm:"column:name"`
	Jobtitle string `gorm:"column:jobtitle"`
	Address  string `gorm:"column:address"`
	Phone    string `gorm:"column:phone"`
	Email    string `gorm:"column:email"`
	Zipcode  uint   `gorm:"column:zipcode"`
	Source   string `gorm:"column:source"`
}

func (dbc *DBConnData) CreateDeltaData(name string, job string, addy string, phone string, email string, zipcode uint) error {
	myDeltadata := &insurdata{Name: name, Jobtitle: job, Address: addy, Phone: phone, Email: email, Zipcode: zipcode, Source: "Delta"}
	return dbc.DB.Create(myDeltadata).Error
}

func (dbc *DBConnData) CreateCignaData(name string, job string, addy string, phone string, email string, zipcode uint) error {
	myDeltadata := &insurdata{Name: name, Jobtitle: job, Address: addy, Phone: phone, Email: email, Zipcode: zipcode, Source: "Cigna"}
	return dbc.DB.Create(myDeltadata).Error
}
