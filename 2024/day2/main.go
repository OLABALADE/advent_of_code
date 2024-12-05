package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var safe int

func main() {
	f, err := os.Open("input.txt")
	check(err)

	defer f.Close()

	r := bufio.NewReader(f)
	for {
		l, _, err := r.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
		}

		line := strings.Split(string(l), " ")
		report := arr_int(line)

		if is_safe(report) || is_tolerable(report) {
			safe += 1
		}
	}

	fmt.Println(safe)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func arr_int(line []string) []int {
	arr := make([]int, len(line))
	for i, s := range line {
		arr[i], _ = strconv.Atoi(s)
	}

	return arr
}

func is_safe(rep []int) bool {
	i, j := -1, 1
	diff := rep[0] - rep[1]

	if !(math.Signbit(float64(diff))) {
		i, j = 1, -1
	}

	count := 0
	for index, level := range rep {
		if index != len(rep)-1 {
			diff = level*i + rep[index+1]*j
			abs_diff := math.Abs(float64(diff))

			if !math.Signbit(float64(diff)) && abs_diff != 0 && abs_diff <= 3 {
				count += 1
			}

		}
	}

	if count == len(rep)-1 {
		return true
	}

	return false
}

func is_tolerable(rep []int) bool {
	for index := range len(rep) {
		c := make([]int, len(rep))
		copy(c, rep)
		c = append(c[:index], c[index+1:]...)

		if is_safe(c) {
			return true
		}
	}

	return false
}
