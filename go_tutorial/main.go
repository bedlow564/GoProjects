package main //package needed

import (
	"booking-app/helper" //import by using module name and then package name (path)
	"fmt"                //needed for print functions
	"time" //for using time methods
	// "strconv" //needed for converting numbers to string
	// "strings" //needed to split strings and other string fucntions
	"sync"
)

//pacakge level variables

var conferenceName = "Go Conference" //use var when defining a variable
// conferenceName2 := "Go Conference" //alt way to define variable
const conferenceTickets uint = 50 //constant variable
var remainingTickets uint = 50

// var names = [50]string{"Brandyn", "Aylah", "Luka"} //array
// var bookings = make([]map[string]string, 0) //slice (a dynmaically growing array i.e. list in Java) //list of maps
var bookings = make([]UserData, 0) //create a list of UserData structs (i.e. java classes)
// var bookings2 = []string{} alt way of defining an array

type UserData struct {
	firstName string
	lastName string
	email string
	numberofTickets uint
}

var userName string //type has to be defined if nothing is assigned at declaration

var wg = sync.WaitGroup{} //variable for making main routine wait for go routines

func main() {

	greetUsers()

	for remainingTickets > 0 { //for can handle boolean expression

		firstName, lastName, email, userTickets := getUserInfo()

		fmt.Printf("Pointer value of userName is %p\n", &userName)

		fmt.Printf("conferenceTickets is %T, remainingTickets is %T, and conferenceName is %T\n\n", conferenceTickets, remainingTickets, conferenceName) //print datatype of variable

		isValidName, isValidEmailAddress, isUserTicketsValid := helper.ValidateUserInput(firstName, lastName, email, uint(userTickets), remainingTickets)

		if isValidName && isValidEmailAddress && isUserTicketsValid {

			bookTicket(firstName, lastName, email, userTickets)
			wg.Add(1) //add only 1 go rountine to wait for (increases counter)
			go sendTIcket(userTickets, firstName, lastName, email) //makes method run concurrently (creates a new thread for every method called)

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

		fmt.Printf("The size of bookings array is %v\n\n", len(bookings))
		wg.Wait()

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
		
		firstNames = append(firstNames, bookings.firstName) //add first name from struct
	}

	return firstNames

}


func getUserInfo() (string, string, string, uint) {

	var userTickets uint
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

func bookTicket(firstName string, lastName string, email string, userTickets uint) {

	var user = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberofTickets: userTickets,
	}

	//add data to map 

	// var user = make(map[string]string)
	// user["firstName"] = firstName
	// user["lastName"] = lastName
	// user["email"] = email
	// user["userTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, user) // add element to slice

	fmt.Printf("List of bookings is %v\n", bookings)

	remainingTickets -= uint(userTickets) //type cast due to different type calculation
	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v shortly!\n\n", firstName, lastName, userTickets, email)
	fmt.Printf("There are %v tickets remaining\n", remainingTickets)
}


func sendTIcket(userTickets uint, firstName string, lastName string, email string) {
	var ticket = fmt.Sprintf("%v tickets for %v %v ", userTickets, firstName, lastName) //Allows to save string in a variable
	fmt.Println("*******************")
	fmt.Printf("Sending ticket:\n%v \nto email address %v\n", ticket, email)
	fmt.Println("*******************")

	time.Sleep(10 * time.Second) //Pause thread for 10 seconds
	wg.Done() //decrements go routine counter
}

