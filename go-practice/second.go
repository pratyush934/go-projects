package main

import "fmt"

/*
11. Create a program to declare and initialise an array of integers.
12. Implement a program that finds the sum of all elements in an array.
13. Write a program to reverse an array.
14. Create a slice of integers and demonstrate adding elements to the slice.
15. Write a program to remove an element from a slice by index.
16. Implement a program that copies one slice into another.
17. Write a Go program to count the frequency of elements in a slice.
18. Implement a map that stores the names of students as keys and their grades as values.
19. Write a program that updates the value associated with a key in a map.
20. Implement a program to check if a key exists in a map.
*/

func question11() {
	var arr [5]int
	arr[1] = 10
	fmt.Println(arr)
}

func question12() {
	arr := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, ele := range arr {
		sum += ele
	}
	fmt.Println(sum)
}

func question13() {
	arr := []int{1, 2, 3, 4, 5}

	fmt.Println(arr)

	left := 0
	right := len(arr) - 1

	for left <= right {
		arr[left], arr[right] = arr[right], arr[left]
		left += 1
		right -= 1
	}

	fmt.Println(arr)

}

func question14() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	var a int
	fmt.Scan(&a)
	arr = append(arr, a)
	fmt.Println(arr)

}

func question15() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	length := len(arr)
	index := 0
	fmt.Scan(&index)
	if index >= length {
		fmt.Println("Nikal Pagle pehle fursat me Nikal")
		return
	}
	arr = append(arr[:index], arr[index+1:]...)
	fmt.Println(arr)
}

func question16() {
	arr := []int{1, 2, 3}
	arr2 := []int{4, 5, 6}
	fmt.Println(arr)

	fmt.Println(arr2)

	copy(arr2, arr)

	fmt.Println(arr)

	fmt.Println(arr2)
}

func question17() {
	arr := []int{12, 12, 23, 61, 53}

	var frequency = make([]int, 100)

	for _, ele := range arr {
		frequency[ele] += 1
	}

	for index, ele := range frequency {
		if ele != 0 {
			fmt.Println("So the element ", index, " is available upto ", ele)
		}
	}
	fmt.Println("Ho gaya ji")
}

func question18() {
	mapmap := make(map[string]int)

	mapmap["Pratyush"] = 10
	mapmap["Smriti"] = 12
	mapmap["Mridula"] = 13

	fmt.Println(mapmap)
}

func main() {
	fmt.Printf("Hello , I am Sword of Morning\n")
	question18()
}
