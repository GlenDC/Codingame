package main

import (
	"fmt"
	"github.com/glendc/cgreader"
	"strings"
)

func main() {
	cgreader.RunAndValidateProgramManual(
		"../input/ascii_1.txt",
		"../output/ascii_%d.txt",
		true,
		func(ch <-chan string) string {
			var width, height int
			var text string

			fmt.Sscanln(<-ch, &width)
			fmt.Sscanln(<-ch, &height)
			fmt.Sscanln(<-ch, &text)

			text = strings.ToUpper(text)

			ascii := make([]string, height)
			for i := 0; i < height; i++ {
				ascii[i] = <-ch
			}

			output := make([]string, height)
			for _, char := range text {
				character := int(char) - 65
				if character < 0 || character > 26 {
					character = 26
				}
				for i := range output {
					position := character * width
					output[i] += ascii[i][position : position+width]
				}
			}

			var program_output string

			for _, line := range output {
				program_output += line + "\n"
			}

			return program_output
		})
}
