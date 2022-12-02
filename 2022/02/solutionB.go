package main

import (
	"fmt"
	"strings"
)

func SolutionB(){
	data, err := ReadLines("input.txt")
	if err != nil{
		panic("cant find input.txt")
	}
	totalScore := 0
	for _, line := range data{
		inputs := strings.Split(line," ")
		game := NewGame(inputs[0],nil,&inputs[1])
		game.Play()
		totalScore += game.GetScore()
	}
	fmt.Println("solution to B:")
	fmt.Println(totalScore)

}
