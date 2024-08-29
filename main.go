package main

import (
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go conference"

const conferenceTickets = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	for {
		//add user input
		firstName, lastName, email, userTickets := getUserInput()

		//validate user input
		isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			//bookings

			bookTickets(remainingTickets, userTickets, firstName, lastName, email)

			wg.Add(1)
			go sendTickets(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of the bookings are %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Printf("Our conf is booked out come back next year!")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("First name or the last name that you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Invalid email.Try again")
			}
			if !isValidTicketNumber {
				fmt.Println("No of tickets that you entered is invalid")
			}
		}

	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v Booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get yo tickets here to attend!")

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

	//add a user
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your gmail:")
	fmt.Scan(&email)

	fmt.Println("How many tickets would you like to buy?")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(remainingTickets uint, userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	//create a map for the user

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thankyou %v %v for booking %v tickets.You will recieve a confirmation mail at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)
}

func sendTickets(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v amt of tickets have been succesfully booked for %v %v", userTickets, firstName, lastName)
	fmt.Println("###############")
	fmt.Printf("Sending ticket:\n %v \nto email address%v\n", ticket, email)
	fmt.Println("###############")
	wg.Done()

}
