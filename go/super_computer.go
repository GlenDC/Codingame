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

func RunProgram(id int) (string, bool) {
	input := fmt.Sprintf("../input/super_computer_%d.txt", id)
	ch, ok := cgreader.RunProgram(input)
	app := program{}
	for <-ok {
		app.Update(<-ch)
	}
	output := app.GetOutput()
	test := fmt.Sprintf("../output/super_computer_%d.txt", id)
	valid := cgreader.TestProgram(test, output)
	return output, valid
}

func main() {
	flag.Parse()
	n := flag.NArg()
	if n > 0 {
		var count int
		fmt.Sscanf(flag.Arg(0), "%d", &count)
		results := make([]bool, count)
		for i := 0; i < count; i++ {
			output, ok := RunProgram(i + 1)
			results[i] = ok
			if n > 1 {
				fmt.Printf("%s\n\n", output)
			}
		}
		for i, result := range results {
			if result {
				fmt.Printf("Program #%d is correct!\n", i+1)
			} else {
				fmt.Printf("Program #%d is incorrect!\n", i+1)
			}
		}
	}
}
