import java.io.File;

fun main(args : Array<String>) {
    val input = File(args.first()).readLines()
    println("Solution 1: ${solution1(input)}")
    println("Solution 2: ${solution2(input)}")
}

private fun solution1(input: List<String>) :Int {
    var x = 0
    var y = 0
    for (direction in input) {
        var split = direction.split(" ")
        when(split[0]) {
            "forward" -> x += split[1].toInt()
            "down" -> y += split[1].toInt()
            else -> {
                y -= split[1].toInt()
            }
        }
    }

    return x * y
}
    
private fun solution2(input: List<String>) :Int {
    var x = 0
    var y = 0
    var aim = 0
    for (direction in input) {
        var split = direction.split(" ")
        when(split[0]) {
            "down" -> aim += split[1].toInt()
            "up" -> aim -= split[1].toInt()
            else -> {
               x += split[1].toInt()
               y += aim * split[1].toInt()
            }
        }
    }

    return x * y
}   
