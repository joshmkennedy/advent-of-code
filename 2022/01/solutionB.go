package main

import (
	"fmt"
	"sort"
)


type ByCal []int32

func (c ByCal) Len() int {
	return len(c);
}

func (c ByCal) Less(i,j int) bool {
	return c[i] > c[j]
}

func (c ByCal) Swap(i, j int){
	a := c[i]
	b := c[j]
	c[i] = b
	c[j] = a
}


func SolutionB(){
	fmt.Println("Solution B:")
	elvesTotals:= TotalAllElfCals("input.txt")
	sort.Sort(ByCal(elvesTotals))

	top3Total := elvesTotals[0] + elvesTotals[1] + elvesTotals[2]
	fmt.Printf("%+v\n",top3Total)
	fmt.Printf("%+v\n",top3Total ==204610)
	fmt.Println("")
}

