package main

import "fmt"

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTicktets uint = 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTicktets, conferenceName)

	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTicktets)
	fmt.Println("Get your tickets here!")

	var userName string
	var userTickets int
	// ask user for name

	userName = "Tom"
	userTickets = 2
	fmt.Println(userName)
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

}
