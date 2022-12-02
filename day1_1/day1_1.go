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
	max_total_cal := 0
	current_cal := int(0)
	current_total_cal := int(0)

	for i := 0; i < len(file_content); i++ {
		if file_content[i] == '\n' {
			current_total_cal += current_cal

			if i == (len(file_content)-1) || file_content[i+1] == '\n' {
				i++
				if current_total_cal > max_total_cal {
					max_total_cal = current_total_cal
				}

				current_total_cal = 0
			}

			current_cal = 0
		} else {
			digit, _ := strconv.Atoi(string(file_content[i]))
			current_cal = current_cal*10 + digit
		}
	}

	fmt.Println(max_total_cal)
}
