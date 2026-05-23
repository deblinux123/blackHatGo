package main

import (
	"errors"
	"fmt"
)

func main() {
	hello()
}

func hello() {
	fmt.Println("Hello, Gophers.")
	fmt.Println(add(12, 12))
	// fmt.Println(devide(12, 2))
	result, err := devide(12, 2)

	if err != nil {
		fmt.Println(result, err)
	}

	a := 12
	b := 21

	swap(&a, &b)

	fmt.Printf("A: %d\n", a)
	fmt.Printf("B: %d\n", b)

	sum := sum(1, 2, 3, 4, 5)
	fmt.Printf("Sum: %d\n", sum)
}

func add(a int, b int) int {
	return a + b
}

// pass by value
func devide(a int, b int) (int, error) {
	if b != 0 {
		return a / b, nil
	}

	return 0, errors.New("Second number can't be 0.")
}

// pass by referense
func swap(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

func sum(num ...int) int {
	sum := 0
	for _, v := range num {
		sum += v
	}
	return sum
}
