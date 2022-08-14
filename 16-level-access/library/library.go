package library

import "fmt"

var Mahasiswa = struct {
	Nama string
	kelas string
}{}

// fungsi init fungsi yang dijalankan pertama kali sebelum fungsi main
func init(){
	Mahasiswa.Nama = "Galuh"
	Mahasiswa.kelas = "Kelas A"

	fmt.Println(Mahasiswa)
	fmt.Println("hello ini merupakan fungsi init")
	fmt.Println()
}


func SayHello(name string){
	fmt.Println("hello ini public function")
	sayName(name)
}

func sayName(name string){
	fmt.Println("namaku", name)
	fmt.Println("hello ini private function")
}