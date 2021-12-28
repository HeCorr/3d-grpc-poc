package main

import (
	"bytes"
	"context"
	"image"
	"image/png"
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

// gRPC RequestRender() function
func (s *requestServer) RequestRender(ctx context.Context, req *pb.RenderRequest) (resp *pb.RenderResponse, err error) {
	// render image
	img, rTime := render(req)

	buf := bytes.Buffer{}

	err = png.Encode(&buf, img)
	if err != nil {
		return resp, err
	}

	return &pb.RenderResponse{
		ImageBytes: buf.Bytes(),
		RenderTime: rTime,
	}, nil
}

func main() {

	l, err := net.Listen("tcp", ":5656")
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

// renders image from request
func render(r *pb.RenderRequest) (image.Image, string) {
	start := time.Now()

	// create rendering context
	context := f.NewContext(width*scale, height*scale)
	context.ClearColor = f.HexColor("#3C3C3C") // bg color
	context.ClearColorBuffer()

	// set some options such as object material, etc
	aspect := float64(width) / float64(height)
	matrix := f.LookAt(eye, center, up).Perspective(fovy, aspect, near, far)
	shader := f.NewPhongShader(matrix, light, eye)
	shader.ObjectColor = color
	shader.DiffuseColor = f.Gray(0.9) // object color
	shader.SpecularColor = f.Gray(0.25)
	shader.SpecularPower = 100
	context.Shader = shader

	// loops through objects and draw them to the context
	for _, o := range r.Objects {
		pos := v2v(o.Position)
		scale := v2v(o.Scale)
		if isVectorZero(scale) {
			scale = f.V(1, 1, 1)
		}
		mesh := parseShape(o.Shape)
		mesh.Transform(f.Scale(scale))
		mesh.Transform(f.Translate(pos))
		context.DrawMesh(&mesh)
	}

	image := context.Image()
	// image = resize.Resize(width, height, image, resize.Bilinear)
	// f.SavePNG("debug.png", image)
	return image, time.Since(start).String()
}
