package main

import (
	"fmt"
)

type position struct {
	x, y int
}

type result struct {
	value     int
	direction string
}

func GetDirection(a, b string, x, y, v int) <-chan result {
	ch := make(chan result)
	go func() {
		difference := x - y
		switch {
		case difference < 0:
			ch <- result{v - 1, a}
		case difference > 0:
			ch <- result{v + 1, b}
		default:
			ch <- result{v, ""}
		}
		close(ch)
	}()
	return ch
}

func main() {
	var target, thor position
	fmt.Scanf("%d %d %d %d\n", &target.x, &target.y, &thor.x, &thor.y)

	for {
		channel_b := GetDirection("N", "S", target.y, thor.y, thor.y)
		channel_a := GetDirection("E", "W", thor.x, target.x, thor.x)

		result_b := <-channel_b
		result_a := <-channel_a

		thor.y = result_b.value
		thor.x = result_a.value

		fmt.Println(result_b.direction + result_a.direction)
	}
}
