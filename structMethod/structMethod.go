package main

import "fmt"

type Mahasiswa struct {
	Nama, Nim, Alamat string
	TanggalLahir      int
}

func (mahasiswa Mahasiswa) sayHello(namaAdmin string) {
	fmt.Println("Hello", namaAdmin, "My Name is ", mahasiswa.Nama)
}
func (a Mahasiswa) sayWelcome() {
	fmt.Println("Selamat Datang", a.Nama)
}

func main() {
	var leo Mahasiswa
	leo.Nama = "Leo"
	leo.TanggalLahir = 23232
	leo.Nim = "12323232"

	leo.sayHello("Admin01")
	leo.sayWelcome()
}
