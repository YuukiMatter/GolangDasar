package main

import "fmt"

// nill hanya bisa digunakan untuk interface, function, map, channel, pointer.
func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	// var person map[string]string = nil
	// fmt.Println(person)
	var person map[string]string = NewMap(("Leo"))
	if person == nil {
		fmt.Println(("Data Kosong"))
	} else {
		fmt.Println((person))
	}
}
