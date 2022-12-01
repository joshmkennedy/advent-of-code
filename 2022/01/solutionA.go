package main

import (
	"fmt"
)

func SolutionA(){
	fmt.Println("Solution A:")
	elvesTotals := TotalAllElfCals("input.txt")
	var biggest int32	= 0
	for _, cal := range elvesTotals {
		if cal > biggest{
			biggest = cal
		}
	}

	fmt.Printf("%+v\n", biggest)
	fmt.Printf("%+v\n", biggest==70374)
}
