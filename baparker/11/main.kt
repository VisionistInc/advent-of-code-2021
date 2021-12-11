import java.io.File

data class Octode(var energyLevel: Int) {
    var flashed = false
}

data class FlashCounter(var value: Int = 0) {
    fun flash() {
        this.value++
    }
}

fun isSynced(octrix: MutableList<MutableList<Octode>>): Boolean {
    var myCount = 0
    octrix.forEach { it.forEach { myCount += it.energyLevel } }
    if (myCount == 0) return true else return false
}

fun unFlash(octrix: MutableList<MutableList<Octode>>) {
    octrix.forEach { it.forEach { it.flashed = false } }
}

fun increaseOctoEnergy(
        octrix: MutableList<MutableList<Octode>>,
        x: Int,
        y: Int,
        flashCounter: FlashCounter
) {
    val octo = octrix.get(y).get(x)
    if (octo.energyLevel == 9) {
        flashCounter.flash()
        octo.flashed = true
        octo.energyLevel = 0
        for (i in -1..1) {
            for (j in -1..1) {
                if (!(i == 0 && j == 0)) {
                    val xIndex = x + j
                    val yIndex = y + i
                    val newPoint = octrix.getOrNull(xIndex)?.getOrNull(yIndex)
                    if (newPoint != null) {
                        increaseOctoEnergy(octrix, xIndex, yIndex, flashCounter)
                    }
                }
            }
        }
    } else {
        if (!octo.flashed) {
            octo.energyLevel++
        }
    }
}

fun main(args: Array<String>) {
    var numIter = -1
    if (args.size > 0) {
        numIter = args[0].toInt() - 1
    }
    val octrix: MutableList<MutableList<Octode>> = mutableListOf()
    File("input.txt").forEachLine { octrix.add(it.map { Octode(it.digitToInt()) }.toMutableList()) }

    val flashCounter = FlashCounter()
    var iter = 0
    var noSync = true
    while (noSync) {
        octrix.forEachIndexed({ y, octrow ->
            octrow.forEachIndexed { x, _ -> increaseOctoEnergy(octrix, x, y, flashCounter) }
        })
        if (isSynced(octrix)) {
            noSync = false
        }
        if (iter == numIter) {
            println(flashCounter.value)
        }
        unFlash(octrix)
        iter++
    }
    println(iter)
}
