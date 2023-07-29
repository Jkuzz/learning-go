package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func sum(arr []int) int {
	sum := 0
	for _, valueInt := range arr {
		sum += valueInt
	}
	return sum
}

func main() {
	readFile, err := os.Open("data.txt")
	check(err)
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	maxCalories := []int{0, 0, 0}
	currentCalories := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" { // Next elf!
			// If this elf's calories are higher than the least of the top 3
			if currentCalories > maxCalories[0] {
				maxCalories[0] = currentCalories
				sort.Ints(maxCalories) // I'm lazy
			}
			currentCalories = 0
		} else {
			parsedCals, convErr := strconv.Atoi(line)
			check(convErr)
			currentCalories += parsedCals
		}
	}

	readFile.Close()
	fmt.Println("Top 3 most calories carried: " + strconv.Itoa(sum(maxCalories)))
}
