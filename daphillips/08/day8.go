package main

import (
	"aoc2021/utils"
	"fmt"
	"strconv"
	"strings"
)

const (
	numLitSegmentsFor1 = 2
	numLitSegmentsFor4 = 4
	numLitSegmentsFor7 = 3
	numLitSegmentsFor8 = 7
)

const (
	top         = 1
	topLeft     = 2
	topRight    = 3
	middle      = 4
	bottomLeft  = 5
	bottomRight = 6
	bottom      = 7
)

type char rune

func (c char) String() string {
	return string(c)
}

type Segments map[int][]char

func filterEasyLitSegments(segString string) int {
	splitSegments := strings.Split(segString, " ")
	filteredSegments := utils.Filter(splitSegments, func(in string) bool {
		segmentLen := len(in)
		return segmentLen == numLitSegmentsFor1 || segmentLen == numLitSegmentsFor4 || segmentLen == numLitSegmentsFor7 || segmentLen == numLitSegmentsFor8
	})
	return len(filteredSegments)
}

func part1(lines []string) {
	displays := utils.Map(lines, func(in string) string {
		return strings.Split(in, " | ")[1]
	})

	easyLitSegments := utils.Map(displays, filterEasyLitSegments)

	fmt.Println(utils.Reduce(easyLitSegments, 0, func(acc int, val int) int {
		return acc + val
	}))
}

func determineSegments(patterns []string) map[int]utils.Set[char] {

	// the segment map is a map of segment position to char
	segmentMap := make(map[int]char)

	// the pattern map is a map of number of lit segments to the segments that are lit
	patternMap := make(map[int][]utils.Set[char])

	// yay some variables and some magic numbers!
	patternMap[numLitSegmentsFor1] = make([]utils.Set[char], 0)
	patternMap[numLitSegmentsFor7] = make([]utils.Set[char], 0)
	patternMap[numLitSegmentsFor4] = make([]utils.Set[char], 0)
	patternMap[5] = make([]utils.Set[char], 0)
	patternMap[6] = make([]utils.Set[char], 0)
	patternMap[numLitSegmentsFor8] = make([]utils.Set[char], 0)

	// populate the pattern map
	for _, pattern := range patterns {
		patternMap[len(pattern)] = append(patternMap[len(pattern)], utils.SetFrom([]char(pattern)))
	}

	// next, we will use a heuristic to determine which letter lights which segment

	// numbers 1, 4, 7, and 8 have a unique number of lit segments, so let's get those now
	oneSegments := utils.SetFrom(patternMap[numLitSegmentsFor1][0].Values())
	sevenSegments := utils.SetFrom(patternMap[numLitSegmentsFor7][0].Values())
	fourSegments := utils.SetFrom(patternMap[numLitSegmentsFor4][0].Values())
	eightSegments := utils.SetFrom(patternMap[numLitSegmentsFor8][0].Values())

	// the top-most segment (value 1) is the only thing in the difference between the segments for 1 and 67
	segmentMap[top] = sevenSegments.Difference(oneSegments).Values()[0]

	// the bottom-right segment is in (read: intersection of) 1, 8, and all the 6-segment values
	bottomRightSegment := oneSegments

	for _, chars := range patternMap[6] {
		bottomRightSegment = bottomRightSegment.Intersection(chars)
	}

	// the top-right is the difference between 1 and the bottom-right
	topRightSegment := oneSegments.Difference(bottomRightSegment)

	segmentMap[bottomRight] = bottomRightSegment.Values()[0]
	segmentMap[topRight] = topRightSegment.Values()[0]

	// the middle is the only one that 4 and the 5-segments have in common
	middleSegment := fourSegments
	for _, chars := range patternMap[5] {
		middleSegment = middleSegment.Intersection(chars)
	}

	segmentMap[middle] = middleSegment.Values()[0]

	// if we take the intersection of the 5-segments, we get the top-middle-bottom segments
	// if we then take the difference of that with what we've found, we got the bottom
	sharedFiveSegments := patternMap[5][0]
	for _, chars := range patternMap[5] {
		sharedFiveSegments = sharedFiveSegments.Intersection(chars)
	}

	currentFoundSegments := make(utils.Set[char], 0)
	for _, v := range segmentMap {
		currentFoundSegments.Add(v)
	}

	bottomSegment := sharedFiveSegments.Difference(currentFoundSegments).Values()[0]

	segmentMap[bottom] = bottomSegment
	currentFoundSegments.Add(bottomSegment)

	// top-left is what 8 and 6-segments have diffed with what we have found
	topLeftSegmentSet := eightSegments
	for _, chars := range patternMap[6] {
		topLeftSegmentSet = topLeftSegmentSet.Intersection(chars)
	}

	topLeftSegmentSet = topLeftSegmentSet.Difference(currentFoundSegments)
	topLeftSegment := topLeftSegmentSet.Values()[0]

	segmentMap[topLeft] = topLeftSegment
	currentFoundSegments.Add(topLeftSegment)

	// last segment is difference of 8 with what we have
	bottomLeftSegment := eightSegments.Difference(currentFoundSegments).Values()[0]
	segmentMap[bottomLeft] = bottomLeftSegment

	// the number map maps a number to the segments needed to light it
	numberMap := make(map[int]utils.Set[char], 0)
	numberMap[0] = utils.SetFrom([]char{segmentMap[1], segmentMap[2], segmentMap[3], segmentMap[5], segmentMap[6], segmentMap[7]})
	numberMap[1] = oneSegments
	numberMap[2] = utils.SetFrom([]char{segmentMap[1], segmentMap[3], segmentMap[4], segmentMap[5], segmentMap[7]})
	numberMap[3] = utils.SetFrom([]char{segmentMap[1], segmentMap[3], segmentMap[4], segmentMap[6], segmentMap[7]})
	numberMap[4] = fourSegments
	numberMap[5] = utils.SetFrom([]char{segmentMap[1], segmentMap[2], segmentMap[4], segmentMap[6], segmentMap[7]})
	numberMap[6] = utils.SetFrom([]char{segmentMap[1], segmentMap[2], segmentMap[4], segmentMap[5], segmentMap[6], segmentMap[7]})
	numberMap[7] = sevenSegments
	numberMap[8] = eightSegments
	numberMap[9] = utils.SetFrom([]char{segmentMap[1], segmentMap[2], segmentMap[3], segmentMap[4], segmentMap[6], segmentMap[7]})

	return numberMap
}

func part2(lines []string) {

	sum := 0

	for _, line := range lines {
		splitLine := strings.Split(line, " | ")
		patterns := strings.Split(splitLine[0], " ")
		outputs := strings.Split(splitLine[1], " ")

		numberMap := determineSegments(patterns)

		var numberStr string
		for _, output := range outputs {
			for num, segments := range numberMap {
				// check if the set of segments equals the set of output segments
				// there should only be 1 match here for each output number
				if segments.Equals(utils.SetFrom([]char(output))) {
					// convert to string for easier concatenation
					numberStr += strconv.Itoa(num)
				}
			}
		}
		// back to number for sum
		number, _ := strconv.Atoi(numberStr)
		sum += number
	}
	fmt.Println(sum)
}

func main() {
	input := utils.ReadFile("input.txt")
	part1(input)
	part2(input)
}
