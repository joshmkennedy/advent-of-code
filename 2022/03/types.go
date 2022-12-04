package main

import "fmt"

var Letters = []string{
	"a",
	"b",
	"c",
	"d",
  "e",
	"f",
	"g",
	"h",
	"i",
	"j",
	"k",
	"l",
	"m",
  "n",
  "o",
  "p",
  "q",
	"r",
	"s",
	"t",
	"u",
	"v",
	"w",
	"x",
	"y",
	"z",
	"A",
	"B",
	"C",
	"D",
  "E",
	"F",
	"G",
	"H",
	"I",
	"J",
	"K",
	"L",
	"M",
  "N",
  "O",
  "P",
  "Q",
	"R",
	"S",
	"T",
	"U",
	"V",
	"W",
	"X",
	"Y",
	"Z",
}

func ItemPriority(letter string) int {
	for index, l := range Letters {
		if l == letter {
			return index + 1
		}
	}	
	fmt.Printf("invalid letter: %s\n", letter)
	panic("Encountered a non-valid letter");
}
