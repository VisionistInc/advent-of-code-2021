import java.io.File

var timerCounts: MutableList<Long> = mutableListOf(0, 0, 0, 0, 0, 0, 0, 0, 0)

fun getInitialTimerCounts() {
    File("input.txt").forEachLine {
        it.split(",")
                .forEach({ timer ->
                    val timerNum = timer.toInt()
                    timerCounts.set(timerNum, timerCounts.get(timerNum).inc())
                })
    }
}

fun main(args: Array<String>) {
    var days = 1
    if (args.size > 0) {
        days = args[0].toInt()
    }

    getInitialTimerCounts()
    for (day in 1..days) {
        val zeroCount = timerCounts.get(0)
        timerCounts.removeAt(0)
        timerCounts.add(zeroCount)
        timerCounts.set(6, timerCounts.get(6) + zeroCount)
    }
    println(timerCounts.sum())
}
