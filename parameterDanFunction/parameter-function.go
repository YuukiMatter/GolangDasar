package main

import "fmt"

func namaMahasiswa(nama_depan string, nama_belakang string) {
	fmt.Println("Nama Mahasiswa : ", nama_depan, nama_belakang, " Ruang kelas 333")
}

func main() {
	//fmt.Println(namaMahasiswa("Leo", "Agustino"))
	namaMahasiswa("Leo", "Agustino")
}
