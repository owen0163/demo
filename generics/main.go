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
//----------------------------------------------------------Defer-------------------------------------
//----------------------------- Goroutine -------------------------------------------

// var i int

// func main() {
// 	start := time.Now()
// 	for range 5 {
// 		go slow(func() {
// 			fmt.Println(i)
// 		})
// 	}
// 	for {
// 		if i >= 5 {
// 			break
// 		}
// 	}
// 	fmt.Println(time.Since(start))
// }

// func slow(fn func()) {
// 	time.Sleep(100 * time.Millisecond)
// 	i++
// 	fn()
// }

//----------------------------Wait goroutines and Data Race---------------------------------------------

// var mux sync.Mutex
// var i int

// func main() {
// 	start := time.Now()
// 	for range 5 {
// 		go slow(func() {
// 			mux.Lock()
// 			fmt.Println(i)
// 			mux.Unlock()
// 		})
// 	}
// 	for {
// 		mux.Lock()
// 		if i >= 5 {
// 			break
// 		}
// 		mux.Unlock()
// 	}
// 	fmt.Println(time.Since(start))
// }

// func slow(fn func()) {
// 	time.Sleep(100 * time.Millisecond)
// 	mux.Lock()
// 	i++
// 	mux.Unlock()
// 	fn()
// }
//-------------------------------------------------------
// type data struct {
// 	i   int
// 	mux sync.Mutex
// }

// var asset data

// func main() {
// 	start := time.Now()
// 	for range 5 {
// 		go slow(func() {
// 			asset.mux.Lock()
// 			fmt.Println(asset.i)
// 			asset.mux.Unlock()
// 		})
// 	}
// 	for {
// 		asset.mux.Lock()
// 		if asset.i >= 5 {
// 			break
// 		}
// 		asset.mux.Unlock()
// 	}
// 	fmt.Println(time.Since(start))
// }

// func slow(fn func()) {
// 	time.Sleep(100 * time.Millisecond)
// 	asset.mux.Lock()
// 	asset.i++
// 	asset.mux.Unlock()
// 	fn()
// }

// --------------------------------------- channel----------------------------------------------------
// bufferchannel---------
// func main() {
// 	var ch chan int
// 	ch = make(chan int, 1)
// 	ch <- 9
// 	<-ch
// 	ch <- 10

// 	fmt.Println(<-ch)
// }

// nobufferchannel------------------
// func main() {
// 	var ch chan int
// 	ch = make(chan int)

// 	go func() {
// 		ch <- 9
// 	}()

// 	fmt.Println(<-ch)
// }
