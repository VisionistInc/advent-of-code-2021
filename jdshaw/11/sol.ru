# frozen_string_literal: true

# check the validity of a spot
def valid_spot(row, col)
  return true if row >= 0 && row < 10 && col >= 0 && col < 10

  false
end

# add one to the octopus and make note of any neighbors
# that got hit with a flash from it
# return 1 if a flash occurred, 0 otherwise
def add_energy(input, neighbors, row, col)
  # increment
  input[row][col] += 1
  # check for flash
  return 0 unless input[row][col] == 10

  # iterate over neighbors
  (row - 1).upto(row + 1) do |y|
    (col - 1).upto(col + 1) do |x|
      # don't add the spot itself
      neighbors << [y, x] unless (y == row && x == col) || !valid_spot(y, x)
    end
  end
  1
end

input = File.read('input').split("\n").map! { |x| x.split('').map(&:to_i) }

flashes = 0
steps = 0

loop do
  # increase steps
  steps += 1
  # track neighbors to add more to after flash
  neighbors = []
  # increment everything by 1, making note of flashes
  input.each_with_index do |row, y|
    row.each_with_index do |_val, x|
      flashes += add_energy(input, neighbors, y, x)
    end
  end

  # now see what spots get bumps from neighbor flashes
  while neighbors.any?
    y, x = neighbors.shift
    flashes += add_energy(input, neighbors, y, x)
  end

  # now clear anything greater than 9
  input.each do |row|
    row.map! { |v| v > 9 ? 0 : v }
  end

  puts format('Part 1: %d', flashes) if steps == 100

  # is everyone 9?
  break if input.flatten.inject(:+).zero?
end

puts format('Part 2: %d', steps)
