import scala.math

val expr1 = (n: Int) => 5 * (math.log10(n) / math.log10(2))
val expr2 = (n: Int) => math.log10(50000 * n) / math.log10(2)

object MyApp {
  def Main(args: Array[String]) = {
    println("|----------|----------|")
  }
}

