package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Calibration struct {
	testValue int
	numbers   []int
}

func readFile(filename string) []Calibration {
	var cals []Calibration
	file, err := os.Open(filename)
	check(err)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		var cal Calibration
		cal.testValue, err = strconv.Atoi(parts[0])
		check(err)
		numbers := strings.Split(parts[1], " ")
		for _, number := range numbers {
			parsed, err := strconv.Atoi(number)
			if err != nil {
				continue
			}
			cal.numbers = append(cal.numbers, parsed)
		}
		cals = append(cals, cal)
	}
	return cals
}

func deepCopy(original Calibration) Calibration {
	newSlice := make([]int, len(original.numbers))
	copy(newSlice, original.numbers)
	original.numbers = newSlice
	return original
}

func concatNumbers(a, b int) int {
	concat := strconv.Itoa(a) + strconv.Itoa(b)
	number, err := strconv.Atoi(concat)
	check(err)
	return number
}

func verifyCalibrationRec(cal Calibration, currentValue int) bool {
	if currentValue > cal.testValue {
		return false
	}
	nextValue := cal.numbers[0]
	addResult := currentValue + nextValue
	mulResult := currentValue * nextValue
	concatResult := concatNumbers(currentValue, nextValue)
	if len(cal.numbers) == 1 && (addResult == cal.testValue || mulResult == cal.testValue || concatResult == cal.testValue) {
		return true
	} else if len(cal.numbers) == 1 {
		return false
	} else {
		cal.numbers = cal.numbers[1:]
		return verifyCalibrationRec(deepCopy(cal), addResult) || verifyCalibrationRec(deepCopy(cal), mulResult) || verifyCalibrationRec(deepCopy(cal), concatResult)
	}
}
func verifyCalibration(cal Calibration, counter *atomic.Uint64) {
	copy := deepCopy(cal)
	currentValue := copy.numbers[0]
	copy.numbers = copy.numbers[1:]
	if verifyCalibrationRec(copy, currentValue) {
		counter.Add(uint64(cal.testValue))
	}
}

func main() {
	calibrations := readFile("input.txt")
	var wg sync.WaitGroup
	var counter atomic.Uint64
	for _, cal := range calibrations {
		wg.Add(1)
		go func() {
			verifyCalibration(cal, &counter)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Result:", counter.Load())

}
