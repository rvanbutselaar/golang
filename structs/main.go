package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// peter := person{"Peter", "Anderson"}
	// alex := person{
	// 	firstName: "Alex",
	// 	lastName:  "Anderson",
	// }

	// fmt.Println(alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	jim.print()
	jim.updateName("Jimmy")
	jim.print()
}

func (pointerToPereson *person) updateName(newFirstName string) {
	(*pointerToPereson).firstName = newFirstName
}

// Een reciever functie die je kan gebruiken op alles van type person struct
func (p person) print() {
	fmt.Printf("%+v", p)
}
