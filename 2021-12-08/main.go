package main

import (
	"fmt"
	"os"
	"strings"
)

func ReadFile() ([]string, []string) {
	f, _ := os.ReadFile("input.txt")
	first_part_arr := []string{}
	second_part_arr := []string{}
	args := strings.Split(string(f), "\n")

	for _, e := range args {
		first_part_arr = append(first_part_arr, e[0:strings.Index(e, "|")])
		second_part_arr = append(second_part_arr, e[strings.Index(e, "|")+2:]) //+2 ignore white space after |
	}

	return first_part_arr, second_part_arr
}

type SSDisplay struct {
	top          bool
	top_left     bool
	top_right    bool
	center       bool
	bottom_left  bool
	bottom_right bool
	bottom       bool
}

type DecodeMap struct {
	top          rune
	top_left     rune
	top_right    rune
	center       rune
	bottom_left  rune
	bottom_right rune
	bottom       rune
}

func main() {

	//input_one, input_two := ReadFile()
	input_one, input_two := ReadFile()

	//fmt.Println("PART ONE")
	//Part_one(input_two)

	//fmt.Println("\nPART TWO")

	//index_char := make([]DecodeMap, 0)

	//order => top,top_left,top_right,center,bottom_left,bottom_right,bottom
	//index_char = append(index_char, DecodeMap{top: 'a', top_left: 'b', top_right: 'c',
	//	center: 'd', bottom_left: 'e', bottom_right: 'f', bottom: 'g'}) //normal map

	index_char := Decode(input_one)
	fmt.Printf("%c", index_char)

	dispaly := SSDisplay{}
	Part_two(input_two, index_char, dispaly) //complicated

}

func Part_one(input_two []string) {
	sum := 0
	for i := 0; i < len(input_two); i++ {
		temp := strings.Replace(input_two[i], "\r", " ", 1) //invalid character \r
		input_two_splits := strings.Split(temp, " ")
		for _, e := range input_two_splits {
			if len(e) == 2 || len(e) == 4 || len(e) == 3 || len(e) == 7 {
				//fmt.Println(e)
				sum += 1
			}
		}
	}
	fmt.Println("result is ", sum)
}

func Part_two(input_two []string, index_char []DecodeMap, dispaly SSDisplay) {
	for i := 0; i < len(input_two); i++ {
		temp := strings.Replace(input_two[i], "\r", " ", 1) //invalid character \r
		input_two_splits := strings.Split(temp, " ")

		for _, e := range input_two_splits {
			for _, char := range e {
				CalculateNumber(char, index_char, i, &dispaly)

				if dispaly.top && dispaly.bottom && dispaly.top_left && dispaly.top_right && dispaly.bottom_right {
					temp := dispaly.top_left
					dispaly.top_left = dispaly.center
					dispaly.center = temp

					CalculateNumber(char, index_char, i, &dispaly) //recalculate
				}
			}
			DispalyToNumber(dispaly)
			Print_screen(dispaly)
			dispaly = SSDisplay{}
		}
	}
}

func CalculateNumber(char rune, index_char []DecodeMap, i int, dispaly *SSDisplay) {
	switch char {
	case index_char[i].top:
		dispaly.top = true
	case index_char[i].top_left:
		dispaly.top_left = true
	case index_char[i].top_right:
		dispaly.top_right = true
	case index_char[i].center:
		dispaly.center = true
	case index_char[i].bottom_left:
		dispaly.bottom_left = true
	case index_char[i].bottom_right:
		dispaly.bottom_right = true
	case index_char[i].bottom:
		dispaly.bottom = true
	}

}

func DispalyToNumber(display SSDisplay) {
	var num int

	if display.top && display.top_left && display.top_right && display.bottom_left && display.bottom_right && display.bottom {
		num = 0
		fmt.Println("number 0")
	} else if display.top && display.top_left && display.center && display.bottom_left && display.bottom {
		num = 1
		fmt.Println("number 1")
	} else if display.top_right && display.bottom_right {
		num = 2
		fmt.Println("number 2")
	} else if display.top && display.top_right && display.center && display.bottom_right && display.bottom {
		num = 3
		fmt.Println("number 3")
	} else if display.top_left && display.top_right && display.center && display.bottom_right {
		num = 4
		fmt.Println("number 4")
	}

	fmt.Println(num)

}

