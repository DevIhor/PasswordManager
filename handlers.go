package main

import (
	"encoding/csv"
	"os"
)

func csvReadAll(filepath string) ([][]string, error) {
	// Open file
	recordsFile, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer recordsFile.Close()

	// Read file
	reader := csv.NewReader(recordsFile)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}

func csvWrite(filepath string, csv_data [][]string) (err error) {
	// Create or open file
	recordsFile, err := os.Create(filepath)
	if err != nil {
		return
	}

	// Close file
	defer func() {
        cerr := recordsFile.Close()
        if err == nil {
            err = cerr
        }
    }()

	// Write to file
	writer := csv.NewWriter(recordsFile)
	err = writer.WriteAll(csv_data)
	return
}