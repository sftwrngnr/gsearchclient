package http

import (
	"fmt"
	"github.com/sftwrngnr/gsearchclient/pkg/crawler"
	"github.com/sftwrngnr/gsearchclient/pkg/html"
	"github.com/sftwrngnr/gsearchclient/pkg/system"
	. "maragu.dev/gomponents"
	html2 "maragu.dev/gomponents/html"
	ghttp "maragu.dev/gomponents/http"
	"net/http"
	"slices"
	"strings"
	"time"
)

func Home2(mux *http.ServeMux) {
	mux.Handle("GET /H2", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (Node, error) {
		return html.HomePage2(), nil
	}))

}

func Home(mux *http.ServeMux) {
	mux.Handle("GET /", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (Node, error) {
		return html.HomePage(), nil
	}))
}

func ZipCodes(mux *http.ServeMux) {
	mux.Handle("GET /zipcodes", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (Node, error) {
		qs := r.URL.Query().Get("state")
		fmt.Printf("Received zipcodes request for state %s\n", qs)
		return html.ZipCodes(qs), nil
	}))
}

func ExecTransfer(mux *http.ServeMux) {
	mux.Handle("POST /exectransfer", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (Node, error) {
		fmt.Printf("Received exectransfer request\n")
		err := r.ParseForm()
		if err != nil {
			return nil, err
		}
		fmt.Printf("%v\n", r.Form)
		turls, err := crawler.TransferURLS(r.Form)
		if err != nil {
			return html2.Div(), err
		}

		var nArr []Node
		for _, turl := range turls {
			nArr = append(nArr, html2.Li(Text(turl)))
		}
		return html2.Nav(html2.Ul(nArr...)), err
	}))
}

func GenQry(mux *http.ServeMux) {
	mux.Handle("POST /genqry", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (Node, error) {
		err := r.ParseForm()
		if err != nil {

			fmt.Printf("Error with ParseForm %s\n", err.Error())
			return nil, err
		}
		return html.GenQry(r.Form)
	}))
}

func QueryTransfer(mux *http.ServeMux) {
	mux.Handle("GET /qrytransfer", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (Node, error) {
		return html.QueryTransfer(nil), nil
	}))
}

func GetCrawlers(mux *http.ServeMux) {
	mux.Handle("GET /getcrawlers", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (Node, error) {
		err := r.ParseForm()
		if err != nil {
			fmt.Printf("Error with ParseForm %s\n", err.Error())
			return nil, err
		}
		fmt.Printf("%v\n", r.Form)
		kl := make([]string, 0)
		for k, _ := range r.Form {
			kl = append(kl, k)
		}
		if slices.Contains(kl, "Company") && !slices.Contains(kl, "Campaign") {
			return html.GetDataForComapny(r.Form), nil
		}
		if slices.Contains(kl, "Company") && slices.Contains(kl, "Campaign") {
			return html.GetDataForCampaign(r.Form), nil
		}
		return nil, fmt.Errorf("could not find Company or Campaign")
	}))
}

func CrawlerExec(mux *http.ServeMux) {
	mux.Handle("GET /crawlexec", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (Node, error) {
		myurls, err := system.GetSystemParams().Dbc.GetUrlsToCrawl(1, 0) // Default to campaign 1
		if err != nil {
			fmt.Printf("Error with GetUrlsToCrawl %s\n", err.Error())
			return nil, err
		}
		for _, url := range myurls {
			fmt.Printf("Crawling url %s\n", url.Url)
			c2 := crawler.NewCrawler2(url.Url, false, "")
			url.Crawldate = time.Now()
			stime := time.Now()
			sc := &crawler.Subcrawler{}
			sc.Procfunc = sc.SCCallback
			c2.Crawl(sc)
			etime := time.Now()
			url.Totalduration = etime.Sub(stime).Seconds()
			url.Crawled = true
			fmt.Printf("Subcrawler results are: %v\n", sc)
			url.Pagescrawled = uint(len(sc.CPages))
			url.Success = true
			err := system.GetSystemParams().Dbc.UpdateCrawlerresults(&url)
			if err != nil {
				fmt.Printf("Error with UpdateCrawlerresults %s\n", err.Error())
			}
		}
		return nil, nil
	}))
}

func CrawlerSetup(mux *http.ServeMux) {
	mux.Handle("GET /crawltest", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (Node, error) {
		var States []string = []string{"AZ", "UT", "IL", "HI"}
		fmt.Printf("Received crawltest request\n")
		fmt.Printf("States: %v\n", States)
		for _, state := range States {
			mySt, serr := system.GetSystemParams().Dbc.GetStateByAbbr(state)
			if serr != nil {
				fmt.Printf("Error with GetStateByAbbr %s\n", serr.Error())
				continue
			}
			zips, zerr := system.GetSystemParams().Dbc.Top10Zipcodes(mySt.ID)
			if zerr != nil {
				fmt.Printf("Error with Top10Zipcodes %s\n", zerr.Error())
				continue
			}
			fmt.Printf("Crawling Delta Dental site for top ten zips in %s\n", state)
			for _, zip := range zips {

				fmt.Printf("Crawling zip %s\n", zip.Zipcode)
				mydc := crawler.NewDeltacrawl()
				err := mydc.Init()
				if err != nil {
					fmt.Printf("Error with Init %s\n", err.Error())
				}
				err = mydc.Run(zip.Zipcode, zip.ID)
				if err != nil {
					fmt.Printf("Error with Run %s\n", err.Error())
				}
			}
		}
		return nil, nil
	}))
}

func Crawler3Exec(mux *http.ServeMux) {
	mux.Handle("GET /crawl3test", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (Node, error) {
		//myurls, err := system.GetSystemParams().Dbc.GetUrlsToCrawl(1, 0) // Default to campaign 1
		myurls, err := system.GetSystemParams().Dbc.GetUrlsToCrawl(1, 0) // Default to campaign 1
		if err != nil {
			fmt.Printf("Error with GetUrlsToCrawl %s\n", err.Error())
			return nil, err
		}

		for _, url := range myurls {
			turl := url.Url
			if !strings.Contains(turl, "www.") {
				turl = fmt.Sprintf("www.%s", turl)
			}
			url.Url = turl
			fmt.Printf("Crawling url %s\n", url.Url)
			c3 := crawler.NewCrawler3(url.Url, false, "")
			url.Crawldate = time.Now()
			stime := time.Now()
			sc := &crawler.Subcrawler{}
			sc.Procfunc = sc.SCCallback
			c3.Crawl(sc)
			etime := time.Now()
			url.Totalduration = etime.Sub(stime).Seconds()
			url.Crawled = true
			fmt.Printf("Subcrawler results are: %v\n", sc)
			err := system.GetSystemParams().Dbc.UpdateCrawlerresults(&url)
			if err != nil {
				fmt.Printf("Error with UpdateCrawlerresults %s\n", err.Error())
			}
			break
		}
		//crawler.Crawl("https://www.arizonaortho.com")
		//c2 := crawler.NewCrawler2("https://www.arizonaortho.com", false)
		return nil, nil
	}))
}

func About(mux *http.ServeMux) {
	mux.Handle("GET /about", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (Node, error) {
		return html.AboutPage(), nil
	}))
}
