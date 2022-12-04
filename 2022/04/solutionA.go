package main

import (
	"fmt"
	"strconv"
	"strings"
)

func SolutionA() {
	pairs, err := ReadLines("input.txt")
	if err != nil {
		panic("couldnt find input.txt")
	}
	dups := 0
	for _, pair := range pairs {
		if hasDuplicateRange(pair) {
			dups++
		}
	}
	fmt.Println("SolutionA:")
	fmt.Println(dups)
	fmt.Printf("%d = %d\n", dups, 582)
}

func hasDuplicateRange(pairString string) bool {
	ranges := strings.Split(pairString, ",")
	rangeA := strings.Split(ranges[0], "-")
	rangeB := strings.Split(ranges[1], "-")

	aLow, err := strconv.ParseInt(rangeA[0], 0, 0)
	if err != nil {
		panic("couldnt conver aLow to a int")
	}
	aHi, err := strconv.ParseInt(rangeA[1], 0, 0)
	if err != nil {
		panic("couldnt conver aHi to a int")
	}
	bLow, err := strconv.ParseInt(rangeB[0], 0, 0)
	if err != nil {
		panic("couldnt conver aLow to a int")
	}
	bHi, err := strconv.ParseInt(rangeB[1], 0, 0)
	if err != nil {
		panic("couldnt conver aHi to a int")
	}

	if (aLow <= bLow && aHi >= bHi) || bLow <= aLow && bHi >= aHi {
		return true
	}

	return false
}
