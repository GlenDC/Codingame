package main

import (
	"fmt"
	"github.com/glendc/cgreader"
)

var max int = 0
var relations map[int][]int

func TravelRecursive(children []int, coverage int) {
	if coverage > max {
		max = coverage
	}

	for _, child := range children {
		TravelRecursive(relations[child], coverage+1)
	}
}

func main() {
	cgreader.RunAndValidateManualPrograms(
		cgreader.GetFileList("../../input/dwarfs_giants_%d.txt", 3),
		cgreader.GetFileList("../../output/dwarfs_giants_%d.txt", 3),
		true,
		func(input <-chan string, output chan string) {
			max, relations = 0, make(map[int][]int)

			var n int
			fmt.Sscanf(<-input, "%d", &n)

			for i := 0; i < n; i++ {
				var r, c int
				fmt.Sscanf(<-input, "%d %d", &r, &c)
				relations[r] = append(relations[r], c)
			}

			for _, children := range relations {
				TravelRecursive(children, 1)
			}

			output <- fmt.Sprintf("%d", max)
		})
}
