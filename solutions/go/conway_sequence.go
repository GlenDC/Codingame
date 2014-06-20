package main

import (
	"fmt"
	"github.com/glendc/cgreader"
)

func main() {
	cgreader.RunAndValidateManualPrograms(
		cgreader.GetFileList("../../input/conway_sequence_%d.txt", 6),
		cgreader.GetFileList("../../output/conway_sequence_%d.txt", 6),
		true,
		func(input <-chan string, output chan string) {
			series, n := make([][]int, 2), 1
			series[0], series[1] = make([]int, 2500), make([]int, 2500)
			var count int
			fmt.Sscanf(<-input, "%d", &series[0][0])
			fmt.Sscanf(<-input, "%d", &count)
			count--

			for i := 0; i < count; i++ {
				pid, cid := i%2, (i+1)%2
				k, c := series[pid][0], 1
				n = 0
				for u := 1; u < len(series[pid]); u++ {
					if v := series[pid][u]; v != k {
						series[cid][n], series[cid][n+1] = c, k
						n += 2
						k, c = v, 1
					} else {
						c++
					}
				}
			}

			id, str := count%2, ""
			for i := 0; i < n; i++ {
				str += fmt.Sprintf("%d ", series[id][i])
			}
			output <- string(str[:len(str)-1])
		})
}
