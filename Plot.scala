import scala.math

val expr1 = (n: Int) => 5 * (math.log10(n) / math.log10(2))
val expr2 = (n: Int) => math.log10(50000 * n) / math.log10(2)

@main
def Main(): Array[Array[Double]] = {
  println("|----------|----------|")
  /*
  val output: Array[Array[Double]] = Seq(10,12,14,16,18,20).map {
    (_ + 1)
  }
  */

  
  val output: Array[Array[Double]] = Seq(10,12,14,16,18,20).map(n => {
    Array(expr1(n), expr2(n))
  })
  dumpJson(output)
  return output
}

def dumpJson(input: Seq[Array[Double]]) = {
  os.write(os.pwd / "data.json", ujson.write(input))  
}

def output(input: Seq[Array[Double]]) = {
  for(line <- input) {
    val v1 = line(0)
    val v2 = line(1)

    println(f"|$v1%10.2f|$v2%10.2f|")    
  }  
}

