package sqldb

import "gorm.io/gorm"

type deltadata struct {
	gorm.Model
	ID       uint   `gorm:"primary_key"`
	Name     string `gorm:"column:name"`
	Jobtitle string `gorm:"column:jobtitle"`
	Address  string `gorm:"column:address"`
	Phone    string `gorm:"column:phone"`
	Email    string `gorm:"column:email"`
	Zipcode  uint   `gorm:"column:zipcode"`
}

func (dbc *DBConnData) CreateDeltaData(name string, job string, addy string, phone string, email string, zipcode uint) error {
	myDeltadata := &deltadata{Name: name, Jobtitle: job, Address: addy, Phone: phone, Email: email, Zipcode: zipcode}
	return dbc.DB.Create(myDeltadata).Error
}
