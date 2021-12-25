package main

import f "github.com/fogleman/fauxgl"

type Object struct {
	Shape    string
	Location f.Vector
}

type RenderRequest struct {
	Objects []Object
}
