import java.io.File;
import kotlin.math.abs;

val folds = mutableListOf<Pair<String, Int>>()

fun main(args : Array<String>) {
    val input = File(args.first()).readLines()
    val points = mutableSetOf<Pair<Int, Int>>()

    var foldingInstruction = false
    for (line in input) {
        if (line.isEmpty()) {
            foldingInstruction = true
            continue
        }

        if (foldingInstruction) {
            var fold = line.replace("fold along ", "").split("=")
            folds.add(Pair<String, Int>(fold[0], fold[1].toInt()))
        } else {
            var point = line.split(",").map{ it.toInt() }
            points.add(Pair<Int, Int>(point[0], point[1]))
        }
    }

    println("Solution 1: ${solution1(points)}")
    println("Solution 2: ")
    solution2(points)
}

private fun solution1(p: MutableSet<Pair<Int, Int>>) :Int {
    var points = p
    if (folds.first().first.equals("y")) {
        points = foldUp(points, folds.first().second)
    } else {
        points = foldLeft(points, folds.first().second)
    }

    return points.size
}

private fun solution2(p: MutableSet<Pair<Int, Int>>) {
    var points = p
    for (f in folds) {
        if (f.first.equals("y")) {
            points = foldUp(points, f.second)
        } else {
            points = foldLeft(points, f.second)
        }
    }
    
    var maxRows = points.map{ it.second }.maxOrNull()!!
    var maxCols = points.map{ it.first }.maxOrNull()!!
    for (row in 0..maxRows) {
        for (col in 0..maxCols) {
            if (points.contains(Pair<Int,Int>(col, row))) {
                print("#")
            } else {
                print(".")
            }
        } 
        println()
    }

}   

private fun foldUp(points: MutableSet<Pair<Int, Int>>, y: Int) :MutableSet<Pair<Int, Int>> {
    val newPoints = mutableSetOf<Pair<Int, Int>>()
    for (p in points) {
        if (p.second > y) {
            newPoints.add(Pair<Int, Int>(p.first, abs((p.second - y) - y)))
        } else {
            newPoints.add(p)
        }
    }
    return newPoints
}

private fun foldLeft(points: MutableSet<Pair<Int, Int>>, x: Int) :MutableSet<Pair<Int, Int>> {
    val newPoints = mutableSetOf<Pair<Int, Int>>()
    for (p in points) {
        if (p.first > x) {
            newPoints.add(Pair<Int, Int>(abs((p.first - x) - x), p.second))
        } else {
            newPoints.add(p)
        }
    }

    return newPoints
}