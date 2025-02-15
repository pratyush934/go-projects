package main

import "fmt"

/*
81. Define an interface for a shape with methods `Area` and `Perimeter`.
82. Write a program that implements the `Shape` interface for a circle and rectangle.
83. Implement a program that demonstrates the use of empty interfaces.
84. Write a program to implement type assertion for an interface.
85. Implement a program that uses type switches with interfaces.
86. Write a program that demonstrates interface embedding.
87. Implement a program that uses an interface to handle different data types.
88. Write a program that demonstrates the `Stringer` interface by customising the `String()`
method for a struct.
89. Implement a program that uses a function to accept any type that implements a specific
interface.
90. Write a program that demonstrates mocking with interfaces in Go.
*/

type Shape interface {
	Area() int
	Perimeter() int
}

type Circle struct {
	radius int
}

type Rect struct {
	l, b int
}

func (c Circle) Area() int {
	return c.radius * c.radius
}

func (r Rect) Area() int {
	return r.l * r.b
}

func (c Circle) Perimeter() int {
	return (int)(c.radius * 2 * 3)
}

func (r Rect) Perimeter() int {
	return 2 * r.l * r.b
}

func question81() {
	var s Shape = Circle{10}
	fmt.Println(s.Area())
	fmt.Println(s.Perimeter())
}

func question83() {
	var a interface{} = "Pratyush"
	fmt.Println(a)
	a = 123
	fmt.Println(a)
}

func question84() {
	var good interface{} = "Pratyush"

	s, ss := good.(string)

	if ss {
		fmt.Println("It is String ", s)
	} else {
		fmt.Println("Not so good")
	}
}

func switchExample(temp interface{}) {
	switch fun := temp.(type) {
	case int:
		fmt.Println("It is an Int", fun)
	case string:
		fmt.Println("It is a string", fun)
	case float64:
		fmt.Println("It is a float64", fun)
	}
}

func question85() {

	var s interface{} = "Mridula"
	switchExample(s)

}

type FirstInterface interface {
	ManyMany()
}

type SecondInterface interface {
	ToMany()
}

type GroupTogether interface {
	FirstInterface
	SecondInterface
}

type BabuRe struct {
	a int
}

func (b BabuRe) ManyMany() {
	fmt.Println("ManyMany", b.a)
}

func (b BabuRe) ToMany() {
	fmt.Println("ToMany", b.a)
}

func question86() {
	var tt GroupTogether = BabuRe{1}
	tt.ToMany()
	tt.ManyMany()
}

type HandleIt interface {
	HandleItBeautiful()
}

type ItIt struct {
	value int
}

func (i ItIt) HandleItBeautiful() {
	fmt.Println(i.value)
}

type StringString struct {
	value string
}

func (s StringString) HandleItBeautiful() {
	fmt.Println(s.value)
}

func question87() {
	ii := ItIt{1}
	ss := StringString{"Hello"}

	var good []HandleIt
	good = append(good, ii, ss)

	fmt.Println(good)
}

type Stringer interface {
	String() string
}

type SayHello struct {
	name    string
	message string
}

func (s SayHello) String() string {
	return "Hello I am " + s.name + " I have a message " + s.message
}

func question88() {
	var goodGood Stringer = SayHello{"Mridula", "Fuck you"}
	fmt.Println(goodGood.String())
}

type ShapeOfYou interface {
	Area() int
}

type Flat struct {
	l int
}

func (f Flat) Area() int {
	return f.l * f.l
}

type Curve struct {
	l int
}

func (c Curve) Area() int {
	return 3 * c.l
}

func question89() {
	var str ShapeOfYou = Flat{19}
	var sstr ShapeOfYou = Curve{12}

	fmt.Println(str.Area())
	fmt.Println(sstr.Area())
}

func main() {
	question89()
}
