package main

import (
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type DataChunk struct {
	id       int //-1 = free
	size     int
	position int
}

type Disk struct {
	data      []DataChunk
	checksum  int
	blockview []int
}

func readFile(filename string) Disk {
	file, err := os.ReadFile(filename)
	check(err)
	var nextId int
	var disk []DataChunk
	for i, position := range file {
		var block DataChunk
		if i%2 == 0 {
			block.id = nextId
			nextId++
		} else {
			block.id = -1
		}
		block.size, err = strconv.Atoi(string(position))
		check(err)
		block.position = i
		disk = append(disk, block)
	}
	return Disk{data: disk}
}

func (d Disk) printData() {
	for _, block := range d.data {
		for i := 0; i < block.size; i++ {
			var symbol string
			if block.id == -1 {
				symbol = "."
			} else {
				symbol = strconv.Itoa(block.id)
			}

			fmt.Printf(symbol)
		}
	}
	fmt.Printf("\n")
}

func (d *Disk) generateBlockView() {
	for _, block := range d.data {
		for i := 0; i < block.size; i++ {
			d.blockview = append(d.blockview, block.id)
		}
	}
}

func (d *Disk) frag() {
	var i int
	j := len(d.blockview) - 1
	for i < j {
		switch {
		case d.blockview[i] != -1:
			i++
			continue
		case d.blockview[j] == -1:
			j--
			continue
		default:
			d.blockview[i], d.blockview[j] = d.blockview[j], d.blockview[i]
		}
	}
}

func (d *Disk) calculateChecksum() {
	for i := 0; d.blockview[i] != -1; i++ {
		d.checksum += i * d.blockview[i]
	}
}

func main() {
	d := readFile("input.txt")
	d.printData()
	d.generateBlockView()
	d.frag()
	d.calculateChecksum()
	fmt.Println("Checksum:", d.checksum)

}
