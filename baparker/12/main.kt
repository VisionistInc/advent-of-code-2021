import java.io.File
import kotlin.collections.mutableListOf

open class Cave(val name: String) {
    val destList: MutableList<String> = mutableListOf()
}

class SmallCave(name: String) : Cave(name) {
    var visited = false
}

var counter1 = 0

fun traverseCaves(
        caveMap: MutableMap<String, Cave>,
        caveName: String,
        path: MutableList<String> = mutableListOf(),
        secondVisit: Boolean = false
) {
    val currentCave = caveMap.get(caveName)

    if (caveName != "end") {
        path.add(caveName)

        currentCave?.destList?.forEach {
            val nextCave = caveMap.get(it)
            if (nextCave is SmallCave) {
                if (!nextCave.visited) {
                    nextCave.visited = true
                    traverseCaves(caveMap, it, path, secondVisit)
                    nextCave.visited = false
                } else if (!secondVisit) {
                    traverseCaves(caveMap, it, path, true)
                }
            } else {
                traverseCaves(caveMap, it, path, secondVisit)
            }
        }
        path.removeLast()
    } else {
        counter1++
    }
}

fun main() {
    val caveMap: MutableMap<String, Cave> = mutableMapOf()
    File("input.txt").forEachLine {
        val input = it.split("-")
        val src = input[0]
        val dest = input[1]

        if (dest != "start") {
            val srcCave =
                    caveMap.getOrPut(src) {
                        if (src[0].isLowerCase()) SmallCave(src) else Cave(src)
                    }
            srcCave.destList.add(dest)
        }
        if (src != "start") {
            val destCave =
                    caveMap.getOrPut(dest) {
                        if (dest[0].isLowerCase()) SmallCave(dest) else Cave(dest)
                    }
            destCave.destList.add(src)
        }
    }

    traverseCaves(caveMap, "start")

    println(counter1)

    // caveMap.forEach {
    //     print(it.key + " ")
    //     println(it.value.destList)
    // }
}
