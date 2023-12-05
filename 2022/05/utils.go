package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
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

func extractArgsFromLine(line string) (int, int, int){
	l := strings.Split(line," ")
	count := parseInt(l[1])
	from := parseInt(l[3])
	to := parseInt(l[5])
	return count, from, to
}

func parseInt(num string) int {
	i, err := strconv.ParseInt(num, 0, 0)
	if err != nil {
		panic("at the disco")
	}
	return int(i)
}

