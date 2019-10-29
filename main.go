package main

import "fmt"

type Person struct {
	EmployeeName string
	EmployeeAge int
}

func main () {
	fmt.Println("No bottles of beer on the wall")
	var per []Person
	per = append(per, Person{
		EmployeeName:"tinh",
		EmployeeAge:17,
	}, Person{
		EmployeeName:"tin1",
		EmployeeAge:23,
	})
	for i := range Count(per){
		fmt.Println("Pass it around, put one up,", i.EmployeeName, "bottles of beer on the wall", i.EmployeeAge)
	}
}

func Count(per []Person) chan Person {
	ch := make(chan Person)

	go func(ch chan Person) {
		for i := 0; i <=  len(per)-1 ; i++ {
			// Blocks on the operation
			ch <- per[i]
		}

		close(ch)
	}(ch)

	return ch
}
