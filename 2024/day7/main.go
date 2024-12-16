package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	values := strings.Split(string(f), "\n")
	values = values[:len(values)-1]
	results := map[uint64][]int{}

	for _, value := range values {
		sp := strings.Split(value, ": ")
		target, _ := strconv.ParseUint(sp[0], 10, 64)
		nums := strings.Fields(sp[1])

		var arr []int
		for _, num := range nums {
			i, _ := strconv.Atoi(num)
			arr = append(arr, i)
		}
		results[target] = arr

	}

	var total_cal uint64
	for target, values := range results {
		if can_obtain(values, target, values[0], 1) {
			total_cal += target
		}
	}

	fmt.Println(total_cal)

}

func can_obtain(nums []int, target uint64, current_value, index int) bool {
	if target == uint64(current_value) {
		return true
	}

	if target < uint64(current_value) {
		return false
	}

	if index >= len(nums) {
		return false
	}

	num := nums[index]

	if can_obtain(nums, target, num+current_value, index+1) {
		return true
	}

	if current_value != 0 && can_obtain(nums, target, num*current_value, index+1) {
		return true
	}

	// Part Two
	s := strconv.Itoa(current_value) + strconv.Itoa(num)
	j, _ := strconv.Atoi(s)

	if can_obtain(nums, target, j, index+1) {
		return true
	}

	return false
}
