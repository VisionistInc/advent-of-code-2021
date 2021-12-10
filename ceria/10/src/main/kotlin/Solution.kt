import java.io.File;
import java.util.Stack;

val matches = mapOf<Char, Char>(
    ')' to '(',
    ']' to '[',
    '}' to '{', 
    '>' to '<'
)
val leftChars = setOf<Char>('(', '[', '{', '<')
val corruptedValues = mapOf<Char, Int>(
    ')' to 3, 
    ']' to 57,
    '}' to 1197, 
    '>' to 25137
)

fun main(args : Array<String>) {
    val input = File(args.first()).readLines()
    println("Solution 1: " + solution1(input))
    
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

// private fun solution2(input: List<String>) :Int {
//    return 0
// }