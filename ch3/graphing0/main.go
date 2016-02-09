package main

import (
    "jonahwilliams/graphs"
    "math/rand"
    "log"
    "net/http"
    "fmt"
)

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}


func handler(w http.ResponseWriter, r *http.Request) {
    x := make([]float64, 50)
    y := make([]float64, 50)
    
    for i := 0; i < 50; i++ {
        x[i] = rand.Float64()
        y[i] = 2 * x[i] + (1 * rand.Float64())
    }
    
    xmax, xmin := graphs.CalcMaxMin(x)
    ymax, ymin := graphs.CalcMaxMin(y)
    
    xscale := graphs.LinearScale(xmin, xmax, 0, 400)
    yscale := graphs.LinearScale(ymin, ymax, 400, 0)
    
    xnot := make([]float64, 50)
    ynot := make([]float64, 50)
    
    for i := range x {
        xnot[i] = xscale(x[i]) 
        ynot[i] = yscale(y[i])
    }
    
    w.Header().Set("content-type","text/html")
    fmt.Fprintf(w, "<html><body>%s</body></html>", graphs.ScatterPlot(xnot, ynot))
}