package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
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

func TotalAllElfCals(path string) []int32 {
	data, err := ReadLines(path)
	if err != nil{
		log.Fatalln("oh no cant find the fileeeeeeeee")
		log.Fatalln(err)
	}
	var elves []int32
	elves	= append(elves, 0)
	cei := 0	//current elf index

	for _,line := range data{
		if line == "" {
			cei++
			var newElf int32 = 0
			elves = append(elves, newElf)
		}else{
			i, err := strconv.ParseInt(line, 10, 32)
			if err!=nil{
				panic(err)
			}
			cals :=int32(i)
			elves[cei] = elves[cei] + cals
		}
	}

	return elves
}
