package main

import "fmt"

func main() {
	var conferenceName = "João"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome, %v!\n", conferenceName)
	fmt.Println("Get your tickets here to attend.")
	fmt.Printf("We have a total of %v tickets and %v tickets available.\n", conferenceTickets, remainingTickets)

}