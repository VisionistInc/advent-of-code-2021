import java.io.File
import java.util.Collections

fun main(args: Array<String>) {
    val timerCounts: MutableList<Long> = MutableList(9) { 0 }
    File("input.txt").forEachLine {
        it.split(",")
                .forEach({ 
                    val timerNum = it.toInt()
                    timerCounts.set(timerNum, timerCounts.get(timerNum).inc())
                })
    }

    for (day in 1..args[0].toInt()) {
        Collections.rotate(timerCounts, -1)
        timerCounts.set(6, timerCounts.get(6) + timerCounts.last())
    }
    println(timerCounts.sum())
}
