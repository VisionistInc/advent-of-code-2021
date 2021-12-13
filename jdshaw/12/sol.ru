# frozen_string_literal: true

input = File.read('input').split("\n")

# create useful lookups
# each neighbor is an empty array to start
neighbors = Hash.new { |h, k| h[k] = [] }
big_caves = []
small_caves = []

# parse input
input.each do |line|
  a, b = line.split('-')
  # assign neighbors
  neighbors[a] << b
  neighbors[b] << a

  # check cave sizes
  if a == a.upcase
    big_caves << a
  else
    small_caves << a
  end
  if b == b.upcase
    big_caves << b
  else
    small_caves << b
  end
end

# remove duplicate values
big_caves.uniq!
small_caves.uniq!

# let's go
paths = [['start']]
final_paths = []

# while we have paths to consider
while paths.any?
  # pull one out
  path = paths.shift
  # got the list of neighbor caves
  next_steps = neighbors[path[-1]]
  # for each neighbor
  next_steps.each do |x|
    # if it's th end, mark it
    if x == 'end'
      final_paths << path + [x]
    # if it's a big cave, we can travel it
    elsif big_caves.include? x
      paths << path + [x]
    # if it's a small cave, we can travel it as long as we never have before
    elsif !path.include? x
      paths << path + [x]
    end
  end
end

puts format('Part 1: %d', final_paths.length)

# let's go again
paths = [['start']]
final_paths = []

# while we have paths to consider
while paths.any?
  # pull one out
  path = paths.shift
  # got the list of neighbor caves
  next_steps = neighbors[path[-1]]
  # for each neighbor
  next_steps.each do |x|
    # if it's th end, mark it
    if x == 'end'
      final_paths << path + [x]
    # if it's a big cave, we can travel it
    elsif big_caves.include? x
      paths << path + [x]
    # if it's a small cave, we can travel it as long as we never have before
    elsif !path.include? x
      paths << path + [x]
    # else, as long as it's not start or end, we may be able to visit the
    # small cave
    elsif x != 'start' && x != 'end'
      # remove all big caves
      no_big = path.select { |a| small_caves.include? a }
      # if the list of small caves on the path is the same size with
      # dupes removed, then we can visit this small cave again
      paths << path + [x] if no_big.length == no_big.uniq.length
    end
  end
end

puts format('Part 2: %d', final_paths.length)
