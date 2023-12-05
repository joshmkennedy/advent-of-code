package main

import "fmt"


type Crate struct {
	Value string
}

type Column struct {
	Crates []*Crate
}

type Columns map[int]*Column

func (cs *Columns) GetTopCrates() map[int]string {
	var vals = make(map[int]string)
	for id,col := range *cs {
		crate,success := col.Peek()
		if success{
			vals[id] =  crate.Value
		} else{
			vals[id] = " "
		}
	}
	return vals
}

func (cs Columns) MoveCrates9000(count int, pullColId, pushColId int){
	pushCol := cs[pushColId]
	for i:=0; i<count; i++ {
		crate, success := cs[pullColId].Pop()
		if !success {
			panic("cant move crates because col is empty")
		}
		pushCol.Push(crate)
	}
}

func (cs Columns) MoveCrates9001(count int, pullColId, pushColId int){
	pushCol := cs[pushColId]
	pullCol := cs[pullColId]
	fmt.Println(count,"from", pullColId,len(pullCol.Crates), "to", pushColId)
	crates:= cs[pullColId].Splice(count)

	// for i:= len(crates ) - 1; i >= 0; i-- {
	for i:= 0; i < len(crates); i++ {
		crate := crates[i]
		pushCol.Push(crate)
	}
}


func (c *Column) IsEmpty()bool{
	return len(c.Crates) == 0
}

func (c *Column) Peek() (Crate,bool) {
	if c.IsEmpty(){
		return Crate{}, false	
	}	
	return *c.Crates[len(c.Crates) - 1], true
}

func (c *Column) Pop() (*Crate, bool){ 
	if c.IsEmpty(){
		return &Crate{},false
	}
	index := len(c.Crates) - 1
	element := c.Crates[index ]
	
	c.Crates = c.Crates[:index ]
	return element, true
}

func (c *Column) Splice(count int) []*Crate {
	index := len(c.Crates) - count
	elements := c.Crates[index:]
	
	c.Crates = c.Crates[:index ]
	return elements
}

func (c *Column) Push(crate *Crate){
	c.Crates = append(c.Crates, crate)
}

func newCrate(val string) *Crate{
	return &Crate{
		Value:val,
	}
}
