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
    var day = 0
    var fish = input.toIntArray()
    var newFishToAdd = 0
    while (day != days) {
        for (f in fish.indices) {
            when(fish[f]) {
                0 -> {
                    fish[f] = 6
                    newFishToAdd++
                }
                else -> fish[f] = fish[f] - 1
            }
        }

        val modifiedFish = fish.copyOf(fish.size + newFishToAdd)
        var x = 0
        while (x < newFishToAdd) {
            modifiedFish[fish.size + x] = 8
            x++
        }

        fish = modifiedFish
        newFishToAdd = 0
        day++
    }

    return fish.size.toLong()
}
