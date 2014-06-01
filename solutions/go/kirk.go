package main

import "fmt"

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

func main() {
	hero := ship{0, 0, -1}

	for {
		fmt.Scanf("%d %d\n", &hero.x, &hero.y)
		for i := range mountains {
			fmt.Scanf("%d\n", &mountains[i])
		}

		fmt.Println(hero.FireOrHold())
	}
}
