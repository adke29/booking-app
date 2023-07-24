package main

import (
	"fmt"
	"strings"
)

func main() {
	var availableTickets int = 50
	var orderedTicket int
	var firstName string
	var lastName string
	var bookings []string
	var email string

	for availableTickets > 0 {
		fmt.Print("How many ticket you want to buy: ")
		fmt.Scan(&orderedTicket)
		if orderedTicket > availableTickets || orderedTicket <= 0 {
			fmt.Println("We only have", availableTickets, " tickets remaining")
			continue
		}

		fmt.Print("First Name: ")
		fmt.Scan(&firstName)
		fmt.Print("Last Name: ")
		fmt.Scan(&lastName)
		if len(firstName) <= 2 || len(lastName) <= 2 {
			fmt.Println("Please input correct name!!")
			continue
		}

		fmt.Print("Email: ")
		fmt.Scan(&email)
		if !strings.Contains(email, "@") {
			fmt.Println("Please input correct email!!")
			continue
		}

		//Booking process
		availableTickets -= orderedTicket
		bookings = append(bookings, (firstName + " " + lastName))
		firstNames := []string{}
		for _, booking := range bookings {
			var name = strings.Fields(booking)
			firstName = name[0]
			firstNames = append(firstNames, firstName)
		}

		results(bookings[len(bookings)-1], availableTickets, firstNames)
		fmt.Println("--------------------------------------------")
		if availableTickets <= 0 {
			fmt.Println("Tickets have been sold out")
		}

	}
}

func results(customer string, availableTickets int, firstNames []string) {
	fmt.Printf("Thank you %v for your booking!\n", customer)
	fmt.Println("Available tickets: ", availableTickets)
	for index, firstName := range firstNames {
		fmt.Print(firstName)
		if index < (len(firstNames) - 1) {
			fmt.Print(", ")
		}
	}
	fmt.Println(" has ordered")

}
