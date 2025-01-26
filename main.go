package main

import (
	"booking-app/shared"
	"fmt"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings []Booking

// Define a struct for bookings
type Booking struct {
	firstName   string
	lastName    string
	email       string
	userTickets uint
}

func main() {
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	greetUsers()

	for {
		// Get user input
		firstName, lastName, email, userTickets := getUserInput()

		// Validate user input
		isValidName, isValidEmail, isValidTicketNumber := shared.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			// Book the ticket
			bookTicket(userTickets, firstName, lastName, email)

			// Display the first names of all bookings
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			// Check if tickets are sold out
			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			// Provide detailed error messages
			if !isValidName {
				fmt.Println("First or last name you entered is too short. Each should have at least 2 characters.")
			}
			if !isValidEmail {
				fmt.Println("Email you entered is invalid. It must contain '@' and end with '@gmail.com'.")
			}
			if !isValidTicketNumber {
				fmt.Println("The number of tickets you entered is invalid. It must be between 1 and the remaining tickets.")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("\nWelcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets, and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// Ask user for their details
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	// Update remaining tickets
	remainingTickets -= userTickets

	// Create a new booking and add it to the list
	newBooking := Booking{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		userTickets: userTickets,
	}
	bookings = append(bookings, newBooking)

	// Confirmation message
	fmt.Printf("\nThank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n\n", remainingTickets, conferenceName)
}
