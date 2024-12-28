package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.ReadFile("input.txt")
	input := strings.Fields(strings.TrimSuffix(string(f), "\n"))
	mp := map[string]int{}
	for _, num := range input {
		mp[num] += 1
	}

	part_1(input)
	part_2(mp)
}

func part_1(stones []string) {
	for a := 0; a < 25; a++ {
		new_stones := []string{}
		for _, stone := range stones {

			if stone == "0" {
				new_stones = append(new_stones, "1")

			} else if len(stone)%2 == 0 {
				k := len(stone) / 2
				l_stone, _ := strconv.Atoi(stone[:k])
				r_stone, _ := strconv.Atoi(stone[k:])

				new_stones = append(new_stones, strconv.Itoa(l_stone), strconv.Itoa(r_stone))
			} else {
				num, _ := strconv.Atoi(stone)
				nw_stone := strconv.Itoa(num * 2024)
				new_stones = append(new_stones, nw_stone)
			}
		}
		stones = new_stones
	}
	fmt.Println(len(stones))
}

func part_2(mp map[string]int) {
	for i := 0; i < 75; i++ {
		new_mp := map[string]int{}

		for key := range mp {
			if key == "0" {
				new_mp["1"] += mp[key]

			} else if len(key)%2 == 0 {
				k := len(key) / 2
				l_stone, _ := strconv.Atoi(key[:k])
				r_stone, _ := strconv.Atoi(key[k:])
				new_mp[strconv.Itoa(l_stone)] += mp[key]
				new_mp[strconv.Itoa(r_stone)] += mp[key]

			} else {
				num, _ := strconv.Atoi(key)
				num *= 2024
				new_mp[strconv.Itoa(num)] += mp[key]
			}
		}
		mp = new_mp
	}

	count := 0
	for _, value := range mp {
		count += value
	}

	fmt.Println(count)
}
