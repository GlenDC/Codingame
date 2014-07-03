package main

import (
	"fmt"
	"github.com/glendc/cgreader"
)

type Jumper struct {
	income, position int
}

func main() {
	cgreader.RunStaticPrograms(
		cgreader.GetFileList("../../input/roller_coaster_%d.txt", 6),
		cgreader.GetFileList("../../output/roller_coaster_%d.txt", 6),
		true,
		func(input <-chan string, output chan string) {
			var L, C, N, p, i, pkg, people, group, total int
			fmt.Sscanf(<-input, "%d %d %d", &L, &C, &N)

			if C == 0 || N == 0 {
				output <- "0"
				return
			}

			groups, jumpers := make([]int, N), make([]Jumper, N)

			for i = range groups {

				fmt.Sscanf(<-input, "%d", &people)
				total += people
				groups[i] = people
			}

			if total <= L {
				output <- fmt.Sprintf("%d", total*C)
				return
			}

			for i = range groups {
				people, p = groups[i], i+1
				for {
					if p == N {
						p = 0
					}

					group = groups[p]

					if pkg = people + group; pkg > L {
						jumpers[i].income, jumpers[i].position = people, p
						break
					} else if pkg == L {
						jumpers[i].income, jumpers[i].position = pkg, (p+1)%N
						break
					} else {
						p++
						people = pkg
					}
				}
			}

			for total, p, i = 0, 0, 0; i < C; i++ {
				total += jumpers[p].income
				p = jumpers[p].position
			}

			output <- fmt.Sprintf("%d", total)
		})
}
