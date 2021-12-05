import java.io.File

fun getPowerConsumption() {
    var countList = IntArray(12)
    var lineCount = 0
    File("input.txt").forEachLine {
        for (index in it.indices) {
            countList[index] += it[index].digitToInt()
        }
        lineCount++
    }
    var gammaRate = ""
    var epsilonRate = ""
    for (count in countList) {
        if (lineCount / count <= 1) {
            gammaRate = gammaRate.plus("1")
            epsilonRate = epsilonRate.plus("0")
        } else {
            gammaRate = gammaRate.plus("0")
            epsilonRate = epsilonRate.plus("1")
        }
    }
    println(gammaRate)
    println(Integer.parseInt(epsilonRate, 2) * Integer.parseInt(gammaRate, 2))
}

fun main() {
    getPowerConsumption()
}
