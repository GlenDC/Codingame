package main

import (
	"fmt"
	"github.com/glendc/cgreader"
	"math"
)

func main() {
	cgreader.RunAndValidateManualProgram(
		"../../input/horse_dual_1.txt",
		"../../output/horse_dual_1.txt",
		true,
		func(ch <-chan string) string {
			var n int
			fmt.Sscanf(<-ch, "%d", &n)

			horses := make([]int, n)
			for i := range horses {
				fmt.Sscanf(<-ch, "%d", &horses[i])
			}

			D := math.MaxInt32
			for i, a := range horses {
				for u, b := range horses {
					if i != u {
						x := a - b

						if x < 0 {
							x *= -1
						}

						if x < D {
							D = x
						}
					}
				}
			}

			return fmt.Sprintf("%d\n", D)
		})
}
