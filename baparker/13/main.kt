import java.io.File

fun main() {
    val dotList: MutableList<Pair<Int, Int>> = mutableListOf()
    val foldList: MutableList<Pair<Int, Int>> = mutableListOf()

    File("input.txt").forEachLine {
        val inputArray = it.split(" ")
        if (inputArray.size > 1) {
            val fold = inputArray[2].split("=")
            if (fold[0] == "x") {
                foldList.add(Pair(fold[1].toInt(), 0))
            } else if (fold[0] == "y") {
                foldList.add(Pair(0, fold[1].toInt()))
            }
        } else {
            val coordinates = it.split(",")
            if (coordinates.size == 2) {
                dotList.add(Pair(coordinates[0].toInt(), coordinates[1].toInt()))
            }
        }
    }

    foldList.forEach { fold ->
        dotList.forEachIndexed { dotIndex, dot ->
            if (fold.first == 0) {
                val y = dot.second - fold.second
                if (y > 0) {
                    val newPair = Pair(dot.first, dot.second - (y * 2))
                    dotList.set(dotIndex, newPair)
                }
            } else if (fold.second == 0) {
                val x = dot.first - fold.first
                if (x > 0) {
                    val newPair = Pair(dot.first - (x * 2), dot.second)
                    dotList.set(dotIndex, newPair)
                }
            }
        }
        print(" -> " + dotList.toSet().size)
    }
    println()

    for (i in 0..dotList.maxOf { it.second }) {
        for (j in 0..dotList.maxOf { it.first }) {
            if (dotList.contains(Pair(j, i))) print("#") else print(".")
        }
        println()
    }
}
