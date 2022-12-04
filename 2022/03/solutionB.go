package main

import "fmt"

func SolutionB(){
	sacks, err := ReadLines("input.txt")
	if err != nil {
		panic("cant read input.txt")
	}

	total := 0
	var group []string
	for _, sack := range sacks {
		group = append(group, sack)
		if len(group) >= 3 {
			total += ItemPriority(elfBadge(group))
			group = []string{}
		}
	}
	fmt.Println(total)
}

func elfBadge(sacks []string) string {
	eA := sacks[0]
	eB := sacks[1]
	eC := sacks[2]
	
	for _,a := range eA {
		for _,b:= range eB {
			if a == b{
				for _, c := range eC {
					if c == a {
						return string(a)
					}
				}
			}
		}
	}
	fmt.Printf("%s\n%s\n%s\n",eA,eB,eC)
	panic("couldnt find a common letter between the 3 sacks")
}
