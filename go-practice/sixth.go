package main

import (
	"fmt"
	"unsafe"
)

/*
51. Write a program that demonstrates pointer declaration and dereferencing.
52. Implement a function that takes a pointer to an integer and doubles its value.
53. Write a program that swaps two variables using pointers.
54. Implement a program that demonstrates passing a pointer to a function.
55. Write a program that creates a pointer to a struct and modifies its fields.
56. Implement a program to demonstrate nil pointers in Go.
57. Write a program to create a pointer slice and modify the elements through dereferencing.
58. Implement a program that shows the difference between passing by value and passing
by reference.
59. Write a program to demonstrate the use of unsafe pointers in Go.
60. Implement a program that demonstrates pointer arithmetic.
*/

func question51() {

	var goodBoy int = 10

	var pratyush *int = &goodBoy

	fmt.Println(pratyush)
	fmt.Println(*pratyush)

}

func question52(goodBoy *int) *int {

	var tototo int = *goodBoy
	tototo = tototo * 2

	return &tototo
}

func question53(a, b *int) (*int, *int) {

	temp := *a
	*a = *b
	*b = temp

	*a, *b = *b, *a

	return a, b
}

func question54(a *int) *int {
	var value = *a + 10
	return &value
}

type Babu struct {
	name string
	age  int
}

func question55(b *Babu) {
	b.name = "Lucy"
	b.age = 100
}

func question56() {

	var a int = 10

	p := &a

	fmt.Println(*p)
	fmt.Println(p)

	p = nil

	fmt.Println(p)

}

func question57() {

	a, b, c := 10, 20, 30
	var pointerSlice []*int = []*int{&a, &b, &c}

	for i, value := range pointerSlice {
		fmt.Println("The value at ", i, " is ", *value)
	}

	for _, value := range pointerSlice {
		*value *= 2
	}

	for i, value := range pointerSlice {
		fmt.Println("The value at ", i, " is ", *value)
	}
}

type Smriti struct {
	name string
	age  int
	love string
}

func question58(s *Smriti) {
	s.name = "Smriti"
	s.age = 25
	s.love = "Pratyush"
}

func question59() {

	//will look into it
	var num int = 42

	ptr := unsafe.Pointer(&num)

	uintPtr := uintptr(ptr)

	offSetPtr := unsafe.Pointer(uintPtr + 0)

	modifiedValue := (*int)(offSetPtr)

	*modifiedValue = 199

	fmt.Println("Original value ", num)
	fmt.Println("Modified value is", *modifiedValue)
}

func question60() {
	vals := []int{10, 20, 30, 40}

	ptrStart := unsafe.Pointer(&vals[0])
	itemSize := unsafe.Sizeof(vals[0])

	for i := 0; i < len(vals); i++ {
		item := *(*int)(unsafe.Add(ptrStart, uintptr(i)*itemSize))
		println(item)
	}
}

func main() {
	fmt.Println("I am Pratyush, I am good boy")
	question60()

}
