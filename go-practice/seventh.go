package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

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

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Good Going It is the first", id)

	time.Sleep(time.Second)

	fmt.Println("Good Back and Back issue is not gone", id)

}

func question62() {

	var wg sync.WaitGroup

	noOfWorkers := 3

	for i := 1; i <= noOfWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("Good Going and Not good Going")

}

func producer(ch chan<- int) {

	for i := 0; i < 5; i++ {
		fmt.Println("Producing :", i)
		ch <- i
		time.Sleep(time.Second)
	}
	close(ch)
}

func consumer(ch <-chan int) {

	for item := range ch {
		fmt.Println("Consuming :", item)
		time.Sleep(2 * time.Second)
	}
}

func question63() {

	ch := make(chan int)

	go producer(ch)
	go consumer(ch)

	time.Sleep(10 * time.Second)
}

func producer1(ch chan<- int) {

	for i := 0; i < 5; i++ {
		fmt.Println("Producing ; ", i)
		ch <- i
		time.Sleep(time.Second)
	}
	close(ch)
}

func consumer1(ch <-chan int) {

	for item := range ch {
		fmt.Println("Consuming : ", item)
		time.Sleep(2 * time.Second)
	}
}

func question64() {

	ch := make(chan int, 5)

	producer1(ch)
	consumer1(ch)

	time.Sleep(15 * time.Second)
}

func question65() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Hello , This message is Message1 from channel1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Hello, This message is Message2 from channel2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:

			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}

	}
}

func sum(arr []int, result chan int) {

	sum := 0

	for _, value := range arr {
		sum += value
	}
	result <- sum
}

func question67() {

	result := make(chan int)

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	mid := len(arr) / 2

	go sum(arr[:mid], result)
	go sum(arr[mid:], result)

	sum1 := <-result
	sum2 := <-result

	fmt.Println(sum1 + sum2)

}

var counter int
var mutex sync.Mutex

func helperQuestion66(wg *sync.WaitGroup) {

	defer wg.Done()
	mutex.Lock()
	counter++
	mutex.Unlock()
}

func question66() {
	var wg sync.WaitGroup

	numberOfRoutines := 10

	wg.Add(numberOfRoutines)

	for i := 0; i < numberOfRoutines; i++ {
		go helperQuestion66(&wg)
	}

	wg.Wait()

	fmt.Println(counter)
}

var once sync.Once

func increment() {
	fmt.Println("Increment with love")
}

func question68() {

	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(increment)
		}()
	}
	wg.Wait()
}

func workers(id int, jobs <-chan int, workers chan<- int, wg *sync.WaitGroup) {

	defer wg.Done()

	for job := range jobs {
		fmt.Println("It is going to be really awesome ", id, "It is not going good ", job)
		time.Sleep(time.Millisecond * time.Duration(rand.IntN(500)))
		workers <- job * job
	}
}

func question69() {

	numberOfWorkers := 3
	numberOfJobs := 10

	jobs := make(chan int)
	results := make(chan int)

	var wg sync.WaitGroup

	for i := 1; i <= numberOfWorkers; i++ {
		wg.Add(1)
		go workers(i, jobs, results, &wg)
	}

	for i := 1; i <= numberOfJobs; i++ {
		jobs <- i
	}

	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	for value := range results {
		fmt.Println("Good Going but it has value ", value)
	}
}

func question70() {

}

func main() {
	question69()
}
