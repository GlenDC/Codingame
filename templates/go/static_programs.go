package main

import (
	"github.com/glendc/cgreader"
)

func main() {
	cgreader.RunStaticPrograms(
		"<input_files>",
		"<output_files>",
		true,
		func(input <-chan string, output chan string) {
			// program logic
		})
}
