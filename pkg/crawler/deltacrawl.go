package crawler

import (
	"fmt"
	"github.com/playwright-community/playwright-go"
	"github.com/sftwrngnr/gsearchclient/pkg/system"
	"log"
	"time"
)

const (
	backButtonId             string = "#backTo-dentistSearchList"
	searchResultsDiv         string = "dentistSearchResults"
	searchControls           string = ".dentistSearchControls"
	singleDentistCtl         string = ".singleDentistControls"
	searchResultsListSection string = "searchResultsList"
	searchResultsClass       string = ".results-list"
	dentistInfoPanel         string = ".info-dentist-0-single"
	dentist_info_card        string = ".info-card__body"
	dentistSearchListCol     string = "#dentistSearchListCol"
	dentist_name             string = ".dentist-name"
	dentist_job_title        string = ".post-jobtitle"
	dentist_contact          string = ".post-contact"
	details_block            string = ".post-details"
	detailsButtonId          string = "disableInList"
	address_block            string = ".post-address"
	paginationSection        string = ".searchPagination"
	nextButton               string = "next"
	onetrustsdk              string = ".onetrust-pc-sdk"
	alertFrame               string = ".ot-text-resize"
	cookieDlg                string = ".alertdialog"
	cookieGroup              string = ".onetrust-button-group-parent"
	cookieButton             string = "#onetrust-accept-btn-handler"
	urlPattern               string = "https://www.deltadental.com/us/en/member/find-a-dentist/dentist-search-results.html?address=%s&specialtyCode=020&plan=Any"
)

type Deltacrawl struct {
	pw      *playwright.Playwright
	pwPage  playwright.Page
	browser playwright.Browser
	zcid    uint
}

func NewDeltacrawl() *Deltacrawl {
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("Failed to start Deltacrawl: %v", err)
	}
	return &Deltacrawl{pw: pw}
}

func (dc *Deltacrawl) Init() (err error) {
	dc.browser, err = dc.pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	if err != nil {
		return err
	}
	return nil
}

func (dc *Deltacrawl) Run(zc string, zcid uint) (err error) {
	dc.zcid = zcid
	page, perr := dc.browser.NewPage()
	if perr != nil {
		err = perr
		return
	}
	dc.pwPage = page
	myResp, gerr := page.Goto(fmt.Sprintf(urlPattern, zc))
	if gerr != nil {
		err = gerr
		return
	}
	fmt.Printf("Waiting for page load\n")
	page.WaitForLoadState()
	fmt.Printf("Loading done\n")
	defer dc.Shutdown()
	dc.AcceptCookies()
	for n := 0; n < 1; n++ {
		dip := page.Locator(dentistInfoPanel).First()
		if dip == nil {
			fmt.Printf("Error locating %s\n", dentistInfoPanel)
			return
		}
		entries, eerr := page.Locator(searchResultsClass).All()
		if eerr != nil {
			fmt.Printf("Error with locating %s:: %s\n", searchResultsClass, eerr.Error())
			err = eerr
			return
		}

		for i, entry := range entries {
			switch i {
			case 0:
				perr := dc.processDentistSearchResults(entry, dip)
				if perr != nil {
					fmt.Printf("Error processing dentist search results for %v: %s\n", entry, perr.Error())
					err = perr
				}
				break
			default:
				fmt.Printf("%v\n", entry)
			}
			fmt.Printf("Entry #%d: %v\n", i, entry)

		}
		fmt.Printf("%v\n", myResp)
		time.Sleep(5 * time.Second)
	}
	return
}

func (dc *Deltacrawl) AcceptCookies() {
	cButt := dc.pwPage.Locator(cookieButton).First()
	if cButt != nil {
		cButt.Click()
	}
}

func (dc *Deltacrawl) Shutdown() {
	var err error
	if err = dc.browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
	if err = dc.pw.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}

}

