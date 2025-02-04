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

type States struct {
	inputfile string
	DB        *gorm.DB
}

func (s *States) Init(dirname string) bool {
	fmt.Printf("States.Init(%s)\n", dirname)
	s.inputfile = filepath.Join(dirname, "50States.csv")
	fmt.Printf("Verifying that %s exists.\n", s.inputfile)
	_, err := os.Stat(s.inputfile)
	if err != nil {
		fmt.Printf("File %s doesn't exist.", s.inputfile)
		return false
	}
	return true
}

func (s *States) Import() (int, error) {
	fmt.Printf("States.Import(%s)\n", s.inputfile)
	readFile, err := os.Open(s.inputfile)
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
			//fmt.Printf("[%d]: %s (%s) [%s,%s]\n", numin, v[1], v[0], v[2], v[3])
			newState := sqldb.States{Abbrev: v[1], Name: v[0], Capitol: v[2], Region: v[3]}
			s.DB.Create(&newState)
			newCity := sqldb.Cities{Name: v[2], State: newState.ID}
			s.DB.Create(&newCity)
		}
		numin++
	}
	return numin, nil
}
