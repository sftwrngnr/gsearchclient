package html

import (
	"fmt"
	"github.com/sftwrngnr/gsearchclient/pkg/system"
	. "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html"
	"strconv"
)

func QueryTransfer(items []string) Node {
	Companies := getCompanies()
	Crawlers := getCrawlers(0)
	Campaigns := getCampaigns(0)
	return page("QueryTransfer",
		Head(Script(Src("https://cdn.tailwindcss.com?plugins=forms,typography")),
			Script(Src("https://unpkg.com/htmx.org"))),
		H1(Text("Market Research Crawler")),
		Form(
			H2(Text("Query Transfer")),
			Text("Company:"), Select(Companies...), Br(),
			Div(ID("qrytransferupdate"),
				Text("Campaigns:"), Select(Campaigns...), Br(),
				Text("Crawlers"), Select(Crawlers...), Br()),
			QueryButton(),
			Img(ID("spinner"), Class("htmx-indicator"), Src("https://unpkg.com/html-spinner")),
			Hr(Style("border: 5px solid blue; border-radius: 5px")),
			Div(Raw("<B>"), Text("Transferred urls to crawler:"), Raw("</B>")),
			Div(ID("transfer_res")),
			Hr(Style("border: 5px solid blue; border-radius: 5px")),
		),
	)

}

func getCrawlers(id uint) []Node {
	var rval []Node
	rval = append(rval, Name("Crawler"))
	if id == 0 {
		rval = append(rval, Option(Value("Name"), Text("None")))
		return rval
	}
	crawlers, err := system.GetSystemParams().Dbc.GetCompanyCrawlers(id)
	if err != nil {
		fmt.Printf("error getting crawlers: %v\n", err)
		rval = append(rval, Option(Value("Name"), Text("None")))
		return rval
	}
	for _, crawler := range crawlers {
		fmt.Printf("%s\n", crawler.Name)
		rval = append(rval, Option(Value(fmt.Sprintf("%d", crawler.ID)), Text(crawler.Name)))
	}
	return rval
}

func getCampaigns(id uint) []Node {
	var rval []Node
	rval = append(rval, Name("Campaign"))
	if id == 0 {
		rval = append(rval, Option(Value("Name"), Text("None")))
		return rval
	}
	campaigns, err := system.GetSystemParams().Dbc.GetCompanyCampaigns(id)
	if err != nil {
		fmt.Printf("error getting campaigns: %v\n", err)
		rval = append(rval, Option(Value("Name"), Text("None")))
		return rval
	}
	for _, campaign := range campaigns {
		fmt.Printf("%s\n", campaign.Name)
		rval = append(rval, Option(Value(fmt.Sprintf("%d", campaign.ID)), Text(campaign.Name)))
	}

	return rval
}

func getCompanies() []Node {
	var rval []Node
	complist, err := system.GetSystemParams().Dbc.GetCompanyList()
	if err != nil {
		fmt.Printf("error getting companies: %v\n", err)
		return nil
	}
	rval = append(rval, Name("Company"), hx.Get("/getcrawlers"), hx.Target("#qrytransferupdate"))
	rval = append(rval, Option(Value("Name"), Text("None")))
	for _, company := range complist {
		fmt.Printf("%s\n", company.Name)
		rval = append(rval, Option(Value(fmt.Sprintf("%d", company.ID)), Text(company.Name)))
	}
	return rval
}

func GetDataForComapny(mymap map[string][]string) Node {
	var rval Node
	ti, err := strconv.Atoi(mymap["Company"][0])
	if err != nil {
		return rval
	}
	var compid uint = uint(ti)
	fmt.Printf("%v\n", mymap["Company"][0])

	Crawlers := getCrawlers(compid)
	Campaigns := getCampaigns(compid)
	rval = Div(ID("qrytransferupdate"),
		Text("Campaigns:"), Select(Campaigns...), Br(),
		Text("Crawlers"), Select(Crawlers...), Br())

	return rval
}

func QueryButton() (rval Node) {
	rval = Button(Type("submit"), Text(`Transfer`),
		Class("rounded-md border border-transparent bg-blue-600 px-4 py-2 text-sm font-medium text-white shadow-sm hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"),
		hx.Target("#transfer_res"),
		hx.Include("#Company"),
		hx.Include("#Campaign"),
		hx.Include("#Crawler"),
		hx.Post("/exectransfer"),
	)
	return
}
