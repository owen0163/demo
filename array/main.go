package main

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
//----------------------------------------------------------------------- variadic func
// func main() {
// 	max := maxInt(1, 10, 100, 1000)

// 	fmt.Println(max)
// }

// func maxInt(nums ...int) int {
// 	var max int
// 	for _, v := range nums {
// 		if v > max {
// 			max = v
// 		}
// 	}
// 	return max
// }

// ---------------------------------------------- map -----------------------------
// func main() {
// 	var m map[string]string

// 	m = make(map[string]string)

// 	m["a"] = "apple"
// 	m["b"] = "banana"
// 	m["c"] = "coconut"

// 	for k, v := range m {
// 		fmt.Println(k, v)
// 	}

// 	fmt.Println(m == nil)
// }
//-------------------------------------------------------------

// func main() {
// 	a := wordCount("Apple Banana Apple Banana apple")

// 	fmt.Println(a["apple"])
// }

// func wordCount(s string) map[string]int {
// 	split := strings.Split(s, " ")
// 	result := map[string]int{}
// 	for i := 0; i < len(split); i++ {
// 		result[split[i]] = result[split[i]] + 1
// 	}
// 	return result
// }

// ------------------------------------------------------- structs-----------------------
// type rect struct {
// 	w, d float64
// }

// func main() {
// 	r := rect{
// 		w: 3,
// 		d: 4,
// 	}

// 	area := r.w * r.d

// 	fmt.Println(area)

// }
//--------------------------------------------------------------------------- methods--------------------------------------

// type rect struct {
// 	w, d float64
// }

// func area(r rect) float64 {
// 	return r.d * r.w
// }
// func main() {
// 	r := rect{
// 		w: 3,
// 		d: 4,
// 	}

// 	fmt.Println(area(r))

// }
//------------------ex-----------------------
// type Book struct {
// 	Name   string
// 	Author string
// }

// // Method แบบ Value Receiver ใช้ struct เดิม
// func (b Book) String() string {
// 	return fmt.Sprintf("%s by %s", b.Name, b.Author)
// }

// // Method แบบ Pointer Receiver สามารถแก้ไขค่าจริงใน struct ได้
// func (b *Book) SetName(name string) {
// 	b.Name = name
// 	fmt.Println("Updated Name:", b.Name)
// }

// func main() {
// 	b := Book{
// 		Name:   "Harry Potter",
// 		Author: "J. K. Rowling",
// 	}

// 	fmt.Println(b.String()) // Harry Potter by J. K. Rowling

// 	b.SetName("New Title")  // เปลี่ยนชื่อหนังสือ
// 	fmt.Println(b.String()) // New Title by J. K. Rowling
// }

//-------------------------------------------Pointer Receiver----------------------

// type rect struct {
// 	w, d float64
// }

// func area(r rect) float64 {
// 	return r.d * r.w
// }

// func (r *rect) setWidth(w float64) {
// 	r.w = w
// }

// func main() {
// 	r := rect{
// 		w: 3,
// 		d: 4,
// 	}
// 	r.setWidth(5)

// 	fmt.Println(area(r))

// }
//----------------------------------------------------------- interfaces ---------------------------------------------------

// func main() {
// 	var a any

// 	a = 10
// 	fmt.Printf("%t %v\n", a, a)
// 	a = "ten"
// 	fmt.Printf("%t %v\n", a, a)

//		var s string
//		s = a
//		fmt.Println(s)
//	}
//-------------------------
// type dog string

// func (d dog) Sound() string {
// 	return "โร่ง"
// }

// type cat string

// func (c cat) Sound() string {
// 	return "เมื้ยว"
// }

// type animal interface {
// 	Sound() string
// }

// func playSound(a animal) {
// 	fmt.Println(a.Sound())
// }

// func main() {
// 	var d dog

// 	playSound(d)

// 	var c cat
// 	playSound(c)
// }
//------------------------- ex-----------------------------------------------------------

// type volumer interface {
// 	Volume() float64
// }

// type cube struct {
// 	edge float64
// } // edge x edge x edge
// func (c cube) Volume() float64 {
// 	return c.edge * c.edge * c.edge
// }

// type triangularPrism struct {
// 	base     float64
// 	attitude float64
// 	height   float64
// } // 0.5 x base x attitude x height
// func (t triangularPrism) Volume() float64 {
// 	return 0.5 * t.base * t.attitude * t.height
// }
// func VolumeOf(v volumer) float64 {
// 	return v.Volume()
// }

// func main() {
// 	var c = cube{edge: 4.46}
// 	fmt.Println("Cube Volume:", VolumeOf(c))

// 	var t = triangularPrism{base: 3, attitude: 4, height: 2}
// 	fmt.Println("Triangular Prism Volume:", VolumeOf(t))
// }
