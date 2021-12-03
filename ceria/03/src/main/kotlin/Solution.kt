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
            val biVal = diagnostic.get(c)
            if (biVal.equals('0')) {
                var count = zerosMap.get(c)?.let{
                    zerosMap.get(c)
                } ?: 0
                zerosMap.put(c, count + 1)
            }
        }
    }

    for (c in input.get(0).indices) {
        var count = zerosMap.get(c)?.let{
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
    var o2Pos = 0
    var o2Input = filter(input, o2Pos, false)
    while (o2Input.size > 1) {
        o2Pos += 1
        o2Input = filter(o2Input, o2Pos, false)
    }

    var co2Pos = 0
    var co2Input = filter(input, co2Pos, true)
    while (co2Input.size > 1) { 
        co2Pos += 1
        co2Input = filter(co2Input, co2Pos, true)
    }

    return o2Input.get(0).toLong(2) * co2Input.get(0).toLong(2)
}  

private fun filter(input: List<String>, pos: Int, isCo2Filter: Boolean): List<String> {
    val threshold = input.size / 2
    var filterVal = '1'

    var count = 0
    for (diagnostic in input) {
        val biVal = diagnostic.get(pos)
        if (biVal.equals('0')) {
            count++
        }
    }

    if (count > threshold) {
        filterVal = '0'
    }

    if (isCo2Filter) {
        if (filterVal.equals('1')){
            filterVal = '0'
        } else {
            filterVal = '1'
        }
    }

    return input.filter { it.get(pos) == filterVal }
}