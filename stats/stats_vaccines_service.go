package stats

import (
	"log"
	"strconv"
)

var VaccineEntries []*VaccineEntry

func init() {
	// Columns
	// country,iso_code,date,total_vaccinations,people_vaccinated,people_fully_vaccinated,daily_vaccinations_raw,daily_vaccinations,total_vaccinations_per_hundred,people_vaccinated_per_hundred,people_fully_vaccinated_per_hundred,daily_vaccinations_per_million,vaccines,source_name,source_website
	vaccineRecords, _ := ReadCsvFile("../stats/data/vaccines/data.csv")

	// Transform preemptively for perfomance reason
	// We can use it afterwards also in REST API as plain JSON
	VaccineEntries = make([]*VaccineEntry, len(vaccineRecords))

	for i, r := range vaccineRecords {
		totalVaccinations, _ := strconv.ParseInt(r[3], 10, 32)
		peopleVaccinated, _ := strconv.ParseInt(r[4], 10, 32)
		peopleFullyVaccinated, _ := strconv.ParseInt(r[5], 10, 32)
		dailyVaccinationsRaw, _ := strconv.ParseInt(r[6], 10, 32)
		dailyVaccinations, _ := strconv.ParseInt(r[7], 10, 32)
		totalVaccinationsPerHundred, _ := strconv.ParseInt(r[8], 10, 32)
		peopleVaccinatedPerHundred, _ := strconv.ParseInt(r[9], 10, 32)
		peopleFullyVaccinatedPerHundred, _ := strconv.ParseInt(r[10], 10, 32)
		DailyVaccinationsPerMillion, _ := strconv.ParseInt(r[11], 10, 32)

		VaccineEntries[i] = &VaccineEntry{
			Country:                         r[0],
			IsoCode:                         r[1],
			Date:                            r[2],
			TotalVaccinations:               uint32(totalVaccinations),
			PeopleVaccinated:                uint32(peopleVaccinated),
			PeopleFullyVaccinated:           uint32(peopleFullyVaccinated),
			DailyVaccinationsRaw:            uint32(dailyVaccinationsRaw),
			DailyVaccinations:               uint32(dailyVaccinations),
			TotalVaccinationsPerHundred:     uint32(totalVaccinationsPerHundred),
			PeopleVaccinatedPerHundred:      uint32(peopleVaccinatedPerHundred),
			PeopleFullyVaccinatedPerHundred: uint32(peopleFullyVaccinatedPerHundred),
			DailyVaccinationsPerMillion:     uint32(DailyVaccinationsPerMillion),
			Vaccines:                        r[12],
			SourceName:                      r[13],
			SourceWebsite:                   r[14],
		}
	}

	log.Printf("Vaccines Service initialized. Total number of vaccine records is %v.", len(VaccineEntries))
}
