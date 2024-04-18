import plotly._
import plotly.element._
                    
val trace1 = Scatter(
  Seq(1, 2, 3, 4),
  Seq(10, 15, 13, 17)
)

val trace2 = Scatter(
  Seq(1, 2, 3, 4),
  Seq(16, 5, 11, 9)
)

val data = Seq(trace1, trace2)

Plotly.plot("div-id", data)
