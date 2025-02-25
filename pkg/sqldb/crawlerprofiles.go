package sqldb

import "gorm.io/gorm"

type Crawlerprofiles struct {
	gorm.Model
	ID                   uint   `gorm:"primaryKey"`
	Maxdepth             uint   `json:"maxdepth"`
	Maxtime              uint   `json:"maxtime"`
	Maxthreads           uint   `json:"maxthreads"`
	Storepages           bool   `json:"storepages"`
	Buildseolist         bool   `json:"buildseolist"`
	Extractaddress       bool   `json:"extractaddress"`
	Extractphone         bool   `json:"extractphone"`
	Specialextract       bool   `json:"specialextract"`
	Extractfunc          string `json:"extractfunc"`
	Name                 string `json:"name"`
	Dailymaxcrawl        uint   `json:"dailymaxcrawl"`
	Extractexternallinks bool   `json:"extractexternallinks"`
	Extractwordcloud     bool   `json:"extractwordcloud"`
	Company              uint   `json:"company"`
	Searchcampaign       uint   `json:"searchcampaign"`
	Multicrawl           bool   `json:"multicrawl"`
	Agenttype            uint   `json:"agenttype"`
}

func (dbc *DBConnData) GetCompanyCrawlers(company uint) (crawlers []Crawlerprofiles, err error) {
	err = dbc.DB.Where("company = ?", company).Find(&crawlers).Error
	return
}
