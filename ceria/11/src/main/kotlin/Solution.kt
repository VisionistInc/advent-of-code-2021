import java.io.File;

var flashCount = 0  // every time flash() gets called, this should be incremented

fun main(args : Array<String>) {
    val input = File(args.first()).readLines()
    val energy = Array<IntArray>(input.size) { IntArray(input.first().length) {0} }
    var i = 0
    for (line in input) {
        var row = energy[i]
        for (idx in line.indices) {
            row[idx] = line.get(idx).digitToInt()
        }
        i++
    }

    // since we're going to mutate energy as we go, clone it before we use it so that
    //  solution 2 can have an original copy
    val energy2 = energy.clone()

    println("Solution 1: ${solution1(energy)}")
    println("Solution 2: ${solution2(energy2)}")
}

private fun solution1(energy: Array<IntArray>) :Int {
    for (x in 1..100) {
        step(energy)
    }
    return flashCount
}

private fun solution2(energy: Array<IntArray>) :Int {
    var stepCount = 0
    var allZeros = false
    while (!allZeros) {
        step(energy)
        stepCount++

        for (rowIdx in energy.indices) {
            if (energy[rowIdx].filter{ it == 0 }.size != energy[rowIdx].size) {
                allZeros = false
                break
            } else {
                allZeros = true
            }
        }
    }
    return stepCount
}

private fun step(energy: Array<IntArray>) {
    var alreadyFlashed = mutableSetOf<Pair<Int, Int>>()
    for (rowIdx in energy.indices) {
        // increase each energy by 1
        energy[rowIdx] = energy[rowIdx].map{ it + 1 }.toIntArray()
    }

    // check for octopuses that are greater than 9
    for (rowIdx in energy.indices) {
        val row = energy[rowIdx]

        for (idx in row.indices) {
            if (row[idx] > 9) {
                // flash the octopus and potentially surrounding octopi 
                // (keep track of the octopuses that have already been flashed throughout the process)
                alreadyFlashed = flash(energy, Pair<Int, Int>(idx, rowIdx), alreadyFlashed)
                flashCount++        
            }
        }
    }

}

private fun flash(energy: Array<IntArray>, octopus: Pair<Int, Int>, alreadyFlashed: MutableSet<Pair<Int, Int>>) :MutableSet<Pair<Int, Int>> { 
    var xStart: Int
    var xEnd: Int
    var yStart: Int
    var yEnd: Int

    if (octopus.first > 0) {
        xStart = octopus.first - 1
    } else {
        xStart = octopus.first
    }

    if (octopus.first < energy[0].size - 1) {
        xEnd =octopus.first + 1
    } else {
        xEnd = octopus.first
    }

    if (octopus.second > 0) {
        yStart = octopus.second - 1
    } else {
        yStart = octopus.second
    }

    if (octopus.second < energy.size - 1) {
        yEnd = octopus.second + 1
    } else {
        yEnd = octopus.second
    }

    // this is the octopus that caused the flash, set it to 0
    energy[octopus.second][octopus.first] = 0
    alreadyFlashed.add(octopus)

    for (x in xStart..xEnd) {
        for (y in yStart..yEnd) {
            val newFlashOctopus = Pair<Int, Int>(x, y)
            if (!alreadyFlashed.contains(newFlashOctopus) ) {
                energy[y][x] += 1
                if (energy[y][x] > 9) {
                    flash(energy, newFlashOctopus, alreadyFlashed)
                    flashCount++
                }
            }
        }
    }

    return alreadyFlashed
}