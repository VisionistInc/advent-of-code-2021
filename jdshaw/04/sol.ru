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

# replace any spots that match num with -1
def mark_spots(board, num)
  board.map! { |x| x == num ? -1 : x }
end

# solve part 1
def part1(calls, boards)
  # for each number in the call
  calls.each do |c|
    # mark spots on all boards
    boards.map! { |board| mark_spots(board, c) }
    # check all boards for winners
    boards.each do |board|
      # remove all the -1 spaces and sum it up
      return c * board.reject! { |x| x == -1 }.sum if winner(board)
    end
  end
end

# solve part 2
def part2(calls, boards)
  # for each number in the call
  calls.each do |c|
    # mark spots on all boards
    boards.map! { |board| mark_spots(board, c) }
    # check each board
    boards.each do |board|
      # check for winner
      next unless winner(board)
      # if there is ony one board left, it's the last one
      # remove all the -1 spaces and sum it up
      return c * board.reject! { |x| x == -1 }.sum if boards.length == 1

      boards.delete(board)
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
