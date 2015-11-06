package main 

import (
		"fmt" 
		"time"
)

func main() {

	var firstName string = "Mochamad"
	var lastName, website string

	lastName 	= "Fathoni"
	website 	= "www.tunathoni.com"

	fmt.Printf("My name is : %s %s \n", firstName, lastName)
	fmt.Printf("My Website is : %s \n", website)
	fmt.Printf("Your time location : %s \n", time.Now())

}