func Decode(lines []string) []DecodeMap {
	decodeArray := make([]DecodeMap, 0)
	decodeMap := DecodeMap{}
	var str string

	for i := 0; i < len(lines); i++ {
		temp := strings.Replace(lines[i], "\r", " ", 1) //invalid character \r
		lines_splits := strings.Split(temp, " ")
		lines_splits = lines_splits[0 : len(lines_splits)-1]
		lines_splits = SortASC(lines_splits)

		for _, e := range lines_splits {
			if len(e) == 2 { //number 1
				decodeMap.top_right = rune(e[0])
				decodeMap.bottom_right = rune(e[1])
			}
			if len(e) == 3 { //number 7
				str = strings.ReplaceAll(e, string(decodeMap.top_right), "")
				str = strings.ReplaceAll(str, string(decodeMap.bottom_right), "")
				decodeMap.top = rune(str[0])
			}
			if len(e) == 4 { //number 4
				str = strings.ReplaceAll(e, string(decodeMap.top_right), "")
				str = strings.ReplaceAll(str, string(decodeMap.bottom_right), "")
				str = strings.ReplaceAll(str, string(decodeMap.top), "")
				decodeMap.top_left = rune(str[0])
				decodeMap.center = rune(str[1])
			}
			if len(e) == 5 {
				for {
					str = e
					str = strings.ReplaceAll(e, string(decodeMap.top_right), "")
					str = strings.ReplaceAll(str, string(decodeMap.bottom_right), "")
					str = strings.ReplaceAll(str, string(decodeMap.top), "")
					str = strings.ReplaceAll(str, string(decodeMap.top_left), "")
					str = strings.ReplaceAll(str, string(decodeMap.center), "")

					if len(str) == 1 {
						decodeMap.bottom = rune(str[0])
						break
					} else {
						break
					}
				}
			}
			if len(e) == 7 {
				str = strings.ReplaceAll(e, string(decodeMap.top_left), "")
				str = strings.ReplaceAll(str, string(decodeMap.bottom_right), "")
				str = strings.ReplaceAll(str, string(decodeMap.top), "")
				str = strings.ReplaceAll(str, string(decodeMap.top_right), "")
				str = strings.ReplaceAll(str, string(decodeMap.center), "")
				str = strings.ReplaceAll(str, string(decodeMap.bottom), "")

				decodeMap.bottom_left = rune(str[0])
			}
		}
	}

	decodeArray = append(decodeArray, decodeMap)

	return decodeArray
}

func SortASC(a []string) []string {
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if len(a[i]) > len(a[j]) {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	return a
}

func Print_screen(dispaly SSDisplay) {
	fmt.Print("\n\n")
	board_l := 6      //ok to modify
	blank_char := " " //ok to modify
	//dont touch
	board_h := board_l + 1
	l := board_l - 1
	h := board_h - 1
	for i := 0; i < board_h; i++ {
		for j := 0; j < board_l; j++ {
			//TOP
			if dispaly.top {
				if i == 0 && j > 0 && j < l {
					fmt.Print("1")
					continue
				}
			}
			if dispaly.top_left {
				if i >= 1 && i <= h/2-1 && j == 0 {
					fmt.Print("1")
					continue
				}
			}
			if dispaly.top_right {
				if i >= 1 && i <= h/2-1 && j == l {
					fmt.Print("1")
					continue
				}
			}
			//END TOP
			//CENTER
			if dispaly.center {
				if i == h/2 && j >= 1 && j <= l-1 {
					fmt.Print("1")
					continue
				}
			}
			//END CENTER
			//BOTTOM
			if dispaly.bottom_left {
				if i > h/2 && i < h && j == 0 {
					fmt.Print("1")
					continue
				}
			}
			if dispaly.bottom_right {
				if i > h/2 && i < h && j == l {
					fmt.Print("1")
					continue
				}
			}
			if dispaly.bottom {
				if i == h && j > 0 && j < l {
					fmt.Print("1")
					continue
				}
			}
			//END BOTTOM
			fmt.Print(blank_char)
		}
		fmt.Print("\n")
	}
}
