package main

import "fmt"

/*
1. Write a program to print "Hello, World!" in Go.
2. Write a program that declares and prints variables of different types (int, float, string, bool).
3. Implement a program that takes user input for a name and prints a greeting.
4. Write a program that swaps two integers without using a temporary variable.
5. Implement a Go function to add two integers and return the result.
6. Write a program that demonstrates the use of constants in Go.
7. Create a program that uses shorthand variable declaration (:=).
8. Write a program that calculates the factorial of a number using recursion.
9. Implement a program that converts Fahrenheit to Celsius.
10. Write a Go program to check if a number is even or odd.
*/

func question1() {
	fmt.Print("Hello World")
}

func question2() {
	var a int = 10
	var b float64 = 45.11
	var c string = "Abc"
	var d bool = 13*18%2 == 0
	fmt.Print(a, " ", b, " ", c, " ", d)
}

func question3() {
	var b string
	fmt.Scan(&b)
	fmt.Print("This is going to be awesome and you must know it", b)
}

func question4() {
	a := 10
	b := 11
	fmt.Print(a, " ", b)
	b, a = a, b
	fmt.Println()
	fmt.Print(a, " ", b)
}

func question5() {
	a := 11
	b := 12
	fmt.Print(a + b)
}

func question6() {
	const a = 10
	const b = 11
	fmt.Print("a : ", a, "b: ", b)
}

func question7() {
	a := 11
	fmt.Print(a)
}

func question8(a int) int {
	if a == 1 {
		return 1
	}

	return question8(a-1) * a
}

func question9(temp int) int {
	return temp * 10
}

func question10(a int) bool {
	return a%2 == 0
}
