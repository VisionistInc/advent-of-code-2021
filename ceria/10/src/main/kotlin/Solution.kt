import java.io.File;
import java.util.Stack;

val matches = mapOf<Char, Char>(
    ')' to '(',
    ']' to '[',
    '}' to '{', 
    '>' to '<'
)
val leftChars = matches.values.toSet()
val corruptedValues = mapOf<Char, Int>(
    ')' to 3, 
    ']' to 57,
    '}' to 1197, 
    '>' to 25137
)
val scoreValues = mapOf<Char, Int>(
    ')' to 1, 
    ']' to 2,
    '}' to 3, 
    '>' to 4
)

fun main(args : Array<String>) {
    val input = File(args.first()).readLines()
    println("Solution 1: " + solution1(input))
    println("Solution 2: " + solution2(input))
    
    // CANNOT USE STRING INTERPLATION BECAUSE OF SPECIAL CHARACTERS IN THE INPUT
    // println("Solution 1: ${solution1(input)}")
    // println("Solution 2: ${solution2(input)}")
}

private fun solution1(input: List<String>) :Int {
    var corrupted = mutableListOf<Char>()
    for (line in input) {
        val syntax = Stack<Char>()
        for (c in line) {
            if (c in leftChars) {   
                syntax.push(c)
            } else {
                val shouldMatch = syntax.pop()
                if (!shouldMatch.equals(matches.get(c))) {
                    corrupted.add(c)
                    break
                }
            }
        }
    }
    var corruptedCount = corrupted.groupingBy { it }.eachCount()
    var sum = 0
    for ((k, v)  in corruptedValues) {
        sum += corruptedCount.get(k)!! * v
    }
    return sum
}

private fun solution2(input: List<String>) :Long {
    val reversedMatches = matches.entries.associate { (key, value) -> value to key }
    val scores = mutableListOf<Long>()

    for (line in input) {
        val syntax = Stack<Char>()
        var broke = false
        for (c in line) {
            if (c in leftChars) {   
                syntax.push(c)
            } else {
                val shouldMatch = syntax.pop()
                if (!shouldMatch.equals(matches.get(c))) {
                    broke = true
                    break
                }
            }
        }

        if (!broke) {
            var completion = ""
            while (!syntax.isEmpty()) {
                val match = syntax.pop()
                completion += reversedMatches.get(match)
            }

            var score = 0L
            for (c in completion) {
                score = (score * 5) + scoreValues.get(c)!!
            }
            scores.add(score)
        }
    }

    scores.sort()
    return scores.get(scores.size/2)  // since the list is 0 based, dividing the size by 2 and ignoring the remainder gives the middle index.
}