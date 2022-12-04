package main

import (
	"fmt"
	"strconv"
	"strings"
)

func SolutionB() {
	pairs, err := ReadLines("input.txt")
	if err != nil {
		panic("couldnt find input.txt")
	}
	dups := 0
	for _, pair := range pairs {
		if hasDuplicate(pair) {
			dups++
		}
	}
	fmt.Println("SolutionB:")
	fmt.Println(dups)
	fmt.Printf("%d = %d\n", dups, 893)
}

func hasDuplicate(pairString string) bool {
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

	if aLow <= bLow && aHi >= bLow {
		return true
	}
	if bLow <= aLow && bHi >= aLow {
		return true
	}

	return false
}

/*
	2,6 3,7
	if 2 < 3 && 2
*/
