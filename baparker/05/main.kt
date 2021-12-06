import java.io.File
import kotlin.math.abs

fun getOverlapCount() {
    var coordinateCount: MutableMap<String, Int> = mutableMapOf()
    File("input.txt").forEachLine {
        val line = it.split(" -> ")
        val start = line[0].split(",")
        val end = line[1].split(",")
        if (start[0].equals(end[0])) {
            val bounds = listOf(start[1].toInt(), end[1].toInt()).sorted()
            for (y in bounds[0]..bounds[1]) {
                val mapKey = start[0].plus(" ").plus(y.toString())
                var mapValue = coordinateCount.get(mapKey)
                if (mapValue != null) {
                    coordinateCount.put(mapKey, mapValue.inc())
                } else {
                    coordinateCount.put(mapKey, 1)
                }
            }
        } else if (start[1].equals(end[1])) {
            // Yay for dupe code
            val bounds = listOf(start[0].toInt(), end[0].toInt()).sorted()
            for (x in bounds[0]..bounds[1]) {
                val mapKey = x.toString().plus(" ").plus(start[1])
                var mapValue = coordinateCount.get(mapKey)
                if (mapValue != null) {
                    coordinateCount.put(mapKey, mapValue.inc())
                } else {
                    coordinateCount.put(mapKey, 1)
                }
            }
        } else {
            val startX = start[0].toInt()
            val startY = start[1].toInt()
            val endX = end[0].toInt()
            val endY = end[1].toInt()
            val diffX = endX - startX
            val diffY = endY - startY
            if (abs(diffX) == abs(diffY)) {
                var xModifier = 1
                if (diffX < 0) {
                    xModifier = -1
                }
                var yModifier = 1
                if (diffY < 0) {
                    yModifier = -1
                }

                for (i in 0..abs(diffX)) {
                    val mapKey =
                            (startX + (i * xModifier))
                                    .toString()
                                    .plus(" ")
                                    .plus((startY + (i * yModifier)).toString())
                    var mapValue = coordinateCount.get(mapKey)
                    if (mapValue != null) {
                        coordinateCount.put(mapKey, mapValue.inc())
                    } else {
                        coordinateCount.put(mapKey, 1)
                    }
                }
            }
        }
    }

    println(coordinateCount.values.filter({ count: Int -> count >= 2 }).size)
}

fun main() {
    getOverlapCount()
}
