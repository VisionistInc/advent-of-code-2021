# frozen_string_literal: true

# scan the dimensions for the cubes
# if none overlap, we are good
def no_overlap(fixed, carve)
  0.upto(2) do |i|
    return true if carve[i][1] < fixed[i][0] ||
                   carve[i][0] > fixed[i][1]
  end

  false
end

# get any pieces of carved cube that don't exist in fixed
def get_unique_cubes(fixed, carve)
  fx, fy, fz = fixed
  cx, cy, cz = carve
  new_cubes = []

  # do x first
  if cx[0] < fx[0]
    # there is some x overlap, so there *may* be overlap
    # lets carve off the part we are sure doesn't overlap

    # copy the existing large cube
    cube = [cx, cy, cz].map(&:clone)
    # set its x end to the start of the fixed
    cube[0][1] = fx[0] - 1
    # add to list
    new_cubes << cube
    # adjust large cube to be smaller
    cx[0] = fx[0]
  end
  if cx[1] > fx[1]
    # carved part ends outside of fixed part, lets carve off the
    # part that is outside of fixed, we know it doesn't overlap

    # copy the existing large cube
    cube = [cx, cy, cz].map(&:clone)
    # set its x start to the end of the fixed
    cube[0][0] = fx[1] + 1
    # add to list
    new_cubes << cube
    # adjust large cube to be smaller
    cx[1] = fx[1]
  end

  # repeat with y and z
  if cy[0] < fy[0]
    cube = [cx, cy, cz].map(&:clone)
    cube[1][1] = fy[0] - 1
    new_cubes << cube
    cy[0] = fy[0]
  end
  if cy[1] > fy[1]
    cube = [cx, cy, cz].map(&:clone)
    cube[1][0] = fy[1] + 1
    new_cubes << cube
    cy[1] = fy[1]
  end
  if cz[0] < fz[0]
    cube = [cx, cy, cz].map(&:clone)
    cube[2][1] = fz[0] - 1
    new_cubes << cube
    cz[0] = fz[0]
  end
  if cz[1] > fz[1]
    cube = [cx, cy, cz].map(&:clone)
    cube[2][0] = fz[1] + 1
    new_cubes << cube
    cz[1] = fz[1]
  end

  new_cubes
end

# read in input
input = File.read('input').split("\n")

# track the state of every x,y,z
cubes = {}

input.each do |line|
  # on/off and cube dimensions
  on_off, dimensions = line.split

  # parse out the dimensions
  x_dim, y_dim, z_dim = dimensions.gsub(/[x-z]=/, '').split(',')
  x_dim = x_dim.split('..').map(&:to_i)
  y_dim = y_dim.split('..').map(&:to_i)
  z_dim = z_dim.split('..').map(&:to_i)
  # assuming no cubes start outside of -50/50 and come into it
  next if x_dim[0] < -50 || x_dim[0] > 50

  # set to true if on, false otherwise
  flag = on_off == 'on'
  x_dim[0].upto(x_dim[1]) do |x|
    y_dim[0].upto(y_dim[1]) do |y|
      z_dim[0].upto(z_dim[1]) do |z|
        cubes[[x, y, z]] = flag
      end
    end
  end
end

puts format('Part 1: %d', cubes.values.count(true))

# well that's not going to work for part 2,
# so lets carve cubes up

# array to track current cubes that are on
cubes = []

input.each do |line|
  # on/off and cube dimensions
  on_off, dimensions = line.split

  # parse out the dimensions
  x, y, z = dimensions.gsub(/[x-z]=/, '').split(',')
  x = x.split('..').map(&:to_i)
  y = y.split('..').map(&:to_i)
  z = z.split('..').map(&:to_i)

  # if the switch is on, we want to add any part of this cube to the
  # cube list that doesn't overlap an existing on cube
  if on_off == 'on'
    # track any new cubes we will add to cubes
    new_cubes = [[x, y, z]]
    cubes.each do |c|
      cube_parts = []
      new_cubes.each do |nc|
        if no_overlap(c, nc)
          cube_parts << nc
        else
          cube_parts += get_unique_cubes(c, nc)
        end
      end
      new_cubes = cube_parts
    end
    cubes += new_cubes
  # if the switch is off we want to remove any part of this cube that
  # overlaps any cube in the cube list
  else
    new_cubes = []
    cubes.each do |c|
      if no_overlap([x, y, z], c)
        new_cubes << c
      else
        new_cubes += get_unique_cubes([x, y, z], c)
      end
    end
    cubes = new_cubes
  end
end

# sum up the dimensions of each cube that is on
total = 0
cubes.each do |c|
  x, y, z = c
  total += ((x[1] - x[0] + 1) * (y[1] - y[0] + 1) * (z[1] - z[0] + 1))
end

puts format('Part 2: %d', total)
