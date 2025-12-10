
content = File.read!("input.txt") 
  |> String.split("\n") 
  |> Enum.chunk_by(&(&1 != "")) 
  |> Enum.reject(&(&1 == [""]))

defmodule Day5 do
  def parse_range(range) do
    [first , last] = String.split(range, "-")
    {String.to_integer(first) , String.to_integer(last)}
  end

  def in_range(range, id) do
    {first, last} = parse_range(range)
    if id in first..last  do
      true
    else
      false
    end
  end

  def is_fresh(ranges, id) do
    if Enum.any?(ranges, fn range -> in_range(range, id) end) do
      1
    else
      0
    end
  end

  def get_range_set(range) do
    {first, last} = parse_range(range)
    MapSet.new(first..last)
  end

  defp merge_step(current_range, []) do
    [current_range]
  end

  defp merge_step({ran_first, ran_last}, [head | tail]) when ran_first <= elem(head,1) + 1 do
    merged_range = {elem(head, 0), max(elem(head,1), ran_last)}
    [merged_range | tail]
  end

  defp merge_step(current_range, merged) do
    IO.inspect merged
    [current_range | merged]
  end

  def merge_ranges(ranges) do
    ranges 
      |> Enum.sort(&(elem(&1, 0) <= elem(&2, 0)))
      |> Enum.reduce([], &merge_step/2)
      # |> Enum.reverse()
  end
end


[ranges | ids] = content

part1 = ids 
 |> Enum.at(0)
 |> Enum.map(&String.to_integer/1)
 |> Enum.reduce(0, fn id, acc -> Day5.is_fresh(ranges, id) + acc end)

merged_ranges = ranges |> Enum.map(&Day5.parse_range/1) |> Day5.merge_ranges()

IO.inspect merged_ranges

part2 = merged_ranges |> Enum.reduce(0, &(elem(&1,1) - elem(&1,0) + &2 + 1))

IO.puts "Part1=#{part1}"
IO.puts "Part2=#{part2}"
