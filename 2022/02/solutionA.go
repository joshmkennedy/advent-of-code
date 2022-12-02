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
		game := NewGameA(inputs[0],inputs[1])
		game.Play()
		totalScore += game.GetScore()
	}
	fmt.Println("solution to A:")
	fmt.Println(totalScore)
}

func NewGameA (o,m string)*GameA{
	return &GameA{
		OChoice:o,
		MChoice:m,
	}
}

type GameA struct {
	OChoice string
	MChoice string
	Outcome int  
	Score int
}

func (g *GameA) Play (){
	if g.OChoice == OPAPER 	{
		if g.MChoice == MPAPER{
			g.Outcome = DRAW
		}
		if g.MChoice == MSCISSORS {
			g.Outcome = WIN
		}
		if g.MChoice == MROCK {
			g.Outcome = LOSS
		}
	}

	if g.OChoice == OROCK 	{
		if g.MChoice == MPAPER{
			g.Outcome = WIN
		}
		if g.MChoice == MSCISSORS {
			g.Outcome = LOSS
		}
		if g.MChoice == MROCK {
			g.Outcome = DRAW
		}
	}

	if g.OChoice == OSCISSORS {
		if g.MChoice == MPAPER{
			g.Outcome = LOSS
		}
		if g.MChoice == MSCISSORS {
			g.Outcome = DRAW
		}
		if g.MChoice == MROCK {
			g.Outcome = WIN
		}
	}
}

func (g *GameA) GetScore() int{
	g.Score = g.Outcome	
	g.Score += g.GetChoiceScore()
	return g.Score
}

func (g *GameA) GetChoiceScore() int{
	choice := g.MChoice
	for _,c := range Choices {
		if c.Kind == choice {
			return c.Value
		}	
	}
	panic("choice not found")
}

type Choice_A struct{
	Kind string;
	Value int;	
}

var Paper Choice_A = Choice_A{
	Kind:MPAPER,
	Value:2,
}
var Rock Choice_A = Choice_A{
	Kind:MROCK,
	Value:1,
}
var Scissors Choice_A = Choice_A{
	Kind:MSCISSORS,
	Value:3,
}

var Choices = []Choice_A{Scissors, Rock, Paper}

