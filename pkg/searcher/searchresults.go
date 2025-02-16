package searcher

type ResultType int

const (
	OrganicResultType ResultType = iota
	KnowledgeGraph
	LocalPack
	Pagination
	MaxResultTypes
)

type SearchResults struct {
	Results map[ResultType]interface{}
}

func NewSearchResults() (rval *SearchResults) {
	rval = &SearchResults{Results: make(map[ResultType]interface{})}
	return
}

func (sr *SearchResults) AddResult(rt ResultType, result interface{}) (rval *SearchResults) {
	rval = sr
	rval.Results[rt] = result
	return
}
