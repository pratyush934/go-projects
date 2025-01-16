package main

import (
	"fmt"
)

type State struct {
	jug1, jug2 int
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func isVisited(state State, visited map[State]bool) bool {

	_, found := visited[state]

	return found
}

func WaterJugProblem() {
	capacity1, capacity2 := 2, 1

	target := 10

	initialState := State{0, 0}
	visited := make(map[State]bool)
	visited[initialState] = true

	queue := []State{initialState}

	for len(queue) > 0 {

		curr := queue[0]
		queue = queue[1:]

		if curr.jug1 == target || curr.jug2 == target {
			fmt.Printf("Hey We got it and we need to get it , we got it bro and the answer is %d\n", target)
			return
		}

		nextStates := []State{
			{capacity1, curr.jug2}, //filling the jug1
			{curr.jug1, capacity2}, //filling the jug2
			{curr.jug1, 0},         //pouring the jug2
			{0, curr.jug2},         //pouring the jug1
			{min(curr.jug1+curr.jug2, capacity2), max(0, curr.jug1+curr.jug2-capacity1)}, //transfer of water from x to y
			{max(0, curr.jug2+curr.jug1-capacity2), min(curr.jug1+curr.jug2, capacity2)}, //transfer of water from y to x
		}

		for _, state := range nextStates {
			if !visited[state] {
				visited[state] = true
				queue = append(queue, state)
			}
		}
	}

	fmt.Println("We didn't find the solution")
}

func main() {
	fmt.Printf("hello , I am Pratyush\n")
	WaterJugProblem()
}
