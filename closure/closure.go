package main

import "fmt"

//Closure : sebuah blok scope bisa berinteraksi dengan data sekitarnya dalam skope yang sama
func main() {
	name := "Leo"
	counter := 0

	increment := func() {
		// name := "budi"
		fmt.Println("Increment")
		counter++
		// fmt.Println(name)
	}

	increment()
	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(name)
}
