import java.io.File;

fun main(args : Array<String>) {
    val input = File(args.first()).readLines().first().split(",").map { it.toInt() }
    println("Solution 1: ${solution1(input)}")
    println("Solution 2: ${solution2(input)}")
}

private fun solution1(input: List<Int>) :Long {
    return copulate(80, input)
}

private fun solution2(input: List<Int>) :Long {
    return copulate(256, input)
}   

private fun copulate(days: Int, input: List<Int>) :Long {
    var fish = input.groupingBy { it }.eachCount().mapValues { it.value.toLong() }
    var newFish = mutableMapOf<Int, Long>()
    var day = 0
    while (day != days) {

        // For each of the possible fish days
        for (x in 8 downTo 0) {
            val count = fish.get(x)?.let{
                fish.get(x)
            } ?: 0

            if (x == 0) {
                newFish.put(8, count)
                val sixCount = newFish.get(6)?.let{
                    newFish.get(6)
                } ?: 0
                newFish.put(6, sixCount + count)
            } else {
                newFish.put(x-1, count)
            }    
        }

        fish = newFish
        newFish = mutableMapOf<Int, Long>()
        day++
    }

    return fish.values.sum()
}
