import java.io.File;

lateinit var heights: Array<IntArray>

fun main(args : Array<String>) {
    val input = File(args.first()).readLines()
    heights = Array<IntArray>(input.size) { IntArray(input.first().length) {0} }
    var i = 0
    for (line in input) {
        var row = heights[i]
        for (idx in line.indices) {
            row[idx] = line.get(idx).digitToInt()
        }
        i++
    }

    println("Solution 1: ${solution1()}")
    println("Solution 2: ${solution2()}")
}

private fun solution1() :Int {
    var sums = mutableListOf<Int>()
    for (rowIdx in heights.indices) {
        val row = heights[rowIdx]

        for (idx in row.indices) {
            val point = row[idx]

            // compare to the left
            if (idx > 0 && row[idx - 1] <= point) {
                continue
            }

            // compare to the right
            if (idx < row.size - 1 && row[idx + 1] <= point) {
                continue
            }

            // compare above
            if (rowIdx > 0 && heights[rowIdx - 1][idx] <= point) {
                continue
            }

            // compare below
            if (rowIdx < heights.size - 1 && heights[rowIdx + 1][idx] <= point) {
                continue
            }

            // add to the list (could just keep a sum, but maybe part 2 will be easier if we keep a list)
            sums.add(point + 1)
        }
    }
    return sums.sum()
}

private fun solution2() :Int {
    var basinSizes = mutableListOf<Int>()
    for (rowIdx in heights.indices) {
        val row = heights[rowIdx]

        for (idx in row.indices) {
            val point = row[idx]

            // compare to the left
            if (idx > 0 && row[idx - 1] <= point) {
                continue
            }

            // compare to the right
            if (idx < row.size - 1 && row[idx + 1] <= point) {
                continue
            }

            // compare above
            if (rowIdx > 0 && heights[rowIdx - 1][idx] <= point) {
                continue
            }

            // compare below
            if (rowIdx < heights.size - 1 && heights[rowIdx + 1][idx] <= point) {
                continue
            }

            // if we made it this far, this is a basin, compute it's size
            var basinBounds = mutableSetOf<Pair<Int, Int>>(Pair<Int, Int>(idx, rowIdx))
            var allBounded = false
        
            while (!allBounded) {
                var newPoints = mutableSetOf<Pair<Int, Int>>()
                for (bp in basinBounds) {
                    newPoints.addAll(checkForBounds(bp, basinBounds))
                }
                basinBounds.addAll(newPoints)

                if (newPoints.isEmpty()) {
                    allBounded = true
                }
            }

            basinSizes.add(basinBounds.size)
        }
    }

    basinSizes.sort()
    // return the product of the 3 largest values
    return basinSizes.takeLast(3).reduce{ acc, i ->  acc * i }
}   

private fun checkForBounds(p: Pair<Int, Int>, basinPoints: Set<Pair<Int, Int>>) :Set<Pair<Int, Int>> {
     var newPoints = mutableSetOf<Pair<Int, Int>>()

    // compare to the point to the left - if it's not 9, and not already in newPoints, add it to newPoints
    if (p.first > 0 && heights[p.second][p.first - 1] != 9) {
       val newPoint = Pair<Int, Int>(p.first - 1, p.second)
        if (!basinPoints.contains(newPoint)) {
            newPoints.add(newPoint)
        }
    }

    // compare to the point to the right - if it's not 9, and not already in newPoints, add it to newPoints
    if (p.first < heights[0].size - 1 && heights[p.second][p.first + 1] != 9) {
        val newPoint = Pair<Int, Int>(p.first + 1, p.second)
        if (!basinPoints.contains(newPoint)) {
            newPoints.add(newPoint)
        }
    }

    // compare to the point above - if it's not 9, and not already in newPoints, add it to newPoints
    if (p.second > 0 && heights[p.second - 1][p.first] != 9) {
        val newPoint = Pair<Int, Int>(p.first, p.second - 1)
        if (!basinPoints.contains(newPoint)) {
            newPoints.add(newPoint)
        }
    }

    // compare to the point above - if it's not 9, and not already in newPoints, add it to newPoints
    if (p.second < heights.size -1 && heights[p.second + 1][p.first] != 9) {
        val newPoint = Pair<Int, Int>(p.first, p.second + 1)
        if (!basinPoints.contains(newPoint)) {
            newPoints.add(newPoint)
        }
    }
    
    return newPoints
}
