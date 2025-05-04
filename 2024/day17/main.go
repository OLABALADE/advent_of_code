package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("ex.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rgx := regexp.MustCompile(`Register \w: (\d+)|Program: ((\d,?)*)`)
	matches := rgx.FindAllStringSubmatch(string(f), -1)

	ra, _ := strconv.Atoi(matches[0][1])
	rb, _ := strconv.Atoi(matches[1][1])
	rc, _ := strconv.Atoi(matches[2][1])
	pg := strings.Split(matches[3][2], ",")
	ip := 0

	combos := map[int]*int{
		4: &ra,
		5: &rb,
		6: &rc,
	}
	part1 := ""
	for {
		if ip >= len(pg) {
			break
		}

		opcode, _ := strconv.Atoi(pg[ip])
		operand, _ := strconv.Atoi(pg[ip+1])
		comOperand := 0
		if operand <= 3 {
			comOperand = operand
		} else if operand != 7 {
			comOperand = *combos[operand]
		}

		switch opcode {
		case 0:
			ra /= (int(math.Pow(2, float64(comOperand))))
		case 1:
			rb ^= operand
		case 2:
			rb = comOperand % 8
		case 3:
			{
				if ra != 0 {
					if ip != operand {
						ip = operand
						continue
					}
					ip = operand
				}
			}
		case 4:
			rb ^= rc
		case 5:
			{
				fmt.Printf("%d,", comOperand%8)
				part1 += fmt.Sprint(comOperand % 8)
			}
		case 6:
			rb = ra / (int(math.Pow(2, float64(comOperand))))
		case 7:
			rc = ra / (int(math.Pow(2, float64(comOperand))))
		}

		ip += 2
	}
	fmt.Println("\n", part1)
}
