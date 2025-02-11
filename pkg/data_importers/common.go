package data_importers

import (
	"bufio"
	"os"
)

// Common routines

func getLineCount(fname string) (linecount int64) {
	readFile, err := os.Open(fname)
	defer readFile.Close()
	if err != nil {
		//fmt.Println(err)
		return
	}
	fs := bufio.NewScanner(readFile)
	fs.Split(bufio.ScanLines)
	for fs.Scan() {
		linecount++
	}
	return
}
