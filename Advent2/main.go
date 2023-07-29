package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type MovePoints struct {
	resultPoints int
	movePoints   map[string]int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	readFile, err := os.Open("data.txt")
	check(err)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	pointInfo := map[string]MovePoints{
		"X": {0, map[string]int{"A": 3, "B": 1, "C": 2}},
		"Y": {3, map[string]int{"A": 1, "B": 2, "C": 3}},
		"Z": {6, map[string]int{"A": 2, "B": 3, "C": 1}},
	}

	totalPoints := 0

	for fileScanner.Scan() {
		line := strings.Split(fileScanner.Text(), " ")
		elfMove, myMove := line[0], line[1]
		totalPoints += pointInfo[myMove].resultPoints + pointInfo[myMove].movePoints[elfMove]
	}

	readFile.Close()
	fmt.Println("Total points earned: ", totalPoints)
}
