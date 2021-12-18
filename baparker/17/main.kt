import kotlin.math.abs
import kotlin.math.sqrt

var xMin = 0
var xMax = 0
var yMin = 0
var yMax = 0

fun step(curX: Int, curY: Int, tarX: Int, tarY: Int): Boolean {
    val nextX = curX + tarX
    val nextY = curY + tarY

    if ((nextX >= xMin && nextX <= xMax) && (nextY >= yMin && nextY <= yMax)) {
        return true
    } else if (nextX > xMax || nextY < yMin) {
        return false
    }
    return step(nextX, nextY, if (tarX > 0) tarX - 1 else 0, tarY - 1)
}

fun getTri(num: Int): Int {
    return num * (num + 1) / 2
}

fun getTriRoot(num: Int): Int {
    return sqrt(2.0 * num).toInt()
}

fun main(args: Array<String>) {
    xMin = args.getOrElse(0) { "241" }.toInt()
    xMax = args.getOrElse(1) { "273" }.toInt()
    yMin = args.getOrElse(2) { "-97" }.toInt()
    yMax = args.getOrElse(3) { "-63" }.toInt()

    val maxHeight = getTri(abs(yMin)) + yMin

    println("Max Height: " + maxHeight)

    val maxYTraj = getTriRoot(maxHeight)
    val minXTraj = getTriRoot(xMin)

    val validTrajectories: MutableList<Pair<Int, Int>> = mutableListOf()

    for (x in minXTraj..xMax) {
        for (y in yMin..maxYTraj) {
            if (step(0, 0, x, y)) {
                validTrajectories.add(Pair(x, y))
            }
        }
    }

    println("Possible Trajectories: " + validTrajectories.size)
}
