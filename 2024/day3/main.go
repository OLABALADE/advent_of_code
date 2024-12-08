package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var result int

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	reg := regexp.MustCompile(`mul\((\d{1,3}),\s*(\d{1,3})\)|do\(\)|don't\(\)`)
	matches := reg.FindAllStringSubmatch(string(f), -1)

	enabled := true
	for _, match := range matches {
		if match[0] == "do()" {
			enabled = true
			continue
		} else if match[0] == "don't()" {
			enabled = false
			continue
		} else if !enabled {
			continue
		}

		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		result += num1 * num2

	}

	fmt.Println(result)
}
