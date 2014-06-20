package main

import (
	"fmt"
	"github.com/glendc/cgreader"
)

func Initialize(input <-chan string) {
	// initial input...
}

func Update(input <-chan string, output chan string) {
	// process input into output
}

func main() {
	cgreader.RunKirkProgram("kirk_1.txt", true, Initialize, Update)
}
