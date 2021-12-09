# frozen_string_literal: true

# janky function to determine if board is a winner
# all marked off squares get turned into -1
def winner(board)
  i = 0
  rows = [0, 0, 0, 0, 0]
  cols = [0, 0, 0, 0, 0]
  while i < 25
    rows[i / 5] += board[i]
    cols[i % 5] += board[i]
    i += 1
  end

  # see if there are any winning rows
  return true if rows.include?(-5) || cols.include?(-5)

  false
end

# solve part 1
def part1(calls, boards)
  # for each number in the call
  calls.map do |c|
    # start at the end so we don't mess up our index if we delete boards
    (boards.length - 1).downto(0) do |i|
      # replace the called number with -1
      boards[i].map! { |x| x == c ? -1 : x }
      # check for winner
      # remove all the -1 spaces and sum it up
      return c * boards[i].reject! { |x| x == -1 }.sum if winner(boards[i])
    end
  end
end

# solve part 2
def part2(calls, boards)
  # for each number in the call
  calls.map do |c|
    # start at the end so we don't mess up our index if we delete boards
    (boards.length - 1).downto(0) do |i|
      # replace the called number with -1
      boards[i].map! { |x| x == c ? -1 : x }
      # check for winner
      next unless winner(boards[i])
      # if there is ony one board left, it's the last one
      # remove all the -1 spaces and sum it up
      return c * boards[i].reject! { |x| x == -1 }.sum if boards.length == 1

      boards.delete_at(i)
    end
  end
end

# read in file, split on double newlines
input = File.read('input').split("\n\n")

# convert all the bingo calls to ints
calls = input[0].split(',').map(&:to_i)

# empty board array
boards = []

# populate boards with a bunch of 25 length number arrays
input[1..].map do |e|
  boards << e.split.map(&:to_i)
end

puts format('Part 1: %d', part1(calls, boards))

boards = []
input[1..].map do |e|
  boards << e.split.map(&:to_i)
end

puts format('Part 2: %d', part2(calls, boards))
