package main

import (
	"fmt"
	"github.com/glendc/cgreader"
)

type Program struct {
	N            int
	Calculations map[int]int
}

func (p *Program) Update(input string) {
	if p.N == 0 {
		fmt.Sscanf(input, "%d\n", &p.N)
		p.Calculations = make(map[int]int, p.N)
	} else {
		var start, duration int
		fmt.Sscanf(input, "%d %d\n", &start, &duration)

		if _, ok := p.Calculations[start]; !ok {
			ok = true
			for s, d := range p.Calculations {
				if !(start+duration < s || start > s+d) {
					ok = false
					break
				}
			}
			if ok {
				fmt.Printf("%d = %d\n", start, duration)
				p.Calculations[start] = duration
			}
		}
	}
}

func (p *Program) GetOutput() string {
	return fmt.Sprintf("%d\n", len(p.Calculations))
}

func main() {
	cgreader.RunAndValidateFlowProgram(
		"../input/super_computer_1.txt",
		"../output/super_computer_1.txt",
		true,
		&Program{})
}
