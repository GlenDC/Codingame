package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var width, height int
	var text string
	fmt.Scanf("%d\n%d\n%s\n", &width, &height, &text)

	text = strings.ToUpper(text)

	reader := bufio.NewReader(os.Stdin)

	ascii := make([]string, height, height)
	for i := 0; i < height; i++ {
		text, _ := reader.ReadString('\n')
		ascii[i] = text
	}

	output := make([]string, height, height)
	for _, char := range text {
		character := int(char) - 65
		if character < 0 || character > 26 {
			character = 26
		}
		for i := range output {
			position := character * width
			output[i] += ascii[i][position : position+width]
		}
	}

	for _, line := range output {
		fmt.Println(line)
	}
}
