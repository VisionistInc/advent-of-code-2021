import java.io.File

fun findMins(
        heightMatrix: MutableList<MutableList<String>>,
        x: Int,
        y: Int,
) {
    val current = heightMatrix.get(y).get(x).get(0)
    listOf(listOf(x - 1, y), listOf(x + 1, y), listOf(x, y - 1), listOf(x, y + 1)).forEach {
        val xIndex = it.get(0)
        val yIndex = it.get(1)
        val newPoint = heightMatrix.getOrNull(yIndex)?.getOrNull(xIndex)
        if (newPoint != null) {
            if (newPoint.get(0) <= current) {
                heightMatrix.get(y).set(x, current + "*")
                if (!newPoint.endsWith("*")) {
                    findMins(
                            heightMatrix,
                            xIndex,
                            yIndex,
                    )
                }
            } else {
                if (!newPoint.endsWith("*")) {
                    heightMatrix.get(yIndex).set(xIndex, newPoint + "*")
                }
            }
        }
    }
}

data class Counter(var value: Int) {
    fun incCounter() {
        this.value++
    }
}

fun findBasins(heightMatrix: MutableList<MutableList<String>>, x: Int, y: Int, counter: Counter) {
    val current = heightMatrix.get(y).get(x).get(0)
    heightMatrix.get(y).set(x, current + "*")
    counter.incCounter()
    listOf(listOf(x - 1, y), listOf(x + 1, y), listOf(x, y - 1), listOf(x, y + 1)).forEach {
        val xIndex = it.get(0)
        val yIndex = it.get(1)
        val newPoint = heightMatrix.getOrNull(yIndex)?.getOrNull(xIndex)
        if (newPoint != null) {
            if (newPoint != "9" && !newPoint.endsWith("*")) {
                findBasins(heightMatrix, xIndex, yIndex, counter)
            }
        }
    }
}

fun main() {
    val heightMatrixMins: MutableList<MutableList<String>> = mutableListOf()
    val heightMatrixBasins: MutableList<MutableList<String>> = mutableListOf()
    File("input.txt").forEachLine {
        heightMatrixMins.add(it.map { it.toString() }.toMutableList())
        heightMatrixBasins.add(it.map { it.toString() }.toMutableList())
    }

    heightMatrixMins.forEachIndexed { y, heightRow ->
        heightRow.forEachIndexed { x, height ->
            if (!height.endsWith("*")) {
                findMins(heightMatrixMins, x, y)
            }
        }
    }
    println(
            heightMatrixMins.flatten().fold(0) { sum, height ->
                if (!height.endsWith("*")) sum + height.toInt() + 1 else sum
            }
    )

    val counterList: MutableList<Int> = mutableListOf()
    var counter = Counter(0)
    heightMatrixBasins.forEachIndexed { y, heightRow ->
        heightRow.forEachIndexed { x, height ->
            if (height != "9" && !height.endsWith("*")) {
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
