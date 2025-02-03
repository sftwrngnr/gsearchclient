package datareaders

// This will read the csv zip code file.

type Zipcode struct {
	zip        string
	city       string
	state      string
	statename  string
	population int64
}

// "zip","lat","lng","city","state_id","state_name","zcta","parent_zcta","population","density","county_fips","county_name","county_weights","county_names_all","county_fips_all","imprecise","military","timezone"
// 0,3, 4, 5 [state name], 8
