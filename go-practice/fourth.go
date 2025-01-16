package main

import (
	"fmt"
	"math"
)

/*
31. Write a program that defines a function to calculate the sum of two numbers.
32. Implement a program that uses a function to calculate the area of a circle.
33. Write a program to create a function that returns multiple values.
34. Implement a recursive function to calculate the Fibonacci sequence.
35. Write a program to demonstrate passing a slice to a function.
36. Implement a function that accepts a variable number of arguments (variadic function).
37. Write a program that demonstrates function closures.
38. Implement a program that uses an anonymous function to double a number.
39. Write a program to create a higher-order function that accepts a function as an
argument.
40. Implement a function that calculates the GCD of two numbers.
*/

func question31(a, b int) int {
	return a + b
}

func question32(r float64) float64 {
	return math.Pi * r * r
}

func question33() (int, int, float64) {
	return 3, 4, 89.33
}

func question34(a int) int {
	if a == 0 || a == 1 {
		return a
	}

	return question34(a-1) + question34(a-2)
}

func question35(a string, b ...int) (int, string) {
	var sum int = 0

	for _, i := range b {
		sum += i
	}

	return sum, "baby"
}

func question36(a ...int) (int, string) {
	var product int = 1
	for _, i := range a {
		product *= i
	}
	return product, "I am ShriShubhang"
}

func question37() func() int {

	x := 0
	return func() int {
		x++
		return x
	}
}

func question38() {

	a := 0

	func() {
		a++
	}()

	fmt.Println(a)

}

func question39(a func(x ...int) int) int {

	b := []int{1, 2, 3}

	a(b...)

	return b[0]
}

func question40(a, b int) int {
	if b == 0 {
		return a
	}

	return question40(b, a%b)
}

func main() {
	fmt.Printf("Hello I am ShriShubhang\n")

	a := func(x ...int) int {
		x[0] = -1

		return x[0]
	}

	tt := question39(a)

	fmt.Println(tt)

}
