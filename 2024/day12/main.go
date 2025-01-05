package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	f, _ := os.ReadFile("input.txt")
	input := strings.Split(string(f), "\n")
	input = input[:len(input)-1]

	visited := [][2]int{}
	cost1 := 0
	cost2 := 0

	for i := range len(input) {
		for j := range len(input[0]) {
			if !slices.Contains(visited, [2]int{i, j}) {
				visited = append(visited, [2]int{i, j})
				ar, per, sd := find_region(i, j, input[i][j], input, &visited)
				cost1 += (ar * per)
				cost2 += (ar * sd)
			}
		}
	}

	fmt.Println("Part 1:", cost1)
	fmt.Println("Part 2:", cost2)
}

func find_region(x, y int, plant byte, mp []string, visited *[][2]int) (int, int, int) {
	area := 0
	perm := 0
	sides := check_corner(x, y, plant, mp)

	dirs := [4][2]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	for _, dir := range dirs {
		dx := x + dir[0]
		dy := y + dir[1]
		if !is_bound(dx, dy, mp) || mp[dx][dy] != plant {
			perm += 1
		}

		if is_bound(dx, dy, mp) && mp[dx][dy] == plant {
			if !slices.Contains(*visited, [2]int{dx, dy}) {
				*visited = append(*visited, [2]int{dx, dy})
				a, p, s := find_region(dx, dy, plant, mp, visited)
				area += a
				perm += p
				sides += s
			}
		}
	}

	area += 1
	return area, perm, sides
}

func check_corner(x, y int, plant byte, mp []string) int {
	dirs := [4][2]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	sides := 0

	for i, dir := range dirs {
		dx := x + dir[0]
		dy := y + dir[1]

		nx := x + dirs[(i+1)%4][0]
		ny := y + dirs[(i+1)%4][1]

		if !is_ok(dx, dy, plant, mp) && !is_ok(nx, ny, plant, mp) {
			sides += 1
		}

		c_x := dir[0] + dirs[(i+1)%4][0]
		c_y := dir[1] + dirs[(i+1)%4][1]

		if (is_ok(dx, dy, plant, mp) && is_ok(nx, ny, plant, mp)) && !is_ok(x+c_x, y+c_y, plant, mp) {
			sides += 1

		}

	}

	return sides
}

func is_bound(x, y int, mp []string) bool {
	m := len(mp)
	n := len(mp[0])
	return x >= 0 && x < m && y >= 0 && y < n
}

func is_ok(x, y int, plant byte, mp []string) bool {
	return is_bound(x, y, mp) && mp[x][y] == plant
}
