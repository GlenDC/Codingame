package main

import (
    "fmt"
    "strings"
    "github.com/glendc/cgreader"
)

func main() {
    cgreader.RunAndValidateManualPrograms(
        []string{
            "../../input/stock_exchange_losses_1.txt",
            "../../input/stock_exchange_losses_2.txt",
            "../../input/stock_exchange_losses_3.txt",
            "../../input/stock_exchange_losses_4.txt",
            "../../input/stock_exchange_losses_5.txt"},
        []string{
            "../../output/stock_exchange_losses_1.txt",
            "../../output/stock_exchange_losses_2.txt",
            "../../output/stock_exchange_losses_3.txt",
            "../../output/stock_exchange_losses_4.txt",
            "../../output/stock_exchange_losses_5.txt"},
        true,
        func(input <-chan string, output chan string) {
            var n int
            fmt.Sscanf(<-input, "%d", &n)

            s := strings.Split(<-input, " ")
            stocks := make([]int, n)
            for i := range s {
                fmt.Sscanf(s[i], "%d", &stocks[i])
            }

            var difference int
            for i := 0; i < len(stocks) - 1; i++ {
                for u := i+1; u < len(stocks); u++ {
                    if i != u {
                        if d := stocks[u] - stocks[i]; d < difference {
                            difference = d
                        }
                    }
                }
            }

            output <- fmt.Sprintf("%d", difference)
        })
}
