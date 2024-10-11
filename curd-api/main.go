package main

import "fmt"

const PORT string = ":8080"

func divide(num, den float64) float64 {
	if den == 0 {
		/*
			panic function is inbuilt function which is defined under the built-in package of golang
			this function terminates the flow of control and starts panicking
			To handle the error as specified the
			panic function also give information about the type of error or the location of the
		*/
		panic("Cannot be divided by 0")
		//fmt.Println("Can not")
	}
	return num / den
}
func main() {
	defer fmt.Println("No No")
	fmt.Println("So Sad")
	defer fmt.Println("Good Going")
	fmt.Println("Not Good Going")
}
