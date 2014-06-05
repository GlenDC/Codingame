package main

import (
	"fmt"
	"github.com/glendc/cgreader"
	"math"
	"sort"
)

type intArray []int

func (s intArray) Len() int {
	return len(s)
}

func (s intArray) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s intArray) Less(i, j int) bool {
	return s[i] < s[j]
}

func main() {
	cgreader.SetTimeout(2.0)
	cgreader.RunAndValidateManualProgram(
		"../../input/horse_dual_3.txt",
		"../../output/horse_dual_3.txt",
		true,
		func(input <-chan string, output chan string) {
			var n int
			fmt.Sscanf(<-input, "%d", &n)

			horses := make([]int, n)
			for i := range horses {
				fmt.Sscanf(<-input, "%d", &horses[i])
			}

			sort.Sort(intArray(horses))

			D := math.MaxInt32
			for i := 1; i < n; i++ {
				x := horses[i-1] - horses[i]

				if x < 0 {
					x *= -1
				}

				if x < D {
					D = x
				}
			}

			output <- fmt.Sprintf("%d", D)
		})
}
