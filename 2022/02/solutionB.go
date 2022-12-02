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
		game := NewGameB(inputs[0],inputs[1])
		game.Play()
		totalScore += game.GetScore()
	}
	fmt.Println("solution to B:")
	fmt.Println(totalScore)

}

func NewGameB(o,outcome string) *GameB{
	return &GameB{
		OChoice: o,
		NeedOutcome:outcome,
	}
}

type GameB struct{
	OChoice string
	MChoice string
	NeedOutcome string
	Outcome int  
	Score int
}

func (g *GameB) Play (){
	if g.OChoice == OPAPER 	{
		if g.NeedOutcome == OUTCOM_LOSS {
			g.MChoice = MROCK
		}
		if g.NeedOutcome == OUTCOM_DRAW {
			g.MChoice = MPAPER
		}
		if g.NeedOutcome == OUTCOM_WIN {
			g.MChoice = MSCISSORS
		}
	}

	if g.OChoice == OROCK 	{
		if g.NeedOutcome == OUTCOM_LOSS{
			g.MChoice = MSCISSORS
		}
		if g.NeedOutcome ==  OUTCOM_DRAW {
			g.MChoice = MROCK
		}
		if g.NeedOutcome == OUTCOM_WIN {
			g.MChoice = MPAPER
		}
	}

	if g.OChoice == OSCISSORS {
		if g.NeedOutcome == OUTCOM_LOSS{
			g.MChoice = MPAPER
		}
		if g.NeedOutcome == OUTCOM_DRAW {
			g.MChoice = MSCISSORS
		}
		if g.NeedOutcome == OUTCOM_WIN {
			g.MChoice = MROCK
		}
	}

	if g.MChoice == ""  {
		panic("Never recieved a choice")
	}

	if g.NeedOutcome == OUTCOM_LOSS {
			g.Outcome = LOSS
	}
	if g.NeedOutcome == OUTCOM_DRAW {
		g.Outcome = DRAW
	}
	if g.NeedOutcome == OUTCOM_WIN {
		g.Outcome = WIN
	}
}

func (g *GameB) GetScore() int{
	g.Score = g.Outcome	
	g.Score += g.GetChoiceScore()
	return g.Score
}

func (g *GameB) GetChoiceScore() int{
	choice := g.MChoice
	for _,c := range ChoicesB {
		if c.Kind == choice {
			return c.Value
		}	
	}
	panic("choice not found")
}

type Choice_B struct{
	Kind string;
	Value int;	
}

var Paper_B Choice_B = Choice_B{
	Kind:MPAPER,
	Value:2,
}
var Rock_B Choice_B = Choice_B{
	Kind:MROCK,
	Value:1,
}
var Scissors_B Choice_B = Choice_B{
	Kind:MSCISSORS,
	Value:3,
}

var ChoicesB = []Choice_B{Scissors_B, Rock_B, Paper_B}

