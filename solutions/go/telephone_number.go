package main

import (
	"fmt"
	"github.com/glendc/cgreader"
)

type Node struct {
	nodes map[int]Node
}

func CreateNode() Node {
	return Node{make(map[int]Node)}
}

func main() {
	cgreader.RunStaticPrograms(
		cgreader.GetFileList("../../input/telephone_number_%d.txt", 5),
		cgreader.GetFileList("../../output/telephone_number_%d.txt", 5),
		true,
		func(input <-chan string, output chan string) {
			var n int
			fmt.Sscanf(<-input, "%d", &n)

			counter, database := 0, CreateNode()
			for i := 0; i < n; i++ {
				var tel string
				fmt.Sscanf(<-input, "%s", &tel)

				it := database
				for u := range tel {
					var number int
					fmt.Sscanf(string(tel[u]), "%d", &number)

					if _, ok := it.nodes[number]; !ok {
						it.nodes[number] = CreateNode()
						counter++
					}

					it = it.nodes[number]
				}
			}

			output <- fmt.Sprintf("%d", counter)
		})
}
