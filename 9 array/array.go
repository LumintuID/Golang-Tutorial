package main

import "fmt"

func main() {
	var fruits [4]string

	// cara horizontal
	fruits  = [4]string{"apple", "grape", "banana", "melon"}

	// cara vertikal
	fruits  = [4]string{
		"apple",
		"grape",
		"banana",
		"melon",
	}

	fmt.Println(fruits)

	// tanpa inisialisasi nilai
	var numbers = [...]int{2, 3, 2, 4, 3}

	fmt.Println("data array \t:", numbers)
	fmt.Println("jumlah elemen \t:", len(numbers))

	fmt.Println()
	fmt.Println("menampilkan data menggunakan for loop")
	for i := 0; i < len(numbers); i++ {
		fmt.Println("angka ",numbers[i])
	}

	fmt.Println()
	fmt.Println("menampilkan data menggunakan range")
	for index, value := range fruits {
		fmt.Println("index", index, "nilai", value)
	}
}