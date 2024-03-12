package main

import (
	"context"
	"google.golang.org/grpc"
	"net"
	pb "test/example/gen"
)

func main() {
	listen, err := net.Listen("tcp", ":8085")
	if err != nil {
		return
	}
	s := grpc.NewServer()
	pb.RegisterMyServiceServer(s, new(DPY))
	err = s.Serve(listen)
	if err != nil {
		return
	}
}

type DPY struct {
	pb.UnimplementedMyServiceServer
}

func (d *DPY) SendMessage(ctx context.Context, in *pb.YourMessage) (*pb.YourMessage, error) {
	return &pb.YourMessage{
		Id:   in.Id,
		Name: in.Name,
	}, nil
}
