package main

import (
	"github.com/glendc/cgreader"
)

func main() {
	cgreader.RunStaticProgram(
		"<input_file>",
		"<output_file>",
		true,
		func(input <-chan string, output chan string) {
			// program logic
		})
}
