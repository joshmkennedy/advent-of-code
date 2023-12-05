package main

import (
	"bufio"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

const (
	OROCK     = "A"
	OPAPER    = "B"
	OSCISSORS = "C"

	MROCK     = "X"
	MPAPER    = "Y"
	MSCISSORS = "Z"

	OUTCOME_LOSS = "X"
	OUTCOME_DRAW = "Y"
	OUTCOME_WIN  = "Z"

	LOSS = 0
	DRAW = 3
	WIN  = 6
)

func NewGame(o string, m, outcome *string) *Game {
	if m != nil {
		return &Game{
			OChoice: o,
			MChoice: *m,
		}
	}
	return &Game{
		OChoice:     o,
		NeedOutcome: *outcome,
	}
}

type Game struct {
	OChoice     string
	MChoice     string
	NeedOutcome string
	Outcome     int
	Score       int
}

func (g *Game) Play() {
	if g.MChoice == "" {
		g.FindChoice()
	}
	g.FindOutcome()
}

func (g *Game) FindChoice() {
	if g.OChoice == OPAPER {
		if g.NeedOutcome == OUTCOME_LOSS {
			g.MChoice = MROCK
		}
		if g.NeedOutcome == OUTCOME_DRAW {
			g.MChoice = MPAPER
		}
		if g.NeedOutcome == OUTCOME_WIN {
			g.MChoice = MSCISSORS
		}
	}

	if g.OChoice == OROCK {
		if g.NeedOutcome == OUTCOME_LOSS {
			g.MChoice = MSCISSORS
		}
		if g.NeedOutcome == OUTCOME_DRAW {
			g.MChoice = MROCK
		}
		if g.NeedOutcome == OUTCOME_WIN {
			g.MChoice = MPAPER
		}
	}

	if g.OChoice == OSCISSORS {
		if g.NeedOutcome == OUTCOME_LOSS {
			g.MChoice = MPAPER
		}
		if g.NeedOutcome == OUTCOME_DRAW {
			g.MChoice = MSCISSORS
		}
		if g.NeedOutcome == OUTCOME_WIN {
			g.MChoice = MROCK
		}
	}
}

func (g *Game) FindOutcome() {
	if g.OChoice == OPAPER {
		if g.MChoice == MPAPER {
			g.Outcome = DRAW
		}
		if g.MChoice == MSCISSORS {
			g.Outcome = WIN
		}
		if g.MChoice == MROCK {
			g.Outcome = LOSS
		}
	}

	if g.OChoice == OROCK {
		if g.MChoice == MPAPER {
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
		if g.MChoice == MPAPER {
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

func (g *Game) GetScore() int {
	g.Score = g.Outcome
	g.Score += g.GetChoiceScore()
	return g.Score
}

func (g *Game) GetChoiceScore() int {
	choice := g.MChoice
	for _, c := range Choices {
		if c.Kind == choice {
			return c.Value
		}
	}
	panic("choice not found")
}

type Choice struct {
	Kind  string
	Value int
}

var Paper Choice = Choice{
	Kind:  MPAPER,
	Value: 2,
}
var Rock Choice = Choice{
	Kind:  MROCK,
	Value: 1,
}
var Scissors Choice = Choice{
	Kind:  MSCISSORS,
	Value: 3,
}

var Choices = []Choice{Scissors, Rock, Paper}
