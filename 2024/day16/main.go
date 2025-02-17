package main

import (
	"container/heap"
	"fmt"
	"os"
	"slices"
	"strings"
)

// //////////////// Priority Queque /////////////////////////
type Node struct {
	index    int
	cost     int
	dir, pos [2]int
	path     [][2]int
}

type Pri_Queue []*Node

func (pq Pri_Queue) Len() int { return len(pq) }
func (pq Pri_Queue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq *Pri_Queue) Push(x any) {
	n := len(*pq)
	item := x.(*Node)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *Pri_Queue) Pop() any {
	old := *pq
	n := len(old)
	item := old[0]
	old[0], old[n-1] = old[n-1], nil
	*pq = old[:n-1]
	heap.Fix(pq, 0)
	return item
}

func (pq Pri_Queue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

//////////////////////////////////////////////////////////

var start Node
var mini_score int = 1000000000000000000
var unique [][2]int

func main() {
	f, _ := os.ReadFile("input.txt")
	input := strings.Split(string(f), "\n")
	input = input[:len(input)-1]

	for i, row := range input {
		for j, ch := range row {
			if ch == 'S' {
				start = Node{
					cost: 0,
					pos:  [2]int{i, j},
					dir:  [2]int{0, 1},
					path: [][2]int{{i, j}},
				}
			}
		}
	}

	solution(input)
	fmt.Println("Part 1:", mini_score)
	fmt.Println("Part 2:", len(unique))
}

func solution(mp []string) {
	visited := [][4]int{}
	pq := make(Pri_Queue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &start)

	dirs := [4][2]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	for len(pq) > 0 {
		n := pq.Pop().(*Node)

		if n.cost > mini_score {
			continue
		}

		a := [4]int{n.pos[0], n.pos[1], n.dir[0], n.dir[1]}
		if !slices.Contains(visited, a) {
			visited = append(visited, a)
		}

		if mp[n.pos[0]][n.pos[1]] == 'E' {
			if n.cost <= mini_score {
				mini_score = n.cost
				for _, point := range n.path {
					if !slices.Contains(unique, point) {
						unique = append(unique, point)
					}
				}
			} else {
				break
			}
		}

		for _, dir := range dirs {
			if dir == [2]int{-n.dir[0], -n.dir[1]} {
				continue
			}

			dx := n.pos[0] + dir[0]
			dy := n.pos[1] + dir[1]

			if mp[dx][dy] != '#' && !slices.Contains(visited, [4]int{dx, dy, dir[0], dir[1]}) {
				nw := Node{pos: [2]int{dx, dy}, dir: dir}
				if dir != n.dir {
					nw.cost += n.cost + 1001
				} else {
					nw.cost += n.cost + 1
				}
				nw.path = append(nw.path, nw.pos)
				nw.path = append(nw.path, n.path...)
				heap.Push(&pq, &nw)
			}
		}
	}
}
