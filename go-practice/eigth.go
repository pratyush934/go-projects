package main

import (
	"errors"
	"fmt"
)

/*
71. Write a program that demonstrates basic error handling in Go.
72. Implement a program that defines a custom error type.
73. Write a program to demonstrate the use of `panic()` and `recover()`.
74. Implement a program that uses the `errors` package to create new errors.
75. Write a program that opens a file and handles errors during file operations.
76. Implement a program that demonstrates defer statements for cleanup tasks.
77. Write a program to demonstrate multiple defer statements and their execution order.
78. Implement a program that imports an external package and uses its functions.
79. Write a program that demonstrates the use of Go modules to manage dependencies.
80. Implement a program that handles multiple errors using Go’s `multierror` package.
*/

func question71(a, b int) (int, error) {

	if b == 0 {
		return 0, errors.New("not so bad")
	}

	return a / b, nil

}

type MyError struct {
	code    int
	message string
}

func question72(a, b int) (int, *MyError) {

	if b == 0 {
		return 0, &MyError{100, "Kya Dekh raha , Halwa hai kya"}
	}

	return a / b, nil
}

func handlePanic() {

	a := recover()
	if a != nil {
		fmt.Println(a)
	}
}

func dividingMyChoice(a, b int) (int, error) {

	defer handlePanic()

	if b == 0 {
		panic("Haye Diaya Haye Daiya")
	}

	return a / b, nil
}

func question73() {
	choice, err := dividingMyChoice(4, 0)

	if err != nil {
		fmt.Println("Error hai ", err)
	}

	fmt.Println(choice)
}

func main() {
	question73()
}
