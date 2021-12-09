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
    i = boards.length - 1
    while i > -1
      # replace the called number with -1
      boards[i].map! { |x| x == c ? -1 : x }
      # check for winner
      if winner(boards[i])
        # remove all the -1 spots
        boards[i].reject! { |x| x == -1 }
        return boards[i].sum * c

      end
      i -= 1
    end
  end
end

# solve part 2
def part2(calls, boards)
  # for each number in the call
  calls.map do |c|
    # start at the end so we don't mess up our index if we delete boards
    i = boards.length - 1
    while i > -1
      # replace the called number with -1
      boards[i].map! { |x| x == c ? -1 : x }
      # check for winner
      if winner(boards[i])
        # if there is ony one board left, it's the last one
        if boards.length == 1
          # remove all the -1 spots
          boards[i].reject! { |x| x == -1 }
          return boards[i].sum * c

        else
          # not the last board, delete
          boards.delete_at(i)
        end
      end
      i -= 1
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
