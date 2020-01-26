package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "grpc-sample/pb/calc"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedCalcServer
}

func (s *server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumReply, error) {
	a := in.GetA()
	b := in.GetB()
	log.Printf("%v, %v", a, b)
	reply := fmt.Sprintf("%d + %d = %d", a, b, a+b)
	return &pb.SumReply{Message: reply}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCalcServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("サーバ起動失敗: %v", err)
	}
}
