package main

import (
	"fmt"
	"io"
	"os"
)

type Direction struct {
	dir    string
	length int
}

type Submarine struct {
	h_position int //horizontal position of the submarine
	depth      int //depth of the subamrine
	aim        int
}

func main() {

	submarine := Submarine{h_position: 0, depth: 0}

	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}

	var temp_dir string
	var temp_length int
	var instruction = []Direction{}
	for {
		_, err := fmt.Fscanf(file, "%s %d\n", &temp_dir, &temp_length) // give a patter to scan
		if err != nil {
			if err == io.EOF {
				break // stop reading the file
			}
			fmt.Println(err)
			os.Exit(1)
		}
		instruction = append(instruction, Direction{dir: temp_dir, length: temp_length})
	}

	//fmt.Println(instruction)

	//part one
	fmt.Println("--------------------------------")
	fmt.Println("            PART ONE            ")
	move(instruction, &submarine)
	fmt.Println("h position: ", submarine.h_position, "\ndepth: ", submarine.depth)
	fmt.Println("final answer: ", submarine.h_position*submarine.depth)

	//parto two
	submarine = Submarine{h_position: 0, depth: 0, aim: 0}
	fmt.Println("--------------------------------")
	fmt.Println("            PART TWO            ")
	move_better(instruction, &submarine)
	fmt.Println("h position: ", submarine.h_position, "\ndepth: ", submarine.depth, "\naim: ", submarine.aim)
	fmt.Println("final answer: ", submarine.h_position*submarine.depth)

}

func move(instruction []Direction, submarine *Submarine) {

	for _, inst := range instruction {
		switch inst.dir {
		case "forward":
			submarine.h_position += inst.length

		case "up":
			submarine.depth -= inst.length

		case "down":
			submarine.depth += inst.length

		}
	}
}

func move_better(instruction []Direction, submarine *Submarine) {

	for _, inst := range instruction {
		switch inst.dir {
		case "forward":
			submarine.h_position += inst.length
			submarine.depth += submarine.aim * inst.length

		case "up":
			submarine.aim -= inst.length

		case "down":
			submarine.aim += inst.length
		}
	}
}
