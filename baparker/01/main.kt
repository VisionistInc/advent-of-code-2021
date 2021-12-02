import java.io.File

fun main() {
  var count = 0
  var last: Int = -1

  File("input.txt").forEachLine {
    val num = it.toInt()

    if (last >= 0 && num > last) {
      count++
    }

    last = num
  }

  println(count)
}
