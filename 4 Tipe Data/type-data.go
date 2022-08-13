package main

import "fmt"

func main() {

	/*
		kumpulan tipe data di golang

		type data	|	nilai cakupan
		____________ _________________
		uint8		|	0 - 255
		uint16		|	0 - 65535
		uint32		|	0 - 4294967295
		uint64		|	0 - 18446744073709551615
		uint		|	sama dengan uint32 atau uint64 (tergantung proc computer 32 atau 64 bit)
		byte		|	sama dengan uint8
		int8		|	-128 - 127
		int16		|	-32768 - 32767
		int32		|	-2147483648 - 2147483647
		int64		|	-9223372036854775808 - 9223372036854775807
		int			|	sama dengan int32 atau int64 (tergantung proc computer 32 atau 64 bit)
		rune		|	sama dengan int32
		bool		|	true atau false
		float		| berisi nilai desimal example 2.62
	*/

	var name string = "abu mandah"
	var age int8 = 26
	var married = false
	var IPK float32 = 3.20

	fmt.Println(name, "bertipe string")
	fmt.Println(age, "bertipe integer")
	fmt.Println(married, "bertipe boolean")
	fmt.Println(IPK, "bertipe float")
}