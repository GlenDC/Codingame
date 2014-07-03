package main

import (
	"fmt"
	"github.com/glendc/cgreader"
	"strings"
)

func main() {
	cgreader.RunStaticProgram(
		"../../input/ascii_4.txt",
		"../../output/ascii_4.txt",
		true,
		func(input <-chan string, output chan string) {
			var width, height int
			var text string

			fmt.Sscanln(<-input, &width)
			fmt.Sscanln(<-input, &height)
			fmt.Sscanln(<-input, &text)

			text = strings.ToUpper(text)

			ascii := make([]string, height)
			for i := 0; i < height; i++ {
				ascii[i] = <-input
			}

			lines := make([]string, height)
			for _, char := range text {
				character := int(char) - 65
				if character < 0 || character > 26 {
					character = 26
				}
				for i := range lines {
					position := character * width
					lines[i] += ascii[i][position : position+width]
				}
			}

			for _, line := range lines {
				output <- fmt.Sprintln(line)
			}
		})
}
