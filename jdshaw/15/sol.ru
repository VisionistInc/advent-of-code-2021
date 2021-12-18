# frozen_string_literal: true

# generate the adjacent spots, checking for map boundries
def next_spots(row, col, max_row, max_col)
  next_list = []
  next_list << [row - 1, col] if row.positive?
  next_list << [row + 1, col] if row < max_row
  next_list << [row, col + 1] if col < max_col
  next_list << [row, col - 1] if col.positive?
  next_list
end

# iterate over all the spots and generate best risk for that spot
def generate_risk(risk_map, risk, end_y, end_x)
  # spots will contain spots we've visited.  start at 0,0
  spots = [[0, 0]]

  # while we have spots to visit
  while spots.any?
    # pull out next spot
    row, col = spots.shift
    # generate new spots to visit
    next_spots(row, col, end_y, end_x).each do |y, x|
      # caluclate new risk and the next spot, if lower,
      # update and add to list of spots to consider
      new_risk = risk[row][col] + risk_map[y][x]
      next if new_risk >= risk[y][x]

      risk[y][x] = new_risk
      spots << [y, x]
    end
  end
end

def solve(risk_map)
  # find the edges of the map
  end_x = risk_map[0].length - 1
  end_y = risk_map.length - 1

  # generate map that holds the optimal risk to get to that spot
  # default each spot to a crazy hgih risk so that any risk
  # will be lower than it
  risk = Array.new(end_y + 1) { Array.new(end_x + 1, (end_x + end_y + 2) * 9) }
  # starting point always risk 0
  risk[0][0] = 0

  generate_risk(risk_map, risk, end_y, end_x)

  risk[end_y][end_x]
end

# gather the map, convert to ints
risk_map = File.read('input').split.map { |x| x.chars.map(&:to_i) }

puts format('Part 1: %d', solve(risk_map))

# make the larget map for part 2
large_risk_map = []

# first grow each row to be 5 times longer
risk_map.each do |row|
  new_row = [].append(*row)
  1.upto(4) do
    row.map! { |n| n < 9 ? n + 1 : 1 }
    new_row.append(*row)
  end
  large_risk_map.append(new_row)
end

# now generate the extra rows
tmp_map = []
1.upto(4) do |i|
  large_risk_map.each do |row|
    tmp_map.append(row.map { |n| n + i < 10 ? n + i : n + i - 9 })
  end
end
# and append to the large map
large_risk_map.append(*tmp_map)

puts format('Part 2: %d', solve(large_risk_map))
