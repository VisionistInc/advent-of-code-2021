import java.io.File;

val startNodes = mutableSetOf<String>()
val endNodes = mutableSetOf<String>()
val nodePaths = mutableListOf<Pair<String, String>>()

fun main(args : Array<String>) {
    val input = File(args.first()).readLines()
    for (line in input) {
        var nodes = line.split("-")
        if (nodes[0].equals("start")) {
            startNodes.add(nodes[1])
        } else if (nodes[1].equals("start")) {
            startNodes.add(nodes[0])
        } else if (nodes[0].equals("end")) {
            endNodes.add(nodes[1])
        } else if (nodes[1].equals("end")) {
            endNodes.add(nodes[0])
        } else {
            nodePaths.add(Pair<String, String>(nodes[0], nodes[1]))
        }
    }

    println("Solution 1: ${solution1()}")
    println("Solution 2: ${solution2()}")
}

private fun solution1() :Int {
    var paths = mutableSetOf<MutableList<String>>()
    for (node in startNodes) {

        // If the start node is also an end node, add it as a single path to paths before
        // executing findPath()
        if (node in endNodes) {
            paths.add(mutableListOf<String>(node))
        }

        paths.addAll(findPath(node, mutableListOf<String>(node), mutableSetOf<MutableList<String>>(), false))
    }
   
    return paths.size
}

private fun solution2() :Int {
    var paths = mutableSetOf<MutableList<String>>()
    for (node in startNodes) {

        // If the start node is also an end node, add it as a single path to paths before
        // executing findPath()
        if (node in endNodes) {
            paths.add(mutableListOf<String>(node))
        }

        paths.addAll(findPath(node, mutableListOf<String>(node), mutableSetOf<MutableList<String>>(), true))
    }
   
    return paths.size
}

private fun findPath(n: String, path: MutableList<String>, paths: MutableSet<MutableList<String>>, extended: Boolean) :MutableSet<MutableList<String>> {
    // find all the nodes we can go to from n
    val possibles = mutableSetOf<String>() 
    for (nodePair in nodePaths) {
        if (nodePair.first.equals(n)) {
            possibles.add(nodePair.second)
        } else if (nodePair.second.equals(n)) {
            possibles.add(nodePair.first)
        }
    }

    // Iterate over all the possible paths
    for (node in possibles) {
        // copy the path thus far
        var pathCopy = ArrayList<String>(path) 

        // make sure we follow the rules about being able to visit a node
        if ((node.all{ it.isLowerCase() } && !pathCopy.contains(node)) || node.all{ it.isUpperCase() }) {
            // if this isn't the same node passed into the algorithm, add it to
            // the path copy
            if (!node.equals(n)) {
                pathCopy.add(node)
            }

            // path copy is complete - add it to the paths
            if (node in endNodes) {
                paths.add(pathCopy)
            }
            
            // continue to look for paths until the enclosing if condition isn't met
            paths.addAll(findPath(node, pathCopy, paths, extended))
        } else {
            if (extended) {
                // check to see if there are other small caves in the path at least twice
                var canUseAgain = true
                var smallCaves = pathCopy.filter{ it.all{ it.isLowerCase() } }.groupingBy{ it }.eachCount()
                if (smallCaves.values.filter{ it == 2 }.size > 0) {
                    canUseAgain = false
                }

                if (canUseAgain) {
                    // if this isn't the same node passed into the algorithm, add it to
                    // the path copy
                    if (!node.equals(n)) {
                        pathCopy.add(node)
                    }  

                    // path copy is complete - add it to the paths
                    if (node in endNodes) {
                        paths.add(pathCopy)
                    }
            
                    // continue to look for paths until the enclosing if condition isn't met
                    paths.addAll(findPath(node, pathCopy, paths, extended))
                }
            }
        }
    }

    return paths
}