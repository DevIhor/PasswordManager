package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

func main() {
	site_url, login, passphrase := credentialsInput()
	generatePassword(site_url, login, passphrase)
}

func generatePassword(site_url string, login string, passphrase string) {
	current_time := time.Now().String()
	data := []byte(site_url + login + passphrase + current_time)
	hash := sha256.Sum256(data)
	password := [8]byte{}
	for i := 0; i < 8; i++ {
		password[i] = hash[4*i] + hash[4*i+1] + hash[4*i+2] + hash[4*i+3]
	}
	fmt.Printf("%x", password)
}

func printCSV(filepath string) {
	csv_data, err := csvReadAll(filepath)
	if err != nil {
		fmt.Println("An error encountered ::", err)
	} else {
		fmt.Println(csv_data)
	}
}

func credentialsInput() (site_url string, login string, passphrase string) {
	fmt.Print("Enter site url: ")
	fmt.Scanln(&site_url)

	fmt.Print("Enter your login: ")
	fmt.Scanln(&login)

	fmt.Print("Enter your password: ")
	fmt.Scanln(&passphrase)

	var answer string
	fmt.Print("Do you want to print saved records? (y/n) ")
	fmt.Scanln(&answer)
	if answer == "y" {
		printCSV("./records.csv")
	}
	return
}