package main

import (
	"fmt"
)

const (
	COLDER        = 67
	WARMER        = 87
	SAME          = 83
	UNKNOWN       = 85
	DEFAULT_SPEED = 1
)

func main() {
	var W, H, X, Y, SX, SY, PX, PY, V, WC int32
	var N, CC uint32
	var INPUT string
	var HINT uint8

	fmt.Scanf("%d %d", &W, &H)
	fmt.Scanf("%d", &N)
	fmt.Scanf("%d %d", &X, &Y)

	SX, SY = DEFAULT_SPEED, DEFAULT_SPEED

	PX, PY = X, Y

	for {
		fmt.Scanf("%s", &INPUT)
		HINT = INPUT[0]

		switch HINT {
		case COLDER:
			CC++
			WC = 0
			switch CC {
			case 1, 3:
				SX = -SX
			case 2, 4:
				SY = -SY
			default:
				CC = 0
				if SX > 1 {
					SX = SX >> 1
				}

				if SY > 1 {
					SY = SY >> 1
				}
			}
		case WARMER:
			CC = 0
			if WC > 1 {
				SX, SY = SX<<1, SY<<1
			}
			WC++
		}

		X, Y = X+SX, Y+SY

		if HINT == WARMER {
			if X == W-1 || X == 0 {
				SX = 0
			}

			if Y == H-1 || H == 0 {
				SY = 0
			}
		} else if HINT == SAME {
			if SX != 0 {
				X = (X + PX) >> 1
			}

			if SY != 0 {
				Y = (Y + PY) >> 1
			}

			if SX < 0 {
				SX = -1
			} else if SX > 0 {
				SX = 1
			}

			if SY < 0 {
				SY = -1
			} else if SY > 0 {
				SY = 1
			}
		}

		if V = W - 1; X > V {
			X, SX = V, -SX
		} else if X < 0 {
			X, SX = 0, -SX
		}

		if V = H - 1; Y > V {
			Y, SY = V, -SY
		} else if Y < 0 {
			Y, SY = 0, -SY
		}

		PX, PY = X, Y

		fmt.Printf("%d %d\n", X, Y)
	}
}
