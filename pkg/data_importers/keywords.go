package data_importers

import (
	"bufio"
	"fmt"
	"github.com/schollz/progressbar/v3"
	"github.com/sftwrngnr/gsearchclient/pkg/sqldb"
	"github.com/sftwrngnr/gsearchclient/pkg/system"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"strings"
)

type KWImporter struct {
	inputfile string `json:"input"`
	DB        *gorm.DB
}

func (kwi *KWImporter) Init(dirname string) bool {
	kwi.inputfile = filepath.Join(dirname, "keywords.csv")
	//fmt.Printf("Verifying that %s exists.\n", z.inputfile)
	_, err := os.Stat(kwi.inputfile)
	if err != nil {
		fmt.Printf("File %s doesn't exist.", kwi.inputfile)
		return false
	}
	return true

}

func (kwi *KWImporter) Import() (numin int, err error) {
	lineCount := getLineCount(kwi.inputfile)
	bar := progressbar.Default(lineCount, "Keywords")
	//fmt.Println("number of lines:", lineCount)
	readFile, _ := os.Open(kwi.inputfile)
	system.GetSystemParams().Dbc.DeleteKeywords()
	defer readFile.Close()
	fs := bufio.NewScanner(readFile)
	fs.Split(bufio.ScanLines)
	defer bar.Close()
	for fs.Scan() {
		_ = bar.Add(1)
		v := strings.Split(fs.Text(), ",")
		for _, kw := range v {
			kw = strings.TrimSpace(kw)
			myKw := &sqldb.Keywords{Keyword: kw}
			kwi.DB.Create(myKw)
			numin++
		}
	}
	return
}
