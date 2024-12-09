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
	id   int //-1 = free
	size int
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

func (d *Disk) findEarliestFreeSpace(size int) (bool, int) {
	for i, chunk := range d.data {
		if chunk.id != -1 {
			continue
		}
		if chunk.size >= size {
			return true, i
		}
	}
	return false, 0
}

func (d *Disk) defrag() {
	j := len(d.data) - 1
	for j != 0 {
		switch {
		case d.data[j].id == -1:
			j--
			continue
		default:
			success, i := d.findEarliestFreeSpace(d.data[j].size)
			if success && i < j {
				d.data[i], d.data[j] = d.data[j], d.data[i] // chunks wechseln
				remainingSpace := d.data[j].size - d.data[i].size
				if remainingSpace > 0 {
					d.data[j].size -= remainingSpace                                                                       // leeren chunk resizen
					d.data = append(d.data[:i+1], append([]DataChunk{{id: -1, size: remainingSpace}}, d.data[i+1:]...)...) // leeren chunk mit rest einfügen
					j++                                                                                                    // j wieder auf den gleichen chunk wie vorher zeigen lassen
				}
			}
			j--
		}
	}
}

func (d *Disk) calculateChecksum() {
	for i := 0; i < len(d.blockview); i++ {
		if d.blockview[i] != -1 {
			d.checksum += i * d.blockview[i]
		}
	}
}

func main() {
	d := readFile("input.txt")
	d.printData()
	d.defrag()
	d.generateBlockView()
	d.calculateChecksum()
	fmt.Println(d.blockview)
	fmt.Println("Checksum:", d.checksum)

}
