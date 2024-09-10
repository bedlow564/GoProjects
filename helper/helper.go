package helper

import "strings"

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) { //put reutrn type here. Multiple values in parenthesis //capitalize function name to use it in other packages (export)
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmailAddress := strings.Contains(email, "@")
	isUserTicketsValid := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmailAddress, isUserTicketsValid //return multiple values
}
