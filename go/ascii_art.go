package main

import (
	"fmt"
	"strings"
	"github.com/glendc/cgreader"
)

func main() {
	progam_id := 1

	input := fmt.Sprintf("../input/ascii_%d.txt", progam_id)
	ch := cgreader.GetManualInput(input)

	var width, height int
	var text string

	fmt.Sscanln(<-ch, &width)
	fmt.Sscanln(<-ch, &height)
	fmt.Sscanln(<-ch, &text)

	text = strings.ToUpper(text)

	ascii := make([]string, height)
	for i := 0; i < height; i++ {
		ascii[i] = <-ch
	}

	output := make([]string, height)
	for _, char := range text {
		character := int(char) - 65
		if character < 0 || character > 26 {
			character = 26
		}
		for i := range output {
			position := character * width
			output[i] += ascii[i][position:position+width]
		}
	}

	var program_output string

	for _, line := range output {
		fmt.Printf("%s\n", line)
		program_output += line + "\n"
	}

	test := fmt.Sprintf("../output/ascii_%d.txt", progam_id)
	result := cgreader.TestOutput(test, program_output)

	if result {
		fmt.Printf("Program #%d is correct!\n", progam_id)
	} else {
		fmt.Printf("Program #%d is incorrect!\n", progam_id)
	}
}
