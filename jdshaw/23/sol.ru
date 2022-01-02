# frozen_string_literal: true

# gross giant function for game rules and not making stupid moves
def can_move(state, value, start, stop)
  # if we are starting in a hall, we need to move to a room ( > 10)
  if start <= 10
    return 0 if stop <= 10

    # depending on its value, we know what entry we need to go to
    # and we can see if there is room for it in the room
    # we only move into a room if it is empty or filled with only
    # the same number
    case value
    when 1
      return 0 unless [11, 12].include?(stop)

      path = if start < 2
               state[start + 1..2]
             else
               state[2..start - 1]
             end
      return 0 if path.sum != 0
      return path.length + 1 if stop == 11 && (state[11]).zero? && state[12] == 1
      return path.length + 2 if stop == 12 && (state[11]).zero? && (state[12]).zero?
    when 10
      return 0 unless [13, 14].include?(stop)

      path = if start < 4
               state[start + 1..4]
             else
               state[4..start - 1]
             end
      return 0 if path.sum != 0
      return path.length + 1 if stop == 13 && (state[13]).zero? && state[14] == 10
      return path.length + 2 if stop == 14 && (state[13]).zero? && (state[14]).zero?
    when 100
      return 0 unless [15, 16].include?(stop)

      path = if start < 6
               state[start + 1..6]
             else
               state[6..start - 1]
             end
      return 0 if path.sum != 0
      return path.length + 1 if stop == 15 && (state[15]).zero? && state[16] == 100
      return path.length + 2 if stop == 16 && (state[15]).zero? && (state[16]).zero?
    else
      return 0 unless [17, 18].include?(stop)

      path = if start < 8
               state[start + 1..8]
             else
               state[8..start - 1]
             end
      return 0 if path.sum != 0
      return path.length + 1 if stop == 17 && (state[17]).zero? && state[18] == 1000
      return path.length + 2 if stop == 18 && (state[17]).zero? && (state[18]).zero?
    end
  else
    # otherwise we are starting in a room and are only going to move to the hall
    # technically we can also move to another room, but that's the same as moving
    # to the hall right next to the room and evaluating again to move into that room
    # so to keep it simple we only move to a hall from a room

    # can't move in entry way
    return 0 if stop > 10 || stop == 2 || stop == 4 || stop == 6 || stop == 8
    # dont move if we are already in our room and aren't blocking anyone
    return 0 if value == 1 && (start == 12 || (start == 11 && state[12] == 1))
    return 0 if value == 10 && (start == 14 || (start == 13 && state[14] == 10))
    return 0 if value == 100 && (start == 16 || (start == 15 && state[16] == 100))
    return 0 if value == 1000 && (start == 18 || (start == 17 && state[18] == 1000))

    # check which space is our door and calculate path to it
    door = 2 if [11, 12].include?(start)
    door = 4 if [13, 14].include?(start)
    door = 6 if [15, 16].include?(start)
    door = 8 if [17, 18].include?(start)
    path = if stop < door
             state[stop..door]
           else
             state[door..stop]
           end
    # can't move if someone in the way
    return 0 if path.sum != 0
    # return the length as long as no one is blocking us in the room
    return path.length if [11, 13, 15, 17].include?(start)
    return path.length + 1 if start == 12 && (state[11]).zero?
    return path.length + 1 if start == 14 && (state[13]).zero?
    return path.length + 1 if start == 16 && (state[15]).zero?
    return path.length + 1 if start == 18 && (state[17]).zero?
  end
  0
end

