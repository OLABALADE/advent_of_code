defmodule Day3 do
  def parse_str(str) do
    str |> String.graphemes() |> Enum.map(&String.to_integer/1)
  end

  def find_max1(str) do
    parse_str(str)
    |> Enum.max()
    |> Integer.to_string()
  end

  def largest_joltage(line) do
    if String.length(line) == 1 do
      String.to_integer(line)
    else
      head = String.first(line)
      tail = String.slice(line, 1..-1//1)
      num = (head <> find_max1(tail)) |> String.to_integer()
      max(num, largest_joltage(tail))
    end
  end

  def find_max2(list, count, length) do
    list
    |>Enum.take(length - count + 1)
    |> Enum.with_index()
    |> Enum.max_by(fn {a,i} -> {a,-i} end)
  end

  def largest_joltage2(line, count) do
    if count == 0 do
      ""
    else
    n = line |> String.graphemes |> length()
    str_list = parse_str(line)

    {d, index} = find_max2(str_list, count, n)

    tail = String.slice(line, (index+1)..-1//1)
    Integer.to_string(d) <> largest_joltage2(tail, count - 1)
    end
  end

end

part1 =
  File.stream!("input.txt")
  |> Stream.map(&String.trim/1)
  |> Enum.reduce(0, fn line, acc -> Day3.largest_joltage(line) + acc end)

part2 =
  File.stream!("input.txt")
  |> Stream.map(&String.trim/1)
  |> Enum.reduce(0, fn line, acc -> String.to_integer(Day3.largest_joltage2(line, 12)) + acc end)

IO.puts("Part 1 : #{part1}")
IO.puts("Part 2 : #{part2}")
