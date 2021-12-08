# determine how much gas is used to move every
# crab to the position.  part2 changes gas
# consumption algorithm
def calculate_gas(pos, crabs, part2=false)
    gas = 0
    crabs.each do |x|
        g = (pos - x).abs
        g = ((g * (g+1))/2) unless (part2 == false)
        gas += g
    end
    return gas
end

# solve one of the parts
def solve(mid, input, part2=false)

    # add each side of the median to future tests
    tests = [mid+1, mid-1]

    # our initial best gas is the median
    best_gas = calculate_gas(mid, input, part2)
    best_pos = mid

    # while we have no more new positions to try
    while tests.length > 0 do

        # pull out a position to test
        x = tests.shift

        # see how much gas it would take
        gas = calculate_gas(x, input, part2)

        # if the gas is better save it off and add its
        # untested neighbor to the list
        # the theory is, if fuel consumption is getting
        # worse as we move, it's not going to get better
        # so we don't have to keep checking.
        if gas <= best_gas
            best_gas = gas
            best_pos = x

            # if moving left of mid
            if x < mid
                tests << x-1
            # or moving right of mid
            else
                tests << x+1
            end
        end
    end

    return best_gas
end


input = File.read("input").split(",").map(&:to_i)

input.sort!

# is this the right time for middle-out?

# find the starting position by finding the median
if input.length % 2 == 0
    mid = (input[input.length/2] + input[(input.length/2)-1])/2
else
    mid = input[input.length/2]
end

puts "Part 1: %d" % [solve(mid,input)]
puts "Part 2: %d" % [solve(mid,input,true)]