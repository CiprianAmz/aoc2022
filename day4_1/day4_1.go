package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../inputs/day4.in")
	if err != nil {
		panic(err)
	}
	sc := bufio.NewScanner(file)

	result := 0

	for sc.Scan() {
		line := sc.Text()
		sections := strings.Split(line, ",")

		sec1 := strings.Split(sections[0], "-")
		sec2 := strings.Split(sections[1], "-")

		sec1_1, _ := strconv.Atoi(sec1[0])
		sec1_2, _ := strconv.Atoi(sec1[1])
		sec2_1, _ := strconv.Atoi(sec2[0])
		sec2_2, _ := strconv.Atoi(sec2[1])

		if (sec1_1 <= sec2_1 && sec1_2 >= sec2_2) ||
			(sec2_1 <= sec1_1 && sec2_2 >= sec1_2) {
			result += 1
		}
	}

	fmt.Println(result)
}
