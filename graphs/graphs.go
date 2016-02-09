package graphs


// ScatterPlot generates an svg scatterplot
func ScatterPlot(x, y []float64) string {
    svg := SVG(400, 400)
    
    for i := range x {
        child := Circle(x[i], y[i], 2)
        svg.AddChild(child)
    }
    
    return svg.String()
}