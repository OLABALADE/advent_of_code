package main

import "core:fmt"
import "core:os"
import "core:strings"

main::proc() {
  filepath := "input.txt"
  data, err := os.read_entire_file_from_path(filepath, context.allocator)

  if err != nil {
    fmt.eprintln("Failed to read file:", filepath)
    return
  }

  defer delete(data)

  part1:=0
  part2:=0
  content := string(strings.trim_right(string(data) ,"\r\n"))

  for i:=0; i + 4 <= len(content); i+=1 {
    slice := content[i:i+4]
    if(!has_duplicates(string(slice))) {
      part1 = i+4
      break
    }
  }

  for i:=0; i + 14 <= len(content); i+=1 {
    slice := content[i:i+14]
    if(!has_duplicates(string(slice))) {
      part2 = i+14
      break
    }
  }

  fmt.println("Part1 -> Characters processed:", part1)
  fmt.println("Part2 -> Characters processed:", part2)
}

has_duplicates :: proc(s:string) -> bool {
  seen := make(map[rune]bool)
  defer delete(seen)

  for r in s {
    if r in seen {
      return true
    }
    seen[r] = true
  }
  return false
}
