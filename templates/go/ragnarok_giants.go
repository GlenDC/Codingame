package main

import (
	"github.com/glendc/cgreader"
)

func Initialize(input <-chan string) {
	// Parse Input
}

func Update(input <-chan string, output chan string) {
	// Define solution Logic
}

func main() {
	cgreader.RunRagnarokGiantsProgram("ragnarok_giants_1.txt", true, Initialize, Update)
}
