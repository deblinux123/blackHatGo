package main

import (
	"fmt"
	"strings"
)

func main() {
	var score [4]int
	score[0] = 34
	score[1] = 65
	score[2] = 88
	score[3] = 98

	fmt.Println(score)

	names := [3]string{"farid", "babak", "mobina"}

	fmt.Println(names)

	fmt.Println(strings.Repeat("=", 30))

	for _, name := range names {
		fmt.Println(name)
	}

	// slice
	// var users []string
	users := []string{"user1", "user2", "user3"}
	fmt.Println(users)

	// add new user to slice
	users = append(users, "User4")
	fmt.Println(users)

	// can make slice with make

	num := make([]float64, 3)

	num = append(num, 12)
	fmt.Println(num)
	num[0] = 12.22
	fmt.Println(num)
}
