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
		rp.ProcessSearchMetaData(rawres)
		break
	default:
		fmt.Printf("Processing of %s skipped\n", key)
	}
}

func (rp *ResultProcessor) processOrganicResults(rawres interface{}) (err error) {
	for _, k := range rawres.([]interface{}) {
		v := k.(map[string]interface{})
		fmt.Printf("Link: %s\n", v["link"].(string))
		fmt.Printf("Position: %d\n", int(v["position"].(float64)))
		fmt.Printf("Source: %s\n", v["source"].(string))
		err = rp.Dbcref.SaveUrlData(rp.Queryid, uint(OrganicResults), 0, uint(v["position"].(float64)), v["link"].(string), v["source"].(string))
	}
	return
}

func (rp *ResultProcessor) ProcessSearchMetaData(rawres interface{}) {

	for k, v := range rawres.(map[string]interface{}) {
		fmt.Printf("k: %s v: %v\n", k, v)
	}

}
