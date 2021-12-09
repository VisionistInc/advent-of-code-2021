import java.io.File

fun findMin(
        heightMatrix: MutableList<MutableList<String>>,
        x: Int,
        y: Int,
) {
    val current = heightMatrix.get(y).get(x).get(0)
    for (i in -1..1 step 2) {
        val xIndex = x + i
        val yIndex = y + i
        if (xIndex >= 0 && xIndex < heightMatrix.get(y).size) {
            val newPoint = heightMatrix.get(y).get(xIndex)
            if (newPoint.get(0) <= current) {
                heightMatrix.get(y).set(x, current + "*")
                if (!newPoint.endsWith("*")) {
                    findMin(
                            heightMatrix,
                            xIndex,
                            y,
                    )
                }
            } else {
                if (!newPoint.endsWith("*")) {
                    heightMatrix.get(y).set(xIndex, newPoint + "*")
                }
            }
        }
        if (yIndex >= 0 && yIndex < heightMatrix.size) {
            val newPoint = heightMatrix.get(yIndex).get(x)
            if (newPoint.get(0) <= current) {
                heightMatrix.get(y).set(x, current + "*")
                if (!newPoint.endsWith("*")) {
                    findMin(
                            heightMatrix,
                            x,
                            yIndex,
                    )
                }
            } else {
                if (!newPoint.endsWith("*")) {
                    heightMatrix.get(yIndex).set(x, newPoint + "*")
                }
            }
        }
    }
}

fun findBasins(
        heightMatrix: MutableList<MutableList<String>>,
        x: Int,
        y: Int,
        counter: MutableList<Int>
) {
    val current = heightMatrix.get(y).get(x).get(0)
    counter.set(0, counter.get(0).inc())
    heightMatrix.get(y).set(x, current + "*")
    for (i in -1..1 step 2) {
        val xIndex = x + i
        val yIndex = y + i
        if (xIndex >= 0 && xIndex < heightMatrix.get(y).size) {
            val newPoint = heightMatrix.get(y).get(xIndex)
            if (newPoint != "9" && !newPoint.endsWith("*")) {
                findBasins(heightMatrix, xIndex, y, counter)
            }
        }
        if (yIndex >= 0 && yIndex < heightMatrix.size) {
            val newPoint = heightMatrix.get(yIndex).get(x)

            if (newPoint != "9" && !newPoint.endsWith("*")) {
                findBasins(heightMatrix, x, yIndex, counter)
            }
        }
    }
}

fun main() {
    var heightMatrix: MutableList<MutableList<String>> = mutableListOf()
    File("input.txt").forEachLine { heightMatrix.add(it.map { it.toString() }.toMutableList()) }

    var counter = mutableListOf(0)
    var counterList: MutableList<Int> = mutableListOf()

    heightMatrix.forEachIndexed { y, heightRow ->
        heightRow.forEachIndexed { x, height ->
            if (height != "9" && !height.endsWith("*")) {
                findBasins(heightMatrix, x, y, counter)
                counterList.add(counter.get(0))
                counter.set(0, 0)
            }
        }
    }

    println(counterList.sorted().reversed().subList(0, 3).forEach { count -> println(count) })

    // heightMatrix.flatten().forEach { height ->
    //     if () {
    //         sum += height.toInt() + 1
    //     }
    // }
}
