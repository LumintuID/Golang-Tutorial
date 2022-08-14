package main

import (
	library "Golang-Tutorial/16-level-access/library"
	f "fmt" // nama alias
)

func main(){
	// public atau dapat diakses diluar packagenya
	library.SayHello("Galuh")

	// private atau tidak dapat diakses diluar packagenya
	// library.sayName("Galuh")

	f.Println()
	siswa1 := library.Siswa{Nama: "Galuh", Nilai: 90}

	f.Println("nama",siswa1.Nama)
	f.Println("nilai", siswa1.Nilai)

}