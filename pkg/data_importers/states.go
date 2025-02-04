package data_importers

import "fmt"

type States struct {
}

func (s States) Init(dirname string) bool {
	fmt.Printf("States.Init(%s)\n", dirname)
	return true
}

func (s States) Import() (int, error) {
	fmt.Printf("States.Import()\n")

}
