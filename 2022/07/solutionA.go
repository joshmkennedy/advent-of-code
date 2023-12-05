package main

func SolutionA(){
	lines, err := ReadLines("input.txt")
	if err != nil {
		panic("Cant find file input.txt")
	}

	tree,err := ProcessStdOut(lines)
	if err != nil {
		panic(err)
	}
	tree.TotalUpDirectories()
	// fmt.Printf("%+v\n",tree)
}
