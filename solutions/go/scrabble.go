package main

import (
	"fmt"
	"github.com/glendc/cgreader"
)

func GetPoints(character string) uint {
	switch character {
	case "e", "a", "i", "o", "n", "r", "t", "l", "s", "u":
		return 1
	case "d", "g":
		return 2
	case "b", "c", "m", "p":
		return 3
	case "f", "h", "v", "w", "y":
		return 4
	case "k":
		return 5
	case "j", "x":
		return 8
	case "q", "z":
		return 10
	}
	cgreader.Tracef("Error, unknown Character: %s", character)
	return 0
}

type WordInformation struct {
	points, amount uint
}

func (w WordInformation) Increase() {
	w.amount++
}

func GetWordInformation(word string) map[string]WordInformation {
	info := make(map[string]WordInformation)
	for i := range word {
		c := string(word[i])
		if _, ok := info[c]; !ok {
			info[c] = WordInformation{GetPoints(c), 1}
		} else {
			info[c].Increase()
		}
	}
	return info
}

func GetCharacterAmount(word string) map[string]uint {
	amount := make(map[string]uint)
	for i := range word {
		c := string(word[i])
		if _, ok := amount[c]; ok {
			amount[c] = 1
		} else {
			amount[c]++
		}
	}
	return amount
}

func main() {
	cgreader.RunAndValidateManualPrograms(
		[]string{
			"../../input/scrabble_1.txt",
			"../../input/scrabble_2.txt",
			"../../input/scrabble_3.txt",
			"../../input/scrabble_4.txt",
			"../../input/scrabble_5.txt",
			"../../input/scrabble_6.txt",
			"../../input/scrabble_7.txt"},
		[]string{
			"../../output/scrabble_1.txt",
			"../../output/scrabble_2.txt",
			"../../output/scrabble_3.txt",
			"../../output/scrabble_4.txt",
			"../../output/scrabble_5.txt",
			"../../output/scrabble_6.txt",
			"../../output/scrabble_7.txt"},
		true,
		func(input <-chan string, output chan string) {
			var n int
			fmt.Sscanf(<-input, "%d", &n)

			dictionary := make([]string, n)
			for i := range dictionary {
				fmt.Sscanf(<-input, "%s", &dictionary[i])
			}

			var w string
			fmt.Sscanf(<-input, "%s", &w)

			word := GetWordInformation(w)

			bw, bwp := "", uint(0)
			for i := range dictionary {
				info, points := GetCharacterAmount(dictionary[i]), uint(0)
				for c, n := range info {
					if value, ok := word[c]; ok {
						if n > value.amount {
							n = value.amount
						}

						points += value.points * n
					} else {
						points = 0
						break
					}
				}
				if points > bwp {
					bwp, bw = points, dictionary[i]
				}
			}

			output <- bw
		})
}
