# frozen_string_literal: true

# generate new coordinates for the variant based on the
# fixed point and variant point that will be anchored to it
def generate_new_coords(f_point, v_point, variant)
  scanner = [f_point[0] - v_point[0],
             f_point[1] - v_point[1],
             f_point[2] - v_point[2]]
  v = variant.clone.map(&:clone)
  0.upto(v.length - 1) do |i|
    v[i][0] += scanner[0]
    v[i][1] += scanner[1]
    v[i][2] += scanner[2]
  end
  [v, scanner]
end

# see if a specific variant overlaps a fixed scanner by 12
def check_points(fixed, variant)
  fixed.each do |f|
    variant.each do |v|
      # for each coordinate in the fixed scanner and variant
      # anchor the variant to the fixed and check for overlaps
      beacons, scanner = generate_new_coords(f, v, variant)
      intersect = fixed & beacons
      return beacons, scanner if intersect.length == 12
    end
  end
  [nil, nil]
end

# see if we have a valid overlap between a fixed scanner and any
# variants from another scanner
def check_overlap(fixed, variants)
  variants.each do |v|
    beacons, scanner = check_points(fixed, v)
    return beacons, scanner unless beacons.nil?
  end
  [nil, nil]
end

# generate all rotations and direction variations for a given point
def generate_variants(scanner)
  variants = []
  direction = [1, -1].repeated_permutation(3).to_a
  rotation = [[], [], [], [], [], []]
  # generate all the rotations
  # for example, [1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]
  scanner.each do |s|
    i = 0
    s.permutation(3).each do |x|
      rotation[i] << x
      i += 1
    end
  end
  # now generate all the directions that could result from each rotation
  # for example [1,2,3],[1,2,-3],[1,-2,3],[1,-2,-3]....
  direction.each do |d|
    rotation.each do |r|
      s = []
      r.each do |p|
        s << [p[0] * d[0], p[1] * d[1], p[2] * d[2]]
      end
      variants << s
    end
  end

  variants
end

# read in input, split on scanners
input = File.read('input').split("\n\n")

# empty list of scanner points
scanners = []

# for each scanner in the input
input.each do |s|
  # grab all the beacons this scanner can see, convert to int
  beacons = s.split("\n")[1..]
  scanner = []
  beacons.each do |beacon|
    beacon = beacon.split(',').map(&:to_i)
    scanner << beacon
  end
  # add this scanner to our total list
  scanners << scanner
end

# how many scanners are there total
total = scanners.length

# track all the beacon positions
final_beacons = []
# scanner 0 will be our anchor
final_beacons << scanners.shift
final_scanners = [[0, 0, 0]]

# iterate over every scanner we will find
0.upto(total - 2) do |i|
  # fnd out how many scanners haven't been anchored yet
  l = scanners.length
  0.upto(l - 1) do |j|
    # this takes a while, printing progress is nice
    print format("%<m>3d / %<t>3d Checking %<i>d (%<j>3d / %<l>3d)\r",
                 m: final_beacons.length, t: total, i: i, j: j, l: l)
    # pull off beacons for one scanner
    s = scanners.shift
    # generate all the twists and turns
    v = generate_variants(s)
    # see if any of them overlap the current scanners beacons
    beacons, scanner = check_overlap(final_beacons[i], v)
    # if not, throw back to the end
    if beacons.nil?
      scanners << s
    # if so, add those beacons and the scanner position to our final trackers
    else
      final_beacons << beacons
      final_scanners << scanner
    end
  end
end

# flatten all the beacons and remove dupes
final = final_beacons.flatten(1).uniq
puts ''
puts format('Part 1: %d', final.length)

# iterate over the pairs, calculating distance
max_dis = 0
pairs = final_scanners.permutation(2)
pairs.each do |p|
  x_dis = p[0][0] - p[1][0]
  y_dis = p[0][1] - p[1][1]
  z_dis = p[0][2] - p[1][2]
  dis = x_dis.abs + y_dis.abs + z_dis.abs
  max_dis = dis if dis > max_dis
end

puts format('Part 2: %d', max_dis)
