package searcher

import "fmt"

type ResultProcessor struct {
	Queryid uint
	Seqid   uint
}

func NewResultProcessor(qry uint, seq uint) *ResultProcessor {
	return &ResultProcessor{
		Queryid: qry,
		Seqid:   seq,
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

func (rp *ResultProcessor) processOrganicResults(rawres interface{}) {
	for _, k := range rawres.([]interface{}) {
		v := k.(map[string]interface{})
		fmt.Printf("Link: %s\n", v["link"].(string))
		fmt.Printf("Position: %d\n", int(v["position"].(float64)))
		fmt.Printf("Source: %s\n", v["source"].(string))
	}
}

func (rp *ResultProcessor) ProcessSearchMetaData(rawres interface{}) {

	for k, v := range rawres.(map[string]interface{}) {
		fmt.Printf("k: %s v: %v\n", k, v)
	}

}
