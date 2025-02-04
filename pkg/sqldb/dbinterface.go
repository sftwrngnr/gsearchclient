package sqldb

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type DBConnData struct {
	DBName   string
	Host     string
	User     string
	Password string
	Port     int16
	db       *sql.DB
}

func (dbc *DBConnData) CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func (dbc *DBConnData) Ping() error {
	return dbc.db.Ping()
}

func (dbc *DBConnData) Connect() (bool, error) {
	var rval bool
	var err error

	if (dbc.DBName == "") || (dbc.User == "") || (dbc.Password == "") {
		err = errors.New("Missing database connection parameters")
	}
	if err == nil {
		psqlconn := fmt.Sprintf("host=%s port=%d database=%s user=%s password=%s  sslmode=disable", dbc.Host, dbc.Port, dbc.DBName, dbc.User, dbc.Password)

		// open database
		dbc.db, err = sql.Open("pgx", psqlconn)
		dbc.CheckError(err)

		// check db
		err = dbc.Ping()
		dbc.CheckError(err)

		fmt.Println("Connected!")
		rval = true
	}

	return rval, err

}

func (dbc *DBConnData) DeleteTableRecs(tablename string) error {
	sqlstmt := fmt.Sprintf("DELETE FROM %s", tablename)
	_, err := dbc.db.Exec(sqlstmt)
	return err
}

func (dbc *DBConnData) Close() {
	dbc.db.Close()
}
