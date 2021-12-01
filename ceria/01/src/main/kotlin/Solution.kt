import java.io.File;

fun main(args : Array<String>) {
    val input = File(args.first()).readLines().map { it.toInt() }
    println("Solution 1: ${solution1(input)}")
    println("Solution 2: ${solution2(input)}")
}

private fun solution1(input: List<Int>) :Int {
    var idx = 1
    var increaseCount = 0
    while (idx < input.size) {
        if (input.get(idx) - input.get(idx-1) > 0) {
            increaseCount++
        }
        idx++
    }
    return increaseCount
}
    
private fun solution2(input: List<Int>) :Int {
    var idx = 3
    var increaseCount = 0
    var prevSum = input.get(2) + input.get(1) + input.get(0)
    while (idx < input.size) {
        val currentSum = input.get(idx) + input.get(idx - 1) + input.get(idx - 2)
        if (currentSum - prevSum > 0) {
            increaseCount++
        }
        prevSum = currentSum
        idx++
    }
    return increaseCount
}   
