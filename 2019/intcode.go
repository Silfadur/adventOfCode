package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type IntCode struct {
	ints   []int
	pc     int
	input  chan int
	output chan int
}

func (ic *IntCode) init() {
	ic.ints = make([]int, 0)
}

func (ic *IntCode) readFromFile(filename string) {
	file, err := os.ReadFile(filename)
	check(err)
	tokens := strings.Split(string(file), ",")
	for _, token := range tokens {
		value, _ := strconv.Atoi(token)
		ic.ints = append(ic.ints, value)
	}
}

func (ic *IntCode) run() {
	for {
		if ic.pc >= len(ic.ints) {
			return
		}
		switch ic.ints[ic.pc] {
		case 1: //add
			parameter1 := ic.ints[ic.pc+1]
			parameter2 := ic.ints[ic.pc+2]
			parameter3 := ic.ints[ic.pc+3]
			ic.ints[parameter3] = ic.ints[parameter1] + ic.ints[parameter2]
			//fmt.Printf("Adding [%d]=%d and [%d]=%d, saving %d to [%d]\n", parameter1, ic.ints[parameter1], parameter2, ic.ints[parameter2], ic.ints[parameter3], parameter3)
			ic.pc += 4
		case 2: //mul
			parameter1 := ic.ints[ic.pc+1]
			parameter2 := ic.ints[ic.pc+2]
			parameter3 := ic.ints[ic.pc+3]
			ic.ints[parameter3] = ic.ints[parameter1] * ic.ints[parameter2]
			//fmt.Printf("Multiplying [%d]=%d and [%d]=%d, saving %d to [%d]\n", parameter1, ic.ints[parameter1], parameter2, ic.ints[parameter2], ic.ints[parameter3], parameter3)
			ic.pc += 4
		case 3: //input
			parameter1 := ic.ints[ic.pc+1]
			ic.ints[parameter1] = <-ic.input
			ic.pc += 2
		case 4: //output
			parameter1 := ic.ints[ic.pc+1]
			ic.output <- ic.ints[parameter1]
			ic.pc += 2
		case 99:
			//fmt.Println("Halting")
			return
		default:
			return
		}

	}
}

const maxRange = 99

func main() {
	for noun := 0; noun <= maxRange; noun++ {
		for verb := 0; verb <= maxRange; verb++ {
			fmt.Printf("Trying noun=%d, verb=%d\n", noun, verb)
			var code IntCode
			code.readFromFile("input_day2.txt")
			code.ints[1] = noun
			code.ints[2] = verb
			code.run()
			if code.ints[0] == 19690720 {
				fmt.Printf("Found output 19690720, noun=%d, verb=%d, result=%d\n", noun, verb, 100*noun+verb)
				return
			}
		}
	}

}
