package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func check_err(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	file_content, err := ioutil.ReadFile("../inputs/day1.in")
	check_err(err)
	current_cal := int(0)
	current_total_cal := int(0)
	top_3 := make([]int, 3)

	for i := 0; i < len(file_content); i++ {
		if file_content[i] == '\n' {
			current_total_cal += current_cal

			if i == (len(file_content)-1) || file_content[i+1] == '\n' {
				i++

				comp_val := current_total_cal
				for i := 0; i < len(top_3); i++ {
					if comp_val > top_3[i] {
						temp := top_3[i]
						top_3[i] = comp_val
						comp_val = temp
					}
				}

				current_total_cal = 0
			}

			current_cal = 0
		} else {
			digit, _ := strconv.Atoi(string(file_content[i]))
			current_cal = current_cal*10 + digit
		}
	}

	max_total_cal := 0
	for i := 0; i < len(top_3); i++ {
		fmt.Println(top_3[i])
		max_total_cal += top_3[i]
	}

	fmt.Println(max_total_cal)
}
