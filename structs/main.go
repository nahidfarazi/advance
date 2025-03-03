package main

import (
	"fmt"
	"struct/project"
)

func main() {
	fmt.Printf("<--- service list --->\n\n1.Bank\n2.Ticket\n\n")
	var option int
	fmt.Printf("Enter your option: ")
	fmt.Scanln(&option)
	if option == 1 {
		project.BankProject()
	} else if option == 2 {
		project.TicketProject()
	}

}
