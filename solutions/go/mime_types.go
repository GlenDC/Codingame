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
		func(ch <-chan string) string {
			var n, m int
			fmt.Sscanf(<-ch, "%d", &n)
			fmt.Sscanf(<-ch, "%d", &m)

			types := make(map[string]string)
			for i := 0 ; i < n ; i++ {
				var key, value string
				fmt.Sscanf(<-ch, "%s %s", &key, &value)
				types[strings.ToLower(key)] = value
			}

			var output string
			for i := 0 ; i < m ; i++ {
				var path string
				fmt.Sscanf(<-ch, "%s", &path)
				path = strings.ToLower(path)

				if strings.Contains(path, ".") {
					sp := strings.Split(path, ".")
					if value, ok := types[sp[len(sp)-1]]; ok {
						output += fmt.Sprintf("%s\n", value)
					} else {
						output += fmt.Sprintln("UNKNOWN")
					}
				} else {
					output += fmt.Sprintln("UNKNOWN")
				}
			}

			return output
		})
}
