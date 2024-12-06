package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"sync/atomic"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Heading int

const (
	None = iota
	Up
	Right
	Down
	Left
)

type Coordinate struct {
	x int
	y int
}

type ManufactoringLab struct {
	guardX            int
	guardY            int
	heading           Heading
	labMap            []string
	tilesVisited      []string
	headingGrid       [][]Heading
	visitedGrid       [][]int
	possibleBlockages []Coordinate
	probablyLoop      bool
}

func (ml *ManufactoringLab) updateVisitedTiles() {
	ml.visitedGrid[ml.guardY][ml.guardX]++
	if ml.visitedGrid[ml.guardY][ml.guardX] > 10 {
		ml.probablyLoop = true
	}

	if ml.tilesVisited[ml.guardY][ml.guardX] == '+' {
		return
	}
	if (ml.tilesVisited[ml.guardY][ml.guardX] == '|') ||
		(ml.tilesVisited[ml.guardY][ml.guardX] == '-') {
		ml.tilesVisited[ml.guardY] = ml.tilesVisited[ml.guardY][:ml.guardX] + "+" + ml.tilesVisited[ml.guardY][ml.guardX+1:]
		return
	}
	ml.headingGrid[ml.guardY][ml.guardX] = ml.heading
	switch ml.heading {
	case Up, Down:
		ml.tilesVisited[ml.guardY] = ml.tilesVisited[ml.guardY][:ml.guardX] + "|" + ml.tilesVisited[ml.guardY][ml.guardX+1:]
	case Left, Right:
		ml.tilesVisited[ml.guardY] = ml.tilesVisited[ml.guardY][:ml.guardX] + "-" + ml.tilesVisited[ml.guardY][ml.guardX+1:]
	}
}

func (ml *ManufactoringLab) nextStepInBounds() bool {
	switch ml.heading {
	case Up:
		return ml.guardY-1 >= 0
	case Down:
		return ml.guardY+1 < len(ml.labMap)
	case Left:
		return ml.guardX-1 >= 0
	case Right:
		return ml.guardX+1 < len(ml.labMap[ml.guardY])
	default:
		return false
	}
}

func (ml *ManufactoringLab) step() bool {
	if ml.nextStepInBounds() {
		switch ml.heading {
		case Up:
			if ml.labMap[ml.guardY-1][ml.guardX] == '#' {
				ml.heading = Right
				ml.updateVisitedTiles()
				return ml.step()
			}
			ml.guardY--
		case Down:
			if ml.labMap[ml.guardY+1][ml.guardX] == '#' {
				ml.heading = Left
				ml.updateVisitedTiles()
				return ml.step()
			}
			ml.guardY++
		case Left:
			if ml.labMap[ml.guardY][ml.guardX-1] == '#' {
				ml.heading = Up
				ml.updateVisitedTiles()
				return ml.step()
			}
			ml.guardX--
		case Right:
			if ml.labMap[ml.guardY][ml.guardX+1] == '#' {
				ml.heading = Down
				ml.updateVisitedTiles()
				return ml.step()
			}
			ml.guardX++
		}
	} else {
		return false
	}
	return true
}
func (ml *ManufactoringLab) countVisitedTiles() int {
	var visited int
	for _, line := range ml.tilesVisited {
		for _, symbol := range line {
			if symbol == '+' || symbol == '-' || symbol == '|' {
				visited++
			}
		}
	}
	return visited
}

func (ml *ManufactoringLab) findPossibleBlockages() int {
	var blockages int
	for y, line := range ml.labMap {
		for x, symbol := range line {
			if symbol != '^' && symbol != 'X' && symbol != '#' {
				ml.possibleBlockages = append(ml.possibleBlockages, Coordinate{x: x, y: y})
			}
		}
	}
	return blockages
}

