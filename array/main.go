package main

import "fmt"

func main() {
	var a = [...]int{-1, 0, 1, 2}
	var s []int

	// s = make([]int, 0)

	// if s == nil {
	// 	fmt.Println("it's nill")
	// } else {
	// 	fmt.Println("it's not nil")
	// }
	//*********************************************************
	// s = append(s, -1, 0, 1, 2)
	// fmt.Println(len(s))
	//--------------------------------------------------------

	s = a[1:3]
	fmt.Println(len(s))
	fmt.Println(s)
}
