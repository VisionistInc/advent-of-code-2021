import java.io.File;

var versions = mutableListOf<Int>()
var packets = mutableListOf<Pair<Int, List<Long>>>()

fun main(args : Array<String>) {
    val input = File(args.first()).readLines().first().trim()
    println("Solution 1: ${solution1(convertHexToBinary(input))}")
    // println("Solution 2: ${solution2(convertHexToBinary(input))}")
}
private fun solution1(binary: String) :Int {
    parse(binary, mutableListOf<Long>())
    return versions.sum()
}

private fun solution2(binary: String) :Long {
    parse(binary, mutableListOf<Long>())
    var nextResult = 0L
    for (x in packets.indices) {
        println("$x  -- ${packets.get(x)}")

        var packet = packets.get(x)
    }

    return nextResult
}

var evaluateImmidiately = setOf<Int>(0, 1)

private fun applyOperators(p :Pair<Int, List<Long>>) :Long {
    when (p.first) {
        0 -> return p.second.sum()
        1 -> return p.second.reduce{ acc, i -> i * acc }
        2 -> return p.second.minOf{ it }  
        3 -> return p.second.maxOf{ it }
        5 -> if (p.second.first() > p.second.last()) { return 1L } else { return 0L }
        6 -> if (p.second.first() < p.second.last()) { return 1L } else { return 0L }
        7 -> if (p.second.first() == p.second.last()) { return 1L } else { return 0L }
    }    
    return 0L
}

var position = 0

private fun parse(binary: String, vals: MutableList<Long>) :Pair<Int, Long> {
    var ptr = 0
    var decimal = -1L
    while (ptr != binary.length - 1 ) {
        var version = binary.slice(ptr..ptr + 2).toInt(2)
        versions.add(version)
        ptr += 3

        var type = binary.slice(ptr..ptr + 2).toInt(2)
        ptr += 3

        if (type == 4) {
            // literal
            var literal = StringBuilder()
            var keepGoing = binary.get(ptr).digitToInt()

            while (keepGoing == 1) {
                literal.append(binary.slice(ptr + 1..ptr + 4))
                ptr += 5
                keepGoing = binary.get(ptr).digitToInt()
            }

            // get the last bit since keepGoing is now a 0
            literal.append(binary.slice(ptr + 1..ptr + 4) )
            ptr += 5

            decimal = literal.toString().toLong(2)
            vals.add(decimal)

        } else {
            // operator
            var lengthTypeId = binary.get(ptr)

            if (lengthTypeId.equals('0')) {
                var values = mutableListOf<Long>()
                // the next 15 bits are a number that represents the total length in bits of the sub-packets contained by this packet
                var packetLength = binary.slice(ptr + 1..ptr + 15).toInt(2)
                ptr += 16
                var newBinary = binary.slice(ptr..ptr + (packetLength - 1))
                ptr += packetLength
                parse(newBinary, values)
                // println("ADDING   $position" + Pair<Int, List<Long>>(type, values) )
                position++

                // if (type in evaluateImmidiately) {
                //     var evaluated = applyOperators(Pair<Int, List<Long>>(type, values))
                //     values = mutableListOf<Long>(evaluated)
                // }

                packets.add( Pair<Int, List<Long>>(type, values) )
            } else {
                // the next 11 bits are a number that represents the number of sub-packets immediately contained by this packet.
                var values = mutableListOf<Long>()
                var packetCount = binary.slice(ptr + 1..ptr + 11).toInt(2)
                ptr += 12
                for (x in 1..packetCount - 1) {
                    var newBinary = binary.slice(ptr..(binary.length - 1))
                    var result = parse(newBinary, values)
                    ptr += result.first
                }
                // println("ADDING   $position" + Pair<Int, List<Long>>(type, values) )
                position++

                // if (type in evaluateImmidiately) {
                //     var evaluated = applyOperators(Pair<Int, List<Long>>(type, values))
                //     values = mutableListOf<Long>(evaluated)
                // }

                packets.add( Pair<Int, List<Long>>(type, values) )
            }
        }

        // are we left with just 0's?
        if (binary.slice(ptr..(binary.length -1)).filter{ it == '1' }.length == 0 ) {
            // no 1's in the remaining binary string, so move the pointer to the end
            ptr = binary.length - 1
        }
    }
    return Pair<Int, Long>(ptr, decimal)
} 

private fun convertHexToBinary(hex: String) :String {
    var i = 0
    val binary = StringBuilder()
    while (i < hex.length) {
        when(hex[i]) {
            '0'  -> binary.append("0000")
            '1'  -> binary.append("0001")
            '2'  -> binary.append("0010")
            '3'  -> binary.append("0011")
            '4'  -> binary.append("0100")
            '5'  -> binary.append("0101")
            '6'  -> binary.append("0110")
            '7'  -> binary.append("0111")
            '8'  -> binary.append("1000")
            '9'  -> binary.append("1001")
            'A', 'a'  -> binary.append("1010")
            'B', 'b'  -> binary.append("1011")
            'C', 'c'  -> binary.append("1100")
            'D', 'd'  -> binary.append("1101")
            'E', 'e'  -> binary.append("1110")
            'F', 'f'  -> binary.append("1111")
        }
        i++
    }
    return binary.toString()
}