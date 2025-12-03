dial = 50
password = 0

defmodule Day1 do 

def zero_rotations(turns, incre, current_number, num_passed) do 
  if current_number >= turns do num_passed
  else 
    zero_rotations(turns, incre, current_number + incre, num_passed + 1)
  end
end

def rotate_dial1(line, current_dial, current_password) do
  if String.starts_with?(line, "R") do
    d = current_dial + (line |> String.slice(1..3) |> String.to_integer()) |> Integer.mod(100)
    pass = if d === 0, do: current_password + 1, else: current_password
    {d, pass}
  else
    d = current_dial + (-1 * (String.slice(line,1..3) |> String.to_integer())) |> Integer.mod(100)
    pass = if d === 0, do: current_password + 1, else: current_password
    {d, pass}
  end
end

def rotate_dial2(line, current_dial, current_password) do
    turn = line |> String.slice(1..3) |> String.to_integer()

  if String.starts_with?(line, "R") do
    d = current_dial + turn |> Integer.mod(100)
    zero_rots = zero_rotations(turn, 100, 100-current_dial, 0)
    pass = if d === 0, do: current_password + 1 + zero_rots , else: current_password + zero_rots
    {d, pass}

  else
    d = current_dial + (-1 * turn) |> Integer.mod(100)
    e = if current_dial === 0, do: 100, else: current_dial
    zero_rots = zero_rotations(turn, 100, e, 0)
    pass = if d === 0, do: current_password + 1 + zero_rots  , else: current_password + zero_rots
    {d, pass}

  end
end
end

part1 = File.stream!("input.txt")
|> Stream.map(&String.trim/1)
|> Enum.reduce({dial, password} , fn (line , {current_dial,current_password}) -> Day1.rotate_dial1(line, current_dial, current_password) end)

IO.puts("Part1: Password=#{elem(part1, 1)}")

part2 = File.stream!("input.txt")
|> Stream.map(&String.trim/1)
|> Enum.reduce({dial, password} , fn (line , {current_dial,current_password}) -> Day1.rotate_dial2(line, current_dial, current_password) end)

IO.puts("Part2: Password=#{elem(part2, 1)}")
