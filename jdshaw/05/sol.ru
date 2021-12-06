# retuen next value based on current and final
# funny thing about ruby, return isn't required
# the last value executed is returned
def get_next(current,final)
    if current < final
        current + 1
    elsif current > final
        current - 1
    else
        current
    end
end

# read in file, make everything comma delimited, and split on newlines
input = File.read("input").gsub(" -> ",",").split("\n")

# any new hash value never before referenced will have a value of 0
collides = Hash.new(0)

# iterate over each entry
input.each do |e|

    # get numbers out
    x,y,x_stop,y_stop = e.split(",").map(&:to_i)

    # if we have a horizontal or vertical line
    if (x == x_stop) || (y == y_stop)

        # keep adding to the hash until both x and y reach their stop
        while true
            collides[[x,y]] += 1
            if (x == x_stop && y == y_stop)
                break
            end

            # get the next value for x and y
            x = get_next(x, x_stop)
            y = get_next(y, y_stop)
        end
    end
end

# create a new array based on every value that is greater than one
# in the hash and then get its length
puts "Part 1: %d" % [collides.select { |k, v| v > 1 }.length]

# any new hash value never before referenced will have a value of 0
collides = Hash.new(0)

# iterate over each entry
input.each do |e|

    # get numbers out
    x,y,x_stop,y_stop = e.split(",").map(&:to_i)

    # keep adding to the hash until both x and y reach their stop
    while true
        collides[[x,y]] += 1
        if (x == x_stop && y == y_stop)
            break
        end

        # get the next value for x and y
        x = get_next(x, x_stop)
        y = get_next(y, y_stop)
    end
end

# create a new array based on every value that is greater than one
# in the hash and then get its length
puts "Part 2: %d" % [collides.select { |k, v| v > 1 }.length]