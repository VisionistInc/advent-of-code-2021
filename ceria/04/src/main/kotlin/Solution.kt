import java.io.File;

lateinit var drawList: List<Int>
var boards = mutableListOf<List<MutableList<Int>>>()

fun main(args : Array<String>) {
    arrangeInput(File(args.first()).readLines())
    println("Solution 1: ${solution1()}")
    println("Solution 2: ${solution2()}")
}

private fun solution1() :Int {
    for (x in drawList.indices) {
        markBoards(drawList.get(x))

        // at least 5 numbers (indexes 0 - 4) have been drawn, check for bingo
        var bingoBoardIndex = -1
        if (x > 3) {
            bingoBoardIndex = checkForBingo() 
        }

        // if bingoBoardIndex is not -1, we have a winning board
        if (bingoBoardIndex != -1) {
            var winningBoard = boards.get(bingoBoardIndex)
            var sum = 0 
            winningBoard.map{ sum += it.filter{ it > 0 }.sum() }
            return sum * drawList.get(x)
        }
    }
    return 0
}
    
private fun solution2() :Int {
    // doesn't matter what we set this to - this is just to initialize it.
    var lastWinningBoard = boards.get(0)
    var lastWinningDraw = drawList.get(0)

    for (x in drawList.indices) {
        markBoards(drawList.get(x))

        // at least 5 numbers (indexes 0 - 4) have been drawn, check for bingo
        if (x > 3) {
            var bingoBoardIndex = -2
            while (bingoBoardIndex != -1) {
                bingoBoardIndex = checkForBingo() 
                if (bingoBoardIndex != -1) {
                    lastWinningBoard = boards.get(bingoBoardIndex)
                    boards.removeAt(bingoBoardIndex)
                    lastWinningDraw = drawList.get(x)
                }
            }
            
        }
    }

    var sum = 0 
    lastWinningBoard.map{ sum += it.filter{ it > 0 }.sum() }

    return sum * lastWinningDraw
}  

private fun arrangeInput(input: List<String>) {
    drawList = input.get(0).split(",").map{ it.toInt() }

    for (x in 2..input.size step 6) {
        boards.add(listOf<MutableList<Int>>(
            input.get(x).trim().split("\\s+".toRegex()).map{ it.toInt() }.toMutableList(),
            input.get(x+1).trim().split("\\s+".toRegex()).map{ it.toInt() }.toMutableList(),
            input.get(x+2).trim().split("\\s+".toRegex()).map{ it.toInt() }.toMutableList(),
            input.get(x+3).trim().split("\\s+".toRegex()).map{ it.toInt() }.toMutableList(),
            input.get(x+4).trim().split("\\s+".toRegex()).map{ it.toInt() }.toMutableList()
        ))
    }
}

private fun markBoards(num: Int) {
    for (board in boards) {
        for (row in board) {
            if (row.contains(num)) {
                // find all occurrances of the number in the row (incase it's more than once) 
                // and mark it by replacing it with -1
                for (x in row.indices) {
                    if (row.get(x) == num) {
                        row.set(x, -1)
                    }
                }
            }
        }
    }
}

private fun checkForBingo() :Int {
    for (boardIdx in boards.indices) {
        val board = boards.get(boardIdx)    
        
        // horizontal check
        for (row in board) {
            if (row.all{ it < 0 }) {
                return boardIdx
            }
        }

        // vertical check
        for ( x in 0..board.get(0).size - 1) {
            var allNegative = true
            for (row in board) {
                if (row.get(x) > 0 ){
                    allNegative = false
                    break
                }
            }

            if (allNegative) {
                return boardIdx
            }
        }

    }
    return -1
}