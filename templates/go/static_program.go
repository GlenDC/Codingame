package main

import (
	"github.com/glendc/cgreader"
)

func main() {
	cgreader.RunStaticProgram(
		"../../input/input.txt",
		"../../output/output.txt",
		true,
		func(input <-chan string, output chan string) {
			// program logic
		})
}
