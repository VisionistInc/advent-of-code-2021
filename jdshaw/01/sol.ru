# iterate over neighbors in a given list, count when second element is larger
def count_greater(vals)
  count = 0
  vals.each_cons(2).map do |a, b|
    count += 1 unless b <= a
  end
  count
end

# read in file, split on new lines, and convert each to an integer
input = File.read('input').split.map(&:to_i)

# solve part 1
puts format('Part 1: %d', count_greater(input))

# create new list that is sum of three neighbors
sum = []
input.each_cons(3).map do |a, b, c|
  sum << (a + b + c)
end

# solve part 2
puts format('Part 2: %d', count_greater(sum))
