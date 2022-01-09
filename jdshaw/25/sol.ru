# frozen_string_literal: true

# check if a move east can occur
def do_move_east(sea_floor, new_floor, coords)
  x, y = coords
  xx = (x + 1) % @width
  if sea_floor[[xx, y]] == '.'
    new_floor[[xx, y]] = '>'
    return true
  else
    new_floor[[x, y]] = '>'
  end
  false
end

# check if a move south can occur
def do_move_south(sea_floor, new_floor, coords)
  x, y = coords
  yy = (y + 1) % @height
  if sea_floor[[x, yy]] == '.'
    new_floor[[x, yy]] = 'v'
    return true
  else
    new_floor[[x, y]] = 'v'
  end
  false
end

# debugging
def print_floor(floor)
  puts('-------')
  0.upto(@height - 1) do |y|
    0.upto(@width - 1) do |x|
      print floor[[x, y]]
    end
    puts ''
  end
end

# read in sea floor, get parameters
input = File.read('input').split.map(&:chars)
@height = input.length
@width = input[0].length

sea_floor = Hash.new('.')

# fill out a hash based on the sea_floor
0.upto(@width - 1) do |x|
  0.upto(@height - 1) do |y|
    sea_floor[[x, y]] = input[y][x]
  end
end

# track if we've moved
moved = true
steps = 0
while moved
  steps += 1
  # map of new floor
  new_floor = Hash.new('.')
  # reset our flag
  moved = false
  # try to move all east facing
  sea_floor.each do |k, v|
    case v
    when '>'
      moved |= do_move_east(sea_floor, new_floor, k)
    when 'v'
      new_floor[k] = v
    end
  end
  # now that they've moved, reset the maps
  sea_floor = new_floor
  new_floor = Hash.new('.')

  # try to move all south facing
  sea_floor.each do |k, v|
    case v
    when 'v'
      moved |= do_move_south(sea_floor, new_floor, k)
    when '>'
      new_floor[k] = v
    end
  end

  # reset the maps
  sea_floor = new_floor
end

puts steps
