# frozen_string_literal: true

XMIN = 0
XMAX = 1
YMIN = 2
YMAX = 3

# see if x,y is in the trench
def in_range(x_pos, y_pos, trench)
  x_pos >= trench[XMIN] &&
    x_pos <= trench[XMAX] &&
    y_pos >= trench[YMIN] &&
    y_pos <= trench[YMAX]
end

# determine is the velocity provided will land in the trench
def hit(x_vel, y_vel, trench)
  # our first position after 0,0 is the velocity
  x = x_vel
  y = y_vel
  # while we haven't overshot the trench
  while x <= trench[XMAX] && y >= trench[YMIN]
    # return 1 if we are currently in the trench
    return 1 if in_range(x, y, trench)

    # otherwise adjust the velocities and positions
    x_vel -= 1 if x_vel.positive?
    y_vel -= 1
    x += x_vel
    y += y_vel
  end

  0 # we overshot the trench, so return 0
end

# read in file and parse out trench range
parts = File.read('input').split('=')
trench = parts[1].split(',')[0].split('..').map(&:to_i)
trench.append(*parts[2].split('..').map(&:to_i))

# the highest point we can ever get is constrained by the lowest point
# in the trench.  Want to hit the bottom of the trench from y=0, so our
# velocity tp get to y=0 should be one less than the distance to the bottom of
# the trench.  Then mas height is the divergent series
y_vel = trench[YMIN].abs - 1
puts format('Part 1: %d', (((y_vel + 1) * y_vel) / 2))

# now collect how many different velocities land in the trench
count = 0

# we can't go past x_max because our first shot will overshoot
1.upto(trench[XMAX]) do |x|
  # we can't shoow lower than y_min for the same reason
  # and we already know the highest we can ever aim from part 1
  trench[YMIN].upto(y_vel) do |y|
    count += hit(x, y, trench)
  end
end

puts format('Part 2: %d', count)
