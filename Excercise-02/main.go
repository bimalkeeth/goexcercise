package main

import "fmt"

type ContactInfo struct {
	email   string
	zipCode string
}

type Person struct {
	firstName string
	lastName  string
	contact   ContactInfo
}

func main() {
	alex := Person{
		firstName: "Alex",
		lastName:  "Anderson",
		contact: ContactInfo{
			email:   "bimalkeeth@gmail.com",
			zipCode: "4563",
		}}

	alex.UpdateName("Jimmy")
	fmt.Println(alex)

}

func (p *Person) UpdateName(name string) {
	(*p).firstName = name
}
