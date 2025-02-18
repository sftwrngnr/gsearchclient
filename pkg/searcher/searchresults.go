package searcher

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type ResultType int

const (
	SerpapiPagination ResultType = iota
	SearchMetaData
	SearchParams
	SearchInfo
	RelatedQuestions
	AIOverview
	OrganicResults
	Pagination
)

type SearchResults struct {
	Results map[ResultType]interface{}
}

func NewSearchResults() (rval *SearchResults) {
	rval = &SearchResults{Results: make(map[ResultType]interface{})}
	return
}

func (sr *SearchResults) ProcessSearchData(rawRes map[string]interface{}) {
	resultkeys := []string{"serpapi_pagination", "search_metadata", "search_parameters", "search_information",
		"related_questions", "ai_overview", "organic_results", "pagination"}
	fmt.Printf("rawRes is %d\n", len(rawRes))
	for i, key := range resultkeys {
		sr.Results[ResultType(i)] = rawRes[key]
	}
	fmt.Printf("ProcessSearchData::sr.Results[SerpapiPagination] is:%v\n", sr.Results[SerpapiPagination])
}

func (sr *SearchResults) StoreResults(searchResults map[string]interface{}) {
	fmt.Printf("StoreResults\n")
	bytes, _ := json.MarshalIndent(searchResults, "", "  ")
	f, err := os.OpenFile("/tmp/searchresults.json", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	n, err := f.Write(bytes)
	fmt.Printf("wrote %d bytes\n", n)
	return
}

func (sr *SearchResults) GetResults() (myRes map[string]interface{}) {
	f, err := os.OpenFile("/tmp/searchresults.json", os.O_RDONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = json.NewDecoder(io.Reader(f)).Decode(&myRes)
	if err != nil {
		panic(err)
	}
	return
}

func (sr *SearchResults) AddResult(rt ResultType, result interface{}) (rval *SearchResults) {
	rval = sr
	rval.Results[rt] = result
	return
}
