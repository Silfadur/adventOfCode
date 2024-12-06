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

type Heading int

const (
	Up = iota
	Down
	Left
	Right
)

type ManufactoringLab struct {
	guardX       int
	guardY       int
	heading      Heading
	labMap       []string
	tilesVisited []string
}

func (ml *ManufactoringLab) run() {
	for ml.step() {
		//fmt.Printf("Guard: %d:%d, tile:%s\n", ml.guardX, ml.guardY, string(ml.labMap[ml.guardY][ml.guardX]))
		ml.updateVisitedTiles()
	}

}

func (ml *ManufactoringLab) updateVisitedTiles() {
	if (ml.tilesVisited[ml.guardY][ml.guardX] == '|') ||
		(ml.tilesVisited[ml.guardY][ml.guardX] == '-') ||
		(ml.tilesVisited[ml.guardY][ml.guardX] == '+') {
		ml.tilesVisited[ml.guardY] = ml.tilesVisited[ml.guardY][:ml.guardX] + "+" + ml.tilesVisited[ml.guardY][ml.guardX+1:]
		return
	}
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
	fmt.Printf("Guard: %d:%d, tile:%s\n", ml.guardX, ml.guardY, string(ml.labMap[ml.guardY][ml.guardX]))
}

func printMap(input []string) {
	for _, line := range input {
		fmt.Println(line)
	}
}

func main() {
	var lab ManufactoringLab
	lab.loadMap("input.txt")
	lab.run()
	printMap(lab.tilesVisited)
	fmt.Println("Tiles visited:", lab.countVisitedTiles())
}
