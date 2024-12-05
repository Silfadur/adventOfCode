package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const ruleRegexString string = `(\d{2}\|\d{2})`
const pagesRegexString string = `(\d{2},)+(\d{2},?)`

func readFile(filename string) string {
	dat, err := os.ReadFile(filename)
	check(err)
	return string(dat)
}

func check(e error) {
	if e != nil {
		panic(e)
	}

}

// 47|53 means 47 must be printed before 53

var rules map[int][]int

func aocSort(a, b int) bool {
	for _, i := range rules[a] {
		if i == b {
			return true
		}
	}
	return false
}

func main() {
	rules = make(map[int][]int)
	input := readFile("input.txt")

	rulesRegex := regexp.MustCompile(ruleRegexString)
	pagesRegex := regexp.MustCompile(pagesRegexString)

	rulesSlice := rulesRegex.FindAllString(input, -1)
	pagesSlice := pagesRegex.FindAllString(input, -1)

	for _, rule := range rulesSlice {
		parts := strings.Split(rule, "|")
		first, _ := strconv.Atoi(parts[0])
		second, _ := strconv.Atoi(parts[1])
		rules[first] = append(rules[first], second)
	}
	var updates [][]int
	for _, pageString := range pagesSlice {
		parts := strings.Split(pageString, ",")
		var update []int
		for _, page := range parts {
			pageNumber, _ := strconv.Atoi(page)
			update = append(update, pageNumber)
		}
		updates = append(updates, update)
	}
	var resultSorted int
	var resultUnsorted int
	for _, update := range updates {
		sorter := func(i, j int) bool {
			return aocSort(update[i], update[j])
		}
		if !sort.SliceIsSorted(update, sorter) {
			sort.Slice(update, sorter)
			resultUnsorted += update[len(update)/2]
		} else {
			resultSorted += update[len(update)/2]
		}
	}
	fmt.Println("Result all sorted:", resultSorted)
	fmt.Println("Result all unsorted:", resultUnsorted)
}
