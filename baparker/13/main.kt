import java.io.File

fun main() {
    val dotList: MutableList<Pair<Int, Int>> = mutableListOf()
    val foldList: MutableList<Pair<Int, Int>> = mutableListOf()
    var highestX = 0
    var highestY = 0
    File("input.txt").forEachLine {
        if (it.startsWith("fold along")) {
            val fold = it.split(" ")[2].split("=")
            if (fold[0] == "x") {
                foldList.add(Pair(fold[1].toInt(), 0))
            } else if (fold[0] == "y") {
                foldList.add(Pair(0, fold[1].toInt()))
            }
        } else {
            val coordinates = it.split(",")
            if (coordinates.size == 2) {
                if (coordinates[0].toInt() > highestX) {
                    highestX = coordinates[0].toInt()
                }
                if (coordinates[1].toInt() > highestY) {
                    highestY = coordinates[1].toInt()
                }
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
    }

    // Probably a much better way to get these numbers for the sizes and loops but I am lazy. To be
    // improved I suppose
    val board: List<MutableList<Char>> = List(40) { MutableList(10) { '.' } }

    dotList.toSet().forEach { board.get(it.first).set(it.second, '#') }

    for (i in 0..5) {
        for (j in 0..38) {
            print(board.get(j).get(i))
        }
        println()
    }
}
