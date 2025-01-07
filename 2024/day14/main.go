package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

const m int = 101
const n int = 103

func main() {
	f, _ := os.ReadFile("input.txt")
	r := regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)
	matches := r.FindAllStringSubmatch(string(f), -1)
	num_robots := len(matches)

	positions := [][2]int{}
	velocities := [][2]int{}

	//Creating empty grid
	mp := []string{}
	for i := 0; i < 103; i++ {
		mp = append(mp, strings.Repeat(".", 101))
	}

	//Stores unique positions of robots
	locations := [][2]int{}

	for _, match := range matches {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		positions = append(positions, [2]int{x, y})

		if !slices.Contains(locations, [2]int{x, y}) {
			locations = append(locations, [2]int{x, y})
		}

		vx, _ := strconv.Atoi(match[3])
		vy, _ := strconv.Atoi(match[4])
		velocities = append(velocities, [2]int{vx, vy})
	}

	np := make([][2]int, len(positions))
	copy(np, positions)

	part_1(positions, velocities, locations)
	part_2(num_robots, mp, np, velocities, locations)
}

func part_1(positions, velocities, locations [][2]int) {
	for i := 0; i < 100; i++ {
		locations = nil
		for index, pos := range positions {
			positions[index][0] = (pos[0] + velocities[index][0] + m) % m
			positions[index][1] = (pos[1] + velocities[index][1] + n) % n

			r := positions[index][1]
			c := positions[index][0]

			if !slices.Contains(locations, [2]int{c, r}) {
				locations = append(locations, [2]int{c, r})
			}
		}
	}

	sx := m / 2
	sy := n / 2

	q1 := 0
	q2 := 0
	q3 := 0
	q4 := 0

	for _, pos := range positions {
		if pos[0] != sx && pos[1] != sy {
			switch {
			case pos[0] < sx && pos[1] < sy:
				q1 += 1
			case pos[0] > sx && pos[1] < sy:
				q2 += 1
			case pos[0] < sx && pos[1] > sy:
				q3 += 1
			case pos[0] > sx && pos[1] > sy:
				q4 += 1
			}
		}
	}

	fmt.Println("Part 1:", q1*q2*q3*q4)
}

func show_tree(mp []string, locations [][2]int) {
	for _, pos := range locations {
		x := pos[0]
		y := pos[1]

		d := []rune(mp[y])
		d[x] = 'o'
		mp[y] = string(d)
	}

	for _, row := range mp {
		fmt.Println(row)
	}
}

func part_2(num_robots int, mp []string, positions, velocities, locations [][2]int) {
	count := 0
	for {
		if len(locations) == num_robots {
			fmt.Println("Part 2:", count)
			show_tree(mp, locations)
			break
		}

		locations = nil
		for index, pos := range positions {
			positions[index][0] = (pos[0] + velocities[index][0] + m) % m
			positions[index][1] = (pos[1] + velocities[index][1] + n) % n

			r := positions[index][1]
			c := positions[index][0]

			if !slices.Contains(locations, [2]int{c, r}) {
				locations = append(locations, [2]int{c, r})
			}
		}
		count += 1
	}
}
