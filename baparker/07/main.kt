import java.io.File
import kotlin.math.abs

fun main() {
    File("input.txt").forEachLine {
        val positions = it.split(",").map({ pos -> pos.toInt() }).sorted()
        val splitPos = positions.chunked(positions.size / 2)
        println(positions.sumOf({ pos -> abs(pos - splitPos[1][0]) }))
        val bestWeightedPos = (positions.sum() / positions.size)
        println(
                positions.sumOf({ pos ->
                    abs(pos - bestWeightedPos) * (abs(pos - bestWeightedPos) + 1) / 2
                })
        )
    }
}
