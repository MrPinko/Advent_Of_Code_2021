package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	f, _ := ioutil.ReadFile("input.txt")
	splitline := strings.Split(strings.ReplaceAll(string(f), "\r\n", "\n"), "\n")

	rule := make(map[string]string)

	template := splitline[0]
	fmt.Println(template)

	for _, e := range splitline[2:] {
		temp := strings.Split(e, " -> ")
		rule[temp[0]] = temp[1]

	}

	var polymer string

	for cycle := 0; cycle < 40; cycle++ {
		polymer = ""
		for i := 0; i < len(template)-1; i++ {
			if i == 0 {
				polymer += string(template[i])
			}
			polymer += rule[string(template[i])+string(template[i+1])]
			polymer += string(template[i+1])
		}
		template = polymer
	}

	result := stringOccurence(polymer)

	//fmt.Println(polymer)
	fmt.Println(result)

	max_num := Max_number(result)
	min_num := Min_number(result)

	fmt.Println(max_num, min_num)

	fmt.Println("part one answer: ", max_num-min_num)
}

func Min_number(result map[string]int) int {
	min_num := 1000000
	for _, e := range result {
		if e < min_num {
			min_num = e
		}
	}
	return min_num
}

func Max_number(result map[string]int) int {
	max_num := 0
	for _, e := range result {
		if e > max_num {
			max_num = e
		}
	}
	return max_num
}

func stringOccurence(polymer string) map[string]int {
	result := make(map[string]int)
	for i := 0; i < len(polymer); i++ {
		result[string(polymer[i])] += 1
	}
	return result
}
