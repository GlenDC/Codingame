package main

import (
	"fmt"
	"github.com/glendc/cgreader"
)

func main() {
	cgreader.RunAndValidateManualProgram(
		"../../input/input.txt",
		"../../output/output.txt",
		true,
		func(input <-chan string, output chan string) {
			// program logic
		})
}
