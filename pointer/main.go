package main

import "fmt"

func main() {
	name := "farid"
	fmt.Println(name)
	fmt.Println(&name)

	number := 12
	pNumber := &number

	*pNumber = 32
	fmt.Printf("The address of the number: %p\n", pNumber)
	fmt.Printf("Value: %d\n", *pNumber)
	fmt.Println("The number now is", number)

	value := 12

	var ptr *int = &value

	fmt.Println(ptr)
	fmt.Println(*ptr)

	result := add(12, 12)
	fmt.Printf("The address of the result: %p\n", result)
	fmt.Printf("The Result Value: %d\n", *result)
}

func add(a, b int) *int {
	result := a + b

	return &result
}
