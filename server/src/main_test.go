package main

import (
	"context"
	"os"
	pb "server/src/proto"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGrpcClient(t *testing.T) {
	file, err := os.OpenFile("render.png", os.O_CREATE, os.ModePerm)
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	conn, err := grpc.Dial("localhost:1313", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Error(err)
	}
	defer conn.Close()

	c := pb.NewRendererClient(conn)

	render, err := c.RequestRender(context.Background(), &pb.RenderRequest{
		Objects: []*pb.Object{
			{
				Shape:    "cube",
				Position: &pb.Vector{X: 0, Y: 0, Z: 0},
				Scale:    &pb.Vector{X: 1, Y: 1, Z: 1},
			}, {
				Shape:    "plane",
				Position: &pb.Vector{X: 0, Y: 0, Z: -0.6},
				Scale:    &pb.Vector{X: 2, Y: 2, Z: 1},
			},
		},
	})
	if err != nil {
		t.Error(err)
	}

	_, err = file.Write(render.ImageBytes)
	if err != nil {
		t.Error(err)
	}

}
