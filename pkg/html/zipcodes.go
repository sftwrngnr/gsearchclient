package html

import (
	"fmt"
	"github.com/sftwrngnr/gsearchclient/pkg/system"
	. "maragu.dev/gomponents"
	"strconv"

	//hx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html"
)

func ZipCodes(qs string) Node {
	fmt.Printf("ZipCode handler called")
	var rval []Node
	fmt.Printf("ZipCode qs: %s\n", qs)
	myZips, err := system.GetSystemParams().Dbc.GetZipsForState(qs)
	if err != nil {
		fmt.Printf("ZipCodes: %v\n", err)
		return nil
	}
	for _, z := range myZips {
		rval = append(rval, Option(Value(strconv.Itoa(int(z.ID))), Text(z.Zipcode)))
	}
	return Select(rval...)
}
