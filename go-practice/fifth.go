package main

import "fmt"

/*
41. Define a struct called `Person` with fields `name` and `age`.
42. Write a method on the `Person` struct that prints the person’s details.
43. Implement a program that demonstrates creating an instance of a struct.
44. Write a method to calculate the area of a rectangle using a `struct`.
45. Implement a program to demonstrate pointer receivers in methods.
46. Write a program that compares two struct instances for equality.
47. Implement a program to create a slice of structs and iterate over it.
48. Write a program that demonstrates embedding one struct into another.
49. Implement a program that initialises a struct with values using a constructor function.
50. Write a program that demonstrates struct tags and using `json.Marshal()` to serialise a
struct.
*/

type Person struct {
	name *string
	age  int
}

func question41() {
	var name string = "Pratyush"
	person := Person{&name, 20}

	fmt.Println(person)

}

func main() {
	fmt.Printf("I am Pratyush, I am most wonderous in the world\n")
	question41()
}
