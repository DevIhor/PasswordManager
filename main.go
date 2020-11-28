package main

import (
	"fmt"
)

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
}

func main() {
	credentialsInput()
}