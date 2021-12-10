# frozen_string_literal: true

input = File.read('input').split("\n")

# helpful lookups
opening = '([{<'
closings = { '(' => ')', '[' => ']', '<' => '>', '{' => '}' }
syntax_points = { ')' => 3, ']' => 57, '>' => 25_137, '}' => 1197 }
complete_points = { ')' => 1, ']' => 2, '>' => 4, '}' => 3 }

# score trackers
pt1_points = 0
pt2_points = []

# for each line
input.each do |line|
  # get all the characters
  chars = line.split('')
  # track order of expected closing characters
  expected = []
  # for each character
  while chars.any?
    # pull it off the list
    c = chars.shift
    # if it's an opening, add it's closing to our tracking list
    if opening.include? c
      expected.push(closings[c])
    # if it's a closing and doesn't match what we expect
    elsif expected.pop != c
      # grab the points and break out
      pt1_points += syntax_points[c]
      break
    end
  end
  # if we have chars left, line was corrupted and can be ignored
  next unless chars.empty?

  points = 0
  # calculate score in order
  expected.reverse.each do |x|
    points = (points * 5) + complete_points[x]
  end
  pt2_points << points
end

puts format('Part 1: %d', pt1_points)

pt2_points.sort!
puts format('Part 2: %d', pt2_points[pt2_points.length / 2])
