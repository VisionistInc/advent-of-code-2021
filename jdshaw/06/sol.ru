# frozen_string_literal: true

###
# D00: 8
# D09: 6       8
# D16: 6 8     1
# D18  4 6     6 8
# D23  6 1     1 3
# D25  4 6 8   6 1 8

# If you look, starting at Day 9, The first fish created a second fish
# that makes the first fish responsible for a new fish 7 and 9 days from
# day 9, day 16 and 18

# so at day 16, we have the same situation, the fish there will be
# responsible for fish at day 23 and 25

# so at day 18, we have the same situation, the fish there will be
# responsible for fish at day 25 and 27

# so any fish that reproduces on a given day is responsible for
# two more fish 7 and 9 days later

# starting with one fish, how many will come
def count_fish(val, days)
  # not enough days to reproduce?
  return 1 unless val < days

  # create new dictionary with 0 default fish at each day
  fish = Hash.new(0)

  # populate first reproduction
  fish[val + 1] = 1

  # starting at day 0
  d = 0
  total = 0
  while d <= days
    # how many fish were born that day
    total += fish[d]

    # see above big comment
    fish[d + 7] += fish[d] unless d + 7 > days
    fish[d + 9] += fish[d] unless d + 9 > days
    d += 1
  end

  total + 1
end

# read in file, split on commas, and convert each to an integer
input = File.read('input').split(',').map(&:to_i)

total = 0
input.each do |x|
  total += count_fish(x, 80)
end

puts "Part 1: #{total}"

total = 0
input.each do |x|
  total += count_fish(x, 256)
end

puts "Part 2: #{total}"
