package main

import (
	"fmt"
	"math"
	"strings"
)

func sayHello(text string, name []string) string {
	var nameString = strings.Join(name, " ")
	return text + nameString
}


/*
	fungsi multiple return

	Penggunaan Fungsi math.Pow()
		Fungsi math.Pow() digunakan untuk nilai pemangkatan. math.Pow(2, 3) berarti 2 pangkat 3, hasilnya 8.
	Penggunaan Konstanta math.Pi
		math.Pi adalah konstanta bawaan package math yang merepresentasikan Pi atau 22/7.
*/
func lingkaran(diameter float64)(float64, float64)  {
	// hitung nilai luas 
	luas := math.Pi * math.Pow(diameter/2, 2)

	// hitung nilai keliling
	keliling := math.Pi * diameter

	return luas, keliling
}

/*
	fungsi variadic
*/
func variadik(n ...int) (res float64){
	total := 0

	for _ ,val := range n{
		total += val
	}

	res = float64(total) / float64(len(n))
	return
}


/*
	fungsi sebagai parameter atau biasa disebut callback function
*/
func filter(data []string, callback func(string) bool) []string {
    var result []string
    for _, each := range data {
        if filtered := callback(each); filtered {
            result = append(result, each)
        }
    }
    return result
}

func main() {
	var name = []string{
		"galuh",
		"prahadi",
	}

	result := sayHello("halo ", name)
	fmt.Println(result)

	fmt.Println()
	luas, keliling := lingkaran(6)
	fmt.Println("luas lingkaran ", luas)
	fmt.Println("keliling lingkaran ", keliling)

	fmt.Println()
	angka := []int{2,3,5,7,8,9,4,2,5,7,8,4,2}
	res := variadik(angka...)
	fmt.Println("rata-rata dari",angka,"adalah",res)


	fmt.Println()
	fmt.Println("hasil nilai dari fungsi closure")
	var getMinMax = func(n []int) (int, int) {
        var min, max int
        for i, e := range n {
            switch {
            case i == 0:
                max, min = e, e
            case e > max:
                max = e
            case e < min:
                min = e
            }
        }
        return min, max
    }

    var numbers = []int{2, 3, 4, 3, 4, 2, 3}
    var min, max = getMinMax(numbers)
    fmt.Printf("data : %v\nmin  : %v\nmax  : %v\n", numbers, min, max)

	fmt.Println()
	fmt.Println("hasil nilai dari fungsi filter")
	var data = []string{"wick", "jason", "ethan"}
    var dataContainsO = filter(data, func(each string) bool {
        return strings.Contains(each, "o")
    })
    var dataLenght5 = filter(data, func(each string) bool {
        return len(each) == 5
    })

    fmt.Println("data asli \t\t:", data)
    // data asli : [wick jason ethan]

    fmt.Println("filter ada huruf \"o\"\t:", dataContainsO)
    // filter ada huruf "o" : [jason]

    fmt.Println("filter jumlah huruf \"5\"\t:", dataLenght5)
    // filter jumlah huruf "5" : [jason ethan]
}