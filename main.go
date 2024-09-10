package main //package needed

import (
	"fmt"     //needed for print functions
	"strings" //needed to split strings and other string fucntions
	"booking-app/helper" //import by using module name and then package name (path)
)

//pacakge level variables

var conferenceName = "Go Conference" //use var when defining a variable
// conferenceName2 := "Go Conference" //alt way to define variable
const conferenceTickets uint = 50 //constant variable
var remainingTickets uint = 50

// var names = [50]string{"Brandyn", "Aylah", "Luka"} //array
var bookings []string //slice (a dynmaically growing array i.e. list in Java)
// var bookings2 = []string{} alt way of defining an array

var userName string //type has to be defined if nothing is assigned at declaration

func main() {

	greetUsers()

	for remainingTickets > 0 { //for can handle boolean expression

		firstName, lastName, email, userTickets := getUserInfo()

		fmt.Printf("The size of bookings array is %v\n", len(bookings))

		fmt.Printf("Pointer value of userName is %p\n", &userName)

		fmt.Printf("conferenceTickets is %T, remainingTickets is %T, and conferenceName is %T\n\n", conferenceTickets, remainingTickets, conferenceName) //print datatype of variable

		isValidName, isValidEmailAddress, isUserTicketsValid := helper.ValidateUserInput(firstName, lastName, email, uint(userTickets), remainingTickets)

		if isValidName && isValidEmailAddress && isUserTicketsValid {

			bookTicket(firstName, lastName, email, userTickets)

			firstNames := getFirstNames()
			fmt.Printf("The following people have booked: %v\n\n", firstNames)

			var noTicketsRemaning bool = remainingTickets == 0 //bool expression an be a variable

			if noTicketsRemaning {
				//end program
				fmt.Println("The conference is booked up. Please come back next year!")
			}
		} else {
			if !isValidName {
				fmt.Println("First or last name is too short")
			}

			if !isValidEmailAddress {
				fmt.Println("Email address entered does not contain @")
			}
			if !isUserTicketsValid {
				fmt.Println("The number of tickets you entered is invalid")
			}

		}

	}

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application!\n\n", conferenceName)
	// fmt.Println("Welcome to", conferenceName, "booking application") //spaces automatically included
	// fmt.Println("There are a total of", conferenceTickets, "and there are", remainingTickets, "available")
	fmt.Printf("There are a total of %v and there are %v available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {

	firstNames := []string{} // alt way of defining a slice

	for _, bookings := range bookings { //for each loop. Get index of slice using range. User underscore to replace index that we are not using
		var names = strings.Fields(bookings)
		firstNames = append(firstNames, names[0])
	}

	return firstNames

}


func getUserInfo() (string, string, string, int) {

	var userTickets int
	var firstName string
	var lastName string
	var email string

	fmt.Print("Enter your firstname: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter your email address: ")
	fmt.Scan(&email)
	fmt.Print("Enter how many tickets you want: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(firstName string, lastName string, email string, userTickets int) {
	bookings = append(bookings, firstName+" "+lastName) // add element to slice

	remainingTickets -= uint(userTickets) //type cast due to different type calculation
	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v shortly!\n\n", firstName, lastName, userTickets, email)
	fmt.Printf("There are %v tickets remaining\n", remainingTickets)
}

