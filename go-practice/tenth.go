package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"os"
	"reflect"
	"sort"
	"time"
)

/*
91. Write a program to read and write to a file.
92. Implement a program that demonstrates working with JSON in Go.
93. Write a program to parse command-line arguments using the `flag` package.
94. Implement a program that reads a large file line by line using Go’s bufio package.
95. Write a program to compress and decompress data using Go’s `gzip` package.
96. Implement a program that demonstrates reflection in Go.
97. Write a program that handles time and date operations using the `time` package.
98. Implement a program to generate random numbers using the `math/rand` package.
99. Write a program to sort an array of integers using Go’s `sort` package.
100. Implement a program that demonstrates unit testing in Go using the `testing` package.
*/

func question91() {

	create, err := os.Create("Haha.txt")

	defer create.Close()
	if err != nil {
		panic(err)
	}

	n, err := create.WriteString("Main Hun Gian")

	if err != nil {
		panic(err)
	}

	fmt.Println(n)

	file, err := os.ReadFile("Haha.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(file))
}

type People struct {
	Name    string
	Age     int
	Address string
}

func question92() {
	people := People{"Pratyush", 10, "Add"}

	marshal, err := json.Marshal(people)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(string(marshal))

	var newPeople People

	err = json.Unmarshal([]byte(marshal), &newPeople)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(newPeople)
}

func question93() {
	s := os.Args[1]
	fmt.Println(s)
}

func question94() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the text ")

	readString, err := reader.ReadString('\n')
	if err != nil {
		return
	}
	fmt.Println(readString)
	fmt.Println(readString)

}

type Details struct {
	name string
	age  int
}

type iD string

func showMe(details, d interface{}) {
	t1 := reflect.TypeOf(details)
	k1 := t1.Kind()
	t2 := reflect.TypeOf(d)
	k2 := t2.Kind()

	fmt.Println("T1 ", t1)
	fmt.Println("K1 ", k1)
	fmt.Println("T2 ", t2)
	fmt.Println("K2 ", k2)

}

func question96() {

	d := iD("12345k")

	details := Details{"Pratyush", 12}

	showMe(details, d)
}

func question97() {
	fmt.Println(time.Now())
}

func question98() {
	n := rand.IntN(10)
	fmt.Println(n)
}

func question99() {
	arr := []int{2, 4, 1, 9, 3}
	sort.Ints(arr)
	fmt.Println(arr)
}

func question100() {
	//look at the main file
}

func main() {
	question100()
}
