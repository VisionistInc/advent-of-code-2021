import java.io.File;
import kotlin.math.abs;

fun main(args : Array<String>) {
    val input = File(args.first()).readLines().first().split(",").map { it.toInt() }
    println("Solution 1: ${solution1(input)}")
    println("Solution 2: ${solution2(input)}")
}

private fun solution1(input: List<Int>) :Int {
    val fuels = mutableListOf<Int>()

    for (possibleXValue in input) {
        var fuel = 0

        for (crabPos in input) {
            fuel += abs(crabPos - possibleXValue)
        }
        fuels.add(fuel)
    }

    return fuels.minOrNull() ?: 0
}
    
private fun solution2(input: List<Int>) :Int {
    val min = input.minOrNull() ?: 0
    val max = input.maxOrNull() ?: 0
    val fuels = mutableListOf<Int>()

    // Have to consider non-listed positions, so the entire range between the min and max
    for (possibleXValue in min..max ) {
        var fuel = 0

        for (crabPos in input) {
            val distance =  abs(crabPos - possibleXValue)
            for (n in 0..distance) {
                fuel += n
            }
        }
        fuels.add(fuel)
    }

    return fuels.minOrNull() ?: 0
}   
