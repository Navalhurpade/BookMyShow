package helper

import (
	"fmt"
	"strings"
)

func GreetUser(totalTickets int, remainingTickets uint) {
	fmt.Printf("Wellcome to Tiket booking system !\nWe have %v tikets, out of which %v are still remaining", totalTickets, remainingTickets)
	fmt.Println("\nBefore proceding furture can you provide few details below !")
}

func ValidateUser(email string, firstName string, lastName string) (bool, bool, bool) {
	isEmailValid := strings.Contains(email, "@")
	isFirstNameValid := len(firstName) > 0
	isLastNameValid := len(lastName) > 0

	return isEmailValid, isFirstNameValid, isLastNameValid
}
