# read in file, split on new lines
input = File.read('input').split

# used to hold how many bits there are in one entry
positions = []

# add a blank string for every bit
input[0].chars.map do
  positions << ''
end

# iterate over every line
input.map do |e|
  # reset to bit 0 each line
  i = 0
  # for each character in the line, add it to the corresponding bit collector
  e.chars.map do |c|
    positions[i] += c
    # increment bit pointer
    i += 1
  end
end

# create the beginnign of a binary string
gamma = '0b'
epsilon = '0b'

# for each bit collector
positions.map do |e|
  # sort the bits
  e = e.chars.sort.join
  # if we have more 1s, add 1 to gamma
  if e.index('1') < (e.length / 2)
    gamma += '1'
    epsilon += '0'
  # otherwise add 0 to gamma
  else
    gamma += '0'
    epsilon += '1'
  end
end

# convert and multiply
puts format('Part 1: %d', Integer(gamma) * Integer(epsilon))

# make a copy in order to abuse
o2 = input.clone.sort

# create the start of the binary string
o2_str = '0b'

# while we have more than one entry
while o2.length > 1
  # check what bit is at one more than half
  v = o2[o2.length / 2][0]
  # add that bit to the final string
  o2_str += v
  # only keep words that start with that bit
  o2 = o2.select { |word| word.start_with?(v) }

  # remove the leading bit so we can compare again
  new_o2 = []
  o2.map do |e|
    new_o2 << e[1..ruby]
  end
  o2 = new_o2
end

# add back on any extra bit that were remaining in the last string
o2_str += o2[0]

# make a copy in order to abuse
co2 = input.clone.sort

# create the start of the binary string
co2_str = '0b'

# while we have more than one entry
while co2.length > 1
  # check what bit is at one more than half
  v = co2[co2.length / 2][0]
  # we want the opposite value added to our string
  v = v == '1' ? '0' : '1'
  co2_str += v
  # only keep words that start with that opposite bit
  co2 = co2.select { |word| word.start_with?(v) }

  # remove the leading bit so we can compare again
  new_co2 = []
  co2.map do |e|
    new_co2 << e[1..]
  end
  co2 = new_co2
end

# add back on any extra bit that were remaining in the last string
co2_str += co2[0]

puts format('Part 2: %d', Integer(o2_str) * Integer(co2_str))
