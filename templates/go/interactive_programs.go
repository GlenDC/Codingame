package main

import (
	"github.com/glendc/cgreader"
)

func Initialize(input <-chan string) {
	// initial input...
}

func Update(input <-chan string, output chan string) {
	// process input into output
}

func main() {
	cgreader.RunInteractivePrograms(
		cgreader.GetFileList("../../input/input_%d.txt", 2),
		"ragnarok",
		true,
		Initialize,
		Update)
}
