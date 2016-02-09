package graphs

import (
    "fmt"
    "strconv"
)

// Element - represents svg node.
type Element string

const (
    svg         Element = "svg"
    rect        Element = "rect"
    circle      Element = "circle"
    ellipse     Element = "ellipse"
    line        Element = "line"
    polygon     Element = "polygon"
    polyline    Element = "polyline"
    path        Element = "path"
    text        Element = "text"
    g           Element = "g"
)

// Attribute - represents valid svg attributes.
type Attribute string

const (
    width       Attribute = "width"
    height      Attribute = "height"
    x           Attribute = "x"
    y           Attribute = "y"
    cx          Attribute = "cx"
    cy          Attribute = "cy"
    r           Attribute = "r"
    stroke      Attribute = "stroke"
    strokeWidth Attribute = "stroke-width"
    fill        Attribute = "fill"
)

// Style TODO Styles
type Style string

const (
    color       Style = "color"
)

// Node - a struct of an svg and its children
type Node struct {
    element         Element
    attributes      map[Attribute]string
    styles          map[Style]string
    children        []Node
}


// The various invariants on svg elements are preserved on their factory
// TODO validate nodes somewhere

// Rect creates an SVG rectangle node
func Rect(xn, yn, w, h float64) Node {
    as := make(map[Attribute]string)
    
    as[x] = strconv.FormatFloat(xn, 'f', -1, 32)
    as[y] = strconv.FormatFloat(yn, 'f', -1, 32)
    as[width] = strconv.FormatFloat(w, 'f', -1, 32)
    as[height] = strconv.FormatFloat(h, 'f', -1, 32)
    return Node{rect, as, nil, nil}
}

// Circle creates an SVG circle node
func Circle(cxn, cyn, rn float64) Node {
    as := make(map[Attribute]string)
    
    as[cx] = strconv.FormatFloat(cxn, 'f', -1, 32)
    as[cy] = strconv.FormatFloat(cyn, 'f', -1, 32)
    as[r] = strconv.FormatFloat(rn, 'f', -1, 32)
    
    return Node{circle, as, nil, nil}
}

// SVG creates an SVG
func SVG(w, h int) Node {
    as := make(map[Attribute]string)
    
    as[width] = strconv.Itoa(w)
    as[height] = strconv.Itoa(h)
    
    return Node{svg, as, nil, nil}
}


// AddChild method modifies the original node by appending a new child to it
func (n *Node) AddChild(c Node) {
    n.children = append(n.children, c)
}



// String conversion method
func (n Node) String() string {
    var cs, as, ss string
    
    // format attribute and style nodes
    for key, val := range n.attributes {
        as += fmt.Sprintf("%s='%s'", key, val)
    }
    
    ss += "style='"
    for key, val := range n.styles {
        ss += fmt.Sprintf("%s:%s;", key, val)
    }
    ss += "'"
    
    // recursively call String() to build any child representation
    for _, child := range n.children {
        cs += child.String()
    }
    if cs == "" {
        return fmt.Sprintf("<%s %s %s />", n.element, as, ss)
    }
    
    return fmt.Sprintf("<%s %s %s>%s</%s>", n.element, as, ss, cs, n.element)    
}
