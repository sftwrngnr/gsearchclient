package data_importers

import (
	"bufio"
	"fmt"
	"github.com/schollz/progressbar/v3"
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
	//fmt.Printf("ZCImport.Init(%s)\n", dirname)
	z.inputfile = filepath.Join(dirname, "uszips.csv")
	//fmt.Printf("Verifying that %s exists.\n", z.inputfile)
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
	//fmt.Printf("ZCImport.Import(%s)\n", z.inputfile)
	lineCount := getLineCount(z.inputfile)
	bar := progressbar.Default(lineCount, "Zipcodes")
	//fmt.Println("number of lines:", lineCount)
	readFile, _ := os.Open(z.inputfile)
	defer readFile.Close()
	fs := bufio.NewScanner(readFile)
	fs.Split(bufio.ScanLines)
	defer bar.Close()
	numin := 0
	for fs.Scan() {
		if numin > 0 {
			// parse csv line
			_ = bar.Add(1)
			v := strings.Split(fs.Text(), ",")
			population, _ := strconv.ParseInt(rmvqts(v[8]), 10, 0)
			latitude, _ := strconv.ParseFloat(rmvqts(v[1]), 32)
			longitude, _ := strconv.ParseFloat(rmvqts(v[2]), 32)
			// Check to see if we have a valid state, if so, check to see if city exists. If city exists, get ID, otherwise insert
			stateid, ferr := z.checkState(rmvqts(v[4]))
			if ferr != nil {
				tState := sqldb.States{Abbrev: rmvqts(v[4]), Name: rmvqts(v[5])}
				z.DB.Create(&tState)
				stateid = tState.ID
			}
			cityid, cerr := z.checkorcreatecity(stateid, rmvqts(v[3]))
			if cerr != nil {
				continue
			}
			myZip := sqldb.Zipcode{Zipcode: rmvqts(v[0]),
				City:       cityid,
				State:      stateid,
				Latitude:   float32(latitude),
				Longitude:  float32(longitude),
				Population: uint(population)}
			z.DB.Create(&myZip)

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
