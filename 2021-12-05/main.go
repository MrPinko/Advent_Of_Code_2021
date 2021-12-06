package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

type Vector struct {
	x1, y1 int
	x2, y2 int
}

func ReadFile() []Vector {
	f, _ := os.Open("input.txt")
	vectors := []Vector{}
	for {
		var temp_x1, temp_x2, temp_y1, temp_y2 int
		_, err := fmt.Fscanf(f, "%d,%d -> %d,%d\n", &temp_x1, &temp_y1, &temp_x2, &temp_y2)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		vectors = append(vectors, Vector{x1: temp_x1, y1: temp_y1, x2: temp_x2, y2: temp_y2})
	}
	return vectors
}

func main() {

	coords := ReadFile()

	max_x, max_y := max_all(coords)
	board := make([][]uint8, max_x+1)
	for i := range board {
		board[i] = make([]uint8, max_y+1)
	}

	Calculate_cords_part_one_better(coords)

	/*bad
	// coords_norm := Calculate_cords(coords)
	// coords = append(coords, coords_norm...)
	// Add_coordinate(board, coords)
	// fmt.Println(Check_board(board))
	*/

	//-------------PART TWO---------//

	fmt.Print("\n\n-------------PART TWO---------\n\n")

	coords = ReadFile()
	Calculate_cords_part_two_better(coords)

}

func Check_board(board [][]uint8) int {
	count := 0
	for _, row := range board {
		for _, val := range row {
			if val >= 2 {
				count++
			}
		}
	}
	return count
}

func Add_coordinate(board [][]uint8, coords []Vector) {
	for _, e := range coords {
		if e.x1 == e.x2 || e.y1 == e.y2 {
			board[e.y1][e.x1] += 1
			board[e.y2][e.x2] += 1
		} else if e.x1 == -1 && e.y1 == -1 {
			board[e.y2][e.x2] += 1
		} else if e.x2 == -1 && e.y2 == -1 {
			board[e.y1][e.x1] += 1
		}
	}
}

func Calculate_cords_part_one_better(coords []Vector) {
	seaFloor := make(map[Vector]int)
	for _, c := range coords {
		addX := 0
		addY := 0
		if c.x1 > c.x2 {
			addX = -1
		}
		if c.x1 < c.x2 {
			addX = 1
		}
		if c.y1 > c.y2 {
			addY = -1
		}
		if c.y1 < c.y2 {
			addY = 1
		}

		if addX != 0 && addY != 0 {
			//skip diagonal for now
			continue
		}

		startX := c.x1
		startY := c.y1
		targetX := c.x2
		targetY := c.y2

		for startX != targetX || startY != targetY {
			seaFloor[Vector{x1: startX, y1: startY}]++

			startX += addX
			startY += addY
		}
		seaFloor[Vector{x1: startX, y1: startY}]++
	}

	overlaps := 0
	for _, v := range seaFloor {
		if v > 1 {
			overlaps++
		}
	}
	fmt.Println("overlaps: ", overlaps)
}

func Calculate_cords(coords []Vector) []Vector {
	var dist int
	var coords2 []Vector

	for _, c := range coords {
		if c.x1 == c.x2 {
			if c.y1 > c.y2 {
				dist = c.y1 - c.y2
				for i := 1; i < dist; i++ {
					coords2 = append(coords2, Vector{x1: -1, y1: -1, x2: c.x2, y2: (c.y2 + i)})
				}
			} else {
				dist = c.y2 - c.y1
				for i := 1; i < dist; i++ {
					coords2 = append(coords2, Vector{x1: c.x1, y1: (c.y1 + i), x2: -1, y2: -1})
				}
			}
		} else if c.y1 == c.y2 {
			if c.x1 > c.x2 {
				dist = c.x1 - c.x2
				for i := 1; i < dist; i++ {
					coords2 = append(coords2, Vector{x1: -1, y1: -1, x2: (c.x2 + i), y2: c.y2})
				}
			} else {
				dist = c.x2 - c.x1
				for i := 1; i < dist; i++ {
					coords2 = append(coords2, Vector{x1: (c.x1 + i), y1: c.y1, x2: -1, y2: -1})
				}
			}
		}
	}

	return coords2
}

func Calculate_cords_part_two_better(coords []Vector) {
	seaFloor := make(map[Vector]int)
	for _, c := range coords {
		addX := 0
		addY := 0
		if c.x1 > c.x2 {
			addX = -1
		}
		if c.x1 < c.x2 {
			addX = 1
		}
		if c.y1 > c.y2 {
			addY = -1
		}
		if c.y1 < c.y2 {
			addY = 1
		}

		startX := c.x1
		startY := c.y1
		targetX := c.x2
		targetY := c.y2

		for startX != targetX || startY != targetY {
			seaFloor[Vector{x1: startX, y1: startY}]++

			startX += addX
			startY += addY
		}
		seaFloor[Vector{x1: startX, y1: startY}]++
	}

	overlaps := 0
	for _, v := range seaFloor {
		if v > 1 {
			overlaps++
		}
	}
	fmt.Println("overlaps: ", overlaps)
}

func IsDiagonal(c Vector) bool {
	//|x1−x2|=|y1−y2|
	// fmt.Print(math.Abs(float64(c.x1)-float64(c.x2)), " ")
	// fmt.Print(math.Abs(float64(c.y1)-float64(c.y2)), " ")
	// fmt.Println(math.Abs(float64(c.x1)-float64(c.x2)) == math.Abs(float64(c.y1)-float64(c.y2)))
	return math.Abs(float64(c.x1)-float64(c.x2)) == math.Abs(float64(c.y1)-float64(c.y2))
}

func max_all(coords []Vector) (int, int) {

	var max_x float64
	var max_y float64

	for _, c := range coords {
		if math.Max(float64(c.x1), float64(c.x2)) > max_x {
			max_x = math.Max(float64(c.x1), float64(c.x2))
		}

		if math.Max(float64(c.y1), float64(c.y2)) > max_y {
			max_y = math.Max(float64(c.y1), float64(c.y2))
		}
	}

	return int(max_x), int(max_y)
}

func print_board(board [][]uint8) {
	for _, row := range board {
		for _, val := range row {
			fmt.Print(val, " ")
		}
		fmt.Println("")
	}
}
