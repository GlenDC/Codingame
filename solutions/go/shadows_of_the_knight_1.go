package main

import (
	"fmt"
)

const (
	U = 85
	D = 68
	L = 76
	R = 82
)

func main() {
	var W, H, X, Y, PLX, PRX, PUY, PDY int
	var N uint32
	var direction string
	var character rune

	fmt.Scanf("%d %d", &W, &H)
	fmt.Scanf("%d", &N)
	fmt.Scanf("%d %d", &X, &Y)

	PLX, PRX = 0, W
	PUY, PDY = 0, H

	for {
		N--

		fmt.Scanf("%s", &direction)

		for _, character = range direction {
			switch character {
			case U:
				PDY = Y
			case D:
				PUY = Y
			case L:
				PRX = X
			case R:
				PLX = X
			}
		}

		X = (PLX + PRX) / 2
		Y = (PUY + PDY) / 2

		if X < 0 {
			X = 0
		} else if X >= W {
			X = W - 1
		}

		if Y < 0 {
			Y = 0
		} else if Y >= H {
			Y = H - 1
		}

		fmt.Println(X, Y)
	}
}
