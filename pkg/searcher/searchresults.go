package searcher

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"os"
	"path/filepath"
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
	LocalResults
	Pagination
)

type SearchResults struct {
	Results map[ResultType]interface{}
}

func NewSearchResults() (rval *SearchResults) {
	rval = &SearchResults{Results: make(map[ResultType]interface{})}
	return
}

func (sr *SearchResults) ProcessSearchData(qid uint, rawRes map[string]interface{}) {
	var rp = NewResultProcessor(qid, 0)
	resultkeys := []string{"serpapi_pagination", "search_metadata", "search_parameters", "search_information",
		"related_questions", "ai_overview", "organic_results", "local_results", "pagination"}
	fmt.Printf("rawRes is %d\n", len(rawRes))
	for i, key := range resultkeys {
		sr.Results[ResultType(i)] = rawRes[key]
		rp.ProcessResults(key, ResultType(i), rawRes[key])

	}
	fmt.Printf("ProcessSearchData::sr.Results[SerpapiPagination] is:%v\n", sr.Results[SerpapiPagination])
}

func (sr *SearchResults) StoreResults(searchResults map[string]interface{}) {
	fmt.Printf("StoreResults\n")
	bytes, _ := json.MarshalIndent(searchResults, "", "  ")
	myUUID := uuid.New()
	f, err := os.OpenFile(fmt.Sprintf("/tmp/searchresults_%s.json", myUUID.String()), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		return
	}
	defer f.Close()
	n, err := f.Write(bytes)
	fmt.Printf("wrote %d bytes\n", n)
	return
}

func (sr *SearchResults) GetJsonFiles(srcdir string) (flist []string, err error) {
	err = filepath.Walk(srcdir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			//fmt.Println(err)
			return nil
		}
		if !info.IsDir() && filepath.Ext(path) == ".json" {
			flist = append(flist, path)
		}
		return nil
	})
	return
}

func (sr *SearchResults) GetResults() (myRes map[string]interface{}, err error) {
	// get first json file that is found.
	fl, ferr := sr.GetJsonFiles("/tmp/")
	if ferr != nil {
		err = ferr
		return
	}
	if len(fl) == 0 {
		fmt.Println("No results found")
		return
	}
	//fmt.Printf("%v\n", fl)
	f, err := os.OpenFile(fl[0], os.O_RDONLY, 0600)
	if err != nil {
		//panic(err)
	}
	defer f.Close()
	err = json.NewDecoder(io.Reader(f)).Decode(&myRes)
	return
}

func (sr *SearchResults) AddResult(rt ResultType, result interface{}) (rval *SearchResults) {
	rval = sr
	rval.Results[rt] = result
	return
}
