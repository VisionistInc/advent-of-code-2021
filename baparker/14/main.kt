import java.io.File

fun main(args: Array<String>) {
    val insertionRules: MutableMap<String, String> = mutableMapOf()
    var pairCount: MutableMap<String, Long> = mutableMapOf()
    val charCount: MutableMap<String, Long> = mutableMapOf()
    var polymerTemplate = ""

    File("input.txt").forEachLine {
        val input = it.split(" -> ")
        if (input.size > 1) {
            insertionRules.put(input[0], input[1])
            pairCount.put(input[0], 0)
        } else if (it.isNotEmpty()) {
            it.forEach { character ->
                val count = charCount.getOrPut(character.toString()) { 0 }
                charCount.set(character.toString(), count.inc())
            }
            polymerTemplate = it
        }
    }
    // Increment characters based on polymerTemplate
    polymerTemplate.windowed(2) { pair ->
        val pairString = pair.toString()
        pairCount.set(pairString, pairCount.getOrElse(pairString) { 0 }.inc())
    }

    for (i in 1..args.getOrElse(0) { "10" }.toInt()) {
        val tempCounterMap: MutableMap<String, Long> = pairCount.toMutableMap()
        pairCount.forEach {
            if (it.value > 0) {
                // subtract count for the current pair
                tempCounterMap.put(it.key, tempCounterMap.getOrElse(it.key) { 0 } - it.value)

                val newChar = insertionRules.getOrElse(it.key) { "" }
                val count = charCount.getOrPut(newChar) { 0 }
                // add number of current pairs to the new character that is added
                charCount.set(newChar, count.plus(it.value))
                val seq1 = it.key.get(0) + newChar
                // Add the count for the current pair to the first new sequence
                tempCounterMap.put(seq1, tempCounterMap.getOrElse(seq1) { 0 } + it.value)
                val seq2 = newChar + it.key.get(1)
                // Add the count for the current pair to the second new sequence
                tempCounterMap.put(seq2, tempCounterMap.getOrElse(seq2) { 0 } + it.value)
            }
        }
        pairCount = tempCounterMap.toMutableMap()
    }
    println(charCount.maxOf { it.value } - charCount.minOf { it.value })
}
