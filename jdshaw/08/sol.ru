# frozen_string_literal: true

# perform the substitution
# split into an array
# and put each element in array in alphabetical order
def clean_strings(str, key)
  str.gsub(/[abcdefg]/,
           'a' => key[0],
           'b' => key[1],
           'c' => key[2],
           'd' => key[3],
           'e' => key[4],
           'f' => key[5],
           'g' => key[6]).split.map { |x| x.chars.sort.join }
end

input = File.read('input').split("\n")

# if translated correctly, we get these strings
solved = %w[abcefg cf acdeg acdfg bcdf abdfg abdefg acf abcdefg abcdfg]

# build lookup relating string to digit 0-9
lookup = Hash[solved.collect { |v| [v, solved.index(v)] }]

# now sort for easy compare
solved.sort!

# generate all the possible combinations
leds = [*'a'..'g']
keys = leds.permutation(7).map(&:join)

count = 0
sum = 0

# for each line of input
input.each do |line|
  # get the code and the output
  code, output = line.split(' | ')

  # try each possible key to see if it
  # results in the solution
  keys.each do |key|
    # get all the chars
    key = key.chars

    # to minimize testing, lets only look at
    # ones that will contain cf or fc
    c = leds[key.index('c')]
    f = leds[key.index('f')]
    s = code.split
    next unless (s.include? c + f) || (s.include? f + c)

    # use the key to decode, standardize, and sort
    decoded = clean_strings(code, key).sort

    # if it matches what we want
    next unless decoded == solved

    # use the key to decode the output
    decoded = clean_strings(output, key)

    # calculate sum for part 2
    sum += lookup[decoded[0]] * 1000 +
           lookup[decoded[1]] * 100 +
           lookup[decoded[2]] * 10 +
           lookup[decoded[3]]

    # count specific numbers for part 1
    count += decoded.count(lookup.key(1))
    count += decoded.count(lookup.key(4))
    count += decoded.count(lookup.key(7))
    count += decoded.count(lookup.key(8))
  end
end

puts format('Part 1: %d', count)
puts format('Part 2: %d', sum)
