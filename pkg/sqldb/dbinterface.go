package sqldb

import (
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConnData struct {
	DBName   string
	Host     string
	User     string
	Password string
	Port     int16
	DB       *gorm.DB
}

func (dbc *DBConnData) CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func (dbc *DBConnData) Ping() error {
	return nil
}

func (dbc *DBConnData) Connect() error {
	var err error

	if (dbc.DBName == "") || (dbc.User == "") || (dbc.Password == "") {
		err = errors.New("Missing database connection parameters")
	}
	if err == nil {
		dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable", dbc.User, dbc.Password, dbc.DBName, dbc.Host, dbc.Port)

		// open database
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			fmt.Println("Error connecting to database")
			return err
		}
		fmt.Println("Connected!")
		dbc.DB = db
	}

	return err

}

func (dbc *DBConnData) DeleteTableRecs(tablename string) error {
	return nil
}

func (dbc *DBConnData) Close() {

}
