import java.io.File;

var ventLines = mutableListOf<Pair<Pair<Int, Int>, Pair<Int, Int>>>()

fun main(args : Array<String>) {
    arrangeInput(File(args.first()).readLines())
    println("Solution 1: ${solution1()}")
    println("Solution 2: ${solution2()}")
}

private fun solution1() :Int {
    var ventCounts = mutableMapOf<Pair<Int, Int>, Int>()
    for (line in ventLines) {
        if (line.first.first == line.second.first) {
            // x of both points are the same, this is a vertical line
            var start = line.first.second
            var end = line.second.second
            if (start > end) {
                // flip the start and end so we don't have to use downTo in the loop
                start = line.second.second
                end = line.first.second
            } 

            for (y in start..end) {
                val newPoint = Pair<Int, Int>(line.first.first, y)
                val count = ventCounts.get(newPoint)?.let{
                    ventCounts.get(newPoint)
                } ?: 0
                ventCounts.put(newPoint, count + 1)    
            }

        } else 
            // Must use an if statement here, because the input looks like it includes diagnols...
            if (line.first.second == line.second.second) {
                var start = line.first.first
                var end = line.second.first
                if (start > end) {
                    // flip the start and end so we don't have to use downTo in the loop
                    start = line.second.first
                    end = line.first.first
                } 

                // y of both points are the same, this is a horizontal line
                for (x in start..end) {
                    val newPoint = Pair<Int, Int>(x, line.first.second)
                    val count = ventCounts.get(newPoint)?.let{
                        ventCounts.get(newPoint)
                    } ?: 0
                    ventCounts.put(newPoint, count + 1)    
                }
            }
    }
    
    return ventCounts.values.filter { it > 1 }.count()
}
    
private fun solution2() :Int {
    var ventCounts = mutableMapOf<Pair<Int, Int>, Int>()
    for (line in ventLines) {
        if (line.first.first == line.second.first) {
            // x of both points are the same, this is a vertical line
            var start = line.first.second
            var end = line.second.second
            if (start > end) {
                // flip the start and end so we don't have to use downTo in the loop
                start = line.second.second
                end = line.first.second
            } 

            for (y in start..end) {
                val newPoint = Pair<Int, Int>(line.first.first, y)
                val count = ventCounts.get(newPoint)?.let{
                    ventCounts.get(newPoint)
                } ?: 0
                ventCounts.put(newPoint, count + 1)    
            }

        } else 
            // Must use an if statement here, because the input looks like it includes diagnols...
            if (line.first.second == line.second.second) {
                var start = line.first.first
                var end = line.second.first
                if (start > end) {
                    // flip the start and end so we don't have to use downTo in the loop
                    start = line.second.first
                    end = line.first.first
                } 

                // y of both points are the same, this is a horizontal line
                for (x in start..end) {
                    val newPoint = Pair<Int, Int>(x, line.first.second)
                    val count = ventCounts.get(newPoint)?.let{
                        ventCounts.get(newPoint)
                    } ?: 0
                    ventCounts.put(newPoint, count + 1)    
                }
            }
        else {
            // This is a diagonal line
            // var startPoint = line.first
            val count = ventCounts.get(line.first)?.let{
                ventCounts.get(line.first)
            } ?: 0
            ventCounts.put(line.first, count + 1) 
            var x = line.first.first
            var y = line.first.second
            while (x != line.second.first) {    
                if (x > line.second.first) {
                    x -= 1
                } else {
                    x += 1
                }

                if (y < line.second.second) {
                    y += 1
                } else {
                    y -= 1
                }

                val newPoint = Pair<Int, Int>(x, y)
                val updatedCount = ventCounts.get(newPoint)?.let{
                    ventCounts.get(newPoint)
                } ?: 0
                ventCounts.put(newPoint, updatedCount + 1) 
            }
        }
    }
    

    // for ((key, value) in ventCounts) {
    //     println("$key = $value")
    // }
    return ventCounts.values.filter { it > 1 }.count()
}  

private fun arrangeInput(input: List<String>) {
    for (line in input) {
        val points = line.split(" -> ")
        val firstPoint = points[0].split(",")
        val secondPoint = points[1].split(",")
        val start = Pair<Int, Int>(firstPoint.first().toInt(), firstPoint.last().toInt())
        val end = Pair<Int, Int>(secondPoint.first().toInt(), secondPoint.last().toInt())
        ventLines.add(Pair<Pair<Int, Int>, Pair<Int, Int>>(start, end))
    }
}