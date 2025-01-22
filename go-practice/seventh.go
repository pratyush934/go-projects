package main

import "time"

/*
61. Write a program that creates a goroutine to print "Hello, Goroutine!".
62. Implement a program that uses a `sync.WaitGroup` to wait for multiple goroutines to
finish.
63. Write a program to pass data between goroutines using channels.
64. Implement a program that demonstrates buffered channels.
65. Write a program that uses a `select` statement with channels.
66. Implement a program that demonstrates the use of a `sync.Mutex` to prevent race
conditions.
67. Write a program that calculates the sum of numbers concurrently using goroutines and
channels.
68. Implement a program that demonstrates the use of `sync.Once` to ensure a function is
called only once.
69. Write a program to demonstrate a worker pool using goroutines.
70. Implement a program that demonstrates the use of `context` for goroutine cancellation.
*/

func question61() {
	println("Hello I am Gunda")
}

func main() {
	println("Hello I am Pratyush , I am good boy")
	time.Sleep(2 * time.Second)
	go question61()
	println("Hello I am Don")
}