# gross giant function for game rules and not making stupid moves...now bigger
def can_move2(state, value, start, stop)
  # if we are starting in a hall, we need to move to a room ( > 10)
  if start <= 10
    return 0 if stop <= 10

    # depending on its value, we know what entry we need to go to
    # and we can see if there is room for it in the room
    # we only move into a room if it is empty or filled with only
    # the same number
    case value
    when 1
      return 0 unless [11, 12, 13, 14].include?(stop)

      path = if start < 2
               state[start + 1..2]
             else
               state[2..start - 1]
             end
      return 0 if path.sum != 0

      path2 = state[11..stop]
      return 0 if path2.sum != 0
      return path.length + path2.length if stop == 14

      filled = state[stop + 1..14]
      return path.length + path2.length if filled.sum == filled.length
    when 10
      return 0 unless [15, 16, 17, 18].include?(stop)

      path = if start < 4
               state[start + 1..4]
             else
               state[4..start - 1]
             end
      return 0 if path.sum != 0

      path2 = state[15..stop]
      return 0 if path2.sum != 0
      return path.length + path2.length if stop == 18

      filled = state[stop + 1..18]
      return path.length + path2.length if filled.sum == filled.length * 10
    when 100
      return 0 unless [19, 20, 21, 22].include?(stop)

      path = if start < 6
               state[start + 1..6]
             else
               state[6..start - 1]
             end
      return 0 if path.sum != 0

      path2 = state[19..stop]
      return 0 if path2.sum != 0
      return path.length + path2.length if stop == 22

      filled = state[stop + 1..22]
      return path.length + path2.length if filled.sum == filled.length * 100
    else
      return 0 unless [23, 24, 25, 26].include?(stop)

      path = if start < 8
               state[start + 1..8]
             else
               state[8..start - 1]
             end
      return 0 if path.sum != 0

      path2 = state[23..stop]
      return 0 if path2.sum != 0
      return path.length + path2.length if stop == 26

      filled = state[stop + 1..26]
      return path.length + path2.length if filled.sum == filled.length * 1000
    end
  else
    # otherwise we are starting in a room and are only going to move to the hall
    # technically we can also move to another room, but that's the same as moving
    # to the hall right next to the room and evaluating again to move into that room
    # so to keep it simple we only move to a hall from a room

    # can't move in entry way
    return 0 if stop > 10 || [2, 4, 6, 8].include?(stop)
    # dont move if we are already in our room and aren't blocking anyone
    return 0 if value == 1 && (start == 14 || state[start + 1..14].sum == (14 - start))
    return 0 if value == 10 && (start == 18 || state[start + 1..18].sum == (18 - start) * 10)
    return 0 if value == 100 && (start == 22 || state[start + 1..22].sum == (22 - start) * 100)
    return 0 if value == 1000 && (start == 26 || state[start + 1..26].sum == (18 - start) * 1000)

    # check which space is our door and calculate path to it
    door = 2 if [11, 12, 13, 14].include?(start)
    door = 4 if [15, 16, 17, 18].include?(start)
    door = 6 if [19, 20, 21, 22].include?(start)
    door = 8 if [23, 24, 25, 26].include?(start)
    path = if stop < door
             state[stop..door]
           else
             state[door..stop]
           end
    # can't move if someone in the way
    return 0 if path.sum != 0

    # if we are at fron of room, we have a clear path
    return path.length if [11, 15, 19, 23].include?(start)

    # get the path out of the room, make sure it's clear
    path2 = state[11..start - 1] if [12, 13, 14].include?(start)
    path2 = state[15..start - 1] if [16, 17, 18].include?(start)
    path2 = state[19..start - 1] if [20, 21, 22].include?(start)
    path2 = state[23..start - 1] if [24, 25, 26].include?(start)
    return 0 if path2.sum != 0

    # return total distance travled
    return path.length + path2.length

  end
  0
end

# get only the letters from the input
input = File.read('input')
input = input.gsub(/\W/, '').chars
# hash to lookup energy values
e = { 'A' => 1, 'B' => 10, 'C' => 100, 'D' => 1000 }

# this is our final state for part 1
win = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 10, 10, 100, 100, 1000, 1000]
# this is our starting state
burrow = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, e[input[0]], e[input[4]], e[input[1]], e[input[5]], e[input[2]],
          e[input[6]], e[input[3]], e[input[7]]]

# states will track states we need to examine for future moves
states = [burrow]

# best energy will track the lowest energy to get to a state
best_energy = Hash.new(99_999_999_999_999_999)

# starting state had 0 energy
best_energy[burrow] = 0

# while there are states to examine
while states.any?
  # pick one off
  s = states.shift

  # try and move from every spot to every spot
  0.upto(s.length - 1) do |start|
    # if nothing is in the starting spot, move on
    next if (s[start]).zero?

    0.upto(s.length - 1) do |stop|
      # check if a valid and smart move
      next if s[stop] != 0

      distance = can_move(s, s[start], start, stop)
      next if distance.zero?

      # make a new state for the move, adjusting start and stop
      new_state = s.clone
      new_state[stop] = new_state[start]
      new_state[start] = 0
      # calculate the energy
      energy = best_energy[s] + (s[start] * distance)
      # if its better, update energy tracker and add new state to
      # list to be examined
      if energy < best_energy[new_state]
        best_energy[new_state] = energy
        states << new_state
      end
    end
  end
end

puts format('Part 1: %d', best_energy[win])

# this is our final state for part 2
win = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 10, 10, 10, 10, 100, 100, 100, 100, 1000, 1000, 1000, 1000]
# this is our starting state with added input
burrow = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, e[input[0]], 1000, 1000, e[input[4]], e[input[1]], 100, 10, e[input[5]],
          e[input[2]], 10, 1, e[input[6]], e[input[3]], 1, 100, e[input[7]]]

# states will track states we need to examine for future moves
states = [burrow]

best_energy = Hash.new(99_999_999_999_999_999)

# starting state had 0 energy
best_energy[burrow] = 0

# while there are states to examine
while states.any?
  # pick one off
  s = states.shift

  # try and move from every spot to every spot
  0.upto(s.length - 1) do |start|
    # if nothing is in the starting spot, move on
    next if (s[start]).zero?

    0.upto(s.length - 1) do |stop|
      # check if a valid and smart move
      next if s[stop] != 0

      distance = can_move2(s, s[start], start, stop)
      next if distance.zero?

      # make a new state for the move, adjusting start and stop
      new_state = s.clone
      new_state[stop] = new_state[start]
      new_state[start] = 0
      # calculate the energy
      energy = best_energy[s] + (s[start] * distance)
      # if its better, update energy tracker and add new state to
      # list to be examined
      if energy < best_energy[new_state]
        best_energy[new_state] = energy
        states << new_state
      end
    end
  end
end

puts format('Part 2: %d', best_energy[win])
