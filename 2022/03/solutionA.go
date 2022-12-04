package main

import "fmt"

func SolutionA(){
	ruckSacks, err := ReadLines("input.txt")
	if err != nil {
		panic("cant find input.txt")
	}
	var total = 0
	for _,sack := range ruckSacks {
		total += ItemPriority(missplacedItem(sack))
	}

	fmt.Println(total)
}

func missplacedItem(sack string) string{
	itemCount := len(sack)
	compA := sack[:(itemCount/2)]
	compB := sack[(itemCount/2):]
	for _, a := range compA {
		for _, b:= range compB {
			if a == b {
				return string(a)
			}
		}
	}
	fmt.Printf("%s | %s\n", compA, compB)
	panic("didnt come across a duplicate letter");
}
