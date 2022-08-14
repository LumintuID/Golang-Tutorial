package main

import (
	"fmt"
)

// interface
type hitung interface{
	keliling() int
	luas() int
}

// struct
type persegiPanjang struct{
	panjang int
	lebar int
}



func (p persegiPanjang) keliling() int  {
	return 2 * (p.panjang + p.lebar)
}

func (p persegiPanjang) luas() int  {
	return p.panjang * p.lebar
}

// var interfaceAny map[string]any{} // penulisan interface kosong pada golang terbaru
	
func main()  {
	var bangunDatar hitung = persegiPanjang{10, 2}
	fmt.Println("===== persegi")
	fmt.Println("luas :", bangunDatar.luas())
	fmt.Println("keliling :", bangunDatar.keliling())

	// interface kosong
	fmt.Println()
	fmt.Println("ini hasil interface kosong")

	var rahasia interface{}
	rahasia = "interface berisi string"
	fmt.Println(rahasia)

	rahasia = []string{"data", "ini", "merupakan", "data", "slice"}
	fmt.Println(rahasia)

	rahasia = 12
	fmt.Println("berisi data integer", rahasia)

	
}