# frozen_string_literal: true

# gather input
input = File.read('input')
# split out dots and folds
dots, folds = input.split("\n\n")
# covert dots to number pairs
dots = dots.split.map { |x| x.split(',').map(&:to_i) }
# convert folds to direction and number
folds = folds.split("\n").map { |x| x.split('=') }

# for part 1
first = true

# for each fold
folds.each do |axis, line|
  # convert the fold line to a number
  line = line.to_i
  # iterate over each dot, adjusting its value based on fold
  dots.each do |dot|
    if axis.include? 'y'
      dot[1] = 2 * line - dot[1] if dot[1] > line
    elsif dot[0] > line
      dot[0] = 2 * line - dot[0]
    end
  end
  # remove any duplicates that now exist
  dots.uniq!
  # print part 1 results after first pass
  if first
    puts format('Part 1: %d', dots.length)
    first = false
  end
end

# find the max value for each direction to bound our pt2 output
max_x = dots.map { |x| x[0] }.max
max_y = dots.map { |y| y[1] }.max

# print out display
puts 'Part 2:'
0.upto(max_y) do |y|
  0.upto(max_x) do |x|
    print(dots.include?([x, y]) ? '#' : ' ')
  end
  puts ''
end
