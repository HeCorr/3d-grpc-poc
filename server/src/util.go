package main

import (
	"strings"

	pb "server/src/proto"

	f "github.com/fogleman/fauxgl"
)

func parseShape(shape string) (m f.Mesh) {
	shape = strings.ToLower(shape)
	switch shape {
	case "cube", "box":
		m = *f.NewCube()
	case "sphere", "ball":
		m = *f.NewSphere(4)
		m.SmoothNormals()
	case "plane":
		m = *f.NewPlane()
	}
	return m
}

func isVectorZero(v f.Vector) bool {
	if v.X == 0 && v.Y == 0 && v.Z == 0 {
		return true
	}
	return false
}

// pb.Vector to f.Vector
func v2v(v1 *pb.Vector) (v2 f.Vector) {
	v2.X = v1.X
	v2.Y = v1.Y
	v2.Z = v1.Z
	return
}
