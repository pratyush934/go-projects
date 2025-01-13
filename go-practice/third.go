package main

import "fmt"

/*
21. Write a program that demonstrates if-else statements by checking if a number is positive,
negative, or zero.
22. Implement a program to print the first 10 natural numbers using a `for` loop.
23. Write a Go program to print all even numbers between 1 and 100.
24. Implement a program that demonstrates the use of a `switch` statement to handle
multiple conditions.
25. Write a program that calculates the sum of all numbers up to a given input using a `for`
loop.
26. Implement a program that breaks out of a loop when a certain condition is met.
27. Write a program that uses a labelled `break` statement.
28. Implement a program that uses `continue` in a loop to skip even numbers.
29. Write a program that calculates the sum of digits of a number using a loop.
30. Implement a program that demonstrates the use of `goto` to jump to a label.
*/

func question21(a int) bool {
	return a%2 == 0
}

func question22() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("Hey Buddy this is me %d\n", i)
	}

}

func question23() {
	for i := 1; i <= 100; i++ {
		if i%13 == 0 {
			fmt.Printf("This one is even and you should know it , damn know it %d\n", i)
		}
	}
}

func question24() {
	a := 10

	switch a {
	case 1, 2, 3, 4:
		fmt.Printf("Hey gandu\n")
	case 10:
		fmt.Printf("Yahi hai to Sahi hai\n")

	}
}

func question25() {
	var n int
	fmt.Scanf("Enter the number %d ", &n)
	var sum int = 0
	for i := 1; i <= n; i++ {
		sum = sum + i
	}
	fmt.Println(sum)
}

func question26() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("hello this is %d\n", i)
		if i == 9 {
			break
		}
	}
}

func main() {
	fmt.Printf("Hello Buddy\n")
	question26()
}
