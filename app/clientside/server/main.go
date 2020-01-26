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
	var sum int32
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			message := fmt.Sprintf("DONE: sum = %d", sum)
			return stream.SendAndClose(&pb.UploadReply{
				Message: message,
			})
		}
		if err != nil {
			return err
		}
		fmt.Println(req.GetValue())
		sum += req.GetValue()
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
