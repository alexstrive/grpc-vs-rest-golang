package stats

import (
	"encoding/csv"
	"log"
	"os"
)

func ReadCsvFile(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Printf("Failed to read file: %v", err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	content, err := csvReader.ReadAll()
	return content, err
}
