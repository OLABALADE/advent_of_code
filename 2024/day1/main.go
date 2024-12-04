package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

var right []int64
var left []int64
var sum int64 = 0
var similarity int64 = 0;

func main() {
	f, err := os.Open("input.txt")
	check(err)

	defer f.Close()

	r := bufio.NewReader(f)

	for {
		line, _, err := r.ReadLine()
        if err != nil {
            if errors.Is(err, io.EOF) {
                break
            }
        }

		ls := string(line)
		arr := strings.Split(ls, "   ")
		r, _ := strconv.ParseInt(arr[0], 10, 64)
		l, err := strconv.ParseInt(arr[1], 10, 64)

		check(err)

		right = append(right, r)
		left = append(left, l)

	}

	slices.Sort(right)
	slices.Sort(left)

	for i := range len(right) {
		rs := right[i] - left[i]
        sum += int64(math.Abs(float64(rs)))
	}

    fmt.Println(sum)

    data := make(map[int64]int64)
    for _, num1 := range left {
        for _, num2 := range right {
            if num1 == num2 {
                data[num2] = data[num2] + 1
            }
        }
    }
    
    for key, value := range data {
        similarity += key * value 
    }

    fmt.Println(similarity)

}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
