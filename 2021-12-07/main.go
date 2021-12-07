package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

func ReadFile() []int {
	f, _ := os.Open("input.txt")
	fileC := []int{}
	for {
		var temp int
		_, err := fmt.Fscanf(f, "%d", &temp)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		fileC = append(fileC, temp)
	}
	return fileC
}

func main() {
	crab := ReadFile()

	fmt.Println("PART ONE RESULT:")
	fmt.Println(calc_fuel(crab, 1))
	fmt.Print("\n\n")

	fmt.Println("PART TWO RESULT:")
	fmt.Println(calc_fuel(crab, 2))
	fmt.Print("\n\n")

}

func calc_fuel(crab []int, part int) int {
	maxPosition := MaxPosition(crab)
	crab2 := make([]int, len(crab))
	sum := 0

	for j := 1; j < maxPosition; j++ {

		for i := 0; i < len(crab); i++ {
			crab2[i] = int(math.Abs(float64(crab[i] - j))) // test every position a crab might go
		}

		if sum != 0 {
			if sum > TotalSum(crab2, part) { //calculate if the current fuel cost is less than the previous
				sum = TotalSum(crab2, part)
			}
		} else {
			sum = TotalSum(crab2, part)
		}
	}
	return sum
}

func MaxPosition(crab []int) int {
	max := crab[0]

	for i := 1; i < len(crab); i++ {
		if crab[i] > max {
			max = crab[i]
		}
	}

	return max
}

func TotalSum(crab2 []int, part int) int {

	// n = (s  * (a+l)) /2
	if part == 1 {
		sum := 0
		for _, e := range crab2 {
			sum += e
		}
		return sum
	} else if part == 2 {
		sum := 0
		for _, e := range crab2 {
			s := e
			a := 1
			l := e
			sum += s * (a + l) / 2
		}
		return sum
	}

	panic(0) //not good

}
