package main

import f "github.com/fogleman/fauxgl"

const (
	scale  = 2    // optional supersampling
	width  = 1920 // output width in pixels
	height = 1080 // output height in pixels
	fovy   = 40   // vertical field of view in degrees
	near   = 1    // near clipping plane
	far    = 10   // far clipping plane
)

var (
	eye    = f.V(3, -3, 2)               // camera position
	center = f.V(0, 0, 0)                // view center position
	up     = f.V(0, 0, 1)                // up vector
	light  = f.V(.6, -.7, 1).Normalize() // light direction
	color  = f.HexColor("#8E9092")       // object color
)

func main() {

}
