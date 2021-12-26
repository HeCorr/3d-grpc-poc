package main

import (
	"context"
	"fmt"
	"image"
	"log"
	"net"
	pb "server/src/proto"
	"time"

	f "github.com/fogleman/fauxgl"
	"google.golang.org/grpc"
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

type requestServer struct {
	pb.UnimplementedRendererServer
}

func (s *requestServer) RequestRender(ctx context.Context, req *pb.RenderRequest) (resp *pb.RenderReply, err error) {

	return &pb.RenderReply{
		ImageBytes: []byte{},
	}, nil
}

func main() {

	l, err := net.Listen("tcp", "localhost:1313")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterRendererServer(s, &requestServer{})

	log.Printf("server listening at %v", l.Addr())
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func render(r *pb.RenderRequest) image.Image {
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

	for _, o := range r.Objects {
		pos := f.V(float64(o.PosX), float64(o.PosY), float64(o.PosZ))
		scale := f.V(float64(o.ScaleX), float64(o.ScaleY), float64(o.ScaleZ))
		if isVectorZero(scale) {
			scale = f.Vector{1, 1, 1}
		}
		mesh := parseShape(o.Shape)
		mesh.Transform(f.Translate(pos))
		mesh.Transform(f.Scale(scale))
		context.DrawMesh(&mesh)
	}

	fmt.Println(time.Since(start))

	image := context.Image()
	// image = resize.Resize(width, height, image, resize.Bilinear)
	f.SavePNG("debug.png", image)
	return image
}
