import java.io.File

fun pt1() {
    var x = 0
    var y = 0

    val directionValues =
            mapOf(
                    "forward" to { value: Int -> x += value },
                    "up" to { value: Int -> y -= value },
                    "down" to { value: Int -> y += value }
            )

    File("input.txt").forEachLine {
        val direction = it.split(" ")
        directionValues[direction[0]]?.invoke(direction[1].toInt())
    }
    println(x * y)
}

fun pt2() {
    var aim = 0
    var x = 0
    var y = 0

    val directionValues =
            mapOf(
                    "forward" to
                            { value: Int ->
                                x += value
                                y += aim * value
                            },
                    "up" to { value: Int -> aim -= value },
                    "down" to { value: Int -> aim += value }
            )

    File("input.txt").forEachLine {
        val direction = it.split(" ")
        directionValues[direction[0]]?.invoke(direction[1].toInt())
    }
    println(x * y)
}

fun main() {
    pt1()
    pt2()
}
