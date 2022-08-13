package main

import "fmt"

type Siswa struct{
	nama string
	nilai int
	Kelas
}

type Kelas struct{
	namaKelas string
}

func main() {
	// inisialiasi struct
	var siswa1 = Siswa{}

	siswa1.nama = "Galuh"
	siswa1.nilai = 80
	siswa1.namaKelas = "B"

	var siswa2 = Siswa{"Prahadi", 70, Kelas{"A"}}

	var siswa3 = Siswa{nama: "Gumelar"}
	siswa3.namaKelas = siswa1.namaKelas

	fmt.Println("siswa 1", siswa1)
	fmt.Println("siswa 2", siswa2)
	fmt.Println("siswa 3", siswa3)
	fmt.Println("siswa 1 kelas", siswa1.namaKelas)
	fmt.Println("siswa 2 kelas", siswa2.Kelas.namaKelas)

	// rubah nama pada object siswa1 menggunakan pointer
	fmt.Println()
	fmt.Println("penggunaan pointer pada struck")
	var siswa4 *Siswa = &siswa1
	siswa4.nama = "Abu"
	fmt.Println("siswa 1", siswa1.nama)
	fmt.Println("siswa 4", siswa4.nama)

	
}