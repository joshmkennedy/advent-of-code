package main

import (
	"fmt"
	"strconv"
	"strings"
)

//import "fmt"

type FileTree struct{
	Files []*FSChild
	Location *Dir
	Home *Dir
}

const (
	HOME =	"/"
	PARENT = ".."
	MAX_FILE_SIZE = 100000
)

func ProcessStdOut(lines []string) (*FileTree, error) {
	tree := &FileTree{}
	tree.Home = setHome(lines[0])
	tree.Location = tree.Home

	for i,line := range lines[1:] {
		fmt.Printf("line %d: CMD %s\n",i,line)
		tree.ProcessLine(line)
	}
	return tree,nil
}


func setHome(line string) *Dir{
	words := strings.Split(line," ")
	name := words[2]
	dir := NewDir(name, &Dir{})
	return dir
}

func (t *FileTree) ProcessLine(line string){
	words := strings.Split(line," ")
	//Process CMD
	if words[0] == "$" {
		if words[1] == "cd" {
			t.ChangeDir(words[2])
			return
		}
		if words[1] == "ls" {
			return
		}
	}

	if words[0] == "dir" {
		name := words[1]
		dir := NewDir(name, t.Location)
		fmt.Printf("%+v\n",dir)

		t.Location.AddDir(dir)
		return
	}
	size,err := strconv.ParseInt(words[0], 0, 0)
	if err != nil {
		panic(err)
	}
	file:= NewFile(words[1], uint64(size), t.Location)
	t.Location.AddFile(file)
	t.Location.IncSize(uint64(size))
}

func (t *FileTree) ChangeDir(dir string) {
	if dir == HOME {
		t.Location = t.Home
		return
	}

	if dir == PARENT {
		t.Location = t.Location.GetParent()
		return
	}
	
	t.Location = t.Location.FindDir(dir)

	
}

func (t *FileTree) TotalUpDirectories() {
	t.Home.IncSizeFromSubDirs()
	fmt.Println("total fs size = ", t.Home.GetSize())
}






