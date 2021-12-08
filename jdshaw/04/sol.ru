# janky function to determine if board is a winner
# all marked off squares get turned into -1
def winner(board)
  # see if any rows have all -1
  i = 0
  while i < 25
    return true if board[i..i + 4].sum == -5

    i += 5
  end

  # see if any columns have -1
  i = 0
  while i < 5
    j = 0
    sum = 0
    while j < 25
      sum += board[j] if j % 5 == i
      j += 1
    end
    return true if sum == -5

    i += 1
  end

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
        # replace all -1 with 0 to not mess up calculations
        boards[i].map! { |x| x == -1 ? 0 : x }
        puts format('Part 1: %s', boards[i].sum * c)
        return
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
          # replace all -1 with 0 to not mess up calculations
          boards[i].map! { |x| x == -1 ? 0 : x }
          puts format('Part 2 %d', boards[i].sum * c)
          return
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
input[1..-1].map do |e|
  boards << e.split.map(&:to_i)
end

part1(calls, boards)

boards = []
input[1..-1].map do |e|
  boards << e.split.map(&:to_i)
end

part2(calls, boards)
