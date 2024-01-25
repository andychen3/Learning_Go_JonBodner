package main

import "fmt"

func main() {
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	e1 := Employee{"Aby", "San", 1}
	e2 := Employee{
		firstName: "TAby",
		lastName:  "Pan",
		id:        2,
	}
	var e3 Employee
	e3.firstName = "MAby"
	e3.lastName = "Kan"
	e3.id = 3

	fmt.Println(e1)
	fmt.Println(e2)
	fmt.Println(e3)

}
