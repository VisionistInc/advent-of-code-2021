# frozen_string_literal: true

# hackey way of converting pixel map to number
def translate(image, x, y)
  total = 0
  total += 256 if image[[x - 1, y - 1]]
  total += 128 if image[[x, y - 1]]
  total += 64 if image[[x + 1, y - 1]]
  total += 32 if image[[x - 1, y]]
  total += 16 if image[[x, y]]
  total += 8 if image[[x + 1, y]]
  total += 4 if image[[x - 1, y + 1]]
  total += 2 if image[[x, y + 1]]
  total += 1 if image[[x + 1, y + 1]]
  total
end

# read in input, split on algorithm/image
algo, image = File.read('input').split("\n\n")

# create a lookup for the algorithm
algo = algo.chars.map { |x| (x == '#') }

# So, one would think that if you had a cluster of 9 dark spots, it would
# remain dark.  But that's not the case with my input, it turns bright, and
# the cluster of 9 bright turn dark.  I guess this is ok if we are doing an
# even number of transformations since all the infinite dark will turn light
# and dark again.  I don't know how to handle the case where they turn light
# and stay light.  Maybe you only count lights in the image area, dunno.
flip = false
default = false
if algo[0]
  if algo[511]
    puts "I don't know what to do with infinite light.  Aborting!"
    return
  end
  flip = true
end

# translate into array of arrays with mappings to boolean
image = image.split.map { |x| x.chars.map { |y| (y == '#') } }

# save off image dimensions for reference
width = image[0].length
height = image.length

# we will track pixels in our hash
# allows us to create a default for the infinity pixels
pixels = Hash.new(default)
0.upto(height - 1) do |y|
  0.upto(width - 1) do |x|
    pixels[[x, y]] = image[y][x]
  end
end

# we need to evaluate one row of pixels beyound the current image
# because the current image can affect it
y_start = -1
y_end = height
x_start = -1
x_end = width

# enhance 50 times!
1.upto(50) do |i|
  # create a new enhanced pixle tracker, updating default if necessary
  default = !default if flip
  enhanced = Hash.new(default)

  # map old pixels to new ones
  x_start.upto(x_end) do |x|
    y_start.upto(y_end) do |y|
      enhanced[[x, y]] = algo[translate(pixels, x, y)]
    end
  end

  # save off enhanced image
  pixels = enhanced

  # grow the area we evaluate
  y_start -= 1
  y_end += 1
  x_start -= 1
  x_end += 1

  puts format('Part 1: %d', pixels.values.count(true)) if i == 2
end

puts format('Part 2: %d', pixels.values.count(true))
