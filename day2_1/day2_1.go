package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var MOVES map[string]int = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
	"X": 1,
	"Y": 2,
	"Z": 3,
}

func main() {
	file, err := os.Open("../inputs/day2.in")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	totalScore := 0
	for sc.Scan() {
		moves := strings.Split(sc.Text(), " ")
		firstPlayer := MOVES[moves[0]]
		secondPlayer := MOVES[moves[1]]
		totalScore += secondPlayer
		if firstPlayer == secondPlayer {
			totalScore += 3
		} else if (firstPlayer == 1 && secondPlayer == 2) ||
			(firstPlayer == 2 && secondPlayer == 3) ||
			(firstPlayer == 3 && secondPlayer == 1) {
			totalScore += 6
		}
	}

	fmt.Println(totalScore)
}
