package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	//alex := person{firstName: "Alex", lastName: "Anderson"}
	//fmt.Println(alex)

	//var alex person
	//alex.firstName = "Alex"
	//alex.lastName = "Anderson"

	//fmt.Println(alex)
	//fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 95000,
		},
	}

	jim.updateName("jimmy")
	jim.print()
}

func (pointertoPerson *person) updateName(newFirstName string) {
	(*pointertoPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
