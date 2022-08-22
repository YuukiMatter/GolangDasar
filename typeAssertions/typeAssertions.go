package main

import "fmt"

func random() interface{} {
	return 123
}

func main() {

	// mengkonversi type data menjadi tipe data yg diinginkan, sering kali digunakan kasus ketika bertemu data interface kosong
	// jika tipe data interface berbeda maka akan terjadi panic dan solusinya menggunakan switch

	var result interface{} = random()
	// var resultString string = result.(string)
	// fmt.Println(resultString)

	switch value := result.(type) {
	case string:
		fmt.Println("value", value, "is string")
	case int:
		fmt.Println("value", value, "is Integer")
	default:
		fmt.Println("unknown type")

	}
}
