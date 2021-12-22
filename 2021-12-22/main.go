package main

import (
	"fmt"
	"io"
	"os"
)

type Reboot_steps struct {
	instr  string
	x1, x2 int
	y1, y2 int
	z1, z2 int
}

type Coords struct {
	x int
	y int
	z int
}

func ReadFile() []Reboot_steps {
	f, _ := os.Open("input.txt")
	fileC := []Reboot_steps{}
	for {
		reboot_steps := Reboot_steps{}
		_, err := fmt.Fscanf(f, "%s x=%d..%d,y=%d..%d,z=%d..%d\n", &reboot_steps.instr, &reboot_steps.x1, &reboot_steps.x2,
			&reboot_steps.y1, &reboot_steps.y2, &reboot_steps.z1, &reboot_steps.z2)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		fileC = append(fileC, reboot_steps)
	}
	return fileC
}

func main() {
	input := ReadFile()

	grid := PartOne(input)
	fmt.Println(len(grid))

}

func PartOne(input []Reboot_steps) map[Coords]string {
	grid := make(map[Coords]string)
	//mmmm yes not optimized
	for _, e := range input {
		max, min := max_min(e.x1, e.x2)
		if max < 50 && min > -50 { //x=-50..50
			for x := 0; x <= diff(e.x1, e.x2); x++ {
				max, min := max_min(e.y1, e.y2)
				if max < 50 && min > -50 { //y=-50..50,
					for y := 0; y <= diff(e.y1, e.y2); y++ {
						max, min := max_min(e.z1, e.z2)
						if max < 50 && min > -50 { //z=-50..50
							for z := 0; z <= diff(e.z1, e.z2); z++ {
								if e.instr == "off" {
									delete(grid, Coords{x: e.x1 + x, y: e.y1 + y, z: e.z1 + z})
								} else {
									grid[Coords{x: e.x1 + x, y: e.y1 + y, z: e.z1 + z}] = e.instr
								}
							}
						}
					}
				}
			}
		}
	}
	return grid
}

func diff(i1, i2 int) int {
	max, min := max_min(i1, i2)
	return max - min
}

func max_min(n1, n2 int) (int, int) {
	if n1 > n2 {
		return n1, n2
	} else {
		return n2, n1
	}
}
