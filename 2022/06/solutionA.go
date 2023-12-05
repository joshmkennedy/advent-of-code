package main

import (
	"fmt"
	"io/ioutil"
)

func SolutionA(){
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic("couldnt read file input.txt")
	}

	processed := 0
	var group []string 
	for _,b := range input {
		letter := string(b)
		group = append(group,letter)
		processed++
		if len(group) >= 4 {
			if unique(group) {
				break
			}
			group=group[1:]
		}
	}
	fmt.Println("processed:",processed)
}

func unique(g []string)bool{
	counts := make(map[string]int)	
	for _,l:= range g{
		counts[l]++
	}
	
	for _,c := range counts {
			if c > 1 {	
				return false
			}
	}
	return true
}
