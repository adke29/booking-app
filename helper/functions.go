package helper

import (
	"fmt"
)

func Results(customer string, availableTickets uint, firstNames []string) {
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
