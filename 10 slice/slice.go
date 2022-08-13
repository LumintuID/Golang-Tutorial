package main

import "fmt"

func main()  {
	member := []string{"galuh", "prahadi", "gumelar"}

	newMember := make([]string, 5)
	
	copy(newMember, member[:1])

	newMember = append(newMember, "adit", "coba")

	fmt.Println(member)
	fmt.Println(newMember)
	fmt.Println("length", len(member))
	fmt.Println("capacity", cap(member))
}