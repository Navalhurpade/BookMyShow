package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
)

func main() {
	const totalTickets int = 100
	var remainingTickets uint = uint(totalTickets)
	var noOfTicketsTOBook uint = 0
	bookings := []string{}

	for {

		firstName := ""
		lastName := ""
		email := ""
		isFirstNameValid := false
		isLastNameValid := false
		isEmailValid := false

		helper.GreetUser(totalTickets, remainingTickets)

		for !(isEmailValid && isFirstNameValid && isLastNameValid) {

			email, firstName, lastName = getUserData()

			isEmailValid, isFirstNameValid, isLastNameValid = helper.ValidateUser(email, firstName, lastName)

			showValidationMessage(isEmailValid, isFirstNameValid, isLastNameValid)
		}

		fmt.Printf("\nHow many tikets would you like to book ?")
		fmt.Scan(&noOfTicketsTOBook)

		remainingTickets = bookTicket(noOfTicketsTOBook, remainingTickets)

		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Congratulations %v, your tiket has been booked succesfully !\nTikect will be sent to your email %v.\n", firstName, email)
		firstNameBookings := []string{}

		for _, name := range bookings {
			firstNameBookings = append(firstNameBookings, strings.Fields(name)[0])
		}

		fmt.Printf("Current bookings %v\n\n", firstNameBookings)
		if remainingTickets < 1 {
			fmt.Printf("Show has been booked out, come back next year !\nBye bye !")
			break
		}
	}
}

func getUserData() (string, string, string) {
	var email, firstName, lastName string

	fmt.Printf("Email :: ")
	fmt.Scan(&email)

	fmt.Printf("firstName :: ")
	fmt.Scan(&firstName)

	fmt.Printf("lastName :: ")
	fmt.Scan(&lastName)

	return email, firstName, lastName
}

func showValidationMessage(isEmailValid bool, isFirstNameValid bool, isLastNameValid bool) {
	if !isEmailValid {
		fmt.Printf("Email is Not valid. Please try again !\n")
	}

	if !isFirstNameValid {
		fmt.Printf("FirstName is Not valid. Please try again !\n")
	}

	if !isLastNameValid {
		fmt.Printf("LastName is Not valid. Please try again !\n")
	}
}

func bookTicket(noOfTicketsTOBook uint, remainingTickets uint) uint {

	if noOfTicketsTOBook > remainingTickets {
		fmt.Printf("Sorry, We only have %v have tikets available. Please try again !", remainingTickets)
		fmt.Scan(&noOfTicketsTOBook)
	}

	return remainingTickets - noOfTicketsTOBook
}
