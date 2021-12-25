package main

import f "github.com/fogleman/fauxgl"

type Object struct {
	Mesh     f.Mesh
	Location f.Vector
}

type RenderRequest struct {
	Objects []Object
}
