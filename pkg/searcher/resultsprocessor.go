package searcher

import (
	"fmt"
	"github.com/sftwrngnr/gsearchclient/pkg/sqldb"
	"github.com/sftwrngnr/gsearchclient/pkg/system"
)

type ResultProcessor struct {
	Queryid uint
	Seqid   uint
	Dbcref  *sqldb.DBConnData
}

func NewResultProcessor(qry uint, seq uint) *ResultProcessor {
	return &ResultProcessor{
		Queryid: qry,
		Seqid:   seq,
		Dbcref:  system.GetSystemParams().Dbc,
	}
}

func (rp *ResultProcessor) ProcessResults(key string, rt ResultType, rawres interface{}) {
	switch rt {
	case OrganicResults:
		//fmt.Printf("Processing key: %s, rawres: %v\n", key, rawres)
		rp.processOrganicResults(rawres)
		break
	case SearchMetaData:
		rp.ProcessSearchMetaData(rawres.(map[string]interface{}))
		break
	default:
		fmt.Printf("Processing of %s skipped\n", key)
	}
}

func (rp *ResultProcessor) processOrganicResults(rawres interface{}) (err error) {
	for _, k := range rawres.([]interface{}) {
		v := k.(map[string]interface{})
		err = rp.Dbcref.SaveUrlData(rp.Queryid, uint(OrganicResults), 0, uint(v["position"].(float64)), v["link"].(string), v["source"].(string))
	}
	return
}

func (rp *ResultProcessor) ProcessSearchMetaData(rawres map[string]interface{}) (err error) {
	fmt.Printf("ProcessedSearchMetaData: %d\n", rp.Queryid)
	err = rp.Dbcref.SaveSearchMetaData(rp.Queryid,
		rawres["status"].(string),
		rawres["id"].(string),
		rawres["total_time_taken"].(float64),
		rawres["created_at"].(string),
		rawres["google_url"].(string),
		rawres["json_endpoint"].(string),
		rawres["processed_at"].(string),
		rawres["raw_html_file"].(string))
	if err != nil {
		fmt.Printf("Error saving search meta data: %s\n", err)
		return err
	}
	return
}

//func (rp *ResultProcessor) ProcessOrganicResults(rawres interface{}) {}
