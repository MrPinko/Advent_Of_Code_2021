package main

import (
	"fmt"
	"io"
	"os"
)

func ReadFile() [][]int {
	f, _ := os.Open("input.txt")
	fileC := [][]int{}
	for {
		var temp int
		var temp2 int
		_, err := fmt.Fscanf(f, "%d,%d\n", &temp, &temp2)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		fileC = append(fileC, []int{temp, temp2})
	}
	return fileC
}

var (
	foldX_array = []int{655, 327, 163, 81, 40}
	foldY_array = []int{447, 223, 111, 55, 27, 13, 6}
	empty_char  = "░"
	full_char   = "█"
)

// var foldX_array = []int{5}
// var foldY_array = []int{7}

func main() {
	input := ReadFile()
	max := []int{}
	for j := range input[0] {
		temp := []int{}
		for i := range input {
			temp = append(temp, input[i][j])
		}
		max = append(max, Max(temp))
	}

	grid := make([][]string, max[1]+1) // y

	for i := range grid {
		grid[i] = make([]string, max[0]+1) //x
	}

	for i := range grid {
		for j := range grid[0] {
			grid[i][j] = empty_char
		}
	}

	for _, e := range input {
		grid[e[1]][e[0]] = full_char //grid [y][x]
	}

	i := 0
	j := 0
	currentLenX := len(grid[0])
	currentLenY := len(grid)

	foldX_arrayCopy := make([]int, len(foldX_array))
	foldY_arrayCopy := make([]int, len(foldY_array))
	copy(foldX_arrayCopy, foldX_array)
	copy(foldY_arrayCopy, foldY_array)

	for {
		if len(foldX_arrayCopy) == 0 && len(foldY_arrayCopy) == 0 {
			break
		}

		if len(foldX_arrayCopy) != 0 {
			FoldX(grid, foldX_array[i], currentLenY, currentLenX)
			currentLenX = currentLenX / 2
			foldX_arrayCopy = foldX_arrayCopy[1:]
			i++
		}
		//! for PART ONE add break
		//break
		if len(foldY_arrayCopy) != 0 {
			FoldY(grid, foldY_array[j], currentLenY, currentLenX)
			currentLenY = currentLenY / 2
			foldY_arrayCopy = foldY_arrayCopy[1:]
			j++
		}
	}

	PrintGrid(currentLenY, currentLenX, grid)

	CountDots(grid, currentLenY, currentLenX) //842

}

func PrintGrid(currentLenY int, currentLenX int, grid [][]string) {
	fmt.Println()
	for i := 0; i < currentLenY; i++ {
		for j := 0; j < currentLenX; j++ {
			fmt.Print(grid[i][j], " ")
		}
		fmt.Println()
	}
}

func CountDots(grid [][]string, indexY int, indexX int) {
	dots := 0
	for i := 0; i < indexY; i++ {
		for j := 0; j < indexX; j++ {
			if grid[i][j] == "#" {
				dots++
			}
		}
	}

	fmt.Println(dots)
}

func FoldY(grid [][]string, index int, indexY int, indexX int) {
	for i := 0; i < indexY; i++ {
		for j := 0; j < indexX; j++ {
			if i > index {
				if grid[i][j] != empty_char {
					grid[index-(i-index)][j] = grid[i][j]
				}
			}
		}
	}
}

func FoldX(grid [][]string, index int, indexY int, indexX int) {
	for i := 0; i < indexY; i++ {
		for j := 0; j < indexX; j++ {
			if j > index {
				if grid[i][j] != empty_char {
					grid[i][index-(j-index)] = grid[i][j]
				}
			}
		}
	}
}

func Max(num []int) int {
	maxnum := 0
	for _, n := range num {
		if n > maxnum {
			maxnum = n
		}
	}
	return maxnum
}
