import java.io.File

val OPENERS = listOf('(', '[', '{', '<')
val CLOSERS = listOf(')', ']', '}', '>')
val PENALTIES = listOf(3, 57, 1197, 25137)
val GOOD_VALUES = listOf(1, 2, 3, 4)

fun parseLineAndFindFirstError(line: String): Int {
    val openerStack: MutableList<Char> = mutableListOf()

    line.forEach {
        if (OPENERS.contains(it)) {
            openerStack.add(it)
        } else {
            val prev = openerStack.last()
            val openerIndex = OPENERS.indexOf(prev)
            val closerIndex = CLOSERS.indexOf(it)
            if (closerIndex != openerIndex) {
                val closerValue = PENALTIES.get(closerIndex)
                return closerValue
            } else {
                openerStack.removeLast()
            }
        }
    }

    return 0
}

fun parseLineAndComplete(line: String): Long {
    val openerStack: MutableList<Char> = mutableListOf()

    line.forEach {
        if (OPENERS.contains(it)) {
            openerStack.add(it)
        } else {
            val prev = openerStack.last()
            if (CLOSERS.indexOf(it) != OPENERS.indexOf(prev)) {
                return 0
            } else {
                openerStack.removeLast()
            }
        }
    }
    return openerStack.reversed().fold(0) { sum, opener ->
        val openerIndex = OPENERS.indexOf(opener)
        (sum * 5) + GOOD_VALUES.get(openerIndex)
    }
}

fun main() {
    var errorCounter = 0
    var pointsList: MutableList<Long> = mutableListOf()
    File("input.txt").forEachLine {
        errorCounter += parseLineAndFindFirstError(it)
        val linePoints = parseLineAndComplete(it)
        if (linePoints > 0) pointsList.add(linePoints)
    }
    println(errorCounter)
    println(pointsList.sorted().get((pointsList.size - 1) / 2).toBigDecimal().toPlainString())
}
