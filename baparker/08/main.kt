import java.io.File

val NON_UNIQUE_CODE_MAP =
        mapOf("abcdeg" to 0, "abcdf" to 2, "acdfg" to 3, "bcdef" to 5, "bcdefg" to 6, "abcdfg" to 9)

fun brokenTrash1() {
    val counter = mutableListOf(0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
    File("input.txt").forEachLine {
        it.split("|")[1]
                .split(" ")
                .forEach({
                    when (it.length) {
                        2 -> counter.set(0, counter.get(0).inc())
                        3 -> counter.set(1, counter.get(1).inc())
                        4 -> counter.set(2, counter.get(2).inc())
                        7 -> counter.set(3, counter.get(3).inc())
                    }
                })
    }
    println(counter.sum())
}

fun brokenTrash2() {
    var counter = 0
    File("input.txt").forEachLine {
        var codeString = ""

        it.split(" | ")[1]
                .split(" ")
                .forEach({
                    println(it.toCharArray().sorted().joinToString(separator = ""))
                    when (it.length) {
                        2 -> codeString += 1
                        3 -> codeString += 7
                        4 -> codeString += 4
                        7 -> codeString += 8
                        else ->
                                codeString +=
                                        NON_UNIQUE_CODE_MAP.get(
                                                it.toCharArray()
                                                        .sorted()
                                                        .joinToString(separator = "")
                                        )
                    }
                })
        counter += codeString.toInt()
    }
    println(counter)
}

fun main() {
    brokenTrash1()
    brokenTrash2()
}
