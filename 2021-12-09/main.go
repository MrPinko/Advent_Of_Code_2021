package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadFile() [][]int {
	f, _ := os.ReadFile("input.txt")
	split := strings.Split(string(f), "\n")

	heatmap := make([][]int, len(split))
	for i := range heatmap {
		heatmap[i] = make([]int, len(split[0])-1)
	}

	for i := 0; i < len(split); i++ {
		for j := 0; j < len(split[0])-1; j++ {
			temp, _ := strconv.Atoi(string(f[i+j+(len(split[0]))*i]))
			heatmap[i][j] = temp
		}
	}

	return heatmap
}

func main() {
	heatmap := ReadFile()

	// first row
	// first row && first column
	//middle
	//last
	// first row && first column
	//middle
	//last
	// first row && first column
	//middle
	//last
	answer := CalculateLowHeatmap(heatmap)

	//fmt.Println(answer)
	sum := 0
	for _, e := range answer {
		sum += e + 1
	}
	fmt.Println(sum)

	fmt.Print("\n\nPART TWO\n\n") //need good recursive method

	//answer = CalculateLowHeatmap_two(heatmap)
	//answer = WatchBasinTop(heatmap, 0, 0)
	answer = WatchBasinTopRight(heatmap, 0, len(heatmap[0])-1)
	fmt.Println(answer)

}

var temp = []int{}

func WatchBasinTop(heatmap [][]int, i int, j int) []int {
	if heatmap[i][j+1] == 9 {
		if heatmap[i+1][0] != 9 {
			temp = append(temp, heatmap[i][j])
			return WatchBasinTop(heatmap, i+1, 0)
		}
	} else {
		fmt.Println(heatmap[i][j])
		temp = append(temp, heatmap[i][j])
		return WatchBasinTop(heatmap, i, j+1)
	}

	return append(temp, heatmap[i][j])
}

func WatchBasinTopRight(heatmap [][]int, i int, j int) []int {

	if heatmap[i][j-1] == 9 {
		temp = append(temp, heatmap[i][j])

		for k := j; heatmap[i][k+1] != 9; k++ {
			if heatmap[i+1][k] != 9 {
				fmt.Println("test ", heatmap[i][j])
				return WatchBasinTopRight(heatmap, i+1, k)
			}
		}

		fmt.Println(heatmap[i][j])
		temp = append(temp, heatmap[i][j])
		return WatchBasinTopRight(heatmap, i+1, len(heatmap[0])-1)
	} else {
		fmt.Println(heatmap[i][j])
		temp = append(temp, heatmap[i][j])
		return WatchBasinTopRight(heatmap, i, j-1)
	}

	return append(temp, heatmap[i][j])

	// if heatmap[i][j] == 9 {
	// 	if heatmap[i+1][0] != 9 {
	// 		return WatchBasinTopRight(heatmap, i+1, len(heatmap[0])-1)
	// } else {
	// 	fmt.Println(heatmap[i][j])
	// 	return WatchBasinTopRight(heatmap, i, j-1)
	// }
	// return nil
}

func CalculateLowHeatmap(heatmap [][]int) []int {
	answer := []int{}
	for i := 0; i < len(heatmap); i++ {
		for j := 0; j < len(heatmap[0]); j++ {
			if i == 0 {
				if j == 0 {
					if WatchDir(heatmap, i, j, "bottom") && WatchDir(heatmap, i, j, "right") {
						answer = append(answer, heatmap[i][j])
					}
				} else if j != len(heatmap[0])-1 {
					if WatchDir(heatmap, i, j, "bottom") && WatchDir(heatmap, i, j, "left") && WatchDir(heatmap, i, j, "right") {
						answer = append(answer, heatmap[i][j])
					}
				} else {
					if WatchDir(heatmap, i, j, "left") && WatchDir(heatmap, i, j, "bottom") {
						answer = append(answer, heatmap[i][j])
					}
				}
			} else if i != len(heatmap)-1 {
				if j == 0 {
					if WatchDir(heatmap, i, j, "top") && WatchDir(heatmap, i, j, "bottom") && WatchDir(heatmap, i, j, "right") {
						answer = append(answer, heatmap[i][j])
					}
				} else if j != len(heatmap[0])-1 {
					if WatchDir(heatmap, i, j, "left") && WatchDir(heatmap, i, j, "top") && WatchDir(heatmap, i, j, "right") && WatchDir(heatmap, i, j, "bottom") {
						answer = append(answer, heatmap[i][j])
					}
				} else {
					if WatchDir(heatmap, i, j, "top") && WatchDir(heatmap, i, j, "bottom") && WatchDir(heatmap, i, j, "left") {
						answer = append(answer, heatmap[i][j])
					}
				}
			} else {
				if j == 0 {
					if WatchDir(heatmap, i, j, "top") && WatchDir(heatmap, i, j, "right") {
						answer = append(answer, heatmap[i][j])
					}
				} else if j != len(heatmap[0])-1 {
					if WatchDir(heatmap, i, j, "top") && WatchDir(heatmap, i, j, "left") && WatchDir(heatmap, i, j, "right") {
						answer = append(answer, heatmap[i][j])
					}
				} else {
					if WatchDir(heatmap, i, j, "top") && WatchDir(heatmap, i, j, "left") {
						answer = append(answer, heatmap[i][j])
					}
				}
			}
		}
	}
	return answer
}

// func CalculateLowHeatmap_two(heatmap [][]int) []int {
// 	answer := []int{}
// 	for i := 0; i < len(heatmap); i++ {
// 		for j := 0; j < len(heatmap[0]); j++ {
// 			if i == 0 {
// 				if j == 0 {
// 					if WatchBasin(heatmap, i, j, "bottom") && WatchBasin(heatmap, i, j, "right") {
// 						answer = append(answer, heatmap[i][j])
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return answer
// }

func WatchDir(heatmap [][]int, i int, j int, dir string) bool {
	switch dir {
	case "top":
		return heatmap[i][j] < heatmap[i-1][j]
	case "left":
		return heatmap[i][j] < heatmap[i][j-1]
	case "right":
		return heatmap[i][j] < heatmap[i][j+1]
	case "bottom":
		return heatmap[i][j] < heatmap[i+1][j]
	}

	return false
}
