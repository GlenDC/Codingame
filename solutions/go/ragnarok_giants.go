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
	thor, dimensions       Vector
	energy, turn, maxTurns int
	giants                 []Vector
}

var WAIT string = "WAIT"
var STRIKE string = "STRIKE"

func GetDirection(a, b string, x, y int) <-chan string {
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
		"%d %d %d \n",
		&ragnarok.dimensions.x,
		&ragnarok.dimensions.y,
		&ragnarok.maxTurns)

	var giants int

	fmt.Sscanf(
		<-ch,
		"%d %d %d %d \n",
		&ragnarok.energy,
		&ragnarok.thor.x,
		&ragnarok.thor.y,
		&giants)

	ragnarok.giants = make([]Vector, giants)

	for i := range ragnarok.giants {
		fmt.Sscanf(
			<-ch,
			"%d %d \n",
			&ragnarok.giants[i].x,
			&ragnarok.giants[i].y)
		ragnarok.giants[i].icon = "G"
	}

	ragnarok.thor.icon = "H"
}

func (ragnarok *Ragnarok) GetInput() (ch chan string) {
	ch = make(chan string)
	go func() {
		ch <- fmt.Sprintf("%d", ragnarok.energy)
	}()
	return
}

func (ragnarok *Ragnarok) Update(ch <-chan string) string {
	return WAIT
}

func (ragnarok *Ragnarok) SetOutput(output string) string {
	if output == STRIKE {
		// do strike....
		ragnarok.energy -= 1
	} else if output != WAIT {
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
	}

	ragnarok.turn++

	giants := append(ragnarok.giants, ragnarok.thor)

	map_info := make([]cgreader.MapObject, len(giants))
	for i, v := range giants {
		map_info[i] = cgreader.MapObject(v)
	}

	cgreader.DrawMap(
		ragnarok.dimensions.x,
		ragnarok.dimensions.y,
		".",
		map_info...)

	return fmt.Sprintf(
		"Amount of Giants = %d\nThor = (%d,%d)\nEnergy = %d",
		len(ragnarok.giants),
		ragnarok.thor.x,
		ragnarok.thor.y,
		ragnarok.energy)
}

func (ragnarok *Ragnarok) LoseConditionCheck() bool {
	if ragnarok.energy <= 0 || ragnarok.turn >= ragnarok.maxTurns {
		return true
	}

	x, y := ragnarok.thor.x, ragnarok.thor.y
	dx, dy := ragnarok.dimensions.x, ragnarok.dimensions.y

	for _, giant := range ragnarok.giants {
		if giant.x == x && giant.y == y {
			return true
		}
	}

	if x < 0 || x >= dx || y < 0 || y >= dy {
		return true
	}

	return false
}

func (ragnarok *Ragnarok) WinConditionCheck() bool {
	return len(ragnarok.giants) == 0
}

func main() {
	cgreader.RunTargetProgram("../../input/ragnarok_giants_10.txt", true, &Ragnarok{})
}
