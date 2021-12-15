import java.io.File

data class CavePosition(val level: Int, val x: Int, val y: Int) {
    var visited = false
    var distanceFromStart = Int.MAX_VALUE
}

fun findPathWithLowestRisk(
        cavePositionMatrix: MutableList<MutableList<CavePosition>>,
        x: Int = 0,
        y: Int = 0,
        counter: Int = -1
) {
    val currentNode = cavePositionMatrix.get(y).get(x)
    val count = currentNode.level + counter

    if (count < currentNode.distanceFromStart &&
                    count < cavePositionMatrix.last().last().distanceFromStart
    ) {
        currentNode.distanceFromStart = count

        if (!(y == cavePositionMatrix.size - 1 && x == cavePositionMatrix.get(0).size - 1)) {
            val top = cavePositionMatrix.getOrNull(y - 1)?.get(x)
            val left = cavePositionMatrix.get(y).getOrNull(x - 1)
            val right = cavePositionMatrix.get(y).getOrNull(x + 1)
            val bottom = cavePositionMatrix.getOrNull(y + 1)?.get(x)
            listOf(top, left, right, bottom)
                    .filter { it != null && it.visited == false }
                    .sortedBy { it!!.level }
                    .forEach {
                        it!!.visited = true
                        findPathWithLowestRisk(cavePositionMatrix, it.x, it.y, count)
                        it.visited = false
                    }
        }
    }
}

fun main() {
    val cavePositionMatrix: MutableList<MutableList<CavePosition>> = mutableListOf()

    var yIndex = 0
    File("input.txt").forEachLine {
        cavePositionMatrix.add(
                it
                        .mapIndexed { xIndex, pos ->
                            CavePosition(pos.digitToInt(), xIndex, yIndex)
                        }
                        .toMutableList()
        )
        yIndex++
    }

    cavePositionMatrix.last().last().distanceFromStart =
            cavePositionMatrix.get(0).sumOf { it.level } +
                    cavePositionMatrix.map { it.last().level }.sum()
    findPathWithLowestRisk(cavePositionMatrix)

    println(cavePositionMatrix.last().last().distanceFromStart)
}
