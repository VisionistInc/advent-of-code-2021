import java.io.File;

fun main(args : Array<String>) {
    val input = File(args.first()).readLines()
    println("Solution 1: ${solution1(input)}")
    println("Solution 2: ${solution2(input)}")
}

private fun solution1(input: List<String>) :Long {
    val threshold = input.size / 2
    var zerosMap = mutableMapOf<Int, Int>()
    var epsilon = ""
    var gamma = ""

    for (diagnostic in input) {
        for (c in diagnostic.indices) {
            if (diagnostic.get(c).equals('0')) {
                val count = zerosMap.get(c)?.let{
                    zerosMap.get(c)
                } ?: 0
                zerosMap.put(c, count + 1)
            }
        }
    }

    for (c in input.get(0).indices) {
        val count = zerosMap.get(c)?.let{
            zerosMap.get(c)
        } ?: 0
        if (count > threshold) {
            gamma += '0'
            epsilon += '1'
        } else {
            gamma += '1'
            epsilon += '0'
        }
    }

    return gamma.toLong(2) * epsilon.toLong(2)
}
    
private fun solution2(input: List<String>) :Long {
    return filter(input, 0, false).get(0).toLong(2) * filter(input, 0, true).get(0).toLong(2)
}  

private fun filter(input: List<String>, pos: Int, isCo2Filter: Boolean): List<String> {
    var filterVal = '1'

    var count = 0
    for (diagnostic in input) {
        if (diagnostic.get(pos).equals('0')) {
            count++
        }
    }

    if (count > input.size / 2) {
        filterVal = '0'
    }

    if (isCo2Filter) {
        if (filterVal.equals('1')){
            filterVal = '0'
        } else {
            filterVal = '1'
        }
    }

    val filtered = input.filter { it.get(pos) == filterVal }
    if (filtered.size > 1) {
        return filter(filtered, pos + 1, isCo2Filter)
    }
    return filtered
}