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

func readFile(filename string) []int {
	file, err := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	check(err)
	var modules []int
	for scanner.Scan() {
		weight, _ := strconv.Atoi(scanner.Text())
		modules = append(modules, weight)
	}
	return modules
}

func calculateFuelRecursive(weight int) int {
	fuel := (weight / 3) - 2
	if fuel <= 0 {
		return 0
	} else {
		return fuel + calculateFuelRecursive(fuel)
	}
}

func main() {
	modules := readFile("input.txt")
	var fuel int
	for _, weight := range modules {
		fuel += (weight / 3) - 2
	}
	fmt.Println("Part 1 Fuel:", fuel)

	fuel = 0
	for _, weight := range modules {
		fuel += calculateFuelRecursive(weight)
	}
	fmt.Println("Part 2 Fuel:", fuel)
}
