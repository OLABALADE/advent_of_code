package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type Point struct {
	x int
	y int
}

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	mp := strings.Split(string(f), "\n")
	mp = mp[:len(mp)-1]

	antennas := map[string][]Point{}
	for i, row := range mp {
		for j, col := range row {
			if col != '.' {
				c := string(col)
				antennas[c] = append(antennas[c], Point{x: i, y: j})
			}
		}
	}

	Part_1(mp, antennas)
	Part_2(mp, antennas)
}

func Part_1(mp []string, antennas map[string][]Point) {
	antinodes := []Point{}

	for _, locations := range antennas {
		for i := 0; i < len(locations); i++ {
			for j := 0; j < len(locations); j++ {
				if j == i {
					continue
				}
				dx := 2*locations[i].x - locations[j].x
				dy := 2*locations[i].y - locations[j].y

				if dx >= 0 && dx <= len(mp)-1 && dy >= 0 && dy <= len(mp[0])-1 {
					antinode := Point{x: dx, y: dy}
					if !slices.Contains(antinodes, antinode) {
						antinodes = append(antinodes, antinode)
					}
				}

			}
		}
	}
	fmt.Println("Part1: ", len(antinodes))
}

func Part_2(mp []string, antennas map[string][]Point) {
	antinodes := []Point{}

	for _, locations := range antennas {
		for i := 0; i < len(locations); i++ {

			if !slices.Contains(antinodes, locations[i]) {
				antinodes = append(antinodes, locations[i])
			}

			for j := i + 1; j < len(locations); j++ {
				dx := locations[i].x - locations[j].x
				dy := locations[i].y - locations[j].y
				point := Point{x: locations[i].x, y: locations[i].y}
				search(&mp, &antinodes, point, dx, dy)
				search(&mp, &antinodes, point, -dx, -dy)
			}
		}
	}
	fmt.Println("Part2: ", len(antinodes))
}

func search(mp *[]string, antinodes *[]Point, current_point Point, dx, dy int) {
	nx := current_point.x + dx
	ny := current_point.y + dy
	m := *mp
	if nx < 0 || nx >= len(m) || ny < 0 || ny >= len(m[0]) {
		return
	}

	current_point = Point{x: nx, y: ny}
	if !slices.Contains(*antinodes, current_point) {
		*antinodes = append(*antinodes, current_point)
	}

	search(mp, antinodes, current_point, dx, dy)
}
