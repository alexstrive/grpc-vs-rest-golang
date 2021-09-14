package stats

import (
	"log"
	"strconv"
)

var Stocks []*Stock

func init() {
	// Columns
	// dateRep,day,month,year,cases,deaths,countriesAndTerritories,geoId,countryterritoryCode,popData2020,continentExp
	stocksRecords, _ := ReadCsvFile("../stats/data/stocks/index.csv")

	// Transform preemptively for perfomance reason
	// We can use it afterwards also in REST API as plain JSON
	Stocks = make([]*Stock, len(stocksRecords))

	for i, s := range stocksRecords {

		open, _ := strconv.ParseFloat(s[2], 32)
		high, _ := strconv.ParseFloat(s[3], 32)
		low, _ := strconv.ParseFloat(s[4], 32)
		close, _ := strconv.ParseFloat(s[5], 32)
		adjClose, _ := strconv.ParseFloat(s[6], 32)
		volume, _ := strconv.ParseFloat(s[7], 32)
		closeUsd, _ := strconv.ParseFloat(s[8], 32)

		Stocks[i] = &Stock{
			Index:    s[0],
			Date:     s[1],
			Open:     float32(open),
			High:     float32(high),
			Low:      float32(low),
			Close:    float32(close),
			AdjClose: float32(adjClose),
			Volume:   float32(volume),
			CloseUSD: float32(closeUsd),
		}
	}

	log.Printf("Stocks Service initialized. Total number of stock records is %v.", len(Stocks))
}
