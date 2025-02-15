package sqldb

import (
	"gorm.io/gorm"
)

type qry_results struct {
	gorm.Model
	ID         uint   `gorm:"primary_key"`
	Query_id   uint   `gorm:"column:query_id"`
	Resultseq  uint   `gorm:"column:resultseq"`
	Resulttype uint   `gorm:"column:result_type"`
	Result     []byte `gorm:"type:jsonb" json:"result"`
}
