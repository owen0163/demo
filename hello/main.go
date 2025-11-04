package main

import "fmt"

// func main() {
//******************* type conversions ******************************
// var price = 9.5
// var paid = 10

// change := float64(paid) - price
// fmt.Println(change)

// var pac = 48
// var say = 30

// tran := float64(pac) / float64(say)
// fmt.Println(tran)
//-------------------------- basic pointer ***********************************

// 	var p *int
// 	var n int = 14

// 	p = &n

// 	*p++
// 	power(p)
// 	fmt.Printf("%p %p\n", &p, p)
// 	fmt.Println(*p)
// }

//	func power(p *int) {
//		fmt.Printf("%p %p\n", &p, p)
//		*p = *p * *p
// }
//-------------------------- basic pointer ***********************************
// func double(n *int) {
// 	// ฟังก์ชันนี้รับตัวชี้ *int และเปลี่ยนค่าของตัวแปรนั้น
// 	*n = *n * 2

// }

// func appendGreeting(s *string) {
// 	// เปลี่ยนค่าของตัวแปร string ที่ถูกชี้ (*s)
// 	*s = fmt.Sprintf("Hi, %s", *s)
// }

// func main() {
// 	var x int = 4
// 	var myName string = "bob"

// 	// 1. ทดสอบ double() กับ x
// 	double(&x) // x ถูกเปลี่ยนจาก 5 เป็น 10

// 	// 2. ทดสอบ appendGreeting() กับ myName
// 	appendGreeting(&myName) // myName ถูกเปลี่ยนจาก "Alice" เป็น "Hi, Alice"

// 	fmt.Println("Value of x after double():", x)
// 	fmt.Println("Value of myName after appendGreeting():", myName)
// }

// func main() {
// 	var n int
// 	var p = &n
// 	m := *p
// 	inc(p)
// 	fmt.Println(m)
// }
// func inc(p *int) {
// 	*p++
// }
////////////////////////////////////////////////// Constants /////////////////////////////////////////////
// const (
// 	byte = 1 << (iota * 10)
// 	kb
// 	mb
// )

//	func main() {
//		fmt.Println(byte)
//		fmt.Println(kb)
//		fmt.Println(mb)
//	}
//
// /////////////////////////////////// functions //////////////////////////////
// func main() {
// 	// fmt.Println(add(1, "owen"))
// 	// fmt.Println(add(1, 2))
// 	fmt.Println(min(1, 2))
// }

// func add(a int, name string) (int, string) {
// 	return a, name
// }

//	func add(a, b int) int {
//		return a + b
//	}
// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }
//****************************************************************************
// func main() {
// 	// fmt.Println(greeting("pallat"))
// 	// fmt.Println(greetingWithAge("pallat", 30))
// 	// fmt.Println(greetingWithAgeAndDrink("pallat", 30, "cola"))
// }

// // greeting("Pallat") should return "Hello, Pallat"
//
//	func greeting(name string) string {
//		return fmt.Sprintf("Hello, %s", name)
//	}
//
// ***********************************************************************
// greetingWithAge("Pallat", 30) should return "Hello, Pallat. You are 30 years old."
//
//	func greetingWithAge(name string, age uint) string {
//		return fmt.Sprintf("Hello, %s. You are %d years old.", name, age)
//	}
//
// ***********************************************************************
// greetingWithAgeAndDrink("Pallat", 30, "Cola") should return "Hello, Pallat. You are 30 years old and your favorite drink is Cola."
// func greetingWithAgeAndDrink(name string, age uint, drink string) string {
// 	return fmt.Sprintf("Hello, %s. You are %d years old and your favorite drink is %s.", name, age, drink)
// }

// ***********************************************************************

// func main() {
// 	var age int
// 	fmt.Scanln(&age)

// 	year := time.Now().Year()
// 	yearOFBirth := year - age

// 	if yearOFBirth > 2012 {
// 		fmt.Println("young")
// 	} else if yearOFBirth >= 1997 {
// 		fmt.Println("GEN Z")
// 	} else {
// 		fmt.Println("HI senior")
// 	}
// }

// ***********************************************************************

// func main() {
// 	var num int
// 	fmt.Print("กรุณาป้อนจำนวนเต็ม: ")
// 	fmt.Scanln(&num)
// 	result := isOdd(num)

// 	fmt.Printf("คุณป้อน: %d\n", num)
// 	fmt.Printf("เป็นเลขคี่หรือไม่: %t\n", result)
// }

// func isOdd(n int) bool {
// 	if n%2 != 0 {
// 		return true
// 	} else {
// 		return false
// 	}

// }

// ***********************************************************************

// func main() {
// 	var age int
// 	fmt.Print("กรุณาป้อนจำนวนเต็ม: ")
// 	fmt.Scanln(&age)

// 	year := time.Now().Year()
// 	yearOFBirth := year - age

// 	switch {
// 	case yearOFBirth > 2012:
// 		fmt.Println("young")
// 	case yearOFBirth >= 1997:
// 		fmt.Println("GENZ")
// 	default:
// 		fmt.Println("Hi senior")
// 	}
// }

// ******************************************************************************

func main() {
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }
	//**********************************************
	// var i := 0
	// for i < 5 {
	// 	fmt.Println(i)
	// 	i++
	// }
	//*******************************************************
	i := 0
loop:
	for {

		if i < 5 {
			fmt.Println(i)
			i++
		} else {
			break loop
		}
	}
}
