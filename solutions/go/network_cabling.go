package main

import (
	"fmt"
	"github.com/glendc/cgreader"
	"math"
)

func main() {
	cgreader.RunAndValidateManualPrograms(
		//"../../input/network_cabling_5.txt",
		//"../../output/network_cabling_5.txt",
		cgreader.GetFileList("../../input/network_cabling_%d.txt", 8),
		cgreader.GetFileList("../../output/network_cabling_%d.txt", 8),
		true,
		func(input <-chan string, output chan string) {
			var n int
			fmt.Sscanf(<-input, "%d", &n)

			if n <= 1 {
				output <- "0"
				return
			}

			yPositions := make(map[int]int)
			averageYPositions := make([]int, n)
			minX, maxX, aY, aN, m := math.MaxInt32, 0, 0, 0, 0
			for i := 0; i < n; i++ {
				var x, y int
				fmt.Sscanf(<-input, "%d %d", &x, &y)

				if x < minX {
					minX = x
				} else if x > maxX {
					maxX = x
				}

				yPositions[y]++
				if yPositions[y] > m {
					aY, m, aN = y, yPositions[y], 1
					averageYPositions[0] = y
				} else if yPositions[y] == m {
					aY += y
					averageYPositions[aN] = y
					aN++
				}
			}

			aY /= aN
			cY, distance := aY, math.MaxInt32
			for i := 0; i < aN; i++ {
				y := averageYPositions[i]
				d := y - aY

				if d < 0 {
					d *= -1
				}

				if d < distance {
					cY, distance = y, d
				}
			}

			distance = maxX - minX
			for k, v := range yPositions {
				d := cY - k
				if d < 0 {
					d *= -1
				}
				distance += d * v
			}

			output <- fmt.Sprintf("%d", distance)
		})
}
