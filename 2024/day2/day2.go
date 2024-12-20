package main

import (
	"bufio"
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

func readfile(filename string) [][]int {
	dat, err := os.Open(filename)
	check(err)
	defer dat.Close()
	scanner := bufio.NewScanner(dat)
	var reports [][]int
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		var report []int
		for _, field := range line {
			number, err := strconv.Atoi(field)
			check(err)
			report = append(report, number)
		}
		reports = append(reports, report)
	}

	return reports
}

func checkReport(report []int) bool {
	var increased int
	var decreased int
	for i := 0; i < (len(report) - 1); i++ {
		ss := report[i : i+2]
		distance := ss[0] - ss[1]
		if distance == 0 {
			return false
		}
		if distance < 0 {
			decreased++
			distance = distance * -1
		} else {
			increased++
		}
		if distance > 3 {
			return false
		}
	}
	if (increased != 0) && (decreased != 0) {
		return false
	}
	return true
}

func checkReportCompensated(report []int) bool {
	if checkReport(report) {
		return true
	}
	for i := 0; i < len(report); i++ {
		reportCopy := make([]int, len(report))
		_ = copy(reportCopy, report)
		compensatedReport := append(reportCopy[:i], reportCopy[i+1:]...)

		if checkReport(compensatedReport) {
			return true
		}
	}
	return false
}

func main() {
	reports := readfile("input.txt")
	var safeReports int
	for _, report := range reports {
		if checkReport(report) {
			safeReports++
		}
	}
	var safeReportsCompensated int
	for _, report := range reports {
		if checkReportCompensated(report) {
			safeReportsCompensated++
		}
	}
	fmt.Println("Part 1, Safe reports: ", safeReports)
	fmt.Println("Part 2, Safe reports compensated: ", safeReportsCompensated)
}
