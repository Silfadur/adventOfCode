package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}

}

func main() {
	dat, err := os.ReadFile("../input.txt")
	check(err)
	//fmt.Println("File input:")
	//fmt.Println(string(dat))
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
	sort.Slice(firstlist, func(i, j int) bool {
		return firstlist[i] < firstlist[j]
	})
	sort.Slice(secondlist, func(i, j int) bool {
		return secondlist[i] < secondlist[j]
	})
	//fmt.Println("First list:", firstlist)
	//fmt.Println("Second list", secondlist)
	totaldistance := 0
	for range firstlist {
		distance := firstlist[0] - secondlist[0]
		if distance < 0 {
			distance = distance * -1
		}
		totaldistance += distance
		firstlist = firstlist[1:]
		secondlist = secondlist[1:]
	}
	fmt.Println("Total distance:", totaldistance)
}
