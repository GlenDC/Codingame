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

func GetInput(input <-chan string) string {
	var line string
	fmt.Sscanf(<-input, "%s", &line)
	return string(line[9:])
}

func ToFloat(str string) (x float64) {
	fmt.Sscanf(str, "%f", &x)
	return
}

func GetDistance(lo_a, lo_b, la_a, la_b float64) float64 {
	x, y := (lo_b-lo_a)*math.Cos((la_a+la_b)/2), la_b-la_a
	return math.Sqrt(math.Pow(x, 2)+math.Pow(y, 2)) * 6731
}

func main() {
	cgreader.RunAndValidateManualPrograms(
		cgreader.GetFileList("../../input/tan_network_%d.txt", 6),
		cgreader.GetFileList("../../output/tan_network_%d.txt", 6),
		true,
		func(input <-chan string, output chan string) {
			start, stop := GetInput(input), GetInput(input)
			var ns, nr uint32

			fmt.Sscanf(<-input, "%u", &ns)
			stations := make(map[string]Station)
			for i := uint32(0); i < ns; i++ {
				station := GetInput(input)
				info := strings.Split(station, ",")
				stations[info[0]] = Station{
					info[1],
					ToFloat(info[3]),
					ToFloat(info[4])}
			}

			fmt.Sscanf(<-input, "%u", &nr)
			routes := make(map[string]float64)
			for i := uint32(0); i < nr; i++ {
				route := GetInput(input)
				ra, ro := string(route[:4]), string(route[13:])

				a, b := stations[ra], stations[ro]
				routes[ra+ro] = GetDistance(
					a.latitude, b.latitude,
					a.longitude, b.longitude)
			}

			output <- start + stop
		})
}
