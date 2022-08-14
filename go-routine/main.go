package main

import (
	"fmt"
	"runtime"
)

func print(index int, text string) {
	for i := 1; i < index; i++ {
		fmt.Println(i, text)
	}
}
func main() {
	runtime.GOMAXPROCS(2) // set proc yang active untuk menjalankan program

	go print(5, "halo")
	print(5, "apa kabar")

	var input string
	fmt.Scanln(&input)
}