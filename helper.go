package main

import "strings"

func inputValidation(
	firstName string,
	lastName string,
	email string,
	userTickets uint,

) (
	bool,
	bool,
	bool,
) {
	validName := len(firstName) >= 2 && len(lastName) >= 2
	validEmail := strings.Contains(email, "@",)
	validTicketNo := userTickets > 0 && userTickets <= remainingTickets

	return validName, validEmail, validTicketNo
}