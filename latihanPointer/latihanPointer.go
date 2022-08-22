package main

import "fmt"

type mahasiswa struct {
	Nama, Nim, Kelas string
}

func main() {

	var mahasiswa1 = mahasiswa{"Leo Agustino", "123456", "C-212"}
	var mahasiswa2 *mahasiswa = &mahasiswa1

	mahasiswa2.Nama = "Budi"

	fmt.Println(mahasiswa1)
	fmt.Println(mahasiswa2)

}
