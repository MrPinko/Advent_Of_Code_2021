package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

//how much fish have that state
func parseData() []int {
	f, _ := os.ReadFile("input.txt")

	args := strings.Split(string(f), ",")
	pop := make([]int, 9) //save how much fish have the same state
	for i := 0; i < len(args); i++ {
		n, _ := strconv.Atoi(args[i])
		pop[n]++
	}
	return pop
}

func main() {
	fishState := parseData()
	days := 256

	for i := 1; i <= days; i++ {
		newFishes := fishState[0]             //fish with state 0 create a new fish
		fishState[0] = 0                      //reset
		for j := 1; j < len(fishState); j++ { //decrement state for every fish
			fishState[j-1] = fishState[j]
		}
		fishState[6] = fishState[6] + newFishes //add new fish to state 6
		fishState[8] = newFishes                //new fish start with state 8
	}

	sum := 0
	for _, c := range fishState {
		sum += c
	}

	fmt.Println("\nafter ", days, " days ")
	fmt.Println("there are a total of", sum, "Lanternfish")
}
