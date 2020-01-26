package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"

	pb "grpc-sample/pb/chat"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedChatServer
}

func (s *server) Chat(stream pb.Chat_ChatServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		fmt.Println("受取：", in.GetMessage())
		message := fmt.Sprintf("%sをうけたったよ", in.GetMessage())
		if err := stream.Send(&pb.ChatReply{
			Message: message,
		}); err != nil {
			return err
		}
		time.Sleep(time.Second * 1)
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterChatServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("サーバ起動失敗: %v", err)
	}
}
