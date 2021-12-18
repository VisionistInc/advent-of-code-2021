import java.io.File;

// Thanks to https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm#Algorithm and comments in
// the slack that it uses this algorithm
fun main(args : Array<String>) {
    val input = File(args.first()).readLines()
        .map { line -> line.map { it -> it.digitToInt() }.toMutableList() }.toMutableList()
    println("Solution 1: ${solution1(input)}")
    println("Solution 2: ${solution2(input)}")
}

fun solution1(input: List<List<Int>>) :Int {
    return doDijkstra(input)
}

private fun solution2(input: MutableList<MutableList<Int>>) :Int {
    // go down
    val largeMap = input
    var modifiedInput = input
    repeat(4) {
        val nextInc = increment(modifiedInput)
        largeMap.addAll(nextInc)
        modifiedInput = nextInc
    }
   
    // now go accross
    modifiedInput = largeMap
    repeat(4) {
        val nextInc = increment(modifiedInput)
        for (x in 0..modifiedInput.size - 1) {
            largeMap[x].addAll(nextInc[x])
        }
        modifiedInput = nextInc
    }

    return doDijkstra(largeMap)
}   

private fun doDijkstra(input: List<List<Int>>) :Int {
    // make the graph from the input
    val vertices = mutableSetOf<Pair<Int, Int>>()
    val edges = mutableListOf<Triple<Pair<Int, Int>, Pair<Int, Int>, Int>>()
    for (y in 0..input.size - 1) {
        for (x in 0..input[0].size - 1) {
            val vertex = Pair(x, y)
            vertices.add(vertex)
            if (x < input[0].size - 1) {
                edges.add(Triple(vertex, Pair(x + 1, y), input[y][x + 1]))
            }
            if (x > 0) {
                edges.add(Triple(vertex, Pair(x - 1, y), input[y][x - 1]))
            }
            if (y < input.size - 1) {
                edges.add(Triple(vertex, Pair(x, y + 1), input[y + 1][x]))
            }
            if (y > 0) {
                edges.add(Triple(vertex, Pair(x, y - 1), input[y - 1][x]))
            }
        }
    }
    
    val neighbors = mutableMapOf<Pair<Int, Int>, MutableMap<Pair<Int, Int>, Int>>()
    for ((from, to, risk) in edges) {
        val neighboursFrom = neighbors[from] ?: mutableMapOf()
        neighboursFrom[to] = risk
        neighbors[from] = neighboursFrom
    }
    
    val start = Pair<Int, Int>(0, 0)
    val moves = mutableSetOf(start)

    // store the selected postion pair and it's associated risk
    val positionsToRisks = vertices.fold(mutableMapOf<Pair<Int, Int>, Int>()) { acc, vertex ->
        acc[vertex] = Int.MAX_VALUE // I don't know how big to make this
        acc
    }
    positionsToRisks.put(start, 0)

    // the algorithm is done when we are out of moves...
    while (moves.isNotEmpty()) {
        val currentPos = moves.minByOrNull{ positionsToRisks.getOrDefault(it, 0) }
        // can no longer move to this position, so remove it from moves -- this is key
        moves.remove(currentPos)

        // assess the risk of each neighbor, and if it's less, add the neighbor to the moves
        for (neighbor in neighbors.getOrDefault(currentPos, emptyMap<Pair<Int, Int>, Int>()).keys ) {
            val newRisk = positionsToRisks.getOrDefault(currentPos, 0) + neighbors.getOrDefault(currentPos, emptyMap<Pair<Int, Int>, Int>()).getOrDefault(neighbor, 0) 
            if (newRisk < positionsToRisks.getOrDefault(neighbor, 0)) {
                positionsToRisks[neighbor] = minOf(positionsToRisks.getOrDefault(neighbor, 0), newRisk)
                moves.add(neighbor)
            }
        }
    }

    // the end position should have the cumulative path
    return positionsToRisks.getOrDefault(Pair<Int, Int>(input[0].size - 1, input.size - 1), 0)
}

private fun increment(input: MutableList<MutableList<Int>>): MutableList<MutableList<Int>> {
    val result = MutableList(input.size) { MutableList(input[0].size) { 0 } }
    for (y in 0..input.size - 1) {
        for (x in 0..input[0].size - 1) {
            if (input[y][x] == 9) {
                result[y][x]  = 1 
            } else {
                result[y][x]  = input[y][x] + 1
            }
        }
    }
    return result
}