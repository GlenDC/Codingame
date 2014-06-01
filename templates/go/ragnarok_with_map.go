package main

import (
	"fmt"
	"github.com/glendc/cgreader"
	"strings"
)

type Vector struct {
	x, y int
	icon string
}

func (v Vector) GetMapCoordinates() string {
	return fmt.Sprintf("%d;%d", v.x, v.y)
}

func (v Vector) GetMapIcon() string {
	return v.icon
}

type Ragnarok struct {
	thor, target, dimensions Vector
	energy                   int
	trail                    []Vector
}

func (ragnarok *Ragnarok) ParseInitialData(ch <-chan string) {
	// parse initial input
}

func (ragnarok *Ragnarok) GetInput() (ch chan string) {
	ch = make(chan string)
	go func() {
		ch <- fmt.Sprintf("%d", ragnarok.energy)
	}()
	return
}

func (ragnarok *Ragnarok) Update(ch <-chan string) string {
	// define your solution update logic
	return "N"
}

func (ragnarok *Ragnarok) SetOutput(output string) string {
	ragnarok.trail = append(ragnarok.trail, Vector{ragnarok.thor.x, ragnarok.thor.y, "+"})

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

	trail := append(ragnarok.trail, ragnarok.thor, ragnarok.target)

	map_info := make([]cgreader.MapObject, len(trail))
	for i, v := range trail {
		map_info[i] = cgreader.MapObject(v)
	}

	cgreader.DrawMap(
		ragnarok.dimensions.x,
		ragnarok.dimensions.y,
		".",
		map_info...)

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
