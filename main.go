package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	credentialsInput()
}

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

func csvWriter(filepath string, csv_data [][]string) (err error) {
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

func printCSV(filepath string) {
	csv_data, err := csvReadAll(filepath)
	if err != nil {
		fmt.Println("An error encountered ::", err)
	} else {
		fmt.Println(csv_data)
	}
}

func credentialsInput() {
	fmt.Println("Enter site url: ")

	var site_url string
	fmt.Scanln(&site_url)

	fmt.Println("Enter your login: ")

	var login string
	fmt.Scanln(&login)

	fmt.Println("Enter your password: ")

	var passphrase string
	fmt.Scanln(&passphrase)

	fmt.Println("Do you want to print saved records? (y/n) ")

	var answer string
	fmt.Scanln(&answer)

	if answer == "y" {
		printCSV("./records.csv")
	}
}