func (dc *Deltacrawl) processDentistSearchResults(ent playwright.Locator, disp playwright.Locator) (err error) {
	indivEnt, iErr := ent.Locator(dentist_info_card).All()

	if iErr != nil {
		err = iErr
		return
	}
	fmt.Printf("There are %d entries in %s\n", len(indivEnt), searchResultsClass)
	for _, entry := range indivEnt {
		myb, berr := entry.TextContent()
		if berr != nil {
			err = berr
		}
		fmt.Printf("%s\n", myb)

		// Get buttons within the info card
		myButtons, berr := entry.Locator("button").All()
		if berr != nil {
			err = berr
			return
		}
		if len(myButtons) > 1 {
			myButtons[2].Click()
			if err = dc.pwPage.WaitForLoadState(); err != nil {
				panic(err)
			}
			dc.ProcessIndivDentistBlock()
			time.Sleep(1 * time.Second)
			bkbutn := dc.pwPage.Locator(backButtonId).First()
			if bkbutn != nil {
				fmt.Printf("Found single dentist button: %v\n", bkbutn)
				bkErr := bkbutn.Click()
				if bkErr != nil {
					fmt.Printf("Click button failed: %v\n", bkErr)
				}
				if err = dc.pwPage.WaitForLoadState(); err != nil {
					panic(err)
				}
			}

		}

	}
	if err = dc.pwPage.WaitForLoadState(); err != nil {
		panic(err)
	}
	time.Sleep(1 * time.Second)
	// Find next buttons

	// Go to next page
	fmt.Printf("Executing Next Button\n")
	myNext := dc.pwPage.GetByText(nextButton)
	fmt.Printf("Next button #%v\n", myNext)
	myNext.GetByRole("button").Click()
	return
}

func (dc *Deltacrawl) ProcessIndivDentistBlock() {
	var (
		dName  string
		dJob   string
		dAddy  string
		dPhone string
		dEmail string
	)

	dcloc := dc.pwPage.Locator(dentistSearchListCol).First()
	if dcloc == nil {
		fmt.Printf("Couldn't find %s\n", dentistSearchListCol)
		return
	}
	dname := dcloc.Locator(dentist_name).First()
	dentist, derr := dname.TextContent()
	if derr != nil {
		fmt.Printf("Couldn't find %s: %v\n", dname, derr)
		return
	}
	dName = dentist
	jobtitle := dcloc.Locator(dentist_job_title).First()
	Job, jerr := jobtitle.TextContent()
	if jerr != nil {
		fmt.Printf("Couldn't find %s: %v\n", jobtitle, jerr)
		return
	}
	dJob = Job
	fmt.Printf("Found %s\n", dentist)
	fmt.Printf("Job title: %v\n", Job)
	addy := dcloc.Locator(address_block).First()
	if addy == nil {
		fmt.Printf("Couldn't find %s\n", address_block)
		return
	}

	Addy, aerr := addy.TextContent()
	if aerr != nil {
		fmt.Printf("Couldn't find %s: %v\n", addy, aerr)
	} else {
		fmt.Printf("Addy: %v\n", Addy)
		dAddy = Addy
	}
	dcont := dcloc.Locator(dentist_contact).First()

	if dcont != nil {
		postDet, pdErr := dcont.Locator(details_block).All()
		if pdErr != nil {
			fmt.Printf("Couldn't find %s: %v\n", details_block, pdErr)
			return
		}
		for i, d := range postDet {
			switch i {
			case 0:
				dive, derr := d.Locator("div").All()
				if derr != nil {
					fmt.Printf("Couldn't find %s: %v\n", "div", derr)
					continue
				}
				for d, di := range dive {
					myTxt, txerr := di.TextContent()
					if txerr != nil {
						fmt.Printf("Couldn't find %s: %v\n", txerr, txerr)
					} else {
						fmt.Printf("%s\n", myTxt)
						if d == 1 {
							dPhone = myTxt
						}
					}

				}
				break
			case 1:
				hrt, hrerr := d.TextContent()
				if hrerr != nil {
					fmt.Printf("Couldn't find text %v\n", hrerr)
					continue
				}
				fmt.Printf("%s\n", hrt)
				dEmail = hrt
				break
			default:
				fmt.Printf("Fuck chocolate shakes.\n")

			}
		}
	}
	dberr := system.GetSystemParams().Dbc.CreateDeltaData(dName, dJob, dAddy, dPhone, dEmail, dc.zcid)
	if dberr != nil {
		fmt.Printf("Couldn't create delta data: %v\n", dberr)
	}
	// Create
}
