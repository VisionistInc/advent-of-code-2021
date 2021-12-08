import java.io.File;

val knownDigits = mapOf<Int, Int>(2 to 1, 3 to 7, 4 to 4, 7 to 8)

fun main(args : Array<String>) {
    var inputMap = mutableMapOf<String, String>()
    File(args.first()).readLines().map { 
        val input = it.split(" | ")
        inputMap.put(input[0], input[1])
    }
    println("Solution 1: ${solution1(inputMap)}")
    println("Solution 2: ${solution2(inputMap)}")
}

private fun solution1(inputMap: Map<String,String>) :Int {
    val knownDigitLengths = knownDigits.keys
    var count = 0
    for (output in inputMap.values) {
        count += output.split(" ").filter{ it.length in knownDigitLengths }.size
    }
    return count
}


/**
 * I realized half way through that I could've created a structure to represent the 7 segments, and then determined 
 * what numbers map to each setgment, and then mapped the segments, but I did the dumb brute force solution.
 * Perhaps I'll come back and do the better solution as time allows
 */
private fun solution2(inputMap: Map<String,String>) :Int {
    var translated = mutableListOf<Int>()

    for ((signalVal, outputVal) in inputMap) {
        var digitsTranslator = mutableMapOf<String, String>()
        val sig = signalVal.split(" ")
        val out = outputVal.split(" ")
        
        // Determine the letters of the known unique numbers (2 letters = 1, 3 letters = 7, 4 letters = 4 7 letters =  8)
        for ((digitLength, digit) in knownDigits) {
            digitsTranslator.put(digit.toString(), sig.filter{ it.length == digitLength }.first())
        }

        // 5 letters could be either 2, 3, or 5.  
        // 3 letters will be in common among all of them, so filter those out
        var fiveLetters = sig.filter{ it.length == 5 }.toMutableList()
        var sixLetters = sig.filter{ it.length == 6 }.toMutableList()
        var common = mutableListOf<String>()
        for (c in fiveLetters.first()) {
            if (fiveLetters[0].contains(c) && fiveLetters[1].contains(c) && fiveLetters[2].contains(c)) {
                common.add(c.toString())
            }
        }
        var leftOversToOrig = mutableMapOf<String, String>()
        for (code in fiveLetters) {
            var removed = code
            for (letter in common) {
                removed.replace(letter, "")
            }
            leftOversToOrig.put(removed, code)
        }

        // the cdoe for 3 will be left with the same two letters as is the code for the 1 - order won't be guaranteed though
        val codeForOne = digitsTranslator.get("1")!!.toList()
        for (leftover in leftOversToOrig.keys) {
            if (leftover.toList().containsAll(codeForOne)) {
                // this is the case for the 3 code
                digitsTranslator.put("3", leftOversToOrig.get(leftover)!!)
                fiveLetters.remove(leftOversToOrig.get(leftover)!!)
                break
            } 
        }

        // the code for 9 will contain all of the code for 3, plus 1 extra
        val codeForThree = digitsTranslator.get("3")!!.toList()
        for (code in sixLetters) {
            if (code.toList().containsAll(codeForThree)) {
                digitsTranslator.put("9", code)
                sixLetters.remove(code)
                break
            }
        }

        // the code for 9 will contain all of the code for 5, but not all of the code for 2,
        // which we will know the code for 2 because it'll the the only code left in fiveLetters
        val codeForNine = digitsTranslator.get("9")!!.toList()
        for (code in fiveLetters) {
            if (codeForNine.containsAll(code.toList())) {
                digitsTranslator.put("5", code)
            } else {
                digitsTranslator.put("2", code)
            }
        }

        // the code for 0 will contain both characters from the code for 1. The code for 6 will not.
        for (code in sixLetters) {
            if (code.toList().containsAll(codeForOne.toList())) {
                digitsTranslator.put("0", code)
            } else {
                digitsTranslator.put("6", code)
            }
        }

        var output = ""
        for (o in out) {
            for ((digit, code) in digitsTranslator) {
                if (code.length == o.length && o.toList().containsAll(code.toList())) {
                    output = output + digit
                    break
                }
            }    
        }
        translated.add(output.toInt())
    }

    return translated.sum()
}   
