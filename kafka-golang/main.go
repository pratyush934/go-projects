package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("Hello I am Pratyush")

	s := "113"

	val, _ := strconv.Atoi(s)

	fmt.Println(val)
}
