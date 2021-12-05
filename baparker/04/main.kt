import java.io.File
import kotlin.collections.mutableListOf

var numberList: List<String> = mutableListOf()
var boardList: MutableList<MutableList<String>> = mutableListOf()
val BOARD_SIZE = 5

fun getNumbersAndBoards() {
    var lineIndex = 0
    var boardIndex = -1
    File("input.txt").forEachLine {
        if (lineIndex == 0) {
            numberList = it.split(",")
        } else {
            if (it.isEmpty()) {
                boardList.add(mutableListOf())
                boardIndex++
            } else {
                boardList.get(boardIndex).addAll(it.trim().split("\\s+".toRegex()))
            }
        }
        lineIndex++
    }
}

fun checkForLineWin(
        board: List<String>,
        start: Int,
        end: Int,
        stepSize: Int,
        indexValue: Int,
): Int {
    var numHits = 0
    for (i in start..end step stepSize) {
        if (board.get(i).startsWith("*")) {
            numHits++
        } else break
    }
    // a win returns the sum
    if (numHits == BOARD_SIZE) {
        var sum = 0
        board.forEach({ value: String ->
            if (!value.startsWith("*")) {
                sum += value.toInt()
            }
        })

        return sum * indexValue.toInt()
    }

    return -1
}

fun traverseBoard(currentNum: String, board: MutableList<String>): MutableList<String>? {
    for (index in board.indices) {
        if (board[index].equals(currentNum)) {
            // Mark the match with an *
            board[index] = "*" + board[index]
            val rowStart = index / BOARD_SIZE * BOARD_SIZE
            val xSum =
                    checkForLineWin(
                            board,
                            rowStart,
                            (rowStart + BOARD_SIZE) - 1,
                            1,
                            currentNum.toInt(),
                    )
            if (xSum > -1) {
                println(xSum)
                return board
            }

            val colStart = index % BOARD_SIZE
            val ySum =
                    checkForLineWin(
                            board,
                            colStart,
                            (BOARD_SIZE * BOARD_SIZE) - 1,
                            BOARD_SIZE,
                            currentNum.toInt(),
                    )
            if (ySum > -1) {
                println(ySum)
                return board
            }
        }
    }
    return null
}

fun traverseAllBoards(currentNum: String) {
    var removeBoards: MutableList<MutableList<String>> = mutableListOf()
    for (board in boardList) {
        val boardToRemove = traverseBoard(currentNum, board)
        if (boardToRemove != null) {
            removeBoards.add(boardToRemove)
        }
    }
    boardList.removeAll(removeBoards)
}

fun getBingo() {
    for (num in numberList) {
        traverseAllBoards(num)
    }
}

fun main() {
    getNumbersAndBoards()
    getBingo()
}
