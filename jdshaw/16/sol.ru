# frozen_string_literal: true

# just a helper funtion for readability
# pull out number of bits specified from packet and convert to integer
def get_num(packet, bits)
  # special case for 1 bit since join will not work
  return packet.shift[0].to_i(2) if bits == 1

  packet.shift(bits).join('').to_i(2)
end

# handles the literal packet
def get_literal(packet)
  # get the more bit and 4 bits of the first value
  more = get_num(packet, 1)
  value = packet.shift(4)

  # while the more bit is set, append 4 bits to value
  while more == 1
    more = get_num(packet, 1)
    value.append(*packet.shift(4))
  end

  # got all the bits that make up value, now make integer out of them
  value.join('').to_i(2)
end

# handle length type 0
def id_zero(packet, ver_sum)
  # array to store all numbers
  numbers = []
  # length of all subpackets is next 15 bits
  sub_packets = packet.shift(get_num(packet, 15))
  # while there is subpacket length left
  # this assumes there is no padding on subpackets!
  while sub_packets.any?
    # parse the subpacket
    v, n = parse_packet(sub_packets)
    # increment version sum
    ver_sum += v
    # save off any number returned from subpacket
    numbers << n
  end

  # return sum and all numbers
  [ver_sum, numbers]
end

# handle length type 1
def id_one(packet, ver_sum)
  # array to store all numbers
  numbers = []
  # total number of subpackets is next 11 bits
  # this assumes sub-sub packets aren't counted in this value!
  num_pkts = get_num(packet, 11)
  # for each subpacket
  1.upto(num_pkts) do
    # parse the packet and increment version sum, add number returned to list
    v, n = parse_packet(packet)
    ver_sum += v
    numbers << n
  end
  # return sum and all numbers
  [ver_sum, numbers]
end

# break out comparison operations to reduce complexity
def compare(type, numbers)
  return 1 if type == 5 && numbers[0] > numbers[1]
  return 1 if type == 6 && numbers[0] < numbers[1]
  return 1 if type == 7 && numbers[0] == numbers[1]

  0
end

# operation packets invovle doing math on a bunch of numbers in sub-packets
# the mathmatical operation is driven by the type
def do_math(type, numbers)
  return numbers.sum if type.zero?
  return numbers.inject(:*) if type == 1
  return numbers.min if type == 2
  return numbers.max if type == 3

  # not one above, so do comparisons
  compare(type, numbers[0..1])
end

# handle an operation subpacket
# get the length type id and call the appropriate handler
def do_operation(packet, type, ver_sum)
  id = get_num(packet, 1)
  ver, nums = id.zero? ? id_zero(packet, ver_sum) : id_one(packet, ver_sum)

  # do the math operation on the numbers returned
  [ver, do_math(type, nums)]
end

# parse the packet, returning any literal value or result of operation
# as well as the cum of all the version fields within this packet
def parse_packet(packet)
  # every packet has a version and type
  version = get_num(packet, 3)
  type = get_num(packet, 3)

  # grab literal value for type 4
  return version, get_literal(packet) if type == 4

  # do operaion for all other types
  do_operation(packet, type, version)
end

# for each char in the file, covert to 4 bit binary number
# then join all the 4 bits into one string, and split each bit out
input = File.read('input').chars.map do |x|
  x.hex.to_s(2).rjust(4, '0')
end.join('').split('')

# kick off the parsing
ver, value = parse_packet(input)

puts format('Part 1: %d', ver)
puts format('Part 2: %d', value)
