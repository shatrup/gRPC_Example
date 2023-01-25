package main

import (
	"context"
	"log"
	"net"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	pb "practice.com/grpc-practice/proto"
)

const (
	PORT = ":50051"
)

type Server struct {
	pb.UnimplementedExampleServiceServer
}

func (s *Server) CreateExample(ctx context.Context, in *pb.NewExample) (*pb.Example, error) {
	log.Printf("Received: %v", in.GetName())
	todo := &pb.Example{
		Name:        in.GetName(),
		Description: in.GetDescription(),
		Done:        false,
		Id:          uuid.New().String(),
	}

	return todo, nil

}

func main() {
	res, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("Failed connection %d ", err)
	}
	s := grpc.NewServer()

	pb.RegisterExampleServiceServer(s, &Server{})
	log.Printf("Server listen at %v ", res.Addr())

	err = s.Serve(res)
	if err != nil {
		log.Fatalf("Failed to listen server %d ", err)
	}

}
