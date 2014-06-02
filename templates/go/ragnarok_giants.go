package main

import (
	"github.com/glendc/cgreader"
)

func Initialize(ch <-chan string) {
	// Parse Input
}

func Update(ch <-chan string) string {
	// Define solution Logic
	return "OUTPUT"
}

func main() {
	cgreader.RunRagnarokGiants("ragnarok_giants_1.txt", true, Initialize, Update)
}
