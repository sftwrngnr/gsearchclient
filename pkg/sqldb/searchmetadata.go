package sqldb

import (
	"fmt"
	"gorm.io/gorm"
)

/*
CREATE TABLE public.search_metadata (

	id bigint NOT NULL,
	query_id bigint,
	status character varying,
	searchid character varying,
	total_time_taken real,
	screated_at date,
	google_url character varying,
	json_endpoint character varying,
	processed_at date,
	raw_html_file character varying,
	created_at date,
	updated_at date,
	deleted_at date

);
*/
type SearchMetadata struct {
	gorm.Model
	ID             uint    `gorm:"primary_key"`
	QueryId        uint    `gorm:"column:query_id"`
	Status         string  `gorm:"column:status"`
	Searchid       string  `gorm:"column:searchid"`
	TotalTimeTaken float64 `gorm:"column:total_time_taken"`
	ScreatedAt     string  `gorm:"column:screated_at"`
	GoogleUrl      string  `gorm:"column:google_url"`
	JsonEndpoint   string  `gorm:"column:json_endpoint"`
	ProcessedAt    string  `gorm:"column:processed_at"`
	RawHtmlFile    string  `gorm:"column:raw_html_file"`
}

func (dbc *DBConnData) SaveSearchMetaData(qryid uint, status string, srchid string, searchtm float64, screated string,
	googleurl string, jsonendpoint string, processedat string, rawhtmlfile string) (err error) {
	myqs := &SearchMetadata{QueryId: qryid, Status: status, Searchid: srchid, TotalTimeTaken: searchtm,
		ScreatedAt: screated, GoogleUrl: googleurl, JsonEndpoint: jsonendpoint,
		ProcessedAt: processedat, RawHtmlFile: rawhtmlfile}
	err = dbc.DB.Create(myqs).Error
	if err != nil {
		fmt.Println(err)
	}
	return
}
