package main

import (
    "fmt"
    "github.com/glendc/cgreader"
)

func main() {
    cgreader.RunAndValidateManualPrograms(
        cgreader.GetFileList("../../input/cgx_formatter_%d.txt", 12),
        cgreader.GetFileList("../../output/cgx_formatter_%d.txt", 12),
        true,
        func(input <-chan string, output chan string) {
            // program logic
        })
}
