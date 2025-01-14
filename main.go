// package main

// import (
// 	// "booking-app/shared"
// 	"fmt"
// 	"strings"
// )

// const conferenceTickets int = 50

// var conferenceName = "Go conference"
// var remainingTickets uint = 50
// var bookings []string

// func main() {
// 	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

// 	greetUsers()

// 	for {
// 		var firstName string
// 		var lastName string
// 		var email string
// 		var userTickets uint

// 		firstName, lastName, email, userTickets = getUserInput()
// 		// isValidName, isValidEmail, isValidTicketNumber := shared.ValidateUserInput(firstName, lastName, email, userTickets)

// 		if isValidName && isValidEmail && isValidTicketNumber {

// 			bookTicket(userTickets, firstName, lastName, email)

// 			firstNames := getFirstNames()
// 			fmt.Printf("The first names of bookings are: %v\n", firstNames)

// 			if remainingTickets == 0 {
// 				// end program
// 				fmt.Println("Our conference is booked out. Come back next year.")
// 				break
// 			}
// 		} else {
// 			if !isValidName {
// 				fmt.Println("First or last name you entered is too short.")
// 			}
// 			if !isValidEmail {
// 				fmt.Println("Email you entered doesn't contain @gmail.com.")
// 			}
// 			if !isValidTicketNumber {
// 				fmt.Println("Ticket number you entered is invalid.")
// 			}
// 		}
// 	}
// }

// func greetUsers() {
// 	fmt.Printf("Welcome to %v booking application", conferenceName)
// 	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
// 	fmt.Println("Get your tickets here to attend")
// }

// func getFirstNames() []string { //inside the bracket (string) is the input parameters and outside of the bracket (string) is the outside parameter
// 	firstNames := []string{}
// 	for _, booking := range bookings {
// 		var names = strings.Fields(booking)
// 		firstNames = append(firstNames, names[0])
// 	}
// 	return firstNames
// }

// func getUserInput() (string, string, string, uint) {
// 	var firstName string
// 	var lastName string
// 	var email string
// 	var userTickets uint

// 	// ask user for their name
// 	fmt.Println("Enter your first name: ")
// 	fmt.Scan(&firstName)

// 	fmt.Println("Enter your last name: ")
// 	fmt.Scan(&lastName)

// 	fmt.Println("Enter your email address: ")
// 	fmt.Scan(&email)

// 	fmt.Println("Enter number of tickets: ")
// 	fmt.Scan(&userTickets)

// 	return firstName, lastName, email, userTickets
// }

// func bookTicket(userTickets uint, firstName string, lastName string, email string) {
// 	remainingTickets -= userTickets
// 	bookings = append(bookings, firstName+" "+lastName)

// 	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
// 	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
// }
