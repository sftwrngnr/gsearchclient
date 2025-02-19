package sqldb

import (
	"fmt"
	"gorm.io/gorm"
)

type Query_results struct {
	gorm.Model
	ID         uint   `gorm:"primary_key"`
	Query_id   uint   `gorm:"column:query_id"`
	Resultseq  uint   `gorm:"column:resultseq"`
	Resulttype uint   `gorm:"column:result_type"`
	Result     []byte `gorm:"type:jsonb" json:"result"`
}

func (dbc *DBConnData) ProcessQry_results(queryid uint, resultseq uint, resulttype uint, result []byte) (myerr error) {
	if result == nil {
		return
	}
	myqr := &Query_results{Query_id: queryid, Resultseq: resultseq, Resulttype: resulttype, Result: result}
	myerr = dbc.DB.Create(myqr).Error
	if myerr != nil {
		fmt.Printf("Database error %s\n", myerr.Error())
	}
	return
}
