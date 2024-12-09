package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var count1 int
var count2 int

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	lines = lines[:len(lines)-1]

	for i, line := range lines {
		for j, char := range line {
			if char == 'X' {
				handle_cases_1(i, j, lines)
			} else if char == 'A' {
				handle_cases_2(i, j, lines)
			}
		}
	}

	fmt.Println(count1)
	fmt.Println(count2)
}

func check_diagonal(x, y, c_x, c_y int, lines []string) {
	s := ""
	i := x + c_x
	j := y + c_y
	length := len(lines[0]) - 1
	for a := 0; a < 3; a++ {
		if i >= 0 && j >= 0 && i <= length && j <= length {
			s += string(lines[i][j])
			i += c_x
			j += c_y
		} else {
			break
		}
	}
	if s == "MAS" {
		count1 += 1
	}
}

func check_for(x, y int, lines []string) {
	s := ""
	j := y + 1
	length := len(lines[0]) - 1
	for a := 0; a < 3; a++ {
		if j <= length {
			s += string(lines[x][j])
			j += 1
		} else {
			break
		}
	}
	if s == "MAS" {
		count1 += 1
	}
}

func check_bac(x, y int, lines []string) {
	s := ""
	j := y - 1
	for a := 0; a < 3; a++ {
		if j >= 0 {
			s += string(lines[x][j])
			j -= 1
		} else {
			break
		}
	}
	if s == "MAS" {
		count1 += 1
	}
}

func check_up(x, y int, lines []string) {
	s := ""
	i := x - 1
	for a := 0; a < 3; a++ {
		if i >= 0 {
			s += string(lines[i][y])
			i -= 1
		} else {
			break
		}
	}
	if s == "MAS" {
		count1 += 1
	}
}

func check_down(x, y int, lines []string) {
	s := ""
	i := x + 1
	for a := 0; a < 3; a++ {
		if i < len(lines) {
			s += string(lines[i][y])
			i += 1
		} else {
			break
		}
	}
	if s == "MAS" {
		count1 += 1
	}
}

func check_all(x, y int, lines []string) {
	check_diagonal(x, y, -1, -1, lines) // top left diagonal
	check_diagonal(x, y, -1, 1, lines)  // top right diagonal
	check_diagonal(x, y, 1, -1, lines)  // bottom left diagonal
	check_diagonal(x, y, 1, 1, lines)   // bottom right diagonal
	check_for(x, y, lines)
	check_bac(x, y, lines)
	check_up(x, y, lines)
	check_down(x, y, lines)
}

func handle_cases_1(x, y int, lines []string) {
	switch {
	case x == 0:
		check_for(x, y, lines)
		check_bac(x, y, lines)
		check_down(x, y, lines)
		check_diagonal(x, y, 1, -1, lines)
		check_diagonal(x, y, 1, 1, lines)
	case x == len(lines)-1:
		check_for(x, y, lines)
		check_bac(x, y, lines)
		check_up(x, y, lines)
		check_diagonal(x, y, -1, -1, lines)
		check_diagonal(x, y, -1, 1, lines)
	case y == 0:
		check_for(x, y, lines)
		check_up(x, y, lines)
		check_down(x, y, lines)
		check_diagonal(x, y, -1, 1, lines)
		check_diagonal(x, y, 1, 1, lines)
	case y == len(lines[0])-1:
		check_up(x, y, lines)
		check_down(x, y, lines)
		check_bac(x, y, lines)
		check_diagonal(x, y, -1, -1, lines)
		check_diagonal(x, y, 1, -1, lines)
	default:
		check_all(x, y, lines)
	}
}

func check_xdiagonal(x, y int, lines []string) {
	a := string(lines[x-1][y-1]) + string(lines[x+1][y+1])
	b := string(lines[x-1][y+1]) + string(lines[x+1][y-1])
	fmt.Println(a, b)

	if (a == "SM" || a == "MS") && (b == "SM" || b == "MS") {
		count2 += 1
	}
}

func handle_cases_2(x, y int, lines []string) {
	if x != 0 && x != len(lines)-1 && y != 0 && y != len(lines[0])-1 {
		check_xdiagonal(x, y, lines)
	}
}
