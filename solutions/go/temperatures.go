package main

import (
	"fmt"
	"github.com/glendc/cgreader"
	"math"
	"strings"
)

func main() {
	cgreader.RunAndValidateManualProgram(
		"../../input/temperatures_1.txt",
		"../../output/temperatures_1.txt",
		true,
		func(ch <-chan string) string {
			var n int
			fmt.Sscanf(<-ch, "%d", &n)

			if n == 0 {
				return "0\n"
			}

			input := strings.SplitAfter(<-ch, " ")

			closest, closest_absolute := math.MaxInt32, math.MaxInt32
			for _, t := range input {
				var temperature int
				fmt.Sscanf(t, "%d", &temperature)

				at := int(math.Abs(float64(temperature)))
				if closest_absolute > at || (closest_absolute == at && temperature > closest) {
					closest, closest_absolute = temperature, at
				}
			}

			return fmt.Sprintf("%d\n", closest)
		})
}
