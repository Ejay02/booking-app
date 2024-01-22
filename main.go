package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets  int = 50
var conferenceName = "Nx Conference"
var remainingTickets uint = 50
var bookings = make([]userData, 0)

type userData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

			firstName, lastName, email, userTickets := getUserInput()
	
			validTicketNo, validName, validEmail := inputValidation(firstName, lastName, email, userTickets, )


			if validTicketNo && validName && validEmail {
				bookTicket(userTickets, firstName, lastName, email)
		
				wg.Add(1)
				go sendTicket(userTickets, firstName, lastName, email)
		
				firstNames := getFirstNames()
				fmt.Printf("These are all our guests: %v\n", firstNames)
		
				if remainingTickets == 0 {
						fmt.Println("We are currently Sold Out, try again next year")
				}
		} else {
				if !validName {
						fmt.Printf("Name too short, try again\n")
				}
				if !validEmail {
						fmt.Printf("Invalid email, try again\n")
				}
				if !validTicketNo {
						if userTickets > remainingTickets {
								fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets 🥹", remainingTickets, userTickets)
						} else {
								fmt.Printf("Invalid ticket number, try again\n")
						}
				}
		}
		
	wg.Wait()

}

func greetUsers(){
	fmt.Printf("Welcome to %v ✨!\nWe have a total of %v tickets, %v are still available.\nGet your tickets here to attend\n", conferenceName, conferenceTickets, remainingTickets)
}

func getFirstNames() []string{
	firstNames := []string{}
	for _, booking := range  bookings{
			firstNames =	append(firstNames, booking.firstName)
	}

	return firstNames
}

func getUserInput()(string, string, string, uint){
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter no of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string){
	remainingTickets = remainingTickets - userTickets


	// map for user
	var userData = userData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings,userData)
	fmt.Printf("List of bookings is %v\n\n",bookings)
	
	fmt.Printf("Thank you %v %v for booking %v tickets, you will receive a confirmation email shortly at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v remaining for this %v\n\n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)

fmt.Println("################################")

fmt.Printf("Sending ticket(s): \n %v \nto email address %v\n", ticket, email)
fmt.Println("################################")
wg.Done()
}