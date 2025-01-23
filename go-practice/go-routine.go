package main

import (
	"fmt"
	"time"
)

func greetings(delay time.Duration) {

	for {
		for _, i := range `-\|/` {
			fmt.Printf("The value exist her is %c\r", i)
			time.Sleep(delay)
		}
	}
}

func fibonacci(a int) int {
	if a < 2 {
		return a
	}

	return fibonacci(a-1) + fibonacci(a-2)
}

func experiment1() {

	go greetings(100 * time.Millisecond)
	const n = 45
	fiboNacci := fibonacci(45)

	fmt.Println("THe fibonacci here is ", fiboNacci)
}

func main() {
	experiment1()
}
