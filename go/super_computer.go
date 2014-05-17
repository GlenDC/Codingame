package main

import (
	"flag"
	"fmt"
	"github.com/glendc/cgreader"
)

type program struct{}

func (p *program) Update(input string) {
	fmt.Println(input)
}

func main() {
	flag.Parse()
	if flag.NArg() == 1 {
		input := fmt.Sprintf("../input/super_computer_%s.txt", flag.Arg(0))
		ch, ok := cgreader.RunProgram(input)
		app := program{}
		for <-ok {
			app.Update(<-ch)
		}
	}
}
