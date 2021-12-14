package main

import (
	"aoc2021/aoc"
	"fmt"
	"strings"
	"unicode"
)

type NodeList map[string]bool
type Graph map[string]NodeList

func (nl NodeList) Contains(n string) bool {
	_, ok := nl[n]
	return ok
}

func main() {

	g := make(Graph)

	aoc.ForLine("input.txt", func(line string) {
		temp := strings.Split(line, "-")

		a := temp[0]
		b := temp[1]

		// "a" -> "b"
		adj, ok := g[a]
		if !ok {
			adj = make(NodeList)
			g[a] = adj
		}
		adj[b] = true

		// "b" -> "a"
		adj, ok = g[b]
		if !ok {
			adj = make(NodeList)
			g[b] = adj
		}
		adj[a] = true

	})

	fmt.Println(findPaths(g, "start", "end", make(NodeList)))
	fmt.Println(findPaths2(g, "start", "end", make(map[string]int), false))
}

func findPaths(g Graph, cur string, dst string, visited NodeList) int {

	if cur == dst {
		return 1
	}

	// Mark current as visited
	if isSmall(cur) {
		visited[cur] = true
	}

	pathCount := 0

	for n := range g[cur] {
		if isSmall(n) {
			if n != "start" && !visited.Contains(n) {
				pathCount += findPaths(g, n, dst, visited)
			}
		} else {
			pathCount += findPaths(g, n, dst, visited)
		}
	}

	delete(visited, cur)

	return pathCount
}

func findPaths2(g Graph, cur string, dst string, visited map[string]int, hasVisitedSmallTwice bool) int {

	if cur == dst {
		return 1
	}

	// Mark current as visited
	if isSmall(cur) && cur != "start" { // excessive checking ,but helps debug

		// // Have we already visisted a small node twice?
		// hasDoubleSmall := false
		// for _, visitCount := range visited {
		// 	if visitCount > 1 {
		// 		hasDoubleSmall = true
		// 	}
		// }

		visited[cur]++
	}

	pathCount := 0

	for n := range g[cur] {
		if isSmall(n) {
			// Don't ever go back through start
			if n == "start" {
				continue
			}

			visitCount, _ := visited[n]

			// Never visisted
			if visitCount == 0 {
				pathCount += findPaths2(g, n, dst, visited, hasVisitedSmallTwice)
			} else if visitCount == 1 && !hasVisitedSmallTwice { // Visited once
				pathCount += findPaths2(g, n, dst, visited, true)
			}

			// if visitCount < 1 {
			// 	pathCount += findPaths2(g, n, dst, visited, hasVisitedSmallTwice)
			// } else if visitCount < 2 {
			// 	pathCount += findPaths2(g, n, dst, visited, hasVisitedSmallTwice)
			// }
		} else {
			// Large node, visit like crazy
			pathCount += findPaths2(g, n, dst, visited, hasVisitedSmallTwice)
		}
	}

	if isSmall(cur) {
		visited[cur]--
	}

	return pathCount
}

func isSmall(s string) bool {
	return !isUpper(s)
}

func isUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
