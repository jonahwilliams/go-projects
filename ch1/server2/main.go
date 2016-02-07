package main

import (
    "image"
    "image/color"
    "image/gif"
    "io"
    "math"
    "math/rand"
    "net/http"
    "log"
    "strconv"
)


const (
    whiteIndex = 0 // First color in palette
    blackIndex = 1 // Next color in palette
)

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
    qs := r.URL.Query()
    cycles := queryDefault(qs, "cycles", 5)
    red := queryDefault(qs, "red", 0)
    green := queryDefault(qs, "green", 0)
    blue := queryDefault(qs, "blue", 0)

    col :=  color.RGBA{uint8(red), uint8(green), uint8(blue), 1}
    lissajous(w, float64(cycles), col)
}


func queryDefault(qs map[string][]string, name string, dv int) int {
    
    // first check if key exists.  return default if not
    q, ok := qs[name]
    
    if !ok {
        return dv
    }
    
    // then try and convert to an int.  return default on err
    res, err := strconv.Atoi(q[0])
    
    if err != nil {
        return dv
    }
    
    // finally return parsed value if it all worked
    return res
}



func lissajous(out io.Writer, c float64, col color.RGBA) {
    var palette = []color.Color{color.White, col}
    const (
        res     = 0.01  // angular resolution
        size    = 100   // image canvas covers [-size ... +size]
        nframes = 64    // number of animation frames
        delay   = 8     // delay between frames in 10ms units
    )
    
    freq := rand.Float64() * 3
    anim := gif.GIF{LoopCount: nframes}
    phase := 0.0
    cycles := float64(c)
    
    for i := 0; i < nframes; i++ {
        rect := image.Rect(0, 0, 2 * size + 1, 2 * size + 1)
        img := image.NewPaletted(rect, palette)
        for t := 0.0; t < cycles * 2 * math.Pi; t += res {
            x := math.Sin(t)
            y := math.Sin(t * freq + phase)
            img.SetColorIndex(size + int(x * size + 0.5), size + int(y * size + 0.5), blackIndex)
        }
        phase += 0.1
        anim.Delay = append(anim.Delay, delay)
        anim.Image = append(anim.Image, img)
    }
    gif.EncodeAll(out, &anim)
    
}