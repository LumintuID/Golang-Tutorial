package main

import "fmt"

func main()  {


	/*
		inisialisasi atau pembuatan map

		var bulan = map[string]int{}
		var bulan = make(map[string]int)
		var bulan = *new(map[string]int)
	*/

	var bulan = map[string]int{
		"januari":  1,
		"februari": 2,
		"maret":    3,
		"april":    3,
	}
	
	for key, val := range bulan {
		fmt.Println(key, "  \t:", val)
	}

	fmt.Println()

	// hapus nilai january
	delete(bulan, "januari")
	fmt.Println("hasil setelah di hapus")
	for key, val := range bulan {
		fmt.Println(key, "  \t:", val)
	}

	fmt.Println()
	fmt.Println("cek nilai ada atau tidak")
	bl := "agustus"
	var value, isExist = bulan[bl]
	if isExist {
		fmt.Println(bl, "mempunyai nilai", value)
	}else{
		fmt.Println(bl, "is not exist")
	}

	// Kombinasi Slice & Map
	fmt.Println()
	var month = []map[string]string{
		{"bulan":"januari", "tanggal": "1",},
		{"bulan":"februari", "tanggal": "20",},
		{"bulan":"maret", "tanggal": "12",},
		{"bulan":"april", "tanggal": "12",},
	}

	for _, val := range month{
		fmt.Println("bulan",val["bulan"],"tanggal",val["tanggal"])
	}
}