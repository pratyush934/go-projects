package main

import (
	"fmt"
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

func main() {
	fmt.Println("I am Pratyush, I am good boy")
	babu := Babu{"Mohini", 20}
	fmt.Println(babu)
	question55(&babu)
	fmt.Println(babu)
}
