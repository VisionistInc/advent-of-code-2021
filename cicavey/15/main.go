package main

import (
	"aoc2021/aoc"
	"container/heap"
	"fmt"
	"math"
)

type Point struct {
	x, y int
}

func (p Point) Add(a Point) Point {
	return Point{p.x + a.x, p.y + a.y}
}

func main() {

	grid := make(map[Point]int)
	lines := aoc.ReadLines("input.txt")

	width := len(lines[0])
	height := len(lines)

	for y, line := range lines {
		for x, v := range line {
			p := Point{x, y}
			grid[p] = int(v) - 48
		}
	}

	path := dijkstra(Point{0, 0}, Point{width - 1, height - 1}, grid)
	var sum int
	for _, p := range path[1:] {
		sum += grid[p]
	}
	fmt.Println(sum)

	// Expand the grid and try again

	// Fill top row, tx = tile x
	for tx := 1; tx < 5; tx++ {

		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				cur := Point{tx*width + x, y}
				prev := Point{cur.x - width, y}
				v := grid[prev] + 1
				if v > 9 {
					v = 1
				}
				grid[cur] = v
			}
		}
	}

	// Fill left col, tx = tile x
	for ty := 1; ty < 5; ty++ {
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				cur := Point{x, ty*height + y}
				prev := Point{x, cur.y - height}
				v := grid[prev] + 1
				if v > 9 {
					v = 1
				}
				grid[cur] = v
			}
		}
	}

	// Fill remaining
	for ty := 1; ty < 5; ty++ {
		for tx := 1; tx < 5; tx++ {
			for y := 0; y < height; y++ {
				for x := 0; x < width; x++ {
					cur := Point{tx*width + x, ty*height + y}
					prev := Point{cur.x, cur.y - height}
					v := grid[prev] + 1
					if v > 9 {
						v = 1
					}
					grid[cur] = v
				}
			}

		}
	}

	width *= 5
	height *= 5

	path = dijkstra(Point{0, 0}, Point{width - 1, height - 1}, grid)
	sum = 0
	for _, p := range path[1:] {
		sum += grid[p]
	}
	fmt.Println(sum)
}

func dijkstra(src Point, end Point, grid map[Point]int) []Point {
	dist := make(map[Point]int)
	prev := make(map[Point]Point)

	dist[src] = 0

	pq := NewEmptyPriorityQueue[Point]()

	for p := range grid {
		if p != src {
			dist[p] = math.MaxInt
		}
		pq.PushValue(p, dist[p])
	}

	for !pq.Empty() {
		item := heap.Pop(&pq).(*Item[Point])
		u := item.value

		// short circuit
		if u == end {
			break
		}

		for _, v := range neighbors(u, grid) {
			alt := dist[u] + grid[v]
			if alt < dist[v] {
				dist[v] = alt
				prev[v] = u
				pq.updateByValue(v, alt)
			}
		}
	}

	var path []Point

	path = append(path, end)

	next := prev[end]
	for next != src {
		path = append(path, next)
		next = prev[next]
	}
	path = append(path, src)

	// Reverse path
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}

func neighbors(p Point, grid map[Point]int) []Point {
	dir := []Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	var n []Point
	for _, d := range dir {
		np := p.Add(d)
		if _, ok := grid[np]; ok {
			n = append(n, np)
		}
	}
	return n
}

// An Item is something we manage in a priority queue.
type Item[T comparable] struct {
	value    T   // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.

	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue[T comparable] []*Item[T]

func NewEmptyPriorityQueue[T comparable]() PriorityQueue[T] {
	pq := make(PriorityQueue[T], 0)
	heap.Init(&pq)
	return pq
}

func (pq PriorityQueue[T]) Empty() bool { return len(pq) == 0 }

func (pq PriorityQueue[T]) Len() int { return len(pq) }

func (pq PriorityQueue[T]) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue[T]) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item[T])
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue[T]) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue[T]) PopValue() T {
	item := heap.Pop(pq).(*Item[T])
	return item.value
}
func (pq *PriorityQueue[T]) PushValue(v T, priority int) {
	item := &Item[T]{
		value:    v,
		priority: priority,
	}
	heap.Push(pq, item)
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue[T]) update(item *Item[T], value T, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func (pq *PriorityQueue[T]) updateByValue(value T, priority int) {
	for _, i := range *pq {
		if i.value == value {
			i.priority = priority
			heap.Fix(pq, i.index)
			break
		}
	}
}
