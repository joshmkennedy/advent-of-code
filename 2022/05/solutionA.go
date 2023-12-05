package main

import (
	"fmt"
)

func SolutionA(){
	columns := getColumns()
	lines, err := ReadLines("input.txt")
	if err != nil {
		panic("cant read input.txt")
	}

	for _,line := range lines {
		count, from, to := extractArgsFromLine(line)
		columns.MoveCrates9000(count,from,to)
	}
	top:= columns.GetTopCrates()	
	var answer string
	for i:=1; i<=len(top); i++{
		answer = fmt.Sprint(answer,top[i])
	}
	fmt.Printf("LJSVLTWQM = %s | solved %t\n", answer, answer == "LJSVLTWQM")
}

//var testColInput = map[int][]string{
// 	1:{"T", "V"},
// 	2:{"V", "C", "A" },
// 	3:{"P"},
// }
//
//
// func testGetCols() Columns {
// 	var columns = make(Columns)
// 	for id,vals := range testColInput{
// 		var col Column
// 		// for i:= len(vals) - 1; i>=0; i-- {
// 		for i:= 0; i<len(vals); i++{
// 			col.Push(&Crate{Value:vals[i]})
// 		}
// 		columns[id] = &col
// 	}
// 	return columns
// }
//
// func _SolutionA(){
// 	columns := testGetCols()
// 	lines:= []string{"move 1 from 2 to 1","move 3 from 1 to 3", "move 2 from 2 to 1", "move 1 from 1 to 2"}
//
// 	for _,line := range lines {
// 		count, from, to := extractArgsFromLine(line)
// 		columns.MoveCrates(count,from,to)
// 	}
// 	top:= columns.GetTopCrates()	
//
// 	var answer string
// 	for i:=1; i<len(top); i++{
// 		answer = fmt.Sprint(answer,top[i])
// 	}
// 	fmt.Printf("LJSVLTWQM = %s", answer)
// }
