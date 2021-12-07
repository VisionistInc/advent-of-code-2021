package main

import (
	"aoc2021/utils"
	"fmt"
	"strconv"
	"strings"
)

// a lanternfish struct knows how many new fish will be created after the timer
// prevents a naive explosion of new fish
type Lanternfish struct {
	Timer      int
	NumNewFish int64
}

const defaultTimer int = 6
const initialTimer int = 8

func (l *Lanternfish) Tick() int64 {
	l.Timer--
	if l.Timer < 0 {
		l.Timer = defaultTimer
		return l.NumNewFish
	}
	return 0
}

func Compact(f []Lanternfish) []Lanternfish {

	// determine how many new fish will be created on a given timer day based on the input
	reducer := make(map[int]int64)

	for _, fish := range f {
		reducer[fish.Timer] += fish.NumNewFish
	}

	shrunkList := make([]Lanternfish, 0)
	// construct a new (packed) list of fish based on the map
	for timer, numNewFish := range reducer {
		shrunkList = append(shrunkList, Lanternfish{timer, numNewFish})
	}
	return shrunkList
}

func runSimulation(input string, numTicks int) {
	fishes := utils.Map(strings.Split(input, ","), func(in string) Lanternfish {
		timer, _ := strconv.Atoi(in)
		return Lanternfish{timer, 1}
	})

	fishes = Compact(fishes)

	for i := 0; i < numTicks; i++ {
		tickChan := make(chan int64)

		for j := 0; j < len(fishes); j++ {
			// haha goroutines go brrrrr
			go func(f *Lanternfish) {
				tickChan <- f.Tick()
			}(&fishes[j])
		}

		numNewFish := int64(0)
		for k := 0; k < len(fishes); k++ {
			newFish := <-tickChan
			numNewFish += newFish
		}

		fishes = append(fishes, Lanternfish{initialTimer, numNewFish})

		// compact in case there are any other overlaps we can make after the bulk add of new fish
		fishes = Compact(fishes)
	}

	totalFish := int64(0)
	for _, packedFish := range fishes {
		totalFish += packedFish.NumNewFish
	}

	fmt.Println(totalFish)
}

func main() {
	input := utils.ReadFile("input.txt")

	runSimulation(input[0], 80)
	runSimulation(input[0], 256)
}
