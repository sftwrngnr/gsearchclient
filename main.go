package main

import (
	"context"
	"fmt"
	googlesearch "github.com/rocketlaunchr/google-search"
)

func main() {
	ctx := context.Background()
	opt := googlesearch.SearchOptions{CountryCode: "us"}
	lret, err := googlesearch.Search(ctx, "cars for sale in queen creek,az", opt)
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Println(lret)
	fmt.Println(googlesearch.Search(ctx, "cars for sale in Toronto, Canada"))
}
