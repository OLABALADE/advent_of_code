package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.ReadFile("input.txt")
	input := string(f)
	input = input[:len(input)-1]
	mp := strings.Split(input, "\n")

	dirs := [5][2]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	score := 0
	rating := 0

	for x := 0; x < len(mp); x++ {
		for y := 0; y < len(mp[0]); y++ {
			if mp[x][y] == '0' {
				search(x, y, 0, dirs, &mp, &[][2]int{}, &score, &rating)
			}
		}
	}
	// Part 1
	fmt.Println(score)
	// Part 2
	fmt.Println(rating)
}

func search(x, y, current_value int, dirs [5][2]int, mp *[]string, found *[][2]int, score, rating *int) {
	m := *mp

	if current_value == 9 {
		if !slices.Contains(*found, [2]int{x, y}) {
			*found = append(*found, [2]int{x, y})
			*score += 1
		}
		*rating += 1
		return
	}

	for _, dir := range dirs {
		dx := x + dir[0]
		dy := y + dir[1]

		if dx >= 0 && dx < len(m) && dy >= 0 && dy < len(m[0]) {
			next_num, _ := strconv.Atoi(string(m[dx][dy]))

			if next_num == current_value+1 {
				search(dx, dy, next_num, dirs, mp, found, score, rating)
			}
		}
	}
}
