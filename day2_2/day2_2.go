package main

import (
	"bufio"
	"fmt"
	"os"
)

var TURN_POINTS map[string]int = map[string]int{
	"A X": 3, // loss + scissors
	"A Y": 4, // draw + rock
	"A Z": 8, // win + paper
	"B X": 1, // loss + rock
	"B Y": 5, // draw + paper
	"B Z": 9, // win + scissors
	"C X": 2, // loss + paper
	"C Y": 6, // draw + scissors
	"C Z": 7, // win + rock
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
		totalScore += TURN_POINTS[sc.Text()]
	}

	fmt.Println(totalScore)
}
