package main

import (
	"fmt"
	"github.com/glendc/cgreader"
)

type SALSquare interface {
	Move([]int, int, int, int)
	GetState() uint8
}

var LAOS int
var Squares []SALSquare
var MoveSquaresPassed []int

func ClipMovement(position, n int) (int, bool) {
	if m := n - 1; position > m {
		return m, false
	} else if position < 0 {
		return 0, false
	}
	return position, true
}

func PositionCheck(mqp []int, position int) ([]int, bool) {
	for i := range mqp {
		if mqp[i] == position {
			return mqp, false
		}
	}
	return append(mqp, position), true
}

func SquareCheck(mqp []int, position, n, turns int) {
	if turns < LAOS {
		switch Squares[position].GetState() {
		case 2:
			LAOS = turns
		default:
			if mqp, ok := PositionCheck(mqp, position); ok {
				Squares[position].Move(mqp, position, n, turns)
			}
		}
	}
}

type RSquare struct{}

func (s RSquare) Move(mpq []int, position, n, turns int) {
	for i := 1; i <= 6; i++ {
		position, ok := ClipMovement(position+i, n)
		if ok {
			SquareCheck(mpq, position, n, turns+1)
		}
	}
}

func (s RSquare) GetState() uint8 {
	return 1
}

type MSquare struct {
	n int
}

func (s MSquare) Move(mpq []int, position, n, turns int) {
	position, ok := ClipMovement(position+s.n, n)
	if ok {
		SquareCheck(mpq, position, n, turns+1)
	}
}

func (s MSquare) GetState() uint8 {
	return 0
}

type ESquare struct{}

func (s ESquare) Move(mpq []int, position, n, turns int) {}

func (s ESquare) GetState() uint8 {
	return 2
}

func main() {
	cgreader.RunAndValidateManualPrograms(
		cgreader.GetFileList("../../input/snakes_and_ladders_%d.txt", 5),
		cgreader.GetFileList("../../output/snakes_and_ladders_%d.txt", 5),
		true,
		func(input <-chan string, output chan string) {
			var n int
			fmt.Sscanf(<-input, "%d", &n)

			Squares = make([]SALSquare, n)
			var start int
			for i := range Squares {
				var t string
				fmt.Sscanf(<-input, "%s", &t)
				switch t {
				case "S", "R":
					Squares[i] = RSquare{}
					if t == "S" {
						start = i
					}
				case "E":
					Squares[i] = ESquare{}
				default:
					var x int
					fmt.Sscanf(t, "%d", &x)
					Squares[i] = MSquare{x}
				}
			}

			max := n
			LAOS = max
			Squares[start].Move([]int{}, start, n, 0)

			if LAOS == max {
				output <- fmt.Sprintf("impossible")
			} else {
				output <- fmt.Sprintf("%d", LAOS)
			}
		})
}
