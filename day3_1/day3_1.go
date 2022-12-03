package main

import (
	"bufio"
	"fmt"
	"os"
)

func findCommonItem(comp1 []byte, comp2 []byte) byte {
	for _, item1 := range comp1 {
		for _, item2 := range comp2 {
			if item1 == item2 {
				return item1
			}
		}
	}
	return 0
}

func findItemPriority(item byte) int {
	if item >= 'a' && item <= 'z' {
		return int(item) - int('a') + 1
	} else if item >= 'A' && item <= 'Z' {
		return int(item) - int('A') + 27
	}
	return 0
}

func main() {
	file, err := os.Open("../inputs/day3.in")
	if err != nil {
		panic(err)
	}
	sc := bufio.NewScanner(file)

	result := 0
	for sc.Scan() {
		items := sc.Bytes()

		firstComp := items[:len(items)/2]
		secondComp := items[len(items)/2:]

		item := findCommonItem(firstComp, secondComp)
		result += findItemPriority(item)
	}
	fmt.Println(result)
}
