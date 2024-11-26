package main

import "fmt"

func main() {
	ageIndra := 90

	var agePointer *int
	agePointer = &ageIndra

	fmt.Println("adult age = ", ageIndra)

	editAdultAge(agePointer)
	fmt.Println("age for getAdultAge = ", ageIndra)
}

func editAdultAge(age *int) {
	*age = *age - 18
}
