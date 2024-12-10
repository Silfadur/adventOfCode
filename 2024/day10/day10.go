package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

type TrailMap struct {
	heightmap  [][]int
	trailheads []Coordinate
	peaks      []Coordinate

	reachablePeaks map[Coordinate][]Coordinate
}

func (t *TrailMap) readFromFile(filename string) {
	dat, err := os.Open(filename)
	check(err)
	defer dat.Close()
	scanner := bufio.NewScanner(dat)
	var y int
	for scanner.Scan() {
		line := scanner.Text()
		var lineslice []int
		for x, symbol := range line {
			number, err := strconv.Atoi(string(symbol))
			check(err)
			switch number {
			case 0:
				t.trailheads = append(t.trailheads, Coordinate{x: x, y: y})
			case 9:
				t.peaks = append(t.peaks, Coordinate{x: x, y: y})
			}
			lineslice = append(lineslice, number)
		}
		y++
		t.heightmap = append(t.heightmap, lineslice)
	}
	t.reachablePeaks = make(map[Coordinate][]Coordinate)
}

func printMap(input [][]int) {
	for _, line := range input {
		fmt.Println(line)
	}
}

func addUnique(slice []Coordinate, value Coordinate) []Coordinate {
	for _, v := range slice {
		if v == value {
			return slice
		}
	}
	return append(slice, value)
}

func (t *TrailMap) findTrails() {
	for _, trailhead := range t.trailheads {
		t.findTrailsFromTrailhead(trailhead, trailhead)
	}

}

func (t *TrailMap) isInBounds(location Coordinate) bool {
	return location.y >= 0 && location.y < len(t.heightmap) && location.x >= 0 && location.x < len(t.heightmap[location.y])
}

func (t *TrailMap) findTrailsFromTrailhead(trailhead Coordinate, current Coordinate) {
	currentHeight := t.heightmap[current.y][current.x]
	if currentHeight == 9 {
		t.reachablePeaks[trailhead] = append(t.reachablePeaks[trailhead], current)
		return
	}
	right := Coordinate{x: current.x + 1, y: current.y}
	left := Coordinate{x: current.x - 1, y: current.y}
	up := Coordinate{x: current.x, y: current.y - 1}
	down := Coordinate{x: current.x, y: current.y + 1}
	if t.isInBounds(right) && t.heightmap[right.y][right.x] == currentHeight+1 {
		t.findTrailsFromTrailhead(trailhead, right)
	}
	if t.isInBounds(left) && t.heightmap[left.y][left.x] == currentHeight+1 {
		t.findTrailsFromTrailhead(trailhead, left)
	}
	if t.isInBounds(up) && t.heightmap[up.y][up.x] == currentHeight+1 {
		t.findTrailsFromTrailhead(trailhead, up)
	}
	if t.isInBounds(down) && t.heightmap[down.y][down.x] == currentHeight+1 {
		t.findTrailsFromTrailhead(trailhead, down)
	}
}

func main() {
	var t TrailMap
	t.readFromFile("input.txt")
	printMap(t.heightmap)
	fmt.Println(t.trailheads)
	fmt.Println(t.peaks)
	t.findTrails()
	var score int
	for _, peaks := range t.reachablePeaks {
		score += len(peaks)
	}
	fmt.Println("Score:", score)

}
