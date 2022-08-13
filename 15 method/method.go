package main

import (
	"fmt"
	"strings"
)


type Siswa struct{
	nama string
	nilai int
}

func (s Siswa) sayHello()  {
	fmt.Println("hello", s.nama)
}

func (s Siswa) getNama(i int) (nama string)  {
	nama = strings.Split(s.nama, " ")[i-1]
	return
}

func (s Siswa) ubahNama1(nama string)  {
	fmt.Println("---> ubahNama1 dijalankan, nama berubah menjadi", nama)
    s.nama = nama
}

func (s *Siswa) ubahNama2(nama string)  {
	fmt.Println("---> ubahNama2 dijalankan, nama berubah menjadi", nama)
    s.nama = nama
}

func main() {
	var siswa1 = Siswa{nama: "Galuh Prahadi", nilai: 70}
	fmt.Println("siswa 1")
	siswa1.sayHello()
	nama := siswa1.getNama(1)
	println("nama panggilan ", nama)

	var siswa2 = Siswa{nama: "Abu Mandah", nilai: 80}
	fmt.Println("siswa 2")
	siswa2.sayHello()
	nama = siswa2.getNama(1)
	println("nama panggilan ", nama)

	fmt.Println()
	fmt.Println("method dengan pointer")
	fmt.Println("nama siswa1 sebelum diganti ", siswa1.nama)

	siswa1.ubahNama1("Gumelar")
	fmt.Println("nama siswa1 sesudah menjalankan fungsi ubahNama1 diganti ", siswa1.nama)

	siswa1.ubahNama2("Abu")
	fmt.Println("nama siswa1 sesudah menjalankan fungsi ubahNama2 diganti ", siswa1.nama)


	
}