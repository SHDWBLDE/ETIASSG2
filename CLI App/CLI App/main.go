package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func clearScreen() {
	if runtime.GOOS == "linux" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else if runtime.GOOS == "window" {
		cmd := exec.Command("cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else if runtime.GOOS == "darwin" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func main() {

	// TODO: API call
	// TODO: Refactor this main function into a separate functions/packages

	fmt.Println("Welcome to ridshare app")
	fmt.Println("1. Create an Account")
	fmt.Println("2. Login")
	fmt.Println()
	fmt.Print("Choice: ")

	var welcomeScreenChoice string
	fmt.Scanln(&welcomeScreenChoice)

	if welcomeScreenChoice == "1" {
		clearScreen()
		fmt.Println("1. Create a passenger account")
		fmt.Println("2. Create a driver account")
		fmt.Println()
		fmt.Print("Choice: ")

		var registerScreenChoice string
		fmt.Scanln(&registerScreenChoice)
		clearScreen()
		if registerScreenChoice == "1" {
			fmt.Println("All fields are required!")
			fmt.Println()
			var (
				firstName    string
				lastName     string
				mobileNumber string
				email        string
				password     string
			)

			fmt.Print("First Name: ")
			fmt.Scanln(&firstName)
			fmt.Print("Last Name: ")
			fmt.Scanln(&lastName)
			fmt.Print("Mobile Number: ")
			fmt.Scanln(&mobileNumber)
			fmt.Print("Email: ")
			fmt.Scanln(&email)
			fmt.Print("Password: ")
			fmt.Scanln(&password)

			fmt.Println(firstName)
		}

	} else if welcomeScreenChoice == "2" {
		fmt.Println("")

		var (
			email    string
			password string
		)

		fmt.Print("Email: ")
		fmt.Scanln(&email)
		fmt.Print("Password: ")
		fmt.Scanln(&password)

		// TODO: call the API

	}

}
