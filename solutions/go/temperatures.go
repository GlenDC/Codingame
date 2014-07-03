package main

import (
	"fmt"
	"github.com/glendc/cgreader"
	"math"
	"strings"
)

func main() {
	cgreader.RunStaticProgram(
		"../../input/temperatures_1.txt",
		"../../output/temperatures_1.txt",
		true,
		func(input <-chan string, output chan string) {
			var n int
			fmt.Sscanf(<-input, "%d", &n)

			if n == 0 {
				output <- "0"
				return
			}

			lines := strings.SplitAfter(<-input, " ")

			closest, closest_absolute := math.MaxInt32, math.MaxInt32
			for _, t := range lines {
				var temperature int
				fmt.Sscanf(t, "%d", &temperature)

				at := int(math.Abs(float64(temperature)))
				if closest_absolute > at || (closest_absolute == at && temperature > closest) {
					closest, closest_absolute = temperature, at
				}
			}

			output <- fmt.Sprintf("%d", closest)
		})
}
