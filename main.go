package main

import (
	"Golang/account"
	"Golang/modules"
	"fmt"
)

func main() {
	modules.Greet()

	meli := account.NewBankAccount("farid", "meli12345", 100)

	fmt.Println(*meli)
	fmt.Println(meli.Balance)
	info, err := meli.Info("meli12345")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(info)
}
