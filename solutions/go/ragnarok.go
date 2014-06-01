package main

import (
	"fmt"
	"github.com/glendc/cgreader"
	"strings"
)

type Vector struct {
	x, y int
}

type Ragnarok struct {
	thor, target, dimensions Vector
	energy                   int
}

func GetDirection(a, b string, x, y, v int) <-chan string {
	ch := make(chan string)
	go func() {
		difference := x - y
		switch {
		case difference < 0:
			ch <- a
		case difference > 0:
			ch <- b
		default:
			ch <- ""
		}
		close(ch)
	}()
	return ch
}

func (ragnarok *Ragnarok) ParseInitialData(ch <-chan string) {
	fmt.Sscanf(
		<-ch,
		"%d %d %d %d %d %d %d \n",
		&ragnarok.dimensions.x,
		&ragnarok.dimensions.y,
		&ragnarok.thor.x,
		&ragnarok.thor.y,
		&ragnarok.target.x,
		&ragnarok.target.y,
		&ragnarok.energy)
}

func (ragnarok *Ragnarok) GetInput() (ch chan string) {
	ch = make(chan string)
	go func() {
		ch <- fmt.Sprintf("%d", ragnarok.energy)
	}()
	return
}

func (ragnarok *Ragnarok) Update(ch <-chan string) string {
	channel_b := GetDirection("N", "S", ragnarok.target.y, ragnarok.thor.y, ragnarok.thor.y)
	channel_a := GetDirection("E", "W", ragnarok.thor.x, ragnarok.target.x, ragnarok.thor.x)

	result_b := <-channel_b
	result_a := <-channel_a

	return fmt.Sprint(result_b + result_a)
}

func (ragnarok *Ragnarok) SetOutput(output string) string {
	if strings.Contains(output, "N") {
		ragnarok.thor.y -= 1
	} else if strings.Contains(output, "S") {
		ragnarok.thor.y += 1
	}

	if strings.Contains(output, "E") {
		ragnarok.thor.x += 1
	} else if strings.Contains(output, "W") {
		ragnarok.thor.x -= 1
	}

	ragnarok.energy -= 1

	return fmt.Sprintf(
		"Target = (%d,%d)\nThor = (%d,%d)\nEnergy = %d",
		ragnarok.target.x,
		ragnarok.target.y,
		ragnarok.thor.x,
		ragnarok.thor.y,
		ragnarok.energy)
}

func (ragnarok *Ragnarok) LoseConditionCheck() bool {
	if ragnarok.energy <= 0 {
		return true
	}

	x, y := ragnarok.thor.x, ragnarok.thor.y
	dx, dy := ragnarok.dimensions.x, ragnarok.dimensions.y

	if x < 0 || x >= dx || y < 0 || y >= dy {
		return true
	}

	return false
}

func (ragnarok *Ragnarok) WinConditionCheck() bool {
	return ragnarok.target == ragnarok.thor
}

func main() {
	cgreader.RunTargetProgram("../../input/ragnarok_1.txt", true, &Ragnarok{})
}
