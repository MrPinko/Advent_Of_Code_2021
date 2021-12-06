package main

import (
	"fmt"
	"io"
	"os"
)

func part_one(number []int) int {
	answer := 0
	pre := 0

	for _, e := range number {
		if pre == 0 {
			pre = int(e)
		} else if e > pre {
			answer += 1
		}
		pre = int(e)
	}
	return answer
}

//How many sums are larger than the previous sum?
func part_two(numbers []int) int {
	var totalsum = []int{}
	answer := 0
	pre := 0

	for i := range numbers {
		totalsum = append(totalsum, sum(numbers[i:i+3]))
	}

	for _, e := range totalsum {
		if pre == 0 {
			pre = e
		} else if e > pre {
			answer += 1
		}
		pre = e
	}

	return answer
}

//sum of a slice of array
func sum(numbers []int) int {
	sum := 0
	for _, e := range numbers {
		sum += e
	}

	return sum
}

func main() {
	file, _ := os.Open("../input.txt")

	var perline int
	var number = []int{}
	//take number as int
	for {
		_, err := fmt.Fscanf(file, "%d\n", &perline) // give a patter to scan
		if err != nil {
			if err == io.EOF {
				break // stop reading the file
			}
			fmt.Println(err)
			os.Exit(1)
		}
		number = append(number, perline)
	}

	fmt.Println(part_one(number))

	fmt.Println(part_two(number))

}