func (ml *ManufactoringLab) loadMap(filename string) {
	dat, err := os.Open(filename)
	check(err)
	defer dat.Close()
	scanner := bufio.NewScanner(dat)
	for scanner.Scan() {
		ml.labMap = append(ml.labMap, scanner.Text())
	}
	ml.tilesVisited = make([]string, len(ml.labMap))
	for i, _ := range ml.tilesVisited {
		var newLine string
		for j := 0; j < len(ml.labMap[i]); j++ {
			newLine += "."
		}
		ml.tilesVisited[i] = newLine
	}
	ml.headingGrid = make([][]Heading, len(ml.labMap))
	ml.visitedGrid = make([][]int, len(ml.labMap))
	for i := 0; i < len(ml.labMap); i++ {
		ml.headingGrid[i] = make([]Heading, len(ml.labMap[i]))
		ml.visitedGrid[i] = make([]int, len(ml.labMap[i]))
	}
	ml.findGuard()
	ml.updateVisitedTiles()
}

func (ml *ManufactoringLab) findGuard() {
	for y, line := range ml.labMap {
		for x, symbol := range line {
			switch symbol {
			case 'v':
				ml.tilesVisited[y] = ml.tilesVisited[y][:x] + "." + ml.tilesVisited[y][x+1:]
				ml.guardX = x
				ml.guardY = y
				ml.heading = Down
			case '^':
				ml.tilesVisited[y] = ml.tilesVisited[y][:x] + "." + ml.tilesVisited[y][x+1:]
				ml.guardX = x
				ml.guardY = y
				ml.heading = Up
			case '<':
				ml.tilesVisited[y] = ml.tilesVisited[y][:x] + "." + ml.tilesVisited[y][x+1:]
				ml.guardX = x
				ml.guardY = y
				ml.heading = Left
			case '>':
				ml.tilesVisited[y] = ml.tilesVisited[y][:x] + "." + ml.tilesVisited[y][x+1:]
				ml.guardX = x
				ml.guardY = y
				ml.heading = Right
			case '#':
				ml.tilesVisited[y] = ml.tilesVisited[y][:x] + "#" + ml.tilesVisited[y][x+1:]
			default:
				continue
			}
		}
	}
	//fmt.Printf("Guard: %d:%d, tile:%s\n", ml.guardX, ml.guardY, string(ml.labMap[ml.guardY][ml.guardX]))
}

func printMap(input []string) {
	fmt.Printf("\\ ")
	for i, _ := range input[0] {
		fmt.Printf("%d", i)
	}
	fmt.Println("")
	for i, line := range input {
		fmt.Println(i, line)
	}
}
func (ml *ManufactoringLab) printHeadings() {
	for _, line := range ml.headingGrid {
		fmt.Println(line)
	}
}

func (ml *ManufactoringLab) run(done chan bool) {
	for ml.step() {
		//fmt.Printf("Guard: %d:%d, tile:%s\n", ml.guardX, ml.guardY, string(ml.labMap[ml.guardY][ml.guardX]))
		ml.updateVisitedTiles()
		if ml.probablyLoop {
			done <- true
			return
		}
	}
	done <- false
}
func main() {
	var lab ManufactoringLab
	testfile := "input.txt"
	lab.loadMap(testfile)
	fmt.Println("Possible blockages found:", lab.findPossibleBlockages())
	var verifiedBlockage atomic.Uint64
	var failedBlockages atomic.Uint64
	var wg sync.WaitGroup
	for _, possibleBlockage := range lab.possibleBlockages {
		wg.Add(1)
		worker := func() {
			//fmt.Println("Testing Blockage:", possibleBlockage)

			var testlab ManufactoringLab
			testlab.loadMap(testfile)

			testlab.labMap[possibleBlockage.y] = testlab.labMap[possibleBlockage.y][:possibleBlockage.x] + "#" + testlab.labMap[possibleBlockage.y][possibleBlockage.x+1:]

			loop := make(chan bool, 1)
			go testlab.run(loop)
			isLoop := <-loop
			if isLoop {
				verifiedBlockage.Add(1)
				wg.Done()
			} else {
				failedBlockages.Add(1)
				wg.Done()
			}
		}
		go worker()
	}
	wg.Wait()
	fmt.Println("PossibleBlockages:", len(lab.possibleBlockages))
	fmt.Println("VerifiedBlockages:", verifiedBlockage)
	fmt.Println("FailedBlockages:", failedBlockages)
}
