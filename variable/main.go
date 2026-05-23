package main

import (
	"fmt"
	"reflect"
)

func main() {
	// typed variable
	var name string = "Farid"
	var age int = 23

	fmt.Println(name)
	fmt.Println(age)

	fmt.Println(reflect.TypeOf(name))
	fmt.Println(reflect.TypeOf(age))

	var score, total int
	score, total = 89, 120
	fmt.Println(score, total)

	// untyped variable
	userName := "babak"
	password := "******"

	fmt.Printf("User Name: %s\nPassword: %s\n", userName, password) // format string
}
