package main

import (
	"flag"
	"fmt"
	"github.com/glendc/cgreader"
)

func main() {
	flag.Parse()
	if flag.NArg() == 1 {
		input := fmt.Sprintf("../input/super_computer_%s.txt", flag.Arg(0))
		ch, ok := cgreader.RunProgram(input)
		for <-ok {
			fmt.Println(<-ch)
		}
	}
}
