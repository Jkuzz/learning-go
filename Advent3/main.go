package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func findCommonItem(firstCompartment string, secondCompartment string) string {
	var common strings.Builder
	common.Grow(len(firstCompartment))
	for _, char := range firstCompartment {
		if strings.Contains(secondCompartment, string(char)) {
			common.WriteRune(char)
		}
	}
	return common.String()
}

func getItemPriority(item string) int {
	val := int(item[0])
	if val > 96 {
		return val - 96
	}
	return val - 38
}

func main() {
	readFile, err := os.Open("data.txt")
	check(err)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	totalPriority := 0
	for fileScanner.Scan() {
		line1 := fileScanner.Text()
		fileScanner.Scan()
		line2 := fileScanner.Text()
		fileScanner.Scan()
		line3 := fileScanner.Text()

		firstCommonItems := findCommonItem(line1, line2)
		securityBadgeItem := findCommonItem(firstCommonItems, line3)
		badgePriority := getItemPriority(securityBadgeItem)
		fmt.Println(string(securityBadgeItem), badgePriority)
		totalPriority += badgePriority
	}

	readFile.Close()
	fmt.Println("Total items priority: ", totalPriority)
}
