package main

import (
	"fmt"
	"github.com/glendc/cgreader"
)

func main() {
	cgreader.RunAndValidateManualPrograms(
		cgreader.GetFileList("../../input/roller_coaster_%d.txt", 6),
		cgreader.GetFileList("../../output/roller_coaster_%d.txt", 6),
		true,
		func(input <-chan string, output chan string) {
			var L, C, N int
            fmt.Sscanf(<-input, "%d %d %d", &L, &C, &N)

            if C == 0 || N == 0 {
                output <- "0"
                return
            }

            groups := make([]int, N)
            for i := range groups {
                fmt.Sscanf(<-input, "%d", &groups[i])
            }

            income, identifier := 0, 0
            for ;C > 0; C-- {
                for l, i := 0, L; l < N && i > 0; l++ {
                    people := groups[identifier%N]
                    if i -= people; i >= 0 {
                        income += people
                        identifier++
                    }
                }
            }

            output <- fmt.Sprintf("%d", income)
		})
}
