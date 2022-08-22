package main

import "fmt"

func main() {
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}
	fmt.Println("/n")
	slice := []string{"Mobil", "Sepeda", "Pesawat"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, value := range slice {

		//menggunakan "_" untum memberitahu bahwa variabel tidak digunakan
		//fmt.Println(value)
		fmt.Println("Index", i, "=", value)
	}

	person := make(map[string]string)
	person["name"] = "Eko"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
