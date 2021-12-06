package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	colorRed   = "\033[31m"
	colorWhite = "\033[37m"
)

type Board struct {
	index  int
	number string
	check  bool
}

func main() {
	f, _ := os.Open("input.txt")
	FContent, _ := io.ReadAll(f)

	str := strings.Split(string(FContent), "\n")
	str = append(str, " ")
	sub_str := str[0]
	drawn_num := strings.Split(sub_str, ",")

	fmt.Println("\ndrawn number", drawn_num)

	var board []Board
	var num string
	count := 0
	for _, e := range str[1:] { //insert only the number in board array
		for _, char := range e {
			if string(char) >= "0" && string(char) <= "9" { //chech if it's a number
				num += string(char)
			} else if len(num) > 0 { //add the number to the board array
				board = append(board, Board{index: count / 5, number: num, check: false})
				num = "" //reset the number to insert in the array
			}
		}
		if e != "\r" {
			count += 1 //index ++
		}
	}

	//draw number
	//we have a winner

	winner, last_drawn_number := Draw_numbers(drawn_num, board)
	printBoard(board)

	fmt.Println("\n\n", winner, "won")

	CalculatePoints(board, winner, last_drawn_number)

	//------------PART TWO-------------//

	fmt.Print("\n\n------------PART TWO-------------\n\n")

	for {
		winner, _ = Draw_numbers(drawn_num, board)
		printBoard(board)
		board = RemoveIndex(board, winner[0].index)

		for i, e := range board {
			if e.index > winner[0].index {
				board[i].index -= 1
			}
		}

		if len(board) == 25 {
			winner, last_drawn_number = Draw_numbers(drawn_num, board)
			printBoard(board)
			CalculatePoints(board, winner, last_drawn_number)
			return
		}
	}
}

func RemoveIndex(s []Board, index int) []Board {
	return append(s[:index*25], s[(index*25)+25:]...)
}

func Draw_numbers(drawn_num []string, board []Board) ([]Board, string) {
	for i := range drawn_num {
		for j := range board {
			if board[j].number == string(drawn_num[i]) {
				board[j].check = true
				check, winner := Check(board)
				if check {
					//fmt.Println("winner ma no", winner)
					return winner, drawn_num[i]
				}
			}
		}
	}
	return nil, "nil"
}

func CalculatePoints(board []Board, winner []Board, last_drawn_number string) {
	sum := 0

	for _, b := range board {
		if b.index == winner[0].index {
			if !b.check {
				number_int, _ := strconv.Atoi(b.number)
				sum += number_int
			}
		}
	}
	last_drawn_number_int, _ := strconv.Atoi(last_drawn_number)
	fmt.Println("sum : ", sum)
	fmt.Println("last drawn number : ", last_drawn_number_int)
	fmt.Println("final answer ", sum*last_drawn_number_int)
}

func Check(boards []Board) (bool, []Board) {

	countH := 0
	countV := 0
	var subH []Board
	var subV []Board

	for i := 0; i < len(boards); i++ {
		if i%5 == 0 {
			countH = 0
			subH = []Board{}
		}
		if boards[i].check {
			subH = append(subH, boards[i])
			countH++
		}
		if countH == 5 {
			break
		}
	}

	if countH == 5 {
		return true, subH
	}

	for i := 0; i < 5; i++ {
		for j := 0; j <= 20; j += 5 {
			if j == 0 {
				countV = 0
				subV = []Board{}
			}
			if boards[i+j].check {
				subV = append(subV, boards[i+j])
				countV++
			}
		}

		if countV == 5 {
			break
		}
	}

	if countV == 5 {
		return true, subV
	}

	return false, nil
}

func printBoard(board []Board) {

	for i := 0; i < len(board); i++ {
		if i%5 == 0 {
			fmt.Print("\n")
			if i%25 == 0 {
				fmt.Print("\n")
			}
		}
		num_int, _ := strconv.Atoi(board[i].number)
		if num_int < 10 {
			if board[i].check {
				fmt.Print(string(colorRed), " ", Strikethrough(board[i].number), ",")
				fmt.Print(string(colorWhite), "")
			} else {
				fmt.Print(string(colorWhite), " ", board[i].number, ",")
			}
		} else {
			if board[i].check {
				fmt.Print(string(colorRed), Strikethrough(board[i].number), ",")
				fmt.Print(string(colorWhite), "")
			} else {
				fmt.Print(string(colorWhite), board[i].number, ",")
			}
		}
	}
	fmt.Println("\n-----------------")

}

func Strikethrough(str string) string {
	newStr := ""
	for _, e := range str {
		newStr += string(e)
		newStr += "\u0336"
	}
	return newStr
}
