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
    y, x = group.shift

    # if it's 9, ignore it
    next unless input[y][x] != 9

    # increment our size and change it to 9
    # so we never consider it again
    size += 1
    input[y][x] = 9

    # add all the neighbors to the group, we don't care
    # what the value is, we'll catch next pass
    add_spots_to_group(input, group, y, x)
  end
  size
end

input = File.read('input').split("\n").map! { |x| x.split('').map(&:to_i) }

risk = 0
# iterate over every spot
input.each_with_index do |row, y|
  row.each_with_index do |val, x|
    # create a list of neighbors to compare with
    # make sure to do bounds checking
    compare = []
    compare << input[y - 1][x] if y.positive?
    compare << input[y + 1][x] if y < input.length - 1
    compare << input[y][x - 1] if x.positive?
    compare << input[y][x + 1] if x < row.length - 1
    # throw away any that the current spot is less than
    compare.reject! { |z| val < z }
    # if none are left, the spot was the smallest
    risk += (1 + val) if compare.none?
  end
end

puts format('Part 1: %d', risk)

# tracks all the basins found
basins = []

# iterate over all the spaces
input.each_with_index do |row, y|
  row.each_with_index do |val, x|
    # if the spot is not 9, see how large the basin is
    basins << measure_basin(input, [[y, x]]) if val != 9
  end
end

# sort the list in descending order
basins.sort! { |a, b| b <=> a }

# multiply first three elements
puts format('Part 2: %d', basins[0..2].inject(:*))
