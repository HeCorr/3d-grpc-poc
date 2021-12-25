package main

import (
	"fmt"
	"image"
	"time"

	f "github.com/fogleman/fauxgl"
)

const (
	scale  = 1    // optional supersampling
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
	render(RenderRequest{
		Objects: []Object{
			{
				Shape: "Cube",
				Location: f.Vector{
					X: 0,
					Y: 0,
					Z: 0,
				},
			},
		},
	})
}

func render(r RenderRequest) image.Image {
	start := time.Now()

	context := f.NewContext(width*scale, height*scale)
	context.ClearColor = f.HexColor("#3C3C3C") // bg color
	context.ClearColorBuffer()

	aspect := float64(width) / float64(height)
	matrix := f.LookAt(eye, center, up).Perspective(fovy, aspect, near, far)
	shader := f.NewPhongShader(matrix, light, eye)
	shader.ObjectColor = color
	shader.DiffuseColor = f.Gray(0.9) // object color
	shader.SpecularColor = f.Gray(0.25)
	shader.SpecularPower = 100
	context.Shader = shader

	cube := f.NewCube()

	context.DrawMesh(cube)

	fmt.Println(time.Since(start))

	image := context.Image()
	// image = resize.Resize(width, height, image, resize.Bilinear)
	f.SavePNG("debug.png", image)
	return image
}
