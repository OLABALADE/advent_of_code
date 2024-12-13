package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"slices"
	"strings"
)

type Point struct {
	x int
	y int
}

func main() {
	con, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	mp := strings.Split(string(con), "\n")
	mp = mp[:len(mp)-1]
	visited := []Point{}
	part_1(mp, &visited)
	part_2(mp, visited)
}

func find_guard(mp []string) (int, int) {
	for x, row := range mp {
		for y, char := range row {
			if char == '^' {
				return x, y
			}
		}
	}
	return -1, -1
}

func part_1(mp []string, visited *[]Point) {
	dir_x := -1
	dir_y := 0
	count := 0
	gx, gy := find_guard(mp)

	for true {
		point := Point{x: gx, y: gy}
		if !slices.Contains(*visited, point) {
			*visited = append(*visited, point)
			count += 1
		}

		if gx >= len(mp)-1 || gx <= 0 || gy >= len(mp[0])-1 || gy <= 0 {
			break
		}

		next := mp[gx+dir_x][gy+dir_y]

		if next == '#' {
			dir_x, dir_y = dir_y, dir_x*-1
		} else {
			gx += dir_x
			gy += dir_y
		}

	}

	fmt.Println(count)
}

func part_2(mp []string, visited []Point) {
	count := 0
	for index, point := range visited {
		if index == 0 {
			continue
		}

		cp := make([]string, len(mp))
		copy(cp, mp)
		cp[point.x] = replace_at_point(cp[point.x], point.y)
		if check_loop(cp) {
			count += 1
		}
	}

	fmt.Println(count)
}

func check_loop(cp []string) bool {
	pvisited := []Point{}
	dir := map[Point][]int{}
	dir_x := -1
	dir_y := 0
	gx, gy := find_guard(cp)

	for true {
		point := Point{x: gx, y: gy}

		v, ok := dir[point]
		if ok {
			if reflect.DeepEqual(v, []int{dir_x, dir_y}) {
				return true
			}
		}

		if !slices.Contains(pvisited, point) {
			pvisited = append(pvisited, point)
			d := []int{dir_x, dir_y}
			dir[point] = d
		}

		if gx >= len(cp)-1 || gx <= 0 || gy >= len(cp[0])-1 || gy <= 0 {
			return false
		}

		next := cp[gx+dir_x][gy+dir_y]

		if next == '#' {
			dir_x, dir_y = dir_y, dir_x*-1

		} else {
			gx += dir_x
			gy += dir_y
		}

	}
	return false
}

func replace_at_point(str string, i int) string {
	b := []rune(str)
	b[i] = rune('#')
	str = string(b)
	return str
}
