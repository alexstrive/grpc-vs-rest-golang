package stats

import (
	"log"
	"strconv"
)

var CovidCases []*CovidCaseStatEntry

func init() {
	log.Println("Initializing Covid Stats Service...")

	// Columns
	// dateRep,day,month,year,cases,deaths,countriesAndTerritories,geoId,countryterritoryCode,popData2020,continentExp
	covidCasesRecords, _ := ReadCsvFile("../stats/data/covid-europe/covid-europe.csv")

	// Transform preemptively for perfomance reason
	// We can use it afterwards also in REST API as plain JSON
	CovidCases = make([]*CovidCaseStatEntry, len(covidCasesRecords))

	for i, s := range covidCasesRecords {
		// Order of fields you can look up in header comment
		// uint64 must be downcased to uint32 due some fields defined so in proto scheme
		// TODO: perhaps there's better way of converting from string to uint32?
		day, _ := strconv.ParseUint(s[1], 10, 32)
		month, _ := strconv.ParseUint(s[2], 10, 32)
		year, _ := strconv.ParseUint(s[3], 10, 32)
		cases, _ := strconv.ParseUint(s[4], 10, 64)
		deaths, _ := strconv.ParseUint(s[5], 10, 64)
		popData2020, _ := strconv.ParseUint(s[9], 10, 64)

		CovidCases[i] = &CovidCaseStatEntry{
			DateRep:                 s[0],
			Day:                     uint32(day),
			Month:                   uint32(month),
			Year:                    uint32(year),
			Cases:                   cases,
			Deaths:                  deaths,
			CountriesAndTerritories: s[6],
			GeoId:                   s[7],
			CountryterritoryCode:    s[8],
			PopData2020:             popData2020,
			ContinentExp:            s[10],
		}
	}

	log.Printf("Covid Stats Service initialized. Total number of covid case records is %v.", len(CovidCases))
}
