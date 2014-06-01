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

func GetDirection(x, y int) <-chan int {
	ch := make(chan int)
	go func() {
		difference := x - y
		switch {
		case difference < 0:
			ch <- -1
		case difference > 0:
			ch <- 1
		default:
			ch <- 0
		}
		close(ch)
	}()
	return ch
}

func GetDirectionLetter(a, b string, v int) string {
	switch v {
	default:
		return ""
	case -1:
		return a
	case 1:
		return b
	}
}

var TX, TY, PX, PY, E int

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

	TX, TY = ragnarok.thor.x, ragnarok.thor.y
	PX, PY = ragnarok.target.x, ragnarok.target.y
	E = ragnarok.energy

	ragnarok.thor.icon, ragnarok.target.icon = "H", "T"
}

func (ragnarok *Ragnarok) GetInput() (ch chan string) {
	ch = make(chan string)
	go func() {
		ch <- fmt.Sprintf("%d", ragnarok.energy)
	}()
	return
}

func (ragnarok *Ragnarok) Update(ch <-chan string) string {
	fmt.Sscanf(<-ch, "%d", &E)

	chx := GetDirection(PX, TX)
	chy := GetDirection(PY, TY)

	dx, dy := <-chx, <-chy
	x := GetDirectionLetter("W", "E", dx)
	y := GetDirectionLetter("N", "S", dy)

	TX, TY = TX+dx, TY+dy

	return y+x
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
	return ragnarok.target.x == ragnarok.thor.x &&
		ragnarok.target.y == ragnarok.thor.y
}

func main() {
	cgreader.RunTargetProgram("../../input/ragnarok_3.txt", true, &Ragnarok{})
}
