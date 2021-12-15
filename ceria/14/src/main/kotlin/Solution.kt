import java.io.File;
import java.util.Stack;

fun main(args : Array<String>) {
    val input = File(args.first()).readLines()
   
    var startPoly = input.first()
    var polymers = mutableMapOf<String, String>()
    for ( p in 2..input.size - 1) {
        val poly = input.get(p).split(" -> ")
        polymers.put(poly[0], poly[1])
    }
    println("Solution 1: ${solution1(startPoly, polymers)}")
    println("Solution 1: ${solution2(startPoly, polymers)}")
}

/* This does not scale for part 2 :( Which, I had a feeling it wouldn't cause this is AOC after all */
private fun solution1(start: String, polymers: Map<String, String>) :Int {
    var steps = 10
    var startPoly = StringBuilder(start)
    for (x in 1..steps) {
        var windows = startPoly.windowed(size = 2, step = 1)
        var insertAt = startPoly.length - 1

         // start with the end of the list so that our indexes for insertion are correct
        val itr: ListIterator<String> = windows.listIterator(windows.size)
        while (itr.hasPrevious()) {
            var p = itr.previous()
            val newP = polymers.get(p)?.let{
                polymers.get(p)
            } ?: ""

            if (!newP.isEmpty()) {
                startPoly = startPoly.insert(insertAt, newP)
            }
            insertAt -= 1
        }
    }
    var polyCounts = startPoly.groupingBy{ it }.eachCount()
    return  polyCounts.values.maxOf { it } - polyCounts.values.minOf { it }
}

private fun solution2(start: String, polymers: Map<String, String>) :Long {
    var steps = 40

    // Loop through the start string, ignoring the beginning and end of it 
    // since those values will never change, and seed the frequency map of windows 
    var frequencies = start.drop(1).dropLast(1).windowed(size = 2, step = 1).groupingBy{ it }.eachCount().mapValues{ it.value.toLong() }

    // The first and last windows are special
    var firstWindow = start.take(2)
    var lastWindow = start.takeLast(2)

    for (i in 0..steps - 1) {
        var newFrequencies = mutableMapOf<String, Long>()

        // Count the frequencies of the windows for this step
        // For each known polymer in the frequency map, look up it's insert value,
        // create two new windows using the insert value and the poylmer and update
        // the frequencies for those new windows
        for (polymer in frequencies.keys) {
            val insertVal = polymers.getOrDefault(polymer, "")

            // Don't know if this can happen... the insertVal was not in our lookup map, i.e it wasn't in the input
            if (insertVal.isEmpty()) {
                continue
            }

            // the new window created by the first character of the current polymer from frequencies and the insertVal
            val leftInsert = polymer[0] + insertVal
            // the new window created by the second character of the current polymer from frequencies and the insertVal 
            val rightInsert = insertVal + polymer[1]

            // sum the frequencies for the forward and backward windows
            newFrequencies[leftInsert] = newFrequencies.getOrDefault(leftInsert, 0) + frequencies.getOrDefault(polymer, 0)
            newFrequencies[rightInsert] = newFrequencies.getOrDefault(rightInsert, 0) + frequencies.getOrDefault(polymer, 0)
        }

        // need to treat the first and last windows special
        // the first character of the last window is he only character of the two that needs to be updated
        val lastInsert = polymers.getOrDefault(lastWindow, "")
        val lastLeftInsert = lastWindow.first() + lastInsert
        newFrequencies[lastLeftInsert] = newFrequencies.getOrDefault(lastLeftInsert, 0) + 1

        // the last character of the first window is he only character of the two that needs to be updated
        val firstInsert = polymers.getOrDefault(firstWindow, "")
        val lastFirstInsert = firstInsert + firstWindow.last()
        newFrequencies[lastFirstInsert] = newFrequencies.getOrDefault(lastFirstInsert, 0) + 1

        // update the first and last windows with the new windows
        firstWindow = firstWindow.first() + firstInsert
        lastWindow = lastInsert + lastWindow.last()

        frequencies = newFrequencies
    }
    
    // count the polymers -- make Pairs using the first character of each window, to their counts
    //i.e. ('B', 72), ('B' 120) and then sum all the pairs with the same "first" character
    val polymerCounts = frequencies
        .map{ Pair<Char, Long>(it.key.first(), it.value) }
        .groupBy { it.first }
        .mapValues { p -> p.value.sumOf { it.second } }
        .toMutableMap()

    // add in the very first character of the very first window -- it was dropped/ignroed in the loop
    polymerCounts[firstWindow.first()] = polymerCounts.getOrDefault(firstWindow.first(), 0) + 1
    // add in the next to last character of the very last window -- it was just not counted in the last loop iteration.
    polymerCounts[lastWindow.first()] = polymerCounts.getOrDefault(lastWindow.first(), 0) + 1
    // add in the very last character of the very last window -- it was dropped/ignored in the loop
    polymerCounts[lastWindow.last()] = polymerCounts.getOrDefault(lastWindow.last(), 0) + 1
    
    return polymerCounts.values.maxOf { it } - polymerCounts.values.minOf { it }
}