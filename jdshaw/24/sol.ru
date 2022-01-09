# ugh, i hand parsed the assembly into ruby to make it run faster
# i can do better than this
def solve14(w, z, z_out)
  x = z % 26
  z /= 1
  x += 10
  x = (x == w ? 0 : 1)
  y = 25 * x
  y += 1
  z *= y
  y = w + 5
  y *= x
  z += y
  z == z_out
end

def solve13(w, z, z_out)
  x = z % 26
  z /= 1
  x += 13
  x = (x == w ? 0 : 1)
  y = 25 * x
  y += 1
  z *= y
  y = w + 9
  y *= x
  z += y
  z == z_out
end

def solve12(w, z, z_out)
  x = z % 26
  z /= 1
  x += 12
  x = (x == w ? 0 : 1)
  y = 25 * x
  y += 1
  z *= y
  y = w + 4
  y *= x
  z += y
  z == z_out
end

def solve11(w, z, z_out)
  x = z % 26
  z /= 26
  x -= 12
  x = (x == w ? 0 : 1)
  y = 25 * x
  y += 1
  z *= y
  y = w + 4
  y *= x
  z += y
  z == z_out
end

def solve10(w, z, z_out)
  x = z % 26
  z /= 1
  x += 11
  x = (x == w ? 0 : 1)
  y = 25 * x
  y += 1
  z *= y
  y = w + 10
  y *= x
  z += y
  z == z_out
end

def solve9(w, z, z_out)
  x = z % 26
  z /= 26
  x -= 13
  x = (x == w ? 0 : 1)
  y = 25 * x
  y += 1
  z *= y
  y = w + 14
  y *= x
  z += y
  z == z_out
end

def solve8(w, z, z_out)
  x = z % 26
  z /= 26
  x -= 9
  x = (x == w ? 0 : 1)
  y = 25 * x
  y += 1
  z *= y
  y = w + 14
  y *= x
  z += y
  z == z_out
end

def solve7(w, z, z_out)
  x = z % 26
  z /= 26
  x -= 12
  x = (x == w ? 0 : 1)
  y = 25 * x
  y += 1
  z *= y
  y = w + 12
  y *= x
  z += y
  z == z_out
end

def solve6(w, z, z_out)
  x = z % 26
  z /= 1
  x += 14
  x = (x == w ? 0 : 1)
  y = 25 * x
  y += 1
  z *= y
  y = w + 14
  y *= x
  z += y
  z == z_out
end

def solve5(w, z, z_out)
  x = z % 26
  z /= 26
  x -= 9
  x = (x == w ? 0 : 1)
  y = 25 * x
  y += 1
  z *= y
  y = w + 14
  y *= x
  z += y
  z == z_out
end

def solve4(w, z, z_out)
  x = z % 26
  z /= 1
  x += 15
  x = (x == w ? 0 : 1)
  y = 25 * x
  y += 1
  z *= y
  y = w + 5
  y *= x
  z += y
  z == z_out
end

def solve3(w, z, z_out)
  x = z % 26
  z /= 1
  x += 11
  x = (x == w ? 0 : 1)
  y = 25 * x
  y += 1
  z *= y
  y = w + 10
  y *= x
  z += y
  z == z_out
end

def solve2(w, z, z_out)
  x = z % 26
  z /= 26
  x -= 16
  x = (x == w ? 0 : 1)
  y = 25 * x
  y += 1
  z *= y
  y = w + 8
  y *= x
  z += y
  z == z_out
end

def solve1(w, z, z_out)
  x = z % 26
  z /= 26
  x -= 2
  x = (x == w ? 0 : 1)
  y = 25 * x
  y += 1
  z *= y
  y = w + 15
  y *= x
  z += y
  z == z_out
end

# put each method in an easily indexed array
@sections = [method(:solve1), method(:solve2), method(:solve3),
             method(:solve4), method(:solve5), method(:solve6),
             method(:solve7), method(:solve8), method(:solve9),
             method(:solve10), method(:solve11), method(:solve12),
             method(:solve13), method(:solve14)]

# We know the last digit has to result in z == 0
# Looking at the code, the only input that drves z is z's previous value from
# the last digit and the actual digit under test
# So we can iterate over 1-9 for the input digit and a range of value for a
# previous z value to find ones that create a condition where z == 0
# Once we identify a digit and z input value, we can move back one digit
# use that z input value as the required z output value, and do it all again
# It looks like there are multiple digit/z input combinations that can result
# in the desired z output, so we need to test them all
# For example, the digit 1 and input z of 3 could result in z out == 0
# as well as digit == 2 and z in == 4
# I decided to do a depth first search to try to get to a number that might
# be 14 digits long quicker.  To determine the largest number, I consider
# the digits from 9 to 1, the smallest from 1 to 9
#
# I have zero confidence that the first 14 digit number I hit is the correct
# answer in all cases but it worked for my input
#
# This is also really slow for the part 1 portion
# part 2 wraps up pretty quickly
#
# There's got to be a better way

def compute_results_smallest(i, num, z_out)
  if i == 14
    @final = num
    return
  end

  section = @sections[i]

  1.upto(9) do |w|
    0.upto(65_535) do |z|
      return unless @final.nil?
      break if i == 13 && z > 0
      next if section.call(w, z, z_out) == false

      compute_results_smallest(i + 1, num + (w * (10**i)), z)
    end
  end
end

def compute_results_largest(i, num, z_out)
  # if we are at 14 digits, save off the number and bail
  if i == 14
    @final = num
    return
  end

  # determine which digit we are analyzing
  section = @sections[i]

  # starting at the highest possible digit
  9.downto(1) do |w|
    # i'm not sure what a good z input rage is,
    # i did 1000 and got no valid numbers
    0.upto(65_535) do |z|
      # if we found an answer, we are done
      return unless @final.nil?
      # if we are on the most significant digit, z can only
      # be 0 because that's its initial state
      break if i == 13 && z > 0
      # see if this digit and z value result in the desired z_out
      next if section.call(w, z, z_out) == false

      # if so, go deeper one more value
      compute_results_largest(i + 1, num + (w * (10**i)), z)
    end
  end
end

@final = nil

compute_results_largest(0, 0, 0)

puts format('Part 1: %d', @final)

@final = nil

compute_results_smallest(0, 0, 0)

puts format('Part 2: %d', @final)
