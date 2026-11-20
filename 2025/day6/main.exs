defmodule Day6 do
  def extract_numbers(line) do
    line 
    |> String.split(" ", trim: true)
    |> Enum.map(&String.to_integer/1)
  end

  def calculate(first, second, operation) do
    if operation == "+" do
      first + second
    else
      first * second
    end
  end

  def perform_operation(current_list, last_list, operations) do
    operations
    |> Enum.with_index()
    |> Enum.reduce([], fn {op, idx}, acc -> 
      first_num = Enum.at(last_list, idx)
      second_num = Enum.at(current_list, idx)
      acc ++ [calculate(first_num, second_num, op)]
    end)
  end

  def get_numbers_at_index(numbers, index) do
    numbers |> Enum.map(&(Enum.at(&1, index)))
  end

  def update_num_map(map, num) do
    IO.puts num
    num |> Integer.to_string()
    |> String.graphemes()
    |> Enum.reverse()
    |> Enum.with_index()
    |> Enum.reduce(map, fn{num_str, index}, acc -> Map.update(acc, index, num_str, &(&1 <> num_str))end)
  end

  def calculate_in_index(op, index, numbers) do
    nums_at_index = get_numbers_at_index(numbers, index)
    IO.puts "Index#{index}"
    IO.inspect nums_at_index
    num_map = nums_at_index |> Enum.reduce(%{}, fn num, mp -> update_num_map(mp, num)end)
    IO.inspect num_map

    {first_key, first_value} = Map.to_list(num_map) |> Enum.at(0)
    first_num = String.to_integer(first_value)

    num_map
    |> Map.delete(first_key)
    |> Enum.reduce(first_num, fn {key, value}, acc -> 
      calculate(acc, String.to_integer(value), op)
    end)

  end
end

content = File.stream!("ex.txt")
  |> Enum.map(&String.trim_trailing/1)
  |> Enum.to_list()
  |> Enum.reverse()

[op_str | num_strs] = content
IO.inspect num_strs
operations = op_str |> String.split(" ", trim: true)
numbers = num_strs |> Enum.map(&Day6.extract_numbers/1)

[first_list | num_list] = numbers
part1 = num_list
  |> Enum.reduce(first_list, &(Day6.perform_operation(&1, &2, operations)))
  |> Enum.sum()

part2 = operations
  |> Enum.with_index()
  |>  Enum.reduce(0, fn{op, index}, acc -> Day6.calculate_in_index(op, index, numbers) + acc end) 

IO.inspect part1
