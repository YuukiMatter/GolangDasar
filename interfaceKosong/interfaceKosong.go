package main

import "fmt"

//sama seperti type apapun terface{
//} <= karena tidak menyebutkan tipe data jadi akan didefinisikan apapun
//func ups(i int, apapun interface{}) interface{} {

func ups(i int) interface{} {
	if i == 1 {
		return 1

	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func main() {
	var data interface{} = ups(3)
	fmt.Println(data)
}
