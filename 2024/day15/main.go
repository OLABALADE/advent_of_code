package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	f, _ := os.ReadFile("input.txt")
	input := strings.Split(string(f), "\n\n")

	store1 := strings.Split(input[0], "\n")
	dirs := strings.Join(strings.Split(input[1], "\n"), "")
	r_pos1 := [2]int{}
	r_pos2 := [2]int{}

	for i := range len(store1) {
		for j := range len(store1[0]) {
			if store1[i][j] == '@' {
				r_pos1[0] = i
				r_pos1[1] = j
				break
			}
		}
	}

	boxes := [][2]int{}
	walls := [][2]int{}

	for i, row := range store1 {
		for j, char := range row {
			switch char {
			case '#':
				walls = append(walls, [2]int{i, j * 2}, [2]int{i, j*2 + 1})
			case 'O':
				boxes = append(boxes, [2]int{i, j * 2})
			case '@':
				r_pos2[0] = i
				r_pos2[1] = j * 2
			}
		}
	}

	part_1(r_pos1, dirs, store1)
	part_2(r_pos2, walls, boxes, dirs)

}

func part_1(r_pos [2]int, dirs string, store []string) {
	dir_mp := map[string][2]int{
		">": {0, 1},
		"<": {0, -1},
		"^": {-1, 0},
		"v": {1, 0},
	}
	for i := 0; i < len(dirs); i++ {
		dir := dir_mp[string(dirs[i])]
		dx := r_pos[0] + dir[0]
		dy := r_pos[1] + dir[1]

		if store[dx][dy] == 'O' {
			c := [2]int{dx, dy}

			for {
				c[0] += dir[0]
				c[1] += dir[1]

				if store[c[0]][c[1]] == '#' {
					break
				}

				if store[c[0]][c[1]] == '.' {
					change_row(&store, c[0], c[1], 'O')
					change_row(&store, dx, dy, '@')
					change_row(&store, r_pos[0], r_pos[1], '.')
					r_pos[0] = dx
					r_pos[1] = dy
					break
				}
			}

		} else if store[dx][dy] == '.' {
			change_row(&store, dx, dy, '@')
			change_row(&store, r_pos[0], r_pos[1], '.')
			r_pos[0] = dx
			r_pos[1] = dy

		}
	}

	gps := 0
	for i, row := range store {
		for j, char := range row {
			if char == 'O' {
				gps += 100*i + j
			}
		}
	}
	fmt.Println("Part 1:", gps)
}

func part_2(r_pos [2]int, walls, boxes [][2]int, dirs string) {
	dir_mp := map[string][2]int{
		">": {0, 1},
		"<": {0, -1},
		"^": {-1, 0},
		"v": {1, 0},
	}

	for i := 0; i < len(dirs); i++ {
		dir := dir_mp[string(dirs[i])]
		dx := r_pos[0] + dir[0]
		dy := r_pos[1] + dir[1]

		stack := [][2]int{}
		if slices.Contains(walls, [2]int{dx, dy}) {
			continue
		}

		if slices.Contains(boxes, [2]int{dx, dy}) {
			stack = append(stack, [2]int{dx, dy})
		} else if slices.Contains(boxes, [2]int{dx, dy - 1}) {
			stack = append(stack, [2]int{dx, dy - 1})
		}

		seen := [][2]int{}
		move := true
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			nx := top[0] + dir[0]
			ny := top[1] + dir[1]

			if slices.Contains(walls, [2]int{nx, ny}) || slices.Contains(walls, [2]int{nx, ny + 1}) {
				move = false
				break
			}

			if slices.Contains(seen, top) {
				continue
			}
			seen = append(seen, top)
			if slices.Contains(boxes, [2]int{nx, ny}) {
				stack = append(stack, [2]int{nx, ny})
			}

			if slices.Contains(boxes, [2]int{nx, ny - 1}) {
				stack = append(stack, [2]int{nx, ny - 1})
			}

			if slices.Contains(boxes, [2]int{nx, ny + 1}) {
				stack = append(stack, [2]int{nx, ny + 1})
			}

		}

		if move {
			for i, box := range boxes {
				if slices.Contains(seen, box) {
					boxes[i][0] += dir[0]
					boxes[i][1] += dir[1]
				}
			}

			r_pos[0] += dir[0]
			r_pos[1] += dir[1]
		}
	}

	gps := 0
	for _, box := range boxes {
		gps += 100*box[0] + box[1]
	}

	fmt.Println("Part 2:", gps)
}

func change_row(store *[]string, row, col int, char byte) {
	st := *store
	n := []rune(st[row])
	n[col] = rune(char)
	st[row] = string(n)
}
