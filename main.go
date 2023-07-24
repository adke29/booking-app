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

	for {
		fmt.Print("How many ticket you want to buy: ")
		fmt.Scan(&orderedTicket)
		if orderedTicket > availableTickets {
			fmt.Println("We only have", availableTickets, " tickets remaining")
			continue
		}
		fmt.Print("First Name: ")
		fmt.Scan(&firstName)
		fmt.Print("Last Name: ")
		fmt.Scan(&lastName)

		availableTickets -= orderedTicket
		bookings = append(bookings, (firstName + " " + lastName))
		fmt.Printf("Thank you %v for your booking!\n", bookings[len(bookings)-1])
		fmt.Println("Available tickets: ", availableTickets)
		firstNames := []string{}
		for _, booking := range bookings {
			var name = strings.Fields(booking)
			firstName = name[0]
			firstNames = append(firstNames, firstName)
		}
		fmt.Println(firstNames)
		fmt.Println("--------------------------------------------")
		if availableTickets <= 0 {
			fmt.Println("Tickets have been sold out")
			break
		}
	}

}
