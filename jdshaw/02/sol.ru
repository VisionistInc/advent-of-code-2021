# read in file, split on new lines
input = File.read('input').split("\n")

# init our position
depth = 0
distance = 0

# iterate over each direction, pulling number out from command
input.map do |e|
  p = e.split(' ')
  if p[0][0] == 'f'
    distance += p[1].to_i
  elsif p[0][0] == 'u'
    depth -= p[1].to_i
  else
    depth += p[1].to_i
  end
end

puts format('Part 1: %d', distance * depth)

# re-init our position
depth = 0
distance = 0
aim = 0

# iterate over each direction, pulling number out from command
input.map do |e|
  p = e.split(' ')
  if p[0][0] == 'f'
    distance += p[1].to_i
    depth += (aim * p[1].to_i)
  elsif p[0][0] == 'u'
    aim -= p[1].to_i
  else
    aim += p[1].to_i
  end
end

puts format('Part 2: %d', distance * depth)
