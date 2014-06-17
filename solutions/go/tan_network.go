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
	hash uint32
	cost uint16
}

var hashMap map[uint32]string
var identifierMap map[string]uint32

func GetInput(input <-chan string) string {
	line := <-input
	return string(line[9:])
}

func ToFloat(str string) (x float64) {
	fmt.Sscanf(str, "%f", &x)
	return
}

func GetCost(lo_a, lo_b, la_a, la_b float64) uint16 {
	x, y := (lo_b-lo_a)*math.Cos((la_a+la_b)/2), la_b-la_a
	return uint16((x*x + y*y) * 100000.0)
}

var minCost uint16 = math.MaxUint16
var routes map[uint32][]Destination
var finalHash, startHash uint32
var finalRoute []uint32
var stationsMC map[uint32]uint16

func TravelRecursive(cost uint16, route []uint32) {
	for _, destination := range routes[route[len(route)-1]] {
		if cost += destination.cost; cost < minCost {
			mcValue, mcOK := stationsMC[destination.hash]
			if (mcOK && cost < mcValue) || !mcOK {
				stationsMC[destination.hash] = cost
				if destination.hash == finalHash {
					minCost = cost
					finalRoute = append(route, destination.hash)
				} else {
					TravelRecursive(cost, append(route, destination.hash))
				}
			}
		}
	}
}

func main() {
	cgreader.RunAndValidateManualPrograms(
		cgreader.GetFileList("../../input/tan_network_%d.txt", 6),
		cgreader.GetFileList("../../output/tan_network_%d.txt", 6),
		true,
		func(input <-chan string, output chan string) {
			// this block could be ommited when solo-running
			minCost = math.MaxUint16
			routes, finalRoute = nil, nil

			start, stop := GetInput(input), GetInput(input)
			hashMap = make(map[uint32]string)
			identifierMap = make(map[string]uint32)
			stationsMC = make(map[uint32]uint16)

			var ns, nr uint32
			fmt.Sscanf(<-input, "%d", &ns)
			stations := make(map[uint32]Station)
			for i := uint32(0); i < ns; i++ {
				station := GetInput(input)
				info := strings.Split(station, ",")
				hashMap[i] = info[0]
				identifierMap[info[0]] = i
				stations[i] = Station{
					info[1][1 : len(info[1])-1],
					ToFloat(info[3]),
					ToFloat(info[4])}
			}

			startHash, finalHash = identifierMap[start], identifierMap[stop]

			if startHash == finalHash {
				output <- stations[startHash].name
				return
			}

			fmt.Sscanf(<-input, "%d", &nr)
			routes = make(map[uint32][]Destination)
			for i := uint32(0); i < nr; i++ {
				route := GetInput(input)
				ra, ro := string(route[:4]), string(route[14:])
				ha, ho := identifierMap[ra], identifierMap[ro]

				a, b := stations[ha], stations[ho]
				cost := GetCost(a.latitude, b.latitude, a.longitude, b.longitude)

				routes[ha] = append(routes[ha], Destination{ho, cost})
			}

			TravelRecursive(0, append(make([]uint32, 0), startHash))

			if finalRoute == nil {
				output <- "IMPOSSIBLE"
			} else {
				for _, hash := range finalRoute {
					output <- stations[hash].name
				}
			}
		})
}
