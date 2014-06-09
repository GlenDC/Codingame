package main

import (
	"fmt"
	"github.com/glendc/cgreader"
)

func main() {
	cgreader.RunAndValidateManualPrograms(
		cgreader.GetFileList("../../input/input_%d.txt", 2),
		cgreader.GetFileList("../../output/output_%d.txt", 2),
		true,
		func(input <-chan string, output chan string) {
			// program logic
		})
}
