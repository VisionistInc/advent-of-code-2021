package main

import (
	"aoc2021/utils"
	"fmt"
	"strings"
)

type Cave struct {
	Name      string
	Small     bool
	Neighbors []*Cave
}

func (c *Cave) AddNeighbor(neighbor *Cave) {
	c.Neighbors = append(c.Neighbors, neighbor)
}

func getNumPaths(caveToVisit string, caves map[string]*Cave, visitedCaves utils.Set[string]) int64 {
	numPaths := int64(0)

	if caveToVisit == "end" {
		return 1
	}

	cave := caves[caveToVisit]

	for _, neighbor := range cave.Neighbors {
		if !neighbor.Small || !visitedCaves.Contains(neighbor.Name) {
			// visit this cave and then check the neighbors
			visitedCaves.Add(neighbor.Name)
			numPaths += getNumPaths(neighbor.Name, caves, visitedCaves)
			visitedCaves.Remove(neighbor.Name)
		}
	}

	return numPaths
}

func hasDoubleVisit(visits map[string]int) bool {
	for cave, numVisits := range visits {
		// if it is a small cave and has been visited twice
		if strings.ToLower(cave) == cave && numVisits == 2 {
			return true
		}
	}
	return false
}

func getDoubleVisit(visits map[string]int) string {
	for cave, numVisits := range visits {
		if numVisits == 2 {
			return cave
		}
	}

	var none string
	return none
}

func getNumPaths2(caveToVisit string, caves map[string]*Cave, visitedCaves map[string]int) int64 {
	numPaths := int64(0)

	if caveToVisit == "end" {
		return 1
	}

	cave := caves[caveToVisit]

	for _, neighbor := range cave.Neighbors {

		// we can only visit start once
		if neighbor.Name != "start" {
			visitedCount := visitedCaves[neighbor.Name]
			// if the neighbor is a large cave, or is the end, or hasn't been visited yet, or has but there aren't any double-visited small caves yet, visit it
			if !neighbor.Small || neighbor.Name == "end" || visitedCount < 1 || !hasDoubleVisit(visitedCaves) {
				visitedCaves[neighbor.Name]++
				numPaths += getNumPaths2(neighbor.Name, caves, visitedCaves)
				visitedCaves[neighbor.Name]--
			}
		}
	}

	return numPaths
}

func main() {
	lines := utils.ReadFile("input.txt")
	caves := make(map[string]*Cave, 0)
	for _, line := range lines {
		c := strings.Split(line, "-")
		sourceCave := c[0]
		destCave := c[1]
		if _, present := caves[sourceCave]; !present {
			caves[sourceCave] = &Cave{sourceCave, strings.ToLower(sourceCave) == sourceCave, make([]*Cave, 0)}
		}

		if _, present := caves[destCave]; !present {
			caves[destCave] = &Cave{destCave, strings.ToLower(destCave) == destCave, make([]*Cave, 0)}
		}

		// the graph is undirected :)
		caves[sourceCave].AddNeighbor(caves[destCave])
		caves[destCave].AddNeighbor(caves[sourceCave])
	}

	visitedCaves := utils.SetFrom([]string{"start"})
	fmt.Println(getNumPaths("start", caves, visitedCaves))

	visitedCaves2 := make(map[string]int, 0)
	visitedCaves2["start"] = 1
	fmt.Println(getNumPaths2("start", caves, visitedCaves2))
}
