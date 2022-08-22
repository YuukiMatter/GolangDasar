package main

import "fmt"

type Filter func(string) string

func filterNama(nama string, filter Filter) {
	namefiltered := filter(nama)
	fmt.Println("Hello", namefiltered)
}

func spamFilter(nama string) string {
	if nama == "Anjing" {
		return "..."
	} else {
		return nama
	}

}

func main() {
	filterNama("Leo", spamFilter)
	//filterNama("Anjing", spamFilter)
}
