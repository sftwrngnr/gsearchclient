package data_importers

import (
	"bufio"
	"fmt"
	"github.com/sftwrngnr/gsearchclient/pkg/sqldb"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// This will read the csv zip code file.

type ZCImport struct {
	inputfile string
	DB        *gorm.DB
}

// "zip","lat","lng","city","state_id","state_name","zcta","parent_zcta","population","density","county_fips","county_name","county_weights","county_names_all","county_fips_all","imprecise","military","timezone"
// 0,3, 4, 5 [state name], 8
func (z *ZCImport) Init(dirname string) bool {
	fmt.Printf("ZCImport.Init(%s)\n", dirname)
	z.inputfile = filepath.Join(dirname, "uszips.csv")
	fmt.Printf("Verifying that %s exists.\n", z.inputfile)
	_, err := os.Stat(z.inputfile)
	if err != nil {
		fmt.Printf("File %s doesn't exist.", z.inputfile)
		return false
	}
	return true
}

func (z *ZCImport) Import() (int, error) {
	rmvqts := func(s string) string {
		if len(s) > 0 && s[0] == '"' {
			s = s[1:]
		}
		if len(s) > 0 && s[len(s)-1] == '"' {
			s = s[:len(s)-1]
		}
		return s
	}
	fmt.Printf("ZCImport.Import(%s)\n", z.inputfile)
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
			population, cerr := strconv.ParseInt(rmvqts(v[8]), 10, 0)
			if cerr != nil {
				fmt.Println(cerr)
				population = 0
			}
			// Check to see if we have a valid state, if so, check to see if city exists. If city exists, get ID, otherwise insert
			stateid, ferr := z.checkState(rmvqts(v[4]))
			if ferr != nil {
				fmt.Println(ferr)
				fmt.Println("Inserting state")
				tState := sqldb.States{Abbrev: rmvqts(v[4]), Name: rmvqts(v[5])}
				z.DB.Create(&tState)
				stateid = tState.ID
			}
			cityid, cerr := z.checkorcreatecity(stateid, rmvqts(v[3]))
			if cerr != nil {
				fmt.Println(cerr)
				fmt.Println("skipping city")
				continue
			}
			myZip := sqldb.Zipcode{Zipcode: rmvqts(v[0]),
				City:       cityid,
				State:      stateid,
				Population: uint(population)}
			z.DB.Create(&myZip)

		} else {
			if numin == 0 {
				fmt.Printf("ZCImport.Import: Header line %s.\n", fs.Text())
			}
		}
		numin++
	}
	return numin, nil

}

func (z *ZCImport) checkState(instate string) (uint, error) {
	myState := sqldb.States{Abbrev: instate}
	result := z.DB.Where("abbrev = ?", instate).First(&myState)
	if result.Error != nil {
		return 0, result.Error
	}
	return myState.ID, nil

}

func (z *ZCImport) checkorcreatecity(stateid uint, cityid string) (uint, error) {
	myCity := sqldb.Cities{State: stateid}
	result := z.DB.Where("name = ?", cityid).First(&myCity, "state = ?", stateid)
	if result.Error != nil {
		// Create city
		myCity := sqldb.Cities{Name: cityid, State: stateid}
		z.DB.Create(&myCity)
		return myCity.ID, nil
	}

	return myCity.ID, nil

}
