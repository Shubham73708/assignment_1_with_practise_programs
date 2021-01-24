// 9 Structure (User)

package main

import (
	"fmt"
)

//Employee is
type Employee struct {
	firstName string
	lastName  string
	age       int
	salary    int
}

func main() {
	emp1 := Employee{
		firstName: "Shubham",
		age:       25,
		salary:    10000,
		lastName:  "Shahi",
	}

	emp2 := Employee{"Shubh", "Thakur", 24, 10000}

	fmt.Println("Employee 1", emp1)
	fmt.Println("Employee 2", emp2)
}
