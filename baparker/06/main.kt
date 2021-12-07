import java.io.File

fun getInitialTimerCounts(timerCounts: MutableList<Long>) {
    File("input.txt").forEachLine {
        it.split(",")
                .forEach({ timer ->
                    val timerNum = timer.toInt()
                    timerCounts.set(timerNum, timerCounts.get(timerNum).inc())
                })
    }
}

fun main(args: Array<String>) {
    val timerCounts: MutableList<Long> = mutableListOf(0, 0, 0, 0, 0, 0, 0, 0, 0)
    getInitialTimerCounts(timerCounts)

    for (day in 1..args[0].toInt()) {
        val zeroCount = timerCounts.get(0)
        timerCounts.removeAt(0)
        timerCounts.add(zeroCount)
        timerCounts.set(6, timerCounts.get(6) + zeroCount)
    }
    println(timerCounts.sum())
}
