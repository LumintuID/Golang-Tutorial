package main

import "fmt"

func kondisi(nilai int8){
	switch {
	case nilai >= 70:
		if nilai >= 85 {
			fmt.Println("Perfect")
		}else{
			fmt.Println("Nice")
		}
	case nilai > 50:
		fmt.Println("not bad")
	default:
		fmt.Println("gak lulus")
	}
}

func main() {
	var nilai int8= 88

	fmt.Println("nilai anda",nilai)
	kondisi(nilai)
}