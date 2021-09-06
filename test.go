package main

import "fmt"

func main() {
	a := make([]int, 4)
	fmt.Printf("%p => ", &a)
	fmt.Println(a)
	a = append(a, 1)
	fmt.Printf("%p => ", &a)
	fmt.Println(a)

}
