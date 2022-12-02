package main

import (
	"fmt"
	"strings"
)

func SolutionA(){
	data, err := ReadLines("input.txt")
	if err != nil{
		panic("cant find input.txt")
	}
	totalScore := 0
	for _, line := range data{
		inputs := strings.Split(line," ")
		game := NewGame(inputs[0],&inputs[1], nil)
		game.FindOutcome()
		totalScore += game.GetScore()
	}
	fmt.Println("solution to A:")
	fmt.Println(totalScore)
}

