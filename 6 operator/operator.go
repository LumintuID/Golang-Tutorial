package main

import (
	"fmt"
)

func aritmatika(a int8, b int8)  {
	/*
		Tanda 	|	Penjelasan
		+ 		|	penjumlahan
		- 		|	pengurangan
		* 		|	perkalian
		/ 		|	pembagian
		% 		|	modulus / sisa hasil pembagian
	*/

	penjumlahan := a + b
	fmt.Println("hasil penjumlahan nilai",a,"dan",b,"adalah", penjumlahan)

	pengurangan := a - b
	fmt.Println("hasil pengurangan nilai",a,"dan",b,"adalah", pengurangan)

	perkalian	:= a * b
	fmt.Println("hasil perkalian nilai",a,"dan",b,"adalah", perkalian)

	pembagian	:= a / b
	fmt.Println("hasil pembagian nilai",a,"dan",b,"adalah", pembagian)

	modulus		:= a % b
	fmt.Println("hasil modulus nilai",a,"dan",b,"adalah", modulus)
}

func perbandingan(a int8, b int8)  {
	/*
		Tanda 	Penjelasan
		------------------------------------------------------------------
		== 		apakah nilai kiri sama dengan nilai kanan
		!= 		apakah nilai kiri tidak sama dengan nilai kanan
		< 		apakah nilai kiri lebih kecil daripada nilai kanan
		<= 		apakah nilai kiri lebih kecil atau sama dengan nilai kanan
		> 		apakah nilai kiri lebih besar dari nilai kanan
		>= 		apakah nilai kiri lebih besar atau sama dengan nilai kanan	
	*/

	sama := a == b
	fmt.Println("nilai",a,"sama dengan nilai",b,  sama)

	tidakSama := a != b
	fmt.Println("nilai",a,"tidak sama dengan nilai",b, tidakSama)

	lebihKecil	:= a < b
	fmt.Println("nilai",a,"lebih kecil dari nilai",b, lebihKecil)

	lebihBesar	:= a > b
	fmt.Println("nilai",a,"lebih besar dari nilai",b, lebihBesar)

}

func logika(left bool, right bool){
	/*
		Tanda 	Penjelasan
		&& 		kiri dan kanan
		|| 		kiri atau kanan
		! 		negasi / nilai kebalikan
	*/

	var leftAndRight = left && right
	fmt.Printf("left && right \t(%t) \n", leftAndRight)

	var leftOrRight = left || right
	fmt.Printf("left || right \t(%t) \n", leftOrRight)

	var leftReverse = !left
	fmt.Printf("!left \t\t(%t) \n", leftReverse)
}

func main() {
	fmt.Println("contoh operator aritmatika ")
	aritmatika(11, 2)
	fmt.Println("--------------------------------")
	fmt.Println("contoh operator perbandingan ")
	perbandingan(11, 2)
	fmt.Println("--------------------------------")
	fmt.Println("contoh operator logika ")
	logika(true, false)
}