package main

import (
	"fmt"
	"github.com/glendc/cgreader"
	"math"
)

func main() {
	cgreader.SetTimeout(60); cgreader.RunAndValidateManualProgram(
		"../../input/horse_dual_3.txt",
		"../../output/horse_dual_3.txt",
		true,
		func(ch <-chan string) string {
			var n int
			fmt.Sscanf(<-ch, "%d", &n)

			horses := make([]int, n)

			D := math.MaxInt32
			for i := range horses {
				fmt.Sscanf(<-ch, "%d", &horses[i])
				for u := 0 ; u < i ; u++ {
					x := horses[u] - horses[i]

					if x < 0 {
						x *= -1
					}

					if x < D {
						D = x
					}
				}
			}

			return fmt.Sprintf("%d\n", D)
		})
}
