import java.io.File

fun brokenTrash1() {
    var counter = 0
    File("input.txt").forEachLine {
        it.split(" | ")[1]
                .split(" ")
                .forEach({
                    when (it.length) {
                        2, 3, 4, 7 -> counter++
                    }
                })
    }
    println(counter)
}

fun brokenTrash2() {
    var counter = 0
    File("input.txt").forEachLine {
        var codeList: MutableList<List<Char>> = MutableList(10) { listOf() }
        val input = it.split(" | ")

        val patterns = input[0].split(" ")
        val fives: MutableList<String> = mutableListOf()
        val sixes: MutableList<String> = mutableListOf()

        // Set the freebies and sort the lengths of 5 and 6
        patterns.forEach({
            when (it.length) {
                2 -> codeList.set(1, it.toCharArray().sorted())
                3 -> codeList.set(7, it.toCharArray().sorted())
                4 -> codeList.set(4, it.toCharArray().sorted())
                5 -> fives.add(it)
                6 -> sixes.add(it)
                7 -> codeList.set(8, it.toCharArray().sorted())
            }
        })
        // Figure out the codes of length 6
        sixes.forEach({ sixCode ->
            val sixCodeAsCharList = sixCode.toCharArray().sorted()
            if (codeList.get(4).intersect(sixCodeAsCharList).size == codeList.get(4).size) {
                codeList.set(9, sixCodeAsCharList)
            } else if (codeList.get(7).intersect(sixCodeAsCharList).size == codeList.get(7).size) {
                codeList.set(0, sixCodeAsCharList)
            } else {
                codeList.set(6, sixCodeAsCharList)
            }
        })

        // Figure out the codes of length 5
        fives.forEach({ fiveCode ->
            val fiveCodeAsCharList = fiveCode.toCharArray().sorted()
            if (codeList.get(7).intersect(fiveCodeAsCharList).size == codeList.get(7).size) {
                codeList.set(3, fiveCodeAsCharList)
            } else if (codeList.get(6).intersect(fiveCodeAsCharList).size == fiveCodeAsCharList.size
            ) {
                codeList.set(5, fiveCodeAsCharList)
            } else {
                codeList.set(2, fiveCodeAsCharList)
            }
        })

        val codeListAsStrings = codeList.map({ code -> code.joinToString(separator = "") })

        var outputString = ""
        input[1].split(" ")
                .forEach({ output ->
                    outputString +=
                            codeListAsStrings.indexOf(
                                    output.toCharArray().sorted().joinToString(separator = "")
                            )
                })

        counter += outputString.toInt()
    }
    println(counter)
}

fun main() {
    brokenTrash1()
    brokenTrash2()
}
