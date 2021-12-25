package main

import (
	"strings"

	f "github.com/fogleman/fauxgl"
)

func parseShape(shape string) (m f.Mesh) {
	shape = strings.ToLower(shape)
	switch shape {
	case "cube", "box":
		m = *f.NewCube()
	case "sphere", "ball":
		m = *f.NewSphere(1)
	case "plane":
		m = *f.NewPlane()
	}
	return m
}
