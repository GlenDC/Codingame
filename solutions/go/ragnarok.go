package main

import (
	"fmt"
	"github.com/glendc/cgreader"
)

func GetDirection(x, y int) <-chan int {
	ch := make(chan int)
	go func() {
		difference := x - y
		switch {
		case difference < 0:
			ch <- -1
		case difference > 0:
			ch <- 1
		default:
			ch <- 0
		}
		close(ch)
	}()
	return ch
}

func GetDirectionLetter(a, b string, v int) string {
	switch v {
	case -1:
		return a
	case 1:
		return b
	default:
		return ""
	}
}

var TX, TY, PX, PY, E int

func Initialize(ch <-chan string) {
	fmt.Sscanf(<-ch, "%d %d %d %d", &PX, &PY, &TX, &TY)
}

func Update(ch <-chan string) string {
	fmt.Sscanf(<-ch, "%d", &E)

	chx := GetDirection(PX, TX)
	chy := GetDirection(PY, TY)

	dx, dy := <-chx, <-chy
	x := GetDirectionLetter("W", "E", dx)
	y := GetDirectionLetter("N", "S", dy)

	TX, TY = TX+dx, TY+dy

	return y + x
}

func main() {
	cgreader.RunRagnarok("../../input/ragnarok_1.txt", true, Initialize, Update)
}
