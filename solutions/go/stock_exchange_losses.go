package main

import (
	"fmt"
	"github.com/glendc/cgreader"
	"strings"
)

func main() {
	cgreader.RunStaticPrograms(
		cgreader.GetFileList("../../input/stock_exchange_losses_%d.txt", 5),
		cgreader.GetFileList("../../output/stock_exchange_losses_%d.txt", 5),
		true,
		func(input <-chan string, output chan string) {
			var ml, cm, v int64 = 0, 0, 0
			var n, i int
			fmt.Sscanf(<-input, "%d", &n)

			if n == 0 {
				output <- fmt.Sprint("0")
				return
			}

			stocks := strings.Split(<-input, " ")
			fmt.Sscanf(stocks[i], "%d", &v)
			cm = v

			for i = 1; i < n; i++ {
				fmt.Sscanf(stocks[i], "%d", &v)

				if d := v - cm; d < ml {
					ml = d
				}

				if v > cm {
					cm = v
				}
			}

			output <- fmt.Sprintf("%d", ml)
		})
}
