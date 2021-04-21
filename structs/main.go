package main

import "fmt"

type ContactInfo struct {
	email string
	zipCode int
}

type Person struct {
	firstName string
	lastName string
	ContactInfo
}

func (p Person) print() {
	fmt.Printf("%+v\n", p)
}

func (pointerToPerson *Person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func main() {
	 alex := Person{
		 firstName: "Alex",
		 lastName: "Anderson", 
		 ContactInfo: ContactInfo{
			 email: "alex.anderson@gmail.com",
			 zipCode: 51403,
			},
		}
	alex.print()

	alex.updateName("Jim")
	alex.print()
}