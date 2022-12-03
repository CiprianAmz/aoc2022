package main

import (
	"bufio"
	"fmt"
	"os"
)

func findCommonItem(comp1 []byte, comp2 []byte, comp3 []byte) byte {
	for _, item1 := range comp1 {
		for _, item2 := range comp2 {
			for _, item3 := range comp3 {
				if item1 == item2 && item1 == item3 {
					return item1
				}
			}
		}
	}
	panic("No common item found.")
}

func findItemPriority(item byte) int {
	if item >= 'a' && item <= 'z' {
		return int(item) - int('a') + 1
	} else if item >= 'A' && item <= 'Z' {
		return int(item) - int('A') + 27
	}
	panic("Invalid character.")
}

func main() {
	file, err := os.Open("../inputs/day3.in")
	if err != nil {
		panic(err)
	}
	sc := bufio.NewScanner(file)

	result := 0

	for sc.Scan() {
		bp1 := sc.Text()
		sc.Scan()
		bp2 := sc.Text()
		sc.Scan()
		bp3 := sc.Text()

		item := findCommonItem([]byte(bp1), []byte(bp2), []byte(bp3))
		result += findItemPriority(item)
	}
	fmt.Println(result)
}
