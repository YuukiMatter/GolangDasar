package main

import "fmt"

type address struct {
	City, Province, Country string
}

func main() {

	var address1 address = address{"Subang", "Jawa Barat", "Indonesia"}
	var address2 *address = &address1
	var address3 *address = &address1

	address2.City = "Bandung"

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	var address4 *address = new(address)
	address4.Country = "America"
	fmt.Println(address4)
}
