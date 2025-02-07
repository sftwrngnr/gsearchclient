package data_importers

import (
	"encoding/json"
	"fmt"
	"github.com/sftwrngnr/gsearchclient/pkg/sqldb"
	"gorm.io/gorm"
	"os"
	"path/filepath"
)

type ACImport struct {
	inputfile string
	DB        *gorm.DB
}

type ACData struct {
	AreaCode  int     `json:"area-code"`
	City      string  `json:"city"`
	State     string  `json:"state"`
	Country   string  `json:"country"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// "zip","lat","lng","city","state_id","state_name","zcta","parent_zcta","population","density","county_fips","county_name","county_weights","county_names_all","county_fips_all","imprecise","military","timezone"
// 0,3, 4, 5 [state name], 8
func (a *ACImport) Init(dirname string) bool {
	fmt.Printf("ACImport.Init(%s)\n", dirname)
	a.inputfile = filepath.Join(dirname, "area-codes-usa.json")
	fmt.Printf("Verifying that %s exists.\n", a.inputfile)
	_, err := os.Stat(a.inputfile)
	if err != nil {
		fmt.Printf("File %s doesn't exist.", a.inputfile)
		return false
	}
	return true
}

func (a *ACImport) Import() (int, error) {
	fmt.Printf("ACImport.Import(%s)\n", a.inputfile)
	var acArr []ACData
	content, err := os.ReadFile(a.inputfile)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	json.Unmarshal(content, &acArr)
	fmt.Printf("Processing JSON data\n")
	var iCount int
	for _, ac := range acArr {
		if ac.Country != "US" {
			continue
		}
		mySt := a.findState(ac.State)
		if mySt == 0 {
			fmt.Printf("Couldn't find %s\n", ac.State)
			continue
		}
		myAc := a.insertAreaCode(fmt.Sprintf("%d", ac.AreaCode), mySt)
		myCity := a.findOrInsertCity(ac.City, mySt)
		// Now insert into correlation table
		a.insertACCity(myAc, myCity)
		iCount++
	}
	return iCount, nil
}

func (a *ACImport) findState(s string) uint {
	mySt := &sqldb.States{Name: s}
	a.DB.Where(mySt).First(mySt)
	return mySt.ID
}

func (a *ACImport) findOrInsertCity(c string, s uint) uint {
	myCity := &sqldb.Cities{Name: c, State: s}
	a.DB.Where(myCity).First(myCity)
	if myCity.ID != 0 {
		fmt.Printf("Found %s\n", myCity.Name)
	} else {
		a.DB.Create(myCity)
	}
	return myCity.ID
}

func (a *ACImport) insertACCity(acode uint, ccode uint) uint {
	myACCity := &sqldb.Cityareacodes{Areacode: acode, City: ccode}
	a.DB.Where(myACCity).First(&myACCity)
	if myACCity.ID == 0 {
		a.DB.Create(&myACCity)
	}
	return myACCity.ID
}

func (a *ACImport) insertAreaCode(acode string, st uint) uint {
	myACode := sqldb.Areacodes{Code: acode, State: st}
	a.DB.Where(myACode).First(&myACode)
	if myACode.ID == 0 {
		a.DB.Create(&myACode)
	}
	return myACode.ID
}
