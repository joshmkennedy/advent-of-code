package main

import (
	"bufio"
	"os"
)

func ReadLines (path string) ([]string, error){
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		lines = append(lines, scanner.Text())	
	}
	return lines, scanner.Err()
}

const (
	OROCK = "A"
	OPAPER = "B"
	OSCISSORS = "C"

	MROCK = "X"
	MPAPER = "Y"
	MSCISSORS = "Z"

	LOSS = 0
	DRAW = 3
	WIN = 6

	OUTCOM_LOSS = "X"
	OUTCOM_DRAW = "Y"
	OUTCOM_WIN = "Z"
)



