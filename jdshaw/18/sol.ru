# frozen_string_literal: true

require 'json'

class Tree
  # used to reference the tree as a flat array for easy lookup of
  # neighboring numbers
  @@flat_array = []

  # member class instance variables
  attr_accessor :value, :left, :right

  # initialize either a blank node or fill in with some info
  def initialize(value = nil)
    if value.is_a?(Array)
      # if we were passed an array, we need to create sub nodes
      @value = nil
      @left = Tree.new(value[0])
      @right = Tree.new(value[1])
    else
      # if we were passed an int or nil, this is a leaf node with no sub nodes
      @value = value
      @left = nil
      @right = nil
    end
  end

  # add another tree to this tree
  # create a new parent node with the two trees as sub nodes
  # then reduce that tree
  def add(tree)
    p = Tree.new
    p.left = self
    p.right = tree
    p.reduce
    p
  end

  # reduce a tree
  # always check for explosions first each loop
  # then check for splits
  # keep going until there is nothing left to reduce
  def reduce
    loop do
      # clear out the array and build it back up
      @@flat_array = []
      # create an ordered list of our nodes for easy access
      # of previous and next nodes
      build_list
      next if explode(0)
      next if check_splits

      break
    end
  end

  # pretty print the current tree
  def print
    if @value.nil?
      format('[%<l>s,%<r>s]', l: @left.print, r: @right.print)
    else
      format('%d', @value)
    end
  end

  # generate the ordered list of leaf nodes, so we can easily
  # find the previous and next node
  def build_list
    if @value.nil?
      @left.build_list
      @right.build_list
    else
      @@flat_array << self
    end
  end

  # perform the actual splitting of a node with value greater than 9
  def do_split
    left = @value / 2
    right = left
    right += 1 if @value.odd?
    @left = Tree.new(left)
    @right = Tree.new(right)
    @value = nil
    true
  end

  # navigate the tree, and return after doing the first split of
  # a number greater than 9
  def check_splits
    # do the split if conditions are met
    return do_split if !@value.nil? && @value > 9
    # otherwise navigate down the tree
    return true if !@left.nil? && @left.check_splits
    return true if !@right.nil? && @right.check_splits

    false
  end

  def do_explosion
    i = @@flat_array.index(@left)
    @@flat_array[i - 1].value += @left.value if i.positive?
    # find the right index in the array and update its later node
    i = @@flat_array.index(@right)
    @@flat_array[i + 1].value += @right.value if i < @@flat_array.length - 1
    # after an explode, the pair becomes a 0
    @value = 0
    @left = nil
    @right = nil
    true
  end

  # navigate the tree, and return after doing the first explosion of
  # a pair that is too deep
  def explode(depth)
    # if we have a node that isn't a leaf, and it's depth is 4, we will
    # consider its children
    if @value.nil? && depth == 4
      # find the left index in the array and update its earlier node
      return do_explosion
    end
    # navigate down the tree
    return true if !@left.nil? && @left.explode(depth + 1)
    return true if !@right.nil? && @right.explode(depth + 1)
  end

  # measure the magnitude of the tree
  def magnitude
    # if a leaf, just return the value
    return @value unless @value.nil?

    # non leafs do the math on their left and right node
    @left.magnitude * 3 + @right.magnitude * 2
  end
end

# read in input like json, create arrays for us
input = File.read('input').split.map! { |a| JSON.parse(a) }

# create the first tree
tree = Tree.new(input[0])

# add in all the additional trees
1.upto(input.length - 1) do |i|
  tree = tree.add(Tree.new(input[i]))
end

puts format('Part 1: %d', tree.magnitude)

# track the max magnitude
maxmag = 0

# iterate over each number pairwise
0.upto(input.length - 1) do |i|
  (0).upto(input.length - 1) do |j|
    # skip when same number
    next if i == j

    # add the two trees
    tree = Tree.new(input[i])
    tree = tree.add(Tree.new(input[j]))

    # check the magnitude
    maxmag = tree.magnitude if tree.magnitude > maxmag
  end
end

puts format('Part 2: %d', maxmag)
