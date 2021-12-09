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

type LowPoint struct {
	i int
	j int
}

var lowPoint []LowPoint
var basins [][]int
var temp_basins = []int{}

func main() {
	heatmap := ReadFile()

	answer := CalculateLowHeatmap(heatmap)

	sum := 0
	for _, e := range answer {
		sum += e + 1
	}
	fmt.Println("anser part one", sum) //answer part one

	fmt.Print("PART TWO\n") ////need good recursive method

	for _, e := range lowPoint {
		temp_basins = []int{}
		basins = append(basins, CalculateBasis(CreateBorder(heatmap), e.i+1, e.j+1))
	}

	basins_len := []int{}
	for _, e := range basins {
		basins_len = append(basins_len, len(e))
	}

	CalculateHighBasins(basins_len) // answer
}

func CreateBorder(heatmap [][]int) [][]int {
	border_heatmap := make([][]int, len(heatmap)+2) //add top and bottom border
	for i := range border_heatmap {
		border_heatmap[i] = make([]int, len(heatmap[0])+2) //add left and right border
	}

	for i := 0; i < len(border_heatmap); i++ {
		for j := 0; j < len(border_heatmap[0]); j++ {
			if i == 0 || i == len(border_heatmap)-1 { //top or bottom
				border_heatmap[i][j] = 9
			} else if j == 0 || j == len(border_heatmap[0])-1 {
				border_heatmap[i][j] = 9
			} else {
				border_heatmap[i][j] = heatmap[i-1][j-1]
			}
		}
	}

	return border_heatmap
}

func CalculateBasis(heatmap [][]int, i int, j int) []int {
	if heatmap[i][j] == 9 {
		return nil
	}
	temp_basins = append(temp_basins, heatmap[i][j])

	//fmt.Println(heatmap[i][j])
	heatmap[i][j] = 9               //delete coords
	CalculateBasis(heatmap, i, j-1) //check right block
	CalculateBasis(heatmap, i, j+1) //check left block
	CalculateBasis(heatmap, i-1, j) //check top block
	CalculateBasis(heatmap, i+1, j) //check bottom block

	return temp_basins
}

func CalculateLowHeatmap(heatmap [][]int) []int { //bad but work
	answer := []int{}
	for i := 0; i < len(heatmap); i++ {
		for j := 0; j < len(heatmap[0]); j++ {
			if i == 0 {
				if j == 0 {
					if WatchDir(heatmap, i, j, "bottom") && WatchDir(heatmap, i, j, "right") {
						lowPoint = append(lowPoint, LowPoint{i: i, j: j})
						answer = append(answer, heatmap[i][j])
					}
				} else if j != len(heatmap[0])-1 {
					if WatchDir(heatmap, i, j, "bottom") && WatchDir(heatmap, i, j, "left") && WatchDir(heatmap, i, j, "right") {
						lowPoint = append(lowPoint, LowPoint{i: i, j: j})
						answer = append(answer, heatmap[i][j])
					}
				} else {
					if WatchDir(heatmap, i, j, "left") && WatchDir(heatmap, i, j, "bottom") {
						lowPoint = append(lowPoint, LowPoint{i: i, j: j})
						answer = append(answer, heatmap[i][j])
					}
				}
			} else if i != len(heatmap)-1 {
				if j == 0 {
					if WatchDir(heatmap, i, j, "top") && WatchDir(heatmap, i, j, "bottom") && WatchDir(heatmap, i, j, "right") {
						lowPoint = append(lowPoint, LowPoint{i: i, j: j})
						answer = append(answer, heatmap[i][j])
					}
				} else if j != len(heatmap[0])-1 {
					if WatchDir(heatmap, i, j, "left") && WatchDir(heatmap, i, j, "top") && WatchDir(heatmap, i, j, "right") && WatchDir(heatmap, i, j, "bottom") {
						lowPoint = append(lowPoint, LowPoint{i: i, j: j})
						answer = append(answer, heatmap[i][j])
					}
				} else {
					if WatchDir(heatmap, i, j, "top") && WatchDir(heatmap, i, j, "bottom") && WatchDir(heatmap, i, j, "left") {
						lowPoint = append(lowPoint, LowPoint{i: i, j: j})
						answer = append(answer, heatmap[i][j])
					}
				}
			} else {
				if j == 0 {
					if WatchDir(heatmap, i, j, "top") && WatchDir(heatmap, i, j, "right") {
						lowPoint = append(lowPoint, LowPoint{i: i, j: j})
						answer = append(answer, heatmap[i][j])
					}
				} else if j != len(heatmap[0])-1 {
					if WatchDir(heatmap, i, j, "top") && WatchDir(heatmap, i, j, "left") && WatchDir(heatmap, i, j, "right") {
						lowPoint = append(lowPoint, LowPoint{i: i, j: j})
						answer = append(answer, heatmap[i][j])
					}
				} else {
					if WatchDir(heatmap, i, j, "top") && WatchDir(heatmap, i, j, "left") {
						lowPoint = append(lowPoint, LowPoint{i: i, j: j})
						answer = append(answer, heatmap[i][j])
					}
				}
			}
		}
	}
	return answer
}

func CalculateHighBasins(num []int) {
	temp := SortDSC(num)[:3]
	sum := 1
	for _, e := range temp {
		sum *= e
	}

	fmt.Println("anser part two", sum) //answer
}

func SortDSC(a []int) []int {
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] < a[j] {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	return a
}

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
