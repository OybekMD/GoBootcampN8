/*
Oybekmd Console DB

You can create your own file DB with this project

## Table of Contents

- [Installation] go run main.go
- [Usage] Go -> "bufio", "fmt", "goexperience/encryption", "os", "strings", "time"
- [Version] v0.1
- [License] MDLicense

## Installation
Provide step-by-step instructions on how to install and set up your project. Include any dependencies that need to be installed.

*/

package main

import (
	"bufio"
	"fmt"
	enc "goexperience/encryption"
	"os"
	"strings"
	"time"
)

// users
// password
// data

func Login(username string, password string) (string, bool) {
	file, err := os.Open("pass.txt")
	if err != nil {
		fmt.Println("Login func file not found!")
		panic(err)
	}
	defer file.Close()

	rd := bufio.NewScanner(file)
	for rd.Scan() { // возвращает true, пока файл не будет прочитан до конца
		check := string(enc.PasswordDecode(username) + "*" + enc.PasswordDecode(password))
		if rd.Text() == check {
			LoginAccountWriteLog(username)
			return "Hello " + username, true
		}
	}
	return "Username or Password is mistake try later!", false
}

func CreateAccountWriteLog(username string) {
	t := time.Now()
	fwrite := string("Created-New-Accound-->Username:" + username + "|Time:" + string(t.Format("02-01-2006-15:04:05")))
	wordsToUserLogfile := strings.Fields(fwrite)
	file, err := os.OpenFile("logs.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	for _, word := range wordsToUserLogfile {
		_, err := file.WriteString(word + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
}

func LoginAccountWriteLog(username string) {
	t := time.Now()
	fwrite := string("Logined-Account-->Username:" + username + "|Time:" + string(t.Format("02-01-2006-15:04:05")))
	wordsToUserLogfile := strings.Fields(fwrite)
	file, err := os.OpenFile("logs.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	for _, word := range wordsToUserLogfile {
		_, err := file.WriteString(word + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
}

func isUsedUsername(newUsername string) bool {
	file, err := os.Open("users.txt")
	if err != nil {
		fmt.Println("isUseredUsername func file not found!")
		panic(err)
	}
	defer file.Close()

	rd := bufio.NewScanner(file)
	for rd.Scan() { // возвращает true, пока файл не будет прочитан до конца
		// name := strings.Split(rd.Text(), "*")
		if rd.Text() == newUsername {
			return true
		}
	}
	return false
}

func CreateAccount(username string, password string) string {
	if isUsedUsername(username) {
		return "This username used"
	}
	wordsToUserfile := strings.Fields(username)
	wordsToPassfile := strings.Fields(enc.PasswordDecode(username) + "*" + enc.PasswordDecode(password))
	ufile, err := os.OpenFile("users.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return "Error CreateAccount Open user file !"
	}
	defer ufile.Close()

	for _, word := range wordsToUserfile {
		_, err := ufile.WriteString(word + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return "Error CreateAccount for pass file!"
		}
	}
	// -------------- PASSWORD -------------

	pfile, err := os.OpenFile("pass.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return "Error CreateAccount Open pass file !"
	}
	defer pfile.Close()

	for _, word := range wordsToPassfile {
		_, err := pfile.WriteString(word + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return "Error CreateAccount for pass file!"
		}
	}
	CreateAccountWriteLog(username)
	return "Create account have been successful!"
}

func main() {
	// Enter to OS
	var choise, name, pass string
	fmt.Println("----Enter to Gosys----")
	fmt.Print("Select-> Login or Sign: ")
	fmt.Scan(&choise)
	if choise == "Login" {
		fmt.Println("Open Login")
		fmt.Println("Input like... Name Password")
		fmt.Print("Input:")
		fmt.Scan(&name, &pass)
		fmt.Println(Login(name, pass))
	} else if choise == "Sign" {
		fmt.Println("Open Sign")
		fmt.Println("Input like... Name Password")
		fmt.Print("Input:")
		fmt.Scan(&name, &pass)
		fmt.Println(CreateAccount(name, pass))

	} else {
		fmt.Println("Undefine input determent!")
		fmt.Println("For exit -> exit or quit")
	}
}
