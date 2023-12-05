package main


var colInput = map[int][]string{
	1:{"T", "V", "J", "W", "N", "R", "M", "S" },
	2:{"V", "C", "P", "Q", "J", "D", "W", "B"},
	3:{"P", "R", "D", "H", "F", "J", "B"},
	4:{ "D", "N", "M", "B", "P", "R", "F"},
	5:{"B", "T", "P", "R", "V", "H"},
	6:{"T", "P", "B", "C"},
	7:{"L", "P", "R", "J", "B"},
	8:{"W", "B", "Z", "T", "L", "S", "C", "N"},
	9:{"G", "S", "L"},
}    
     
func getColumns() Columns {
	var columns = make(Columns)
	for id,vals := range colInput{
		var col Column
		for i:= len(vals) - 1; i>=0; i-- {
			col.Push(&Crate{Value:vals[i]})
		}
		columns[id] = &col
	}
	return columns
}    

