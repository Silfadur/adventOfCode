package main

import (
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}

}

func readfile(filename string) ([]int, []int) {
	dat, err := os.ReadFile(filename)
	check(err)
	lines := strings.Split(string(dat), "\r\n")
	var firstlist []int
	var secondlist []int
	for _, line := range lines {
		elements := strings.Split(line, "   ")
		var number int
		number, err = strconv.Atoi(elements[0])
		check(err)
		firstlist = append(firstlist, number)
		number, err = strconv.Atoi(elements[1])
		check(err)
		secondlist = append(secondlist, number)
	}
	return firstlist, secondlist
}
func countOccurences(numberToFind int, numbers []int) int {
	var occurences int
	for _, number := range numbers {
		if number == numberToFind {
			occurences++
		}
	}
	return occurences
}

func main() {
	firstlist, secondlist := readfile("../input.txt")
	var similarity int
	for _, number := range firstlist {
		similarity += number * countOccurences(number, secondlist)
	}
	println("Similarity:", similarity)
}
