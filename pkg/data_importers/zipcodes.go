package data_importers

import (
	"bufio"
	"fmt"
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
			fmt.Printf("%v\n", v)
		} else {
			fmt.Printf("Zipcodes.Import: Header line %s.\n", fs.Text())
		}
		numin++
	}
	return numin, nil

}
