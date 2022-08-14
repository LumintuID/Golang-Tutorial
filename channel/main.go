package main

import "fmt"

func printMessage(text chan string)  {
	fmt.Println(<-text)
}

func main()  {
	// buat channel dengan variable data
	var data = make(chan string)

	var sayHello = func (nama string)  {
		var text = fmt.Sprintf("hello %s", nama)
		data <- text
	}

	go sayHello("galuh")
	go sayHello("prahadi")
	go sayHello("gumelar")

	var msg1 = <- data
	fmt.Println(msg1)

	var msg2 = <- data
	fmt.Println(msg2)

	var msg3 = <- data
	fmt.Println(msg3)

	fmt.Println()
	fmt.Println("chanel sebagai parameter")

	var msg = make(chan string)

	for _, val := range []string{"galuh", "prahadi", "gumelar"} {
		go func (d string)  {
			var data = fmt.Sprintf("hello %s", d)
			msg<-data
		}(val)
	}

	for i := 0; i < 3; i++ {
		printMessage(msg)
	}

	
}