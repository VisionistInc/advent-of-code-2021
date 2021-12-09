# frozen_string_literal: true

def add_spots_to_group(input, group, row, col)
  group << [row - 1, col] if row.positive?
  group << [row + 1, col] if row < input.length - 1
  group << [row, col - 1] if col.positive?
  group << [row, col + 1] if col < input[row].length - 1
end

# measures how big a basin is by counting all the non-9
# spaces that touch each other
def measure_basin(input, group)
  size = 0

  # while we have spots in our group
  while group.length.positive?
    # grab the next spot
    row, col = group.shift

    # if it's 9, ignore it
    next unless input[row][col] != 9

    # increment our size and change it to 9
    # so we never consider it again
    size += 1
    input[row][col] = 9

    # add all the neighbors to the group, we don't care
    # what the value is, we'll catch next pass
    add_spots_to_group(input, group, row, col)
  end
  size
end

input = File.read('input').split("\n").map! { |x| x.split('').map(&:to_i) }

risk = 0
# iterate over every spot
0.upto(input.length - 1) do |row|
  0.upto(input[row].length - 1) do |col|
    # create a list of neighbors to compare with
    # make sure to do bounds checking
    compare = []
    compare << input[row - 1][col] if row.positive?
    compare << input[row + 1][col] if row < input.length - 1
    compare << input[row][col - 1] if col.positive?
    compare << input[row][col + 1] if col < input[row].length - 1
    # throw away any that the current spot is less than
    compare.reject! { |z| input[row][col] < z }
    # if none are left, the spot was the smallest
    risk += (1 + input[row][col]) if compare.none?
  end
end

puts format('Part 1: %d', risk)

# tracks all the basins found
basins = []

# iterate over all the spaces
0.upto(input.length - 1) do |row|
  0.upto(input[row].length - 1) do |col|
    # if the spot is not 9, see how large the basin is
    basins << measure_basin(input, [[row, col]]) if input[row][col] != 9
  end
end

# sort the list in descending order
basins.sort! { |a, b| b <=> a }

# multiply first three elements
puts format('Part 2: %d', basins[0..2].inject(:*))
