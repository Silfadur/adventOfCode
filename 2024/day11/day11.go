package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readInput(filename string) map[int]int {
	data := make(map[int]int)
	file, err := os.ReadFile(filename)
	check(err)
	strings := strings.Split(string(file), " ")
	for _, numberString := range strings {
		number, err := strconv.Atoi(numberString)
		check(err)
		data[number]++
	}
	return data
}

func getNumberOfDigits(number int) int {
	count := 0
	for number > 0 {
		number /= 10
		count++
	}
	return count
}

func splitDigits(number int) (int, int) {
	left, right := 0, 0
	numberofDigits := getNumberOfDigits(number)
	for i := 0; i < numberofDigits/2; i++ {
		digit := number % 10
		right = right + digit*int(math.Pow(10, float64(i)))
		number /= 10
	}
	for i := 0; i < numberofDigits/2; i++ {
		digit := number % 10
		left = left + digit*int(math.Pow(10, float64(i)))
		number /= 10
	}
	return left, right
}

func blink(data map[int]int) map[int]int {
	newData := make(map[int]int)
	for k, v := range data {
		newData[k] = v
	}
	for stone, count := range data {
		if count == 0 {
			continue
		}
		switch {
		case stone == 0: // if stone is engraved with zero, replace with stone engraved with 1
			newData[0] -= count
			newData[1] += count
		case getNumberOfDigits(stone)%2 == 0: // if stone has even number of digits, replace with two stones, one with the right half of digits and the other with the left half
			left, right := splitDigits(stone)
			newData[stone] -= count
			newData[left] += count
			newData[right] += count
		default:
			newData[stone] -= count
			newData[stone*2024] += count
		}
	}
	return newData
}

func countStones(data map[int]int) int {
	count := 0
	for _, value := range data {
		count += value
	}
	return count
}

func printData(data map[int]int) {
	keys := make([]int, 0, len(data))
	for key := range data {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for _, key := range keys {
		count := data[key]
		if count != 0 {
			fmt.Printf("%d: %d ", key, count)
		}
	}
	fmt.Println()
}

func main() {
	data := readInput("input.txt")
	for i := 0; i < 75; {
		//fmt.Printf("B Blink %d: %d stones ", i, countStones(data))
		//printData(data)
		data = blink(data)
		//fmt.Printf("A Blink %d: %d stones ", i, countStones(data))
		//printData(data)
		i++
	}
	fmt.Println(countStones(data))

}
