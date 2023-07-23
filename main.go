package main //everything is a package. So we need to declare it first. Before that use go mod init booking-app to initialize.

import (
	"fmt" //important for print
	"strings"
)

func main() {

	conferenceName := "Go Conference" //Alternative way to declare. But does not work with const and specified types.
	const conferenceTickets = 50
	var remainingTickets = 50
	var bookings []string //array, []is slice
	//or use bookings :[]string{}

	fmt.Println("Welcome to", conferenceName, "Booking Application. Get your Tickets here")
	fmt.Printf("We have a total of %v tickets and %v remaining tickets.\n", conferenceTickets, remainingTickets) //fortmat using placeholder, %T if used returns type.

	for remainingTickets > 0 {
		var userName string
		var lastName string
		var email string
		var userTickets int

		fmt.Print("Enter your First Name:")
		fmt.Scan(&userName)
		fmt.Print("Enter your Last Name:")
		fmt.Scan(&lastName)
		fmt.Print("Enter your e-Mail address:")
		fmt.Scan(&email)
		fmt.Print("Enter the number of Tickets:")
		fmt.Scan(&userTickets)

		//checking user inputs
		isValidName := len(userName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicket := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicket {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, userName+" "+lastName)
			fmt.Printf("The whole booking: %v\n", bookings)
			fmt.Printf("The first Element: %v\n", bookings[0])
			fmt.Printf("The length of slice: %v\n", len(bookings))

			fmt.Println(userName, "booked", userTickets, "tickets. You will recieve your confirmation at", email)
			fmt.Println("Remaining tickets:", remainingTickets)

			firstNames := []string{}
			for _, booking := range bookings { //_is a blank identifier
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Tickets are Sold Out!")
				break
			}
		} else {
			if !isValidName{
				fmt.Println("First Name or Last Name you entered is too Short")
			}
			if !isValidEmail{
				fmt.Println("Invalid Email Address")
			}
			if !isValidTicket{
				fmt.Println("Invalid Ticket")
			}
			fmt.Println("User Data Invalid")
		}
	}
}
