/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"github.com/sftwrngnr/gsearchclient/cmd"
)

var GitCommit string

func main() {
	fmt.Printf("gsearch version %s\n", GitCommit)
	cmd.Execute()
}
