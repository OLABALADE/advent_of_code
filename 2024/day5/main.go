package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	safe_updates := []string{}
	fixed_updates := []string{}
	rules := map[string][]string{}
	updates := []string{}
	sum1 := 0
	sum2 := 0

	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	content := string(f)

	sp := strings.Split(content, "\n\n")
	r := strings.Split(sp[0], "\n")
	u := strings.Split(sp[1], "\n")
	load_rules(r, rules)
	load_updates(u, &updates)

	for _, update := range updates {
		ul := strings.Split(update, ",")
		if is_safe(ul, rules) {
			safe_updates = append(safe_updates, update)
			continue
		} else {
			fixed := strings.Join(fix_update(update, rules), ",")
			fixed_updates = append(fixed_updates, fixed)
		}
	}

	if len(safe_updates) != 0 && len(fixed_updates) != 0 {
		for _, update := range safe_updates {
			ul := strings.Split(update, ",")
			mid := len(ul) / 2
			n, _ := strconv.Atoi(ul[mid])
			sum1 += n
		}

		for _, update := range fixed_updates {
			ul := strings.Split(update, ",")
			mid := len(ul) / 2
			n, _ := strconv.Atoi(ul[mid])
			sum2 += n
		}
	}
	fmt.Println(sum1)
	fmt.Println(sum2)
}

func load_rules(list []string, store map[string][]string) {
	for _, rule := range list {
		rule = strings.TrimRight(rule, "\n")
		a := strings.Split(rule, "|")
		x := a[0]
		value, ok := store[x]

		if ok {
			store[x] = append(value, a[1])
		} else {
			store[x] = []string{a[1]}
		}
	}
}

func load_updates(list []string, store *[]string) {
	for _, update := range list {
		update = strings.TrimRight(update, "\n")
		*store = append(*store, update)
	}
}

func is_safe(update []string, store map[string][]string) bool {
	p := []string{}
	for i, num := range update {
		if i == 0 {
			p = append(p, num)
			continue
		}
		ys, ok := store[num]
		if ok {
			for _, f := range p {
				if slices.Contains(ys, f) {
					return false
				}
			}
		}
		p = append(p, num)
	}
	return true
}

func fix_update(update string, store map[string][]string) []string {
	ul := strings.Split(update, ",")
	for i := 0; i < len(ul); i++ {
		j := i
		for j > 0 && !slices.Contains(store[ul[j-1]], ul[j]) {
			ul[j-1], ul[j] = ul[j], ul[j-1]
			j -= 1
		}
	}
	return ul
}
