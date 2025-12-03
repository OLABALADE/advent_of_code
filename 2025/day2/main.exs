count = 0
defmodule Day2 do 
    def parse(str) do
      [first, last] = String.split(str, "-")
      {String.to_integer(first), String.to_integer(last)}
  end

  def count_invalid_id1(first, last) do 
      Enum.filter(first..last, &(is_invalid1(&1)))
      |> Enum.sum()
  end

  def count_invalid_id2(first, last) do 
      Enum.filter(first..last, &(is_invalid2(&1)))
      |> Enum.sum()
  end

  def is_invalid1(id) do 
     Regex.match?(~r/^(\d+)\1$/, Integer.to_string(id))
  end

  def is_invalid2(id) do 
     Regex.match?(~r/^(\d+)\1+$/, Integer.to_string(id))
  end

  def num_of_invalid_id(range, func) do 
    {first, last} = parse(range)
    func.(first, last)
  end
end

{:ok , content} = File.read("input.txt")
part1 = content
|> String.split(",")
|> Enum.map(&String.trim/1)
|>Enum.reduce(0, fn (range, acc) -> Day2.num_of_invalid_id(range, &Day2.count_invalid_id1/2) + acc end)

part2 = content
|> String.split(",")
|> Enum.map(&String.trim/1)
|>Enum.reduce(0, fn (range, acc) -> Day2.num_of_invalid_id(range, &Day2.count_invalid_id2/2) + acc end)

IO.puts("Part1=#{part1}")
IO.puts("Part2=#{part2}")
