package sqldb

import (
	"gorm.io/gorm"
)

type Urls struct {
	gorm.Model
	ID          uint   `gorm:"primary_key"`
	QueryId     uint   `gorm:"column:query_id"`
	QuerySrc    uint   `gorm:"column:query_src"`
	SeqId       uint   `gorm:"column:seq_id"`
	Position    uint   `gorm:"column:position"`
	Url         string `gorm:"column:url"`
	Source      string `gorm:"column:source"`
	Transferred bool   `gorm:"column:transferred"`
}

func (dbc *DBConnData) SaveUrlData(qryid uint, qrysrc uint, seqid uint, posit uint, url string, source string) error {
	Url := &Urls{QueryId: qryid, QuerySrc: uint(qrysrc), SeqId: seqid, Position: posit, Url: url, Source: source}
	return dbc.DB.Create(Url).Error
}

func (dbc *DBConnData) GetUrls(transfrd bool) (urls []Urls, err error) {
	err = dbc.DB.Where("transferred = ?", transfrd).Find(&urls).Error
	return
}

func (dbc *DBConnData) UpdateRec(url *Urls) error {
	return dbc.DB.Save(url).Error
}
