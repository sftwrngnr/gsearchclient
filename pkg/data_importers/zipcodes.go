package data_importers

import (
	"bufio"
	"fmt"
	"github.com/sftwrngnr/gsearchclient/pkg/sqldb"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"strings"
)

// This will read the csv zip code file.

type Zipcodes struct {
	inputfile string
	DB        *gorm.DB
}

// "zip","lat","lng","city","state_id","state_name","zcta","parent_zcta","population","density","county_fips","county_name","county_weights","county_names_all","county_fips_all","imprecise","military","timezone"
// 0,3, 4, 5 [state name], 8
func (z *Zipcodes) Init(dirname string) bool {
	fmt.Printf("Zipcodes.Init(%s)\n", dirname)
	z.inputfile = filepath.Join(dirname, "uszips.csv")
	fmt.Printf("Verifying that %s exists.\n", z.inputfile)
	_, err := os.Stat(z.inputfile)
	if err != nil {
		fmt.Printf("File %s doesn't exist.", z.inputfile)
		return false
	}
	return true
}

func (z *Zipcodes) Import() (int, error) {
	rmvqts := func(s string) string {
		if len(s) > 0 && s[0] == '"' {
			s = s[1:]
		}
		if len(s) > 0 && s[len(s)-1] == '"' {
			s = s[:len(s)-1]
		}
		return s
	}
	fmt.Printf("Zipcodes.Import(%s)\n", z.inputfile)
	readFile, err := os.Open(z.inputfile)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	fs := bufio.NewScanner(readFile)
	fs.Split(bufio.ScanLines)
	numin := 0
	for fs.Scan() {
		if numin > 0 {
			// parse csv line

			v := strings.Split(fs.Text(), ",")
			// Check to see if we have a valid state, if so, check to see if city exists. If city exists, get ID, otherwise insert
			stateid := z.checkState(rmvqts(v[4]))
			if stateid == -1 {
				cityid := checkorcreatecity(stateid, rmvqts(v[3]))
				myZip := sqldb.Zipcode{Zipcode: rmvqts(v[0]),
					City:  cityid,
					State: stateid}
				z.DB.Create(&myZip)
			}

		} else {
			if numin == 0 {
				fmt.Printf("Zipcodes.Import: Header line %s.\n", fs.Text())
			}
		}
		numin++
	}
	return numin, nil

}

func (z *Zipcodes) checkState(instate string) int {
	myState := sqldb.States{Abbrev: instate}
	myState := z.DB.First(&myState)

}
