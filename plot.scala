import scala.math

val expr1 = (n: Int) => 5 * (math.log10(n) / math.log10(2))
val expr2 = (n: Int) => math.log10(50000 * n) / math.log10(2)

@main
def Main() = {
  println("|----------|----------|")
  for(n <- Seq(10, 12,14,16,18,20)) {
    val v1 = expr1(n)
    val v2 = expr2(n)
    println(f"|$v1%10.2f|$v2%10.2f|")
  }
}
