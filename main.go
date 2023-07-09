package main

import (
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go Conference"
const conferenceTickets = 50
var remainingTickets uint = 50

var bookings = make([]userData, 0)

type userData struct{
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}
func main(){

	greetUsers()	
	
		// ask user for their name
	firstName, lastName, email, userTickets := getUserInput()

	isValidName, isValidEmail, isValidTickets := validateUserInput(firstName, lastName,email,userTickets,remainingTickets)
		
	if isValidName && isValidEmail && isValidTickets{
		bookTickets( userTickets,firstName, lastName,email)
		go sendTicket( userTickets,firstName, lastName,email)
	
        wg.Add(1)
        firstNames := printFirstNames()
		fmt.Printf("These are all our bookings: %v\n", firstNames)

		
		if remainingTickets == 0{
			fmt.Println("Our conference is booked out. Come back next year.")
			// break
		}
		}else{
		if !isValidName{
			fmt.Println("first name or last name is too short")
		}
        if !isValidEmail{
			fmt.Println("email address you entered is not valid")
		}
		if !isValidTickets{
			fmt.Println("number of tickets you entered is invalid")
		}
	}
    wg.Wait()
	}
	
func greetUsers(){
	fmt.Println()
	fmt.Printf("Welcome to %v booking application\n",conferenceName)
	fmt.Printf("we have total of %vtickets and %v tickets are still available\n", conferenceTickets,remainingTickets)
	fmt.Println("Get your tickets here to attend !!")
}

func printFirstNames() []string{
	firstNames := []string{}
	for _, booking := range bookings{
		firstNames = append(firstNames,booking.firstName )
	}
    return firstNames
}

func bookTickets(userTickets uint,firstName string, lastName string,email string){
	remainingTickets = remainingTickets - userTickets

	var userData = userData{
		firstName:		  firstName,
		lastName:		  lastName,
		email:    		  email,
		numberOfTickets:  userTickets,
		
	}
	bookings = append(bookings,userData)
	fmt.Println()
	fmt.Printf("list of bookings is : %v\n", bookings)
	
	fmt.Printf("Thank you %v %v for booking %v tickets.\nYou will receive a confirmation email at %v.\n", firstName,lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for %v conference.\n\n", remainingTickets, conferenceName)

}
func getUserInput() (string, string, string, uint){
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("\nEnter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your Email address: ")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets : ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func sendTicket(userTickets uint, firstName string, lastName string, email string){
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###################################################")
	fmt.Printf("Sending ticket: \n %v \nto email address %v\n", ticket, email)
	fmt.Println("###################################################")
	wg.Done()
}
