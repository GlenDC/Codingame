package main

import (
	"fmt"
	"github.com/glendc/cgreader"
)

const (
	WATER_MASK  = 1
	GROUP_SHIFT = 1
	WATER       = 79
)

type waterInfo struct {
	value, position uint32
}

func main() {
	cgreader.RunAndValidateManualPrograms(
		cgreader.GetFileList("../../input/surface_%d.txt", 2),
		cgreader.GetFileList("../../output/surface_%d.txt", 2),
		true,
		func(input <-chan string, output chan string) {
			var G, h, v, g, o, i, t uint32
			var L, H, N, X, Y int
			var LINE string
			var V rune
			var w waterInfo

			fmt.Sscanf(<-input, "%d", &L)
			fmt.Sscanf(<-input, "%d", &H)

			WORLD, GROUPS := make([]uint32, L*H), make(map[uint32][]waterInfo)

			for Y = 0; Y < H; Y++ {
				LINE = <-input
				for X, V = range LINE {
					h, v = uint32(Y*L), 0
					i = h + uint32(X)
					if V == WATER {
						if t = i - 1; X != 0 && WORLD[t]&WATER_MASK == 1 {
							v = WORLD[t]
						}

						if t = i - uint32(L); h != 0 && WORLD[t]&WATER_MASK == 1 {
							if v&WATER_MASK == 1 {
								g, o = v>>GROUP_SHIFT, WORLD[t]>>GROUP_SHIFT

								for _, w = range GROUPS[o] {
									WORLD[w.position] = v
								}

								GROUPS[g] = append(GROUPS[g], GROUPS[o]...)
								delete(GROUPS, o)
							} else {
								v = WORLD[t]
							}

							GROUPS[g] = append(GROUPS[g], waterInfo{v, i})
						} else if v&WATER_MASK == 1 {
							g = v >> GROUP_SHIFT
							GROUPS[g] = append(GROUPS[g], waterInfo{v, i})
						} else {
							g = G
							G++

							v = (g << GROUP_SHIFT) | 1

							GROUPS[g] = append(GROUPS[g], waterInfo{v, i})
						}

						WORLD[i] = v
					}
				}
			}

			fmt.Sscanf(<-input, "%d", &N)

			for ; N > 0; N-- {
				fmt.Sscanf(<-input, "%d %d", &X, &Y)
				t = uint32(Y*L + X)

				if WORLD[t]&WATER_MASK == 0 {
					output <- "0"
				} else {
					g = WORLD[t] >> GROUP_SHIFT
					output <- fmt.Sprintf("%d", len(GROUPS[g]))
				}
			}
		})
}
