package main

import (
	ticket_window "cli-app/pkg/ticket-window"
	"fmt"
	"strconv"
)

func main() {
	user := ""
	selectedWindowId := 0

	windows := ticket_window.GenerateTicketWindows()
	// concerts := concert.GenerateConcerts()
	greeting(&user)
	listWindows(windows, &selectedWindowId)
}

func greeting(user *string) {
	fmt.Println("Welcome to the concert ticketing system!")
	fmt.Println("Please enter your name:")
	fmt.Scan(user)
	fmt.Printf("Hello, %s!\n", *user)
}

func listWindows(windows []ticket_window.TicketWindow, selectedWindow *int) {
	input := ""
	openWindows := ticket_window.FilterOpenWindows(windows)
	if len(openWindows) == 0 {
		fmt.Println("Sorry, there are no available windows to buy from.")
		return
	}

	fmt.Printf("We have %v available windows you can buy form: \n", len(openWindows))
	for _, window := range openWindows {
		fmt.Printf("%d. %s\n", window.ID, window.Name)
	}
	fmt.Println("Please select a window to buy from:")
	fmt.Scan(input)
	id, err := strconv.Atoi(input)
	if err != nil || id < 0 || id >= len(openWindows) {
		fmt.Println("Invalid input. Please try again.")
		listWindows(windows, selectedWindow)
		return
	}
	*selectedWindow = id
}
