package main

import (
	"fmt"
)

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

func (p *Person) GetName() string { // question42
	return *p.name
}

func (p *Person) SetName(name *string) { //question42
	p.name = name
}

func question41() {
	var name string = "Pratyush"
	person := Person{&name, 20}

	fmt.Println(*person.name)
	fmt.Println(person.age)

}

func question43() {
	var name string
	name = "Pratyush"
	person := Person{&name, 10}
	fmt.Println(person)
}

type Rectangle struct {
	length, breadth int
}

func (r *Rectangle) GetDimension() (int, int) {
	return r.length, r.breadth
}

func (r *Rectangle) SetDimension(length, breadth int) {
	r.length = length
	r.breadth = breadth
}

func question44() {
	rectangle := Rectangle{10, 20}

	fmt.Println(rectangle.length * rectangle.breadth)
}

type GeometricFigures struct {
	dimension1, dimension2 int
}

func (g *GeometricFigures) SetDimensions(length, breadth int) {
	g.dimension1 = length
	g.dimension2 = breadth
}

func (g *GeometricFigures) GetDimensions() (int, int) {
	return g.dimension1, g.dimension2
}

func (g *GeometricFigures) Area() int {
	return g.dimension1 * g.dimension2
}

func question45() {
	geometricFigure := GeometricFigures{}
	geometricFigure.SetDimensions(10, 20)
	area := geometricFigure.Area()
	fmt.Println("The area exist for the following rectangle is", area)
}

type Struct1 struct {
	dimension1, dimension2 int
}

type Struct2 struct {
	dimension1, dimension2 int
}

func (s1 Struct1) Equals(s2 Struct2) bool {
	return s1.dimension1 == s2.dimension1 && s1.dimension2 == s2.dimension2
}

func question46() {
	struct1 := Struct1{10, 20}
	struct2 := Struct2{dimension1: 10, dimension2: 20}

	fmt.Println(struct1.Equals(struct2))
}

func main() {
	fmt.Printf("I am Pratyush, I am most wonderous in the world\n")
	question46()
}
