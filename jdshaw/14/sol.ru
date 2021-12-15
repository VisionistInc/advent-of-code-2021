# frozen_string_literal: true

# Each key in the Hash @pairs is 2 chars.  Create a new Hash to track chars.
# Add the value from @pairs to each char key.
def count_chars(pairs)
  # track counts of individual characters
  counts = Hash.new(0)
  # for each XY key
  pairs.each do |k, v|
    # add value to both X and Y
    counts[k[0]] += v
    counts[k[1]] += v
  end

  counts
end

# Determine how many of each character show up in @pairs and then subtract
# the least frequent from the most frequent.
def sub_max_min(pairs, template)
  # get hash for each individual character hit rate
  counts = count_chars(pairs)

  # every character is doubled since it got added twice
  # example AB -> C results in AB, BC, so Bshows up twice
  # the only exception is the first and last character, so we plus them up
  # to match all the other characters
  counts[template[0]] += 1
  counts[template[-1]] += 1

  # now we divide every count by 2, since that is the real hit rate
  counts.each { |k, v| counts[k] = v / 2 }

  # return difference
  counts.values.max - counts.values.min
end

# Given a hash @pairs of pair occurrences
# generate a new hash of pairs according to @rules hash
def generate_all_pairs(pairs, rules)
  # create temp hash to track newly created pairs
  new_pairs = Hash.new(0)
  # each pair creates two new pairs.
  # if we have 5 AB, we add 5 to AC and 5 to CB
  pairs.each do |k, v|
    new_pairs[rules[k][0]] += v
    new_pairs[rules[k][1]] += v
  end

  new_pairs
end

# Solve the problem.  Track pairs and how often they occur.  Each pair will
# generate two new pairs based on @rules.
def solve(template, rules, steps)
  # create hash to track pair occurrences
  pairs = Hash.new(0)

  # seed initially with the pairs from the template in input
  template.chars.each_cons(2) { |a, b| pairs[a + b] += 1 }

  # use the current pairs to generate new pairs for the desired number of steps
  1.upto(steps) do
    pairs = generate_all_pairs(pairs, rules)
  end

  # get the difference in hit rates needed
  sub_max_min(pairs, template)
end

# gather input
input = File.read('input')

# get parts of input
template, rules = input.split("\n\n")

# all the rules are on their own line
rules = rules.split("\n")
# this hash will hold the two pairs one pair will become
xfer = {}
rules.each do |x|
  pair, mid = x.split(' -> ')
  xfer[pair] = [x[0] + mid, mid + x[1]]
end

puts format('Part 1: %d', solve(template, xfer, 10))
puts format('Part 2: %d', solve(template, xfer, 40))
