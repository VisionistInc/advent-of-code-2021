import java.io.File

val VERSION_TYPEID_LENGTH = 3
val BIT_GROUP_SIZE = 4
val LITERAL_GROUP_SIZE = 5
val LENGTH_TYPE_ZERO_SIZE = 15
val LENGTH_TYPE_ONE_SIZE = 11

fun binFromHex(hex: String): String {
    return String.format(
                    "%" + BIT_GROUP_SIZE + "s",
                    Integer.toBinaryString(Integer.parseUnsignedInt(hex, 16))
            )
            .replace(" ".toRegex(), "0")
}

// Current spot in the binary string
var pointer = 0
var binString = ""
var versionSum = 0

fun literalValue(): Long {
    var literalString = ""

    var groupFlag = true
    var totalGroupSize = 0
    while (groupFlag) {
        val group = binString.substring(pointer, pointer + LITERAL_GROUP_SIZE)
        if (group.startsWith("0")) {
            groupFlag = false
        }

        literalString += group.substring(1)
        pointer += LITERAL_GROUP_SIZE
        totalGroupSize += LITERAL_GROUP_SIZE
    }
    return literalString.toLong(radix = 2)
}

fun versionAndTypeId(): Int {
    val value =
            Integer.parseUnsignedInt(
                    binString.substring(pointer, pointer + VERSION_TYPEID_LENGTH),
                    2
            )
    pointer += VERSION_TYPEID_LENGTH
    return value
}

fun traversePackets(): Long {
    val version = versionAndTypeId()
    versionSum += version
    val typeId = versionAndTypeId()

    if (typeId == 4) {
        return literalValue()
    } else {
        val lengthTypeId = binString.substring(pointer, pointer + 1)
        pointer++

        val valueList: MutableList<Long> = mutableListOf()

        if (lengthTypeId == "0") {
            var targetPointer =
                    Integer.parseUnsignedInt(
                            binString.substring(pointer, pointer + LENGTH_TYPE_ZERO_SIZE),
                            2
                    )

            pointer += LENGTH_TYPE_ZERO_SIZE
            targetPointer += pointer
            while (pointer < targetPointer) {
                valueList.add(traversePackets())
            }
        } else {
            val numPackets =
                    Integer.parseUnsignedInt(
                            binString.substring(pointer, pointer + LENGTH_TYPE_ONE_SIZE),
                            2
                    )
            pointer += LENGTH_TYPE_ONE_SIZE

            for (i in 0 until numPackets) {
                valueList.add(traversePackets())
            }
        }

        var result: Long = 0
        when (typeId) {
            0 -> result = valueList.sum()
            1 -> result = valueList.fold(1) { prod, value -> value * prod }
            2 -> result = valueList.minOrNull() ?: 0
            3 -> result = valueList.maxOrNull() ?: 0
            5 -> result = if (valueList.get(0).toInt().compareTo(valueList.get(1)) == 1) 1 else 0
            6 -> result = if (valueList.get(1).toInt().compareTo(valueList.get(0)) == 1) 1 else 0
            7 -> result = if (valueList.get(0).toInt().compareTo(valueList.get(1)) == 0) 1 else 0
        }

        if (binString.substring(pointer).length < VERSION_TYPEID_LENGTH * 2) {
            pointer += binString.substring(pointer).length
        }
        return result
    }
}

fun main() {
    File("input.txt").forEachLine { it.forEach { binString += binFromHex(it.toString()) } }
    println("Expression Result: " + traversePackets())
    println("Version Sum: " + versionSum)
}
