package main

import (
	"fmt"
	"io"
	"log"
	"net"

	pb "grpc-sample/pb/upload"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedUploadServer
}

func (s *server) Upload(stream pb.Upload_UploadServer) error {
	for {
		point, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.UploadReply{
				Message: "OK",
			})
		}
		if err != nil {
			return err
		}
		fmt.Println(point.GetValue())
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUploadServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("サーバ起動失敗: %v", err)
	}
}
