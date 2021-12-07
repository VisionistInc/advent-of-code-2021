import java.io.File
import kotlin.math.abs
import kotlin.math.ceil
import kotlin.math.floor

fun getDist(p1: Int, p2: Int): Int {
    return abs(p1 - p2)
}

fun getTriangularNum(num: Int): Int {
    return num * (num + 1) / 2
}

fun getSmallestSumOfFuelCost(list: List<Int>, center: Double): Int {
    var floorSum = 0
    var ceilSum = 0
    list.forEach({ pos ->
        val floorCenter = floor(center).toInt()
        floorSum += getTriangularNum(getDist(pos, floorCenter))
        val ceilCenter = ceil(center).toInt()
        ceilSum += getTriangularNum(getDist(pos, ceilCenter))
    })
    return minOf(floorSum, ceilSum)
}

fun main() {
    File("input.txt").forEachLine {
        val positions = it.split(",").map({ pos -> pos.toInt() }).sorted()
        val splitPos = positions.chunked(positions.size / 2)
        println(positions.sumOf({ pos -> abs(pos - splitPos[1][0]) }))
        val bestWeightedPos = (positions.sum().toDouble() / positions.size)
        println(getSmallestSumOfFuelCost(positions, bestWeightedPos))
    }
}
