package main

import (
	"fmt"
	"github.com/glendc/cgreader"
)

var mountains [8]int

type ship struct {
	x, y, target int
}

func (s *ship) GetTarget() {
	y := 0
	s.target = -1
	for i, mountain := range mountains {
		if mountain >= y {
			y = mountain
			s.target = i
		}
	}
}

func (s *ship) FireOrHold() string {
	if s.target == -1 {
		s.GetTarget()
	}

	if s.x == s.target {
		s.target = -1
		return "FIRE"
	}

	return "HOLD"
}

var hero ship

func Initialize(input <-chan string) {
	hero = ship{0, 0, -1}
}

func Update(input <-chan string, output chan string) {
	fmt.Sscanf(<-input, "%d %d", &hero.x, &hero.y)
	for i := range mountains {
		fmt.Sscanf(<-input, "%d", &mountains[i])
	}

	output <- hero.FireOrHold()
}

func main() {
	//cgreader.SetFrameRate(10)
	//cgreader.RunKirkProgram("../../input/kirk_6.txt", true, Initialize, Update)
	cgreader.RunKirkPrograms(
		cgreader.GetFileList("../../input/kirk_%d.txt", 6),
		false,
		Initialize,
		Update)
}
