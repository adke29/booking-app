package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {

	var availableTickets uint = 50
	var orderedTicket uint
	var firstName string
	var lastName string
	var bookings []helper.UserModel
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
		var userData = helper.UserModel{
			FirstName:     firstName,
			LastName:      lastName,
			Email:         email,
			OrderedTicket: orderedTicket,
		}

		bookings = append(bookings, userData)
		firstNames := []string{}
		for _, booking := range bookings {
			firstNames = append(firstNames, booking.FirstName)
		}

		helper.Results(firstName, availableTickets, firstNames)
		wg.Add(1)
		go sendTicket(email, orderedTicket) //create go routine
		fmt.Println("--------------------------------------------")
		if availableTickets <= 0 {
			fmt.Println("Tickets have been sold out")
		}
		fmt.Println(bookings)
		wg.Wait()

	}
}

func sendTicket(email string, orderedTickets uint) {
	time.Sleep(5 * time.Second)
	fmt.Println("######################")
	fmt.Printf("Sending %v tickets to %v\n", orderedTickets, email)
	fmt.Println("Tickets have been sent")
	fmt.Println("######################")
	wg.Done()
}
