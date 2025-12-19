package main

// type number interface {
// 	~int | float64
// }

// func max[t number](a, b t) t {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// type numeric int

// func main() {
// 	var a, b numeric = 100, 99
// 	fmt.Println(max(a, b))
// }
//----------------------------   First Class Functions ------------------------------

// type number interface {
// 	~int | float64
// }

// func max[t number](a, b t) t {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// type numeric int

// func main() {
// 	var a, b float64 = 100, 99

// 	var maximun = max[float64]

// 	fmt.Println(maximun(a, b))

// 	fn := func() int {
// 		return 10
// 	}

// 	fmt.Println(fn())
// }

//------------------------------------------ Closure Function ---------------------------------------------

// func oddEven(n int, fn func(int) bool) {
// 	fmt.Println(fn(n))
// }

// func newCounter() func() int {
// 	i := 0
// 	return func() int {
// 		i++
// 		return i
// 	}
// }
// func main() {
// isOdd := func(n int) bool {
// 	return n%2 == 1
// }
// isEven := func(n int) bool {
// 	return n&1 == 0
// }
// oddEven(5, isEven)
// 	fn := newCounter()

// 	fmt.Println(fn())
// 	fmt.Println(fn())
// 	fmt.Println(fn())
// }
//-------------------------------------Type Assertions and Switch type----------------------------------------------------

// func main() {
// 	var a any

// 	a = 10

// 	decision(a)
// }

// func decision(a any) {
// if v, ok := a.(int); ok {
// 	fmt.Println("it's int", v)
// 	var n int
// 	n = v
// 	fmt.Println(n)
// }
// 	switch a.(type) {
// 	case int:
// 		fmt.Println("it's int")
// 	default:
// 		fmt.Println("unknown")
// 	}
// }

//-----------------------------------------Defer-------------------------------------------------------
