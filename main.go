package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

PASSWORDS_FILEPATH := "./records.csv"

func main() {
	showMenu()
}

func showMenu() {
	option = chooseMenuOption()

	if option == 1 {
		printCSV()
		showMenu()
	} else if option == 2 {
		addNewPassword()
		showMenu()
	}
}

func chooseMenuOption() option int {
	fmt.Println("Please type number of option, you want to choose:")
	fmt.Println("1 - Show passwords")
	fmt.Println("2 - Add new password")
	fmt.Scanln(&option)
	return
}

func printCSV() {
	csv_data, err := csvReadAll(PASSWORDS_FILEPATH)
	if err != nil {
		fmt.Println("An error encountered ::", err)
	} else {
		fmt.Println(csv_data)
	}
}

func addNewPassword() {
	site_url, login, passphrase := inputData()
	current_time := time.Now().String()
	password := generatePassword(site_url, login, current_time, passphrase)
	saveData(site_url, login, current_time, password)
}

func inputData() (site_url string, login string, passphrase string) {
	fmt.Print("Enter site url: ")
	fmt.Scanln(&site_url)

	fmt.Print("Enter your login: ")
	fmt.Scanln(&login)

	fmt.Print("Enter your password: ")
	fmt.Scanln(&passphrase)

	return
}

func generatePassword(site_url string, login string, now string, passphrase string) password string {
	data := []byte(site_url + login + passphrase + now)
	hash := sha256.Sum256(data)
	password := [8]byte{}
	for i := 0; i < 8; i++ {
		password[i] = hash[4*i] + hash[4*i+1] + hash[4*i+2] + hash[4*i+3]
	}
	fmt.Printf("%x", password)
	return
}

func saveData(site_url string, login string, current_time string, password string) {
	csv_data, err := csvReadAll(PASSWORDS_FILEPATH)
	if err != nil {
		fmt.Println("An error encountered ::", err)
	} else {
		fmt.Println(csv_data)
	}
}