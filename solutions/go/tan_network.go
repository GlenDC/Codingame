package main

import (
	"fmt"
	"github.com/glendc/cgreader"
	"math"
	"strings"
)

type Station struct {
	name                string
	longitude, latitude float64
}

type Destination struct {
	identifier string
	cost       float64
}

func GetInput(input <-chan string) string {
	line := <-input
	return string(line[9:])
}

func ToFloat(str string) (x float64) {
	fmt.Sscanf(str, "%f", &x)
	return
}

func GetCost(lo_a, lo_b, la_a, la_b float64) float64 {
	x, y := (lo_b-lo_a)*math.Cos((la_a+la_b)/2), la_b-la_a
	return x*x + y*y
}

var minCost float64 = math.MaxFloat64
var routes map[string][]Destination
var finalStation, startStation string
var finalRoute []string

func TraceDestinations(station string, destinations []Destination) {
	var stops string
	for _, stop := range destinations {
		stops += stop.identifier + ", "
	}
	cgreader.Tracef("%s => Stops: %s", station, stops)
}

func TravelRecursive(cost float64, route []string, counter int) {
	TraceDestinations(route[len(route)-1], routes[route[len(route)-1]])
	for _, station := range routes[route[len(route)-1]] {
		if cost += station.cost; cost < minCost {
			if station.identifier == finalStation {
				minCost = cost
				finalRoute = append(route, station.identifier)
			} else {
				isOK := true
				for _, stop := range route {
					if station.identifier == stop {
						isOK = false
						break
					}
				}
				if isOK {
					cgreader.Tracef("%d. len = %d, to %s costs %f", counter, len(route), station.identifier, cost)
					TravelRecursive(cost, append(route, station.identifier), counter+1)
				}
			}
		}
	}
}

/*
cgreader.RunAndValidateManualPrograms(
		cgreader.GetFileList("../../input/tan_network_%d.txt", 4),
		cgreader.GetFileList("../../output/tan_network_%d.txt", 4),*/

func main() {
	cgreader.RunAndValidateManualProgram(
		"../../input/tan_network_4.txt",
		"../../output/tan_network_4.txt",
		true,
		func(input <-chan string, output chan string) {
			minCost = math.MaxFloat64
			finalStation, startStation = "", ""
			routes, finalRoute = nil, nil
			startStation, finalStation = GetInput(input), GetInput(input)

			var ns, nr uint32

			fmt.Sscanf(<-input, "%d", &ns)
			stations := make(map[string]Station)
			for i := uint32(0); i < ns; i++ {
				station := GetInput(input)
				info := strings.Split(station, ",")
					stations[info[0]] = Station{
						info[1][1 : len(info[1])-1],
						ToFloat(info[3]),
						ToFloat(info[4])}
			}

			if startStation == finalStation {
				output <- stations[startStation].name
				return
			}

			fmt.Sscanf(<-input, "%d", &nr)
			routes = make(map[string][]Destination)
			for i := uint32(0); i < nr; i++ {
				route := GetInput(input)
				ra, ro := string(route[:4]), string(route[14:])

				// maybe this can be omitted.
				// do a/b test when all solutions work
				isOK := true
				for _, destination := range routes[ra] {
					if destination.identifier == ro {
						isOK = false
						break
					}
				}

				if isOK {
					a, b := stations[ra], stations[ro]
					cost := GetCost(
						a.latitude, b.latitude,
						a.longitude, b.longitude)

					if ra == startStation {
						cgreader.Tracef("%s -> %s", ra, ro)
					}

					routes[ra] = append(routes[ra], Destination{ro, cost})
				}
			}

			var startStops string
			for _, stop := range routes[startStation] {
				startStops += stop.identifier + ", "
			}
			cgreader.Tracef("Original destinations: %s", startStops)
			TravelRecursive(0, append(make([]string, 0), startStation), 0)

			if finalRoute == nil {
				output <- "IMPOSSIBLE"
			} else {
				for _, identifier := range finalRoute {
					output <- stations[identifier].name
				}
			}
		})
}
