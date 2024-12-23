package main

import (
	"fmt"
	"os"
	"strconv"
)

type File struct {
	id    int
	block int
	index int
}

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	input_num := string(f)
	input_num = input_num[:len(input_num)-1]

	frag := []string{}
	files := []File{}
	ni := 0

	for index, char := range input_num {
		id := "."
		num, _ := strconv.Atoi(string(char))

		if (index+1)%2 != 0 {
			id = strconv.Itoa(ni)
			file := File{
				id:    ni,
				block: num,
				index: len(frag),
			}
			files = append(files, file)
			ni++
		}

		for i := 0; i < num; i++ {
			frag = append(frag, string(id))
		}
	}

	fr := make([]string, len(frag))
	copy(fr, frag)

	Part_1(frag)
	Part_2(fr, files)
}

func Part_1(frag []string) {
	k := len(frag) - 1
Out:
	for i := 0; i < len(frag); i++ {
		if frag[i] == "." {
			for j := k; j > 0; j-- {
				a := frag[j]
				if a != "." {
					if j == i-1 {
						break Out
					}
					frag[i] = a
					frag[j] = "."
					k = j - 1
					break
				}
			}

		}
	}
	fmt.Println(cal_checksum(frag))
}

func Part_2(frag []string, files []File) {
	for i := len(files) - 1; i > 0; i-- {
		free_space := []string{}
		file := files[i]

		r := frag[:]
		if i != len(files)-1 {
			r = frag[:file.index+1]
		}

		for j, mem := range r {
			if mem == "." {
				free_space = append(free_space, mem)

			} else if len(free_space) > 0 {
				if file.block <= len(free_space) {
					f := j - len(free_space)

					for k := 0; k < file.block; k++ {
						frag[f] = strconv.Itoa(file.id)
						frag[file.index+k] = "."
						f++
					}
					break
				} else {
					free_space = []string{}
				}
			}
		}
	}
	fmt.Println(cal_checksum(frag))
}

func cal_checksum(frag []string) int {
	checksum := 0
	for idx, ch := range frag {
		if ch != "." {
			num, _ := strconv.Atoi(ch)
			checksum += (idx * num)
		}
	}
	return checksum
}
