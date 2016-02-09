package graphs

// Scale is the function signature of scale functions used internally to fit data to a given svg size
type Scale func(float64) float64

// LinearScale creates a function which scales a dataset to a given range.
// dmin, dmax - the max/min of the dataset
// min, max   - the min/max of the dimension
func LinearScale(dmin, dmax, min, max float64) Scale {
    return func(x float64) float64 {
        return ((x - dmin) * (max - min) / (dmax - dmin)) + min
    }
}

// CalcMaxMin the max, min of a slice of floats
func CalcMaxMin(xs []float64) (float64, float64) {
    min, max := xs[0], xs[0]
    
    for _, x := range xs {
        if x < min {
            min = x
            continue
        }
        
        if x > max {
            max = x
        }
    }
    
    return max, min
}