package main

import (
	"fmt"
	"log"
	"net"
	"time"

	pb "grpc-sample/pb/notification"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedNotificationServer
}

func (s *server) Notification(req *pb.NotificationRequest, stream pb.Notification_NotificationServer) error {
	fmt.Println("リクエスト受け取った")
	for i:=int32(0);i<req.GetNum();i++ {
		message := fmt.Sprintf("%d", i)
		if err := stream.Send(&pb.NotificationReply{
			Message: message,
		}); err != nil {
			return err
		}
		time.Sleep(time.Second * 1)
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterNotificationServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("サーバ起動失敗: %v", err)
	}
}
