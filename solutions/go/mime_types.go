package main

import (
	"fmt"
	"github.com/glendc/cgreader"
	"strings"
)

func main() {
	cgreader.RunAndValidateManualProgram(
		"../../input/mime_types_5.txt",
		"../../output/mime_types_5.txt",
		true,
		func(input <-chan string, output chan string) {
			var n, m int
			fmt.Sscanf(<-input, "%d", &n)
			fmt.Sscanf(<-input, "%d", &m)

			types := make(map[string]string)
			for i := 0; i < n; i++ {
				var key, value string
				fmt.Sscanf(<-input, "%s %s", &key, &value)
				types[strings.ToLower(key)] = value
			}

			for i := 0; i < m; i++ {
				var path string
				fmt.Sscanf(<-input, "%s", &path)
				path = strings.ToLower(path)

				if strings.Contains(path, ".") {
					sp := strings.Split(path, ".")
					if value, ok := types[sp[len(sp)-1]]; ok {
						output <- fmt.Sprintf("%s", value)
					} else {
						output <- fmt.Sprint("UNKNOWN")
					}
				} else {
					output <- fmt.Sprint("UNKNOWN")
				}
			}
		})
}
