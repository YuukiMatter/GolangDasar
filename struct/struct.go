package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var leo Customer
	leo.Name = "leo"
	leo.Age = 30
	leo.Address = "jl.Pintu Kemana saja"

	fmt.Println(leo)

	//literals struct

	budi := Customer{
		Name:    "Budi",
		Address: "Jakarta",
		Age:     20,
	}
	fmt.Println(budi)
}
