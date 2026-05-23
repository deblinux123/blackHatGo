package main

import (
	"fmt"
	"strings"
)

func main() {
	for i := 1; i < 100; i++ {
		if i%15 == 0 {
			fmt.Printf("%d is FizzBuzz\n", i)
		} else if i%5 == 0 {
			fmt.Printf("%d is Fizz\n", i)
		} else if i%3 == 0 {
			fmt.Printf("%d is Buzz\n", i)
		} else {
			fmt.Println(i)
			continue
		}
	}

	// list of score numbers
	fmt.Println(strings.Repeat("=", 30))

	score := [5]int{88, 45, 97, 23, 12}

	min := score[0]
	max := score[0]
	for _, v := range score {
		if v < min {
			min = v
		} else if v > max {
			max = v
		}
	}

	fmt.Printf("Min: %d\nMax: %d\n", min, max)

	// infinit loop

	for {
		fmt.Println("What is your name please enter your name or q for quit: ")
		userInput := ""
		fmt.Scan(&userInput)

		if strings.ToLower(userInput) == "q" || strings.ToLower(userInput) == "quit" {
			fmt.Println("Good By")
			break
		}
	}
}
