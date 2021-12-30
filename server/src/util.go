package main

import (
	"strings"

	pb "server/src/proto"

	f "github.com/fogleman/fauxgl"
)

// parses shape into a FauxGL Mesh
func parseShape(shape string) (m f.Mesh) {
	shape = strings.ToLower(shape)
	switch shape {
	case "cube", "box":
		m = *f.NewCube()
	case "sphere", "ball":
		m = *f.NewSphere(4)
		m.SmoothNormals()
		m.Transform(f.Scale(f.V(.5, .5, .5)))
	case "cylinder", "cilinder":
		m = *f.NewCylinder(4, true)
		m.SmoothNormalsThreshold(1)
		m.Transform(f.Scale(f.V(.5, .5, 1)))
	case "cone":
		m = *f.NewCone(4, true)
		m.SmoothNormalsThreshold(1)
		m.Transform(f.Scale(f.V(.5, .5, 1)))
	case "plane":
		m = *f.NewPlane()
	case "icosahedron", "ico", "icosphere":
		m = *f.NewIcosahedron()
		m.Transform(f.Scale(f.V(.5, .5, .5)))
	}
	return m
}

// returns true if vector is {0, 0, 0}
func isVectorZero(v f.Vector) bool {
	if v.X == 0 && v.Y == 0 && v.Z == 0 {
		return true
	}
	return false
}

// pb.Vector to f.Vector converter
func v2v(v1 *pb.Vector) (v2 f.Vector) {
	v2.X = v1.X
	v2.Y = v1.Y
	v2.Z = v1.Z
	return
}
