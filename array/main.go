package main

import "fmt"

// func main() {
// 	var a = [...]int{-1, 0, 1, 2, 3, 4}
// 	var s []int

// 	// s = make([]int, 0)

// 	// if s == nil {
// 	// 	fmt.Println("it's nill")
// 	// } else {
// 	// 	fmt.Println("it's not nil")
// 	// }
// 	//*********************************************************
// 	// s = append(s, -1, 0, 1, 2)
// 	// fmt.Println(len(s))
// 	//--------------------------------------------------------

// 	s = a[1:3]
// 	fmt.Println(len(s))
// 	fmt.Println(s)
// 	fmt.Println(cap(s))
// }

// func abc() {
// 	s := []string{"apple", "banana", "coconut", "durian", "elderberries", "figs", "grapes"}

// 	// TODO: slice the s to the new variable, only "apple", "banana" and "coconut" needed
// 	fmt.Print(s[0:3])
// 	// [apple banana coconut]
// }

// func efg() {
// 	s := []string{"apple", "banana", "coconut", "durian", "elderberries", "figs", "grapes"}
// 	// TODO: slice the s to the new variable, only "elderberries", "figs" and "grapes" needed

// 	fmt.Print(s[4:7])
// 	// * [elderberries figs grapes]
// }

// func cde() {
// 	s := []string{"apple", "banana", "coconut", "durian", "elderberries", "figs", "grapes"}

// 	// TODO: slice the s to the new variable, only "coconut", "durian" and "elderberries" needed

// 	fmt.Print(s[2:5])
// 	// * [coconut durian elderberries]
// }

func main() {
	max := maxInt(1, 10, 100, 1000)

	fmt.Println(max)
}

func maxInt(nums ...int) int {
	var max int
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}
