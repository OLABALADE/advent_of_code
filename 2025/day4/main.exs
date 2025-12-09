grid =  
  File.stream!("input.txt")
  |> Stream.map(&String.trim_trailing/1)
  |> Enum.to_list()


n = grid |> Enum.at(0) |> String.graphemes() |> length()
m = length(grid)

defmodule Day4 do
  @dirs [
    {0,-1},
    {1,-1},
    {1,0},
    {1,1},
    {0,1},
    {-1,1},
    {-1,0},
    {-1,-1},
  ]

  def is_accessible(i, j, mp, m, n, dirs) do
    if Enum.at(mp, i) |> String.at(j) != "@" do
      0
    else

    count = Enum.filter(dirs, fn {dx, dy} ->  
      is_adjacent(i + dx, j + dy, mp, m, n) end)
    |> length()


    if count < 4, do: 1, else: 0
    end
  end

  def is_adjacent(i, j, mp, m, n) do
    if i in 0..m-1 and j in 0..n-1 do
      Enum.at(mp, i) |> String.at(j) == "@"
    else
      false
    end
  end

  def find_removed(mp, num_row, num_col, positions) do 
      positions
      |> Enum.reduce([], fn {row, col}, rmv -> 
        value = is_accessible(row, col, mp, num_row, num_col, @dirs) 
        if value == 1 , do: rmv ++ [{row, col}], else: rmv
    end) 
  end

  def replace_row(row, col, mp) do
    r = mp |> Enum.at(row)
    r |> String.graphemes() |> List.replace_at(col, ".") |> Enum.join()
  end

  def replace_map(removed, mp)do
    removed 
    |> Enum.reduce(mp, fn {row, col}, mp -> 
      new_row = replace_row(row, col, mp)
      mp |> List.replace_at(row, new_row)
    end)
  end

  def count_all_removed(mp, count, num_row, num_col, positions) do
    removed_pos = find_removed(mp, num_row, num_col, positions)
    num_removed = removed_pos |> length()
    if num_removed == 0 do
      count
    else
      new_map = replace_map(removed_pos, mp)
      count_all_removed(new_map, count + num_removed, num_row, num_col, positions)
    end
  end

end

positions = for i <- 0..m-1 do
  for j <- 0..n-1 do
    {i,j}
  end
end
|> Enum.reduce([], fn list, acc -> list ++ acc end) 
|> Enum.reverse()

part1 = Day4.find_removed(grid, m, n, positions) |> length()
part2 = Day4.count_all_removed(grid,0, m, n, positions)

IO.puts "Part1=#{part1}"
IO.puts "Part2=#{part2}"
