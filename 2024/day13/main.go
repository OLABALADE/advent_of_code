package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

func main() {
	f, _ := os.ReadFile("input.txt")
	r := regexp.MustCompile(`X\+(\d+), Y\+(\d+)|X=(\d+), Y=(\d+)`)
	matches := r.FindAllStringSubmatch(string(f), -1)

	targets := [][2]float64{}
	buttons := [][2][2]float64{}

	for i := 2; i < len(matches); i += 3 {
		x, _ := strconv.ParseFloat(matches[i][3], 64)
		y, _ := strconv.ParseFloat(matches[i][4], 64)
		targets = append(targets, [2]float64{x, y})
	}

	for j := 0; j < len(matches); j += 3 {
		ax, _ := strconv.ParseFloat(matches[j][1], 64)
		ay, _ := strconv.ParseFloat(matches[j][2], 64)
		bx, _ := strconv.ParseFloat(matches[j+1][1], 64)
		by, _ := strconv.ParseFloat(matches[j+1][2], 64)

		buttons = append(buttons, [2][2]float64{{ax, ay}, {bx, by}})
	}
	part_1(targets, buttons)
	part_2(targets, buttons)

}

func part_1(targets [][2]float64, buttons [][2][2]float64) {
	var cost float64 = 0
	for index, target := range targets {
		ax := buttons[index][0][0]
		ay := buttons[index][0][1]
		bx := buttons[index][1][0]
		by := buttons[index][1][1]

		d := (ax*by - ay*bx)
		x := (bx*-target[1] - by*-target[0]) / d
		y := (ay*-target[0] - ax*-target[1]) / d

		if x == math.Trunc(x) && y == math.Trunc(y) && x <= 100 && y <= 100 {
			cost += x*3 + y
		}
	}
	fmt.Printf("Part1: %f\n", cost)
}

func part_2(targets [][2]float64, buttons [][2][2]float64) {
	var cost float64 = 0
	for index, target := range targets {
		target[0] += 10000000000000
		target[1] += 10000000000000

		ax := buttons[index][0][0]
		ay := buttons[index][0][1]
		bx := buttons[index][1][0]
		by := buttons[index][1][1]

		d := (ax*by - ay*bx)
		x := (bx*-target[1] - by*-target[0]) / d
		y := (ay*-target[0] - ax*-target[1]) / d

		if x == math.Trunc(x) && y == math.Trunc(y) {
			cost += x*3 + y
		}
	}
	fmt.Printf("Part2: %f\n", cost)
}
