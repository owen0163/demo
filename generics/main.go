package main

import "fmt"

type number interface {
	~int | float64
}

func max[t number](a, b t) t {
	if a > b {
		return a
	}
	return b
}

type numeric int

func main() {
	var a, b numeric = 10, 99
	fmt.Println(max(a, b))
}
