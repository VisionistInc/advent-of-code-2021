# frozen_string_literal: true

def roll_training(dice)
  sum = dice
  dice = dice == 100 ? 1 : dice + 1
  sum += dice
  dice = dice == 100 ? 1 : dice + 1
  sum += dice
  dice = dice == 100 ? 1 : dice + 1
  [sum, dice]
end

# read in input
input = File.read('input').split("\n")

# grab out starting positions
p1 = input[0].split[-1].to_i
p2 = input[1].split[-1].to_i

# set initial score
p1_score = 0
p2_score = 0

# set first roll value
dice = 1
num_rolls = 0

# play until someone wins
loop do
  ### Do player 1 first

  # roll the dice, get your move and new dice
  sum, dice = roll_training(dice)
  # count rolls
  num_rolls += 3
  # determine new position
  p1 += sum
  p1 -= 10 while p1 > 10
  # add position to score
  p1_score += p1
  # see if we won
  puts format('Part 1: %d', p2_score * num_rolls) if p1_score >= 1000

  ### Do player 2 next

  # roll the dice, get your move and new dice
  sum, dice = roll_training(dice)
  # count rolls
  num_rolls += 3
  # determine new position
  p2 += sum
  p2 -= 10 while p2 > 10
  # add position to score
  p2_score += p2
  # see if we won
  if p2_score >= 1000
    puts format('Part 1: %d', p1_score * num_rolls)
    break
  end
end

# onto part 2

# generate a hit rate for roll values based on all permutations
counts = [1, 2, 3].repeated_permutation(3)
hit_rate = Hash.new(0)
counts.each do |x|
  hit_rate[x.sum] += 1
end

# get the initial positions again
p1 = input[0].split[-1].to_i
p2 = input[1].split[-1].to_i

# use a hash to track states [pos1, pos2, score1, score2] to reduce size
# count how many paths get us to that state
# seed with initial state
states = Hash.new(0)
states[[p1, p2, 0, 0]] = 1

# track wins
p1_wins = 0
p2_wins = 0

# while we have states to consider
while states.any?
  # pull one off
  state, count = states.shift
  # pull out useful values
  p1, p2, p1_score, p2_score = state

  # now for each possible roll p1 can make
  hit_rate.each do |k, v|
    # determine new position
    p1_next = k + p1
    p1_next -= 10 if p1_next > 10
    # determine new score
    p1_score_next = p1_score + p1_next
    # check for win
    if p1_score_next >= 21
      # if win, all current states * all possible p1 rolls led us to this win
      p1_wins += (count * v)
      next
    end

    # no win, so have p2 take all their turns
    hit_rate.each do |kk, vv|
      # determine new position
      p2_next = p2 + kk
      p2_next -= 10 if p2_next > 10
      # determine new score
      p2_score_next = p2_score + p2_next
      # check for win
      if p2_score_next >= 21
        # if win, all current states * all possible p1 rolls * all
        # possible p2 rolls led us to this win
        p2_wins += (count * v * vv)
        next
      end

      # no winners, so the number of paths that gets us to this next state
      # is the number to current state * all possible p1 rolls * all possible p2 rolls
      states[[p1_next, p2_next, p1_score_next, p2_score_next]] += (count * v * vv)
    end
  end
end

puts format('Part 2: %d', p1_wins > p2_wins ? p1_wins : p2_wins)
