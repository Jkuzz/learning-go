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

type CleaningRange struct {
	start int
	end   int
}

func rangeContains(range1 CleaningRange, range2 CleaningRange) bool {
	if range1.start >= range2.start && range1.end <= range2.end {
		return true
	}
	if range1.start <= range2.start && range1.end >= range2.end {
		return true
	}
	return false
}

func rangeOverlaps(range1 CleaningRange, range2 CleaningRange) bool {
	if range1.start <= range2.start && range1.end >= range2.start {
		return true
	}
	if range1.start <= range2.end && range1.end >= range2.end {
		return true
	}
	return rangeContains(range1, range2)
}

func parseLineToRanges(line string) (CleaningRange, CleaningRange) {
	splitRanges := strings.Split(line, ",")
	r1, r2 := splitRanges[0], splitRanges[1]
	r1Start, r1StartErr := strconv.Atoi(strings.Split(r1, "-")[0])
	r1End, r1EndErr := strconv.Atoi(strings.Split(r1, "-")[1])
	r2Start, r2StartErr := strconv.Atoi(strings.Split(r2, "-")[0])
	r2End, r2EndErr := strconv.Atoi(strings.Split(r2, "-")[1])

	if r1StartErr != nil || r1EndErr != nil || r2StartErr != nil || r2EndErr != nil {
		panic("Parsing error")
	}
	range1 := CleaningRange{r1Start, r1End}
	range2 := CleaningRange{r2Start, r2End}
	return range1, range2	
}

func main() {
	readFile, err := os.Open("data.txt")
	check(err)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	overlappingRanges := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		range1, range2 := parseLineToRanges(line)
		// if rangeContains(range1, range2) {
		if rangeOverlaps(range1, range2) {
			overlappingRanges += 1
		}
		// fmt.Println(range1, range2, rangeOverlaps(range1, range2))
	}

	readFile.Close()
	fmt.Println("Total overlapping ranges: ", overlappingRanges)
}
