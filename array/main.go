package main

import "fmt"

func main() {
	var s []int

	s = make([]int, 0)

	if s == nil {
		fmt.Println("it's nill")
	} else {
		fmt.Println("it's not nil")
	}

}
