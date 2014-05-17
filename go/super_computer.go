package main

import (
	"flag"
	"fmt"
	"github.com/glendc/cgreader"
)

type program struct {
	N            int
	Calculations map[int]int
}

func (p *program) Update(input string) {
	if p.N == 0 {
		fmt.Sscanf(input, "%d\n", &p.N)
		p.Calculations = make(map[int]int, p.N)
	} else {
		var start, duration int
		fmt.Sscanf(input, "%d %d\n", &start, &duration)

		if _, ok := p.Calculations[start]; !ok {
			p.Calculations[start] = duration
		}

		//for k,v := range p.Calculations {
		//}
	}
}

func (p *program) GetOutput() string {
	return fmt.Sprintf("%d\n", len(p.Calculations))
}

func main() {
	flag.Parse()
	n := flag.NArg()
	if n > 0 {
		input := fmt.Sprintf("../input/super_computer_%s.txt", flag.Arg(0))
		ch, ok := cgreader.RunProgram(input)
		app := program{}
		for <-ok {
			app.Update(<-ch)
		}
		output := app.GetOutput()
		fmt.Printf("Output:\n%s\n", output)
		if n > 1 {
			test := fmt.Sprintf("../output/super_computer_%s.txt", flag.Arg(0))
			fmt.Print("Test: ")
			if cgreader.TestProgram(test, output) {
				fmt.Print("Program is correct")
			} else {
				fmt.Print("Program is incorrect")
			}
		}
	}
}
