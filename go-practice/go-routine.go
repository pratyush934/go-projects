package main

import (
	"fmt"
	"io"
	"log"
	"net"
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

func networkConnection() {

	listen, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		log.Fatal("There is error and we need to fix this", err)
	}

	for {
		conn, err := listen.Accept()

		if err != nil {
			log.Fatal("There was an error and program seems to be aborted", err)
		}

		handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	for {
		_, err := io.WriteString(conn, time.Now().Format("02:01:2006"))

		if err != nil {
			log.Fatal("Not so good but , Not so good", err)
			return
		}

		time.Sleep(1 * time.Second)
	}
}

func helloWorld() {
	fmt.Printf("Hello World from Functionk\n")
}

func main() {
	go helloWorld()
	time.Sleep(1000 * time.Millisecond)
	fmt.Printf("Hello World babu")
}
