import java.io.File
import kotlin.collections.mutableListOf

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
        if (count != 0 && lineCount / count <= 1) {
            gammaRate = gammaRate.plus("1")
            epsilonRate = epsilonRate.plus("0")
        } else {
            gammaRate = gammaRate.plus("0")
            epsilonRate = epsilonRate.plus("1")
        }
    }
    println(Integer.parseInt(epsilonRate, 2) * Integer.parseInt(gammaRate, 2))
}

fun greaterThanEqualTo(num1: Int, num2: Int): Boolean {
    return num1 >= num2
}

fun lessThan(num1: Int, num2: Int): Boolean {
    return num1 < num2
}

fun getRating(list: MutableList<MutableList<Int>>, compareFunc: (Int, Int) -> Boolean): Int {
    var ratingList = list.toList()
    for (column in list.indices) {
        if (ratingList.size > 1) {
            var zeroList: MutableList<MutableList<Int>> = mutableListOf()
            var oneList: MutableList<MutableList<Int>> = mutableListOf()
            for (row in ratingList.indices) {
                val bit = ratingList.get(row).get(column)
                if (bit == 0) {
                    zeroList.add(ratingList.get(row))
                } else {
                    oneList.add(ratingList.get(row))
                }
            }
            if (compareFunc(oneList.size, zeroList.size)) {
                ratingList = oneList
            } else {
                ratingList = zeroList
            }
        }
    }

    return Integer.parseInt(ratingList.get(0).joinToString(""), 2)
}

fun getLifeSupportAndOxygenGenerator() {
    var lineCount = 0
    var dataList: MutableList<MutableList<Int>> = mutableListOf()

    File("input.txt").forEachLine {
        var lineList: MutableList<Int> = mutableListOf()
        for (index in it.indices) {
            val bit = it[index].digitToInt()
            lineList.add(bit)
        }
        dataList.add(lineList)
        lineCount++
    }

    var o2GenRating = getRating(dataList, ::greaterThanEqualTo)
    var cO2ScrubberRating = getRating(dataList, ::lessThan)

    println(o2GenRating * cO2ScrubberRating)
}

fun main() {
    getPowerConsumption()
    getLifeSupportAndOxygenGenerator()
}
