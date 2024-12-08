package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Coordinate struct {
	x int
	y int
}

type AntennaMap struct {
	grid      []string
	antennas  map[rune][]Coordinate
	antinodes map[Coordinate]bool
}

func printMap(input []string) {
	fmt.Printf("\\\t ")
	for i, _ := range input[0] {
		fmt.Printf("%d", i)
	}
	fmt.Println("")
	for i, line := range input {
		fmt.Println(i, "\t", line)
	}

}

func (am *AntennaMap) readFile(filename string) {
	dat, err := os.Open(filename)
	check(err)
	defer dat.Close()
	scanner := bufio.NewScanner(dat)
	for scanner.Scan() {
		am.grid = append(am.grid, scanner.Text())
	}
	am.antennas = make(map[rune][]Coordinate)
	am.antinodes = make(map[Coordinate]bool)
	for y, line := range am.grid {
		for x, frequency := range line {
			if frequency != '.' {
				am.antennas[frequency] = append(am.antennas[frequency], Coordinate{x: x, y: y})
			}
		}
	}

}

func isInBounds(location Coordinate, grid []string) bool {
	return location.x >= 0 && location.y >= 0 && location.y < len(grid) && location.x < len(grid[location.y])
}

func (am *AntennaMap) findAntinodes() {
	var possibleLocations [][2]Coordinate
	for _, list := range am.antennas {
		for i := 0; i < len(list); i++ {
			for j := 0; j < len(list); j++ {
				if j == i {
					continue
				}
				possibleLocations = append(possibleLocations, [2]Coordinate{list[i], list[j]})
			}
		}
	}
	for _, locations := range possibleLocations {
		a1 := locations[0]
		a2 := locations[1]
		d := Coordinate{x: a2.x - a1.x, y: a2.y - a1.y}
		for mul := -70; mul <= 70; mul++ {
			antinode := Coordinate{x: a2.x + (mul * d.x), y: a2.y + (mul * d.y)}
			if isInBounds(antinode, am.grid) {
				am.antinodes[antinode] = true
			}
		}
	}
	fmt.Println("Number of unique Antinode locations:", len(am.antinodes))
}

func main() {
	var am AntennaMap
	am.readFile("input.txt")
	printMap(am.grid)
	am.findAntinodes()

}
