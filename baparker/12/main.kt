import java.io.File
import kotlin.collections.mutableListOf

open class Cave(val name: String) {
    val destList: MutableList<String> = mutableListOf()
}

class SmallCave(name: String) : Cave(name) {
    var visited = false
}

var counter = 0

fun traverseCaves(
        caveMap: MutableMap<String, Cave>,
        caveName: String,
        secondVisit: Boolean = false,
        path: MutableList<String> = mutableListOf(),
) {
    val currentCave = caveMap.get(caveName)
    path.add(caveName)

    if (caveName != "end") {

        currentCave?.destList?.forEach {
            val nextCave = caveMap.get(it)
            if (nextCave is SmallCave) {
                if (!nextCave.visited) {
                    nextCave.visited = true
                    traverseCaves(caveMap, it, secondVisit, path)
                    nextCave.visited = false
                } else if (!secondVisit) {
                    traverseCaves(caveMap, it, true, path)
                }
            } else {
                traverseCaves(caveMap, it, secondVisit, path)
            }
        }
    } else {
        counter++
    }
    path.removeLast()
}

fun addNodes(caveMap: MutableMap<String, Cave>, first: String, second: String) {
    if (second != "start") {
        val cave =
                caveMap.getOrPut(first) {
                    if (first[0].isLowerCase()) SmallCave(first) else Cave(first)
                }
        cave.destList.add(second)
    }
}

fun main(args: Array<String>) {
    var useSecondVisit = !args.getOrElse(0) { "false" }.toBoolean()

    val caveMap: MutableMap<String, Cave> = mutableMapOf()
    File("input.txt").forEachLine {
        val input = it.split("-")

        addNodes(caveMap, input[0], input[1])
        addNodes(caveMap, input[1], input[0])
    }

    traverseCaves(caveMap, "start", useSecondVisit)

    println(counter)
}
