package main

import (
    "fmt"
    "github.com/glendc/cgreader"
)

func main() {
    cgreader.RunAndValidateManualPrograms(
        []string{
            "../../input/input_1.txt",
            "../../input/input_2.txt"},
        []string{
            "../../output/output_1.txt",
            "../../output/output_2.txt"},
        true,
        func(input <-chan string, output chan string) {
            // program logic
        })
}
