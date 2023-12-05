package main

import (
	"fmt"
)

func SolutionB(){
	columns := getColumns()
	lines, err := ReadLines("input.txt")
	if err != nil {
		panic("cant read input.txt")
	}

	for _,line := range lines {
		count, from, to := extractArgsFromLine(line)
		columns.MoveCrates9001(count,from,to)
	}
	top:= columns.GetTopCrates()	
	var answer string
	for i:=1; i<=len(top); i++{
		answer = fmt.Sprint(answer,top[i])
	}
	fmt.Printf( "???? = %s | solved %t\n", answer, answer == "?????")
}

