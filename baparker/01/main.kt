import java.io.File

fun main(args: Array<String>) {
  var windowSize = 1
  if (args.size > 0) {
    windowSize = args[0].toInt()
  }

  var count = 0
  val lineList = arrayListOf<Int>()

  File("input.txt").forEachLine {
    val num = it.toInt()
    lineList.add(0, num)

    if (lineList.size > windowSize && lineList.get(windowSize) < lineList.get(0)) {
      count++
    }
  }

  println(count)
}
