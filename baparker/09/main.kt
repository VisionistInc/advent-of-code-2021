import java.io.File

data class Node(val value: Int) {
    var visited = false
}

data class Counter(var value: Int) {
    fun incCounter() {
        this.value++
    }
}

fun findMins(
        heightMatrix: MutableList<MutableList<Node>>,
        x: Int,
        y: Int,
) {
    val current = heightMatrix.get(y).get(x).value
    listOf(listOf(x - 1, y), listOf(x + 1, y), listOf(x, y - 1), listOf(x, y + 1)).forEach {
        val xIndex = it.get(0)
        val yIndex = it.get(1)
        val newPoint = heightMatrix.getOrNull(yIndex)?.getOrNull(xIndex)
        if (newPoint != null) {
            if (newPoint.value <= current) {
                heightMatrix.get(y).get(x).visited = true
                if (!newPoint.visited) {
                    findMins(
                            heightMatrix,
                            xIndex,
                            yIndex,
                    )
                }
            } else {
                if (!newPoint.visited) {
                    heightMatrix.get(yIndex).get(xIndex).visited = true
                }
            }
        }
    }
}

fun findBasins(heightMatrix: MutableList<MutableList<Node>>, x: Int, y: Int, counter: Counter) {
    heightMatrix.get(y).get(x).visited = true
    counter.incCounter()
    listOf(listOf(x - 1, y), listOf(x + 1, y), listOf(x, y - 1), listOf(x, y + 1)).forEach {
        val xIndex = it.get(0)
        val yIndex = it.get(1)
        val newPoint = heightMatrix.getOrNull(yIndex)?.getOrNull(xIndex)
        if (newPoint != null) {
            if (newPoint.value != 9 && !newPoint.visited) {
                findBasins(heightMatrix, xIndex, yIndex, counter)
            }
        }
    }
}

fun main() {
    val heightMatrixMins: MutableList<MutableList<Node>> = mutableListOf()
    val heightMatrixBasins: MutableList<MutableList<Node>> = mutableListOf()
    File("input.txt").forEachLine {
        heightMatrixMins.add(it.map { Node(it.digitToInt()) }.toMutableList())
        heightMatrixBasins.add(it.map { Node(it.digitToInt()) }.toMutableList())
    }

    heightMatrixMins.forEachIndexed { y, heightRow ->
        heightRow.forEachIndexed { x, height ->
            if (!height.visited) {
                findMins(heightMatrixMins, x, y)
            }
        }
    }
    println(
            heightMatrixMins.flatten().fold(0) { sum, height ->
                if (!height.visited) sum + height.value + 1 else sum
            }
    )

    val counterList: MutableList<Int> = mutableListOf()
    var counter = Counter(0)
    heightMatrixBasins.forEachIndexed { y, heightRow ->
        heightRow.forEachIndexed { x, height ->
            if (height.value != 9 && !height.visited) {
                findBasins(heightMatrixBasins, x, y, counter)
                counterList.add(counter.value)
                counter.value = 0
            }
        }
    }
    println(
            counterList.sortedDescending().subList(0, 3).fold(1) { product, count ->
                product * count
            }
    )
}
