package main

import (
	"fmt"
)

type Player struct {
	pos    int
	points int
}

var dice = 0

//Player 1 starting position: 4
//Player 2 starting position: 6
func main() {

	//0 to 9 and i can use %10 after add 1
	p1 := Player{pos: 4 - 1}
	p2 := Player{pos: 6 - 1}
	pointsRequired := 1000

	Part_One(p1, p2, pointsRequired)

	fmt.Print("\n\n\n")
	win := Part_Two(multiverse{p1: p1, p2: p2})
	fmt.Printf("player 1 total win : %d\nplayer 2 total win : %d\n", win[0], win[1])
}

// keep track of state of the multiverse.
type multiverse struct {
	p1, p2 Player
}

func Part_One(p1 Player, p2 Player, pointsRequired int) {
	var winner int

	for {
		sum := 0
		for i := 0; i < 3; i++ {
			dice++
			sum += dice
		}

		p1.pos = (p1.pos + sum) % 10

		p1.points += p1.pos + 1
		if p1.points >= pointsRequired {
			winner = 1
			break
		}

		sum = 0
		for i := 0; i < 3; i++ {
			dice++
			sum += dice
		}

		p2.pos = (p2.pos + sum) % 10

		p2.points += p2.pos + 1
		if p2.points >= pointsRequired {
			winner = 2
			break
		}
	}

	switch winner {
	case 1:
		fmt.Printf("player 1 won with: %d points and in position %d", p1.points, p1.pos+1) // because 0 to 9
		fmt.Printf("\nand player 2 lost with: %d points and in position %d", p2.points, p2.pos+1)
		fmt.Printf("\nanswer : %d", p2.points*dice)
	case 2:
		fmt.Printf("player 2 won with: %d points and in position %d", p2.points, p2.pos+1)
		fmt.Printf("\nplayer 1 lost with: %d points and in position %d", p1.points, p1.pos+1)
		fmt.Printf("\nanswer : %d", p1.points*dice)
	}
}

var dp = make(map[multiverse][]int64)

func Part_Two(g multiverse) []int64 {
	if g.p1.points >= 21 {
		return []int64{1, 0}
	}
	if g.p2.points >= 21 {
		return []int64{0, 1} // memoization to keep track and count of who won a game
	}
	v, ok := dp[g] //if alredy exist skip
	if ok {
		return v
	}
	win := []int64{0, 0}
	for d1 := 1; d1 < 4; d1++ {
		for d2 := 1; d2 < 4; d2++ {
			for d3 := 1; d3 < 4; d3++ {
				p1 := (g.p1.pos + d1 + d2 + d3) % 10
				s1 := g.p1.points + p1 + 1

				// do the switch
				w := Part_Two(multiverse{
					p1: Player{pos: g.p2.pos, points: g.p2.points},
					p2: Player{pos: p1, points: s1},
				})
				win[0] += w[1]
				win[1] += w[0]
			}
		}
	}
	dp[g] = win
	return win
}
