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

  
  val output: Seq[Double] = Seq(10,12,14,16,18,20).map {
    // val v: Array[Double] = Array(expr1(_))
    // val result: Array[Double] = Array(expr1(_));
    // println(result(0));
    expr1(_)
  }
  /*
  val output: Array[Array[Double]] = Seq(10,12,14,16,18,20).map {
    val v1 = expr1(_)
    val v2 = expr2(_)
    (v1, v2)    
  }
  */
  
  for(line <- output) {
    println(line)
  }


  val ret: Array[Array[Double]] = Array();  
  return ret;
}

def output(output: Array[Array[Double]]) = {
  for(line <- output) {
    val v1 = line(0)
    val v2 = line(1)

    println(f"|$v1%10.2f|$v2%10.2f|")    
  }  
}

// os.write(os.pwd / "raw-updated.json", ujson.write(json))