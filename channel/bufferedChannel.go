package main

import "fmt"

func main()  {
	channel := make(chan int, 1)

	go func(){
		for {
			i := <-channel
			fmt.Println("terima data", i)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("kirim data", i)
		channel <- i
	}
}