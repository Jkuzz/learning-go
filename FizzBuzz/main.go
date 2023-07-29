package main

import (
	"fmt"
	"strconv"
)

type printConditional struct {
	output    string
	condition func(int) bool
}

func getPrintConditional() []printConditional {
	return []printConditional{
		{
			"Fizz", func(i int) bool {
				return i%3 == 0
			},
		},
		{
			"Buzz", func(i int) bool {
				return i%5 == 0
			},
		},
		{
			"Bam", func(i int) bool {
				return i%7 == 0
			},
		},
		{
			"Nice!", func(i int) bool {
				return i == 69
			},
		},
	}
}

func main() {
	for i := 0; i < 100; i++ {
		out := ""
		for _, conditional := range getPrintConditional() {
			if conditional.condition(i) {
				out += conditional.output
			}
		}
		if out == "" {
			out = strconv.Itoa(i)
		}
		fmt.Println(out)
	}
}
