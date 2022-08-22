package main

import "fmt"

func getFullName() (string, string, string) {
	return "Leo", "Agustino", "Nama akhir"
}

func main() {
	// "_" untuk mengabaikan variabel
	namaDepan, namaTengah, NamaAkhir := getFullName()
	//namaDepan, _, _ := getFullName()

	fmt.Println(namaDepan)
	fmt.Println(namaTengah)
	fmt.Println(NamaAkhir)

}
