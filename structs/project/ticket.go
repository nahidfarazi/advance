package project

import (
	"fmt"
)

type Ticket struct {
	TicketID      int
	PassengerName string
	Destination   string
	IsBooked      bool
}

func (t *Ticket) TicketBooked(booked bool) {
	if t.IsBooked {
		fmt.Println("Ticket is already Booked")
	} else {
		t.IsBooked = true
		fmt.Println("Ticket Booked successful")
	}
}
func (t *Ticket) TicketCancel(booked bool) {
	if t.IsBooked {
		t.IsBooked = false
		fmt.Println("Ticket cancel successful")
		return
	} else {
		fmt.Println("Ticket already canceled")
	}
}

func (t *Ticket) TicketInfo() {
	if t.IsBooked {
		fmt.Printf("\n*** Ticket Information ***\n")
		fmt.Println("Ticket Id: ", t.TicketID)
		fmt.Println("Passenger Name: ", t.PassengerName)
		fmt.Println("Destination: ", t.Destination)
		fmt.Println("Ticket is Booked\n")
	} else {

		fmt.Printf("\nFirst booked a ticket\n")
	}
}

func createTicket(pName, dName string, booked bool) Ticket {
	t_Id := Grand()
	ticket := Ticket{
		TicketID:      int(t_Id),
		PassengerName: pName,
		Destination:   dName,
		IsBooked:      booked,
	}
	return ticket
}
func TicketProject() {
	var pName, dName string
	var booked bool
	var option int
	fmt.Printf("Enter your name: ")
	fmt.Scanln(&pName)
	fmt.Printf("Enter your Destination: ")
	fmt.Scanln(&dName)
	fmt.Printf("1.Booked\n2.Cancel\n")
	fmt.Printf("Enter your option: ")
	fmt.Scanln(&option)
	if option == 1 {
		booked = true
	} else if option == 2 {
		booked = false
	} else {
		fmt.Println("Invalid option")
	}

	tInfo := createTicket(pName, dName, booked)
	tInfo.TicketInfo()

